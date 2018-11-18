package service

import (
	"sync"

	"github.com/astaxie/beego/logs"
)

type ProductCountMgr struct {
	productCount map[int]int
	lock         sync.RWMutex
}

func (p *ProductCountMgr) Count(productId int) (count int) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	count = p.productCount[productId]
	return
}

func (p *ProductCountMgr) Add(productId, count int) {
	p.lock.Lock()
	defer p.lock.Unlock()
	cur, ok := p.productCount[productId]
	if !ok {
		logs.Debug("product_id: %v, cur: %v, map: %v", productId, cur, p.productCount)
		cur = count
	} else {
		logs.Debug("else product_id: %v, cur: %v, map: %v", productId, cur, p.productCount)
		cur += count
	}
	p.productCount[productId] = cur
}

func NewProductCountMgr() (productMgr *ProductCountMgr) {
	productMgr = &ProductCountMgr{
		productCount: make(map[int]int, 128),
	}
	return
}
