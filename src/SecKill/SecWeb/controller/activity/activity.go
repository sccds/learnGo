package activity

import (
	"SecKill/SecWeb/model"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ActivityController struct {
	beego.Controller
}

func (p *ActivityController) ListActivity() {
	activityModel := model.NewActivityModel()
	activityList, err := activityModel.GetActivityList()
	if err != nil {
		logs.Warn("get activity list failed, err: %v", err)
		return
	}
	p.Data["activity_list"] = activityList
	p.TplName = "activity/list.html"
	p.Layout = "layout/layout.html"
}

func (p *ActivityController) CreateActivity() {
	p.TplName = "activity/create.html"
	p.Layout = "layout/layout.html"
}

func (p *ActivityController) SubmitActivity() {
	activityModel := model.NewActivityModel()
	var activity model.Activity
	p.TplName = "activity/list.html"
	p.Layout = "layout/layout.html"
	var err error
	var Error string = "succ"
	defer func() {
		if err != nil {
			p.Data["Error"] = Error
			p.TplName = "activity/error.html"
		}
	}()

	// submit data
	name := p.GetString("activity_name")
	if len(name) == 0 {
		Error = "activity name is empty"
		err = fmt.Errorf(Error)
		return
	}
	productId, err := p.GetInt("product_id")
	if err != nil {
		err = fmt.Errorf("invalid product_id, err: %v", err)
		Error = err.Error()
		return
	}
	startTime, err := p.GetInt64("start_time")
	if err != nil {
		err = fmt.Errorf("invalid start_time, err: %v", err)
		Error = err.Error()
		return
	}
	endTime, err := p.GetInt64("end_time")
	if err != nil {
		err = fmt.Errorf("invalid end_time, err: %v", err)
		Error = err.Error()
		return
	}
	total, err := p.GetInt("total")
	if err != nil {
		err = fmt.Errorf("invalid total number, err: %v", err)
		Error = err.Error()
		return
	}
	speed, err := p.GetInt("speed")
	if err != nil {
		err = fmt.Errorf("invalid speed, err: %v", err)
		Error = err.Error()
		return
	}
	limit, err := p.GetInt("buy_limit")
	if err != nil {
		err = fmt.Errorf("invalid buy_limit, err: %v", err)
		Error = err.Error()
		return
	}
	buyRate, err := p.GetFloat("buy_rate")
	if err != nil {
		err = fmt.Errorf("invalid buy_rate, err: %v", err)
		Error = err.Error()
		return
	}
	activity.ActivityName = name
	activity.ProductId = productId
	activity.StartTime = startTime
	activity.EndTime = endTime
	activity.Total = total
	activity.Speed = speed
	activity.BuyLimit = limit
	activity.BuyRate = buyRate
	err = activityModel.CreateActivity(&activity)
	if err != nil {
		err = fmt.Errorf("create activity failed, err: %v", err)
		Error = err.Error()
		return
	}
	logs.Debug("create and submit activity[%v] succ", activity)
	return
}
