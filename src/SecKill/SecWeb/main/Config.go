package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MysqlConfig struct {
	UserName string
	Passwd   string
	Port     int
	Database string
	Host     string
}

type EtcdConf struct {
	Addr          string
	EtcdKeyPrefix string
	ProductKey    string
	Timeout       int
}

var (
	AppConf Config
)

type Config struct {
	mysqlConf MysqlConfig
	etcdConf  EtcdConf
	LogPath   string
	LogLevel  string
}

func initConfig() (err error) {
	// init logger
	AppConf.LogPath = beego.AppConfig.String("log_path")
	AppConf.LogLevel = beego.AppConfig.String("log_level")

	// init mysql
	username := beego.AppConfig.String("mysql_user_name")
	if len(username) == 0 {
		logs.Error("load config of mysql_user_name failed is null")
		return
	}
	AppConf.mysqlConf.UserName = username

	mysql_pw := beego.AppConfig.String("mysql_passwd")
	if len(mysql_pw) == 0 {
		logs.Error("load config of mysql_passwd failed is null")
		return
	}
	AppConf.mysqlConf.Passwd = mysql_pw

	mysql_host := beego.AppConfig.String("mysql_host")
	if len(mysql_host) == 0 {
		logs.Error("load config of mysql_host failed is null")
		return
	}
	AppConf.mysqlConf.Host = mysql_host

	mysql_database := beego.AppConfig.String("mysql_database")
	if len(mysql_database) == 0 {
		logs.Error("load config of mysql_database failed is null")
		return
	}
	AppConf.mysqlConf.Database = mysql_database

	port, err := beego.AppConfig.Int("mysql_port")
	if err != nil {
		logs.Error("load config of mysql_port failed, err: %v", err)
		return
	}
	AppConf.mysqlConf.Port = port

	// etcd config
	addr := beego.AppConfig.String("etcd_addr")
	if len(mysql_database) == 0 {
		logs.Error("load config of etcd_addr failed is null")
		return
	}
	AppConf.etcdConf.Addr = addr

	prefix := beego.AppConfig.String("etcd_sec_key_prefix")
	if len(prefix) == 0 {
		logs.Error("load config of etcd_sec_key_prefix failed is null")
		return
	}
	AppConf.etcdConf.EtcdKeyPrefix = prefix

	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0 {
		logs.Error("load config of etcd_product_key failed is null")
		return
	}
	AppConf.etcdConf.ProductKey = productKey

	timeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		logs.Error("load config of etcd_timeout failed, err: %v", err)
		return
	}
	AppConf.etcdConf.Timeout = timeout

	return
}
