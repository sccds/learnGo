package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ProductId   int    `db:"id"`
	ProductName string `db:"name"`
	Total       int    `db:"total"`
	Status      int    `db:"status"`
}

type ProductModel struct {
	Db *sqlx.DB
}

func NewProductModel() *ProductModel {
	productModel := &ProductModel{}
	return productModel
}

func (p *ProductModel) GetProductList() (list []Product, err error) {
	sql := "SELECT id, name, total, status FROM product"
	err = Db.Select(&list, sql)
	if err != nil {
		logs.Warn("select from mysql failed, err: %v", err)
		return
	}
	return
}

func (p *ProductModel) CreateProduct(product *Product) (err error) {
	sql := "INSERT INTO product(name, total, status) VALUES(?, ?, ?)"
	_, err = Db.Exec(sql, product.ProductName, product.Total, product.Status)
	if err != nil {
		logs.Warn("insert into mysql failed, err: %v", err)
		return
	}
	logs.Debug("insert into database succ")
	return
}
