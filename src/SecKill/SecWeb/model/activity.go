package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

const (
	ActivityStatusNormal  = 0
	ActivityStatusDisable = 1
)

type SecProductInfoConf struct {
	ProductId         int
	StartTime         int64
	EndTime           int64
	Status            int
	Total             int
	Left              int
	OnePersonBuyLimit int
	BuyRate           float64
	SoldLimit         int
	//secLimit         *SecLimit
}

type Activity struct {
	ActivityId   int    `db:"id"`
	ActivityName string `db:"name"`
	ProductId    int    `db:"product_id"`
	StartTime    int64  `db:"start_time"`
	EndTime      int64  `db:"end_time"`
	Total        int    `db:"total"`
	Status       int    `db:"status"`
	StartTimeStr string
	EndTimeStr   string
	StatusStr    string
	Speed        int     `db:"sec_speed"`
	BuyLimit     int     `db:"buy_limit"`
	BuyRate      float64 `db:"buy_rate"`
}

type ActivityModel struct {
}

func NewActivityModel() *ActivityModel {
	return &ActivityModel{}
}

func (p *ActivityModel) GetActivityList() (activityList []*Activity, err error) {
	sql := "select id, name, product_id, start_time, end_time, total, status, sec_speed, buy_limit, buy_rate from activity order by id desc"
	err = Db.Select(&activityList, sql)
	if err != nil {
		logs.Error("select activity list from database failed, err: %v", err)
		return
	}
	for _, v := range activityList {
		st := time.Unix(v.StartTime, 0)
		v.StartTimeStr = st.Format("2006-01-02 15:04:05")
		et := time.Unix(v.EndTime, 0)
		v.EndTimeStr = et.Format("2006-01-02 15:04:05")
		now := time.Now().Unix()
		if now > v.EndTime {
			v.StatusStr = "activity end"
			continue
		}
		if v.Status == ActivityStatusDisable {
			v.StatusStr = "activity disabled"
		} else if v.Status == ActivityStatusNormal {
			v.StatusStr = "activity normal"
		}
	}
	logs.Debug("get activity succ, activity list is [%v]", activityList)
	return
}

func (p *ActivityModel) ProductValid(productId int, total int) (valid bool, err error) {
	sql := "select id, name, total, status from  product where id=?"
	var productList []*Product
	err = Db.Select(&productList, sql, productId)
	if err != nil {
		logs.Warn("select product by id failed, err: %v", err)
		return
	}
	if len(productList) == 0 {
		err = fmt.Errorf("product[%v] is not exists", productId)
		return
	}
	if total > productList[0].Total {
		err = fmt.Errorf("product[%v] total number is invalid", productId)
		return
	}
	valid = true
	return
}

func (p *ActivityModel) CreateActivity(activity *Activity) (err error) {
	valid, err := p.ProductValid(activity.ProductId, activity.Total)
	if err != nil {
		logs.Error("product exists failed, err: %v", err)
		return
	}
	if !valid {
		logs.Error("product id [%v] err: %v", activity.ProductId, err)
		return
	}
	if activity.StartTime <= 0 || activity.EndTime <= 0 {
		logs.Error("start_time[%v] or endtime[%v] err", activity.StartTime, activity.EndTime)
		return
	}
	if activity.StartTime >= activity.EndTime {
		logs.Error("start_time[%v] greater than endtime[%v] err", activity.StartTime, activity.EndTime)
		return
	}
	now := time.Now().Unix()
	//if activity.EndTime < now || activity.StartTime <= now {
	if activity.EndTime < now {
		logs.Error("start_time[%v] endtime[%v] less than now[%v]", activity.StartTime, activity.EndTime, now)
		return
	}
	sql := "INSERT INTO activity(name, product_id, start_time, end_time, total, sec_speed, buy_limit, buy_rate) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = Db.Exec(sql, activity.ActivityName, activity.ProductId, activity.StartTime, activity.EndTime, activity.Total, activity.Speed, activity.BuyLimit, activity.BuyRate)
	if err != nil {
		logs.Warn("insert into mysql failed, err: %v", err)
		return
	}
	logs.Debug("insert into database activity succ")

	// 添加成功后，把 activity 写入 etcd
	err = p.SyncToEtcd(activity)
	if err != nil {
		logs.Warn("sync to etcd failed, err: %v, data: %v", err, activity)
		return
	}
	return
}

func (p *ActivityModel) SyncToEtcd(activity *Activity) (err error) {
	if strings.HasSuffix(EtcdPrefix, "/") == false {
		EtcdPrefix = EtcdPrefix + "/"
	}
	etcdKey := fmt.Sprintf("%s%s", EtcdPrefix, EtcdProdcutKey)
	secProductInfoList, err := loadProductFromEtcd(etcdKey)
	var secProductInfo SecProductInfoConf
	secProductInfo.StartTime = activity.StartTime
	secProductInfo.EndTime = activity.EndTime
	secProductInfo.ProductId = activity.ProductId
	secProductInfo.SoldLimit = activity.Speed
	secProductInfo.OnePersonBuyLimit = activity.BuyLimit
	secProductInfo.Status = activity.Status
	secProductInfo.Total = activity.Total
	secProductInfo.BuyRate = activity.BuyRate
	secProductInfoList = append(secProductInfoList, secProductInfo)
	data, err := json.Marshal(secProductInfoList)
	if err != nil {
		logs.Error("json marshal failed, err: %v", err)
		return
	}
	_, err = EtcdClient.Put(context.Background(), etcdKey, string(data))
	if err != nil {
		logs.Error("put data[%v] to etcd failed, err %v", string(data), err)
		return
	}
	logs.Debug("put to etcd succ, data: %v", string(data))
	return
}

func loadProductFromEtcd(key string) (secProductInfo []SecProductInfoConf, err error) {
	logs.Debug("start get product from etcd")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := EtcdClient.Get(ctx, key)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err: %v", key, err)
		return
	}
	logs.Debug("get from etcd succ, resp: %v", resp)
	for k, v := range resp.Kvs {
		logs.Debug("key[%v] value[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("unmarshal from etcd failed, err: %v", err)
			return
		}
		logs.Debug("sec info conf is [%v]", secProductInfo)
	}
	return
}
