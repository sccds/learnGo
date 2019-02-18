package tailf

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

const (
	StatusNormal = "normal"
	StatusDelete = "delete"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TextMsg struct {
	Msg   string
	Topic string
}

// 一个日志收集器对应一个tail
type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   string
	exitChan chan int
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func InitTail(conf []CollectConf, chanSize int) (err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invalid conf")
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}
	for _, v := range conf {
		createNewTask(v)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename: %s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailObj.conf.Topic,
			}
			tailObjMgr.msgChan <- textMsg
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf: %v", tailObj.conf)
			return
		}
	}
}

func UpdateConfig(confs []CollectConf) (err error) {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()
	for _, oneConf := range confs {
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}
		if isRunning {
			continue
		}
		createNewTask(oneConf)
	}

	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}
	tailObjMgr.tailObjs = tailObjs
	return
}

func createNewTask(conf CollectConf) {
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}
	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if errTail != nil {
		logs.Error("collect filename[%s] failed, err: %v", conf.LogPath, errTail)
		return
	}
	obj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)
	go readFromTail(obj)
}
