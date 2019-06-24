package models

import (
	"github.com/astaxie/beego/orm"
)

func (os *OrderSell) TableName() string {
    return "order_sell"
}

func (os *OrderSell) Insert() error {
	if _, err := orm.NewOrm().Insert(os); err != nil {
		return err
	}
	return nil
}

func (os *OrderSell) Read(fields ...string) error {
	if err := orm.NewOrm().Read(os, fields...); err != nil {
		return err
	}
	return nil
}

func (os *OrderSell) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(os, field, fields...)
}

func (os *OrderSell) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(os, fields...); err != nil {
		return err
	}
	return nil
}

func (os *OrderSell) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(os, fields ...); err != nil {
		return err
	}
	return nil
}

func OneSellOrder(currencyId int64, rateCurrencyId int64) (int64, []OrderSell, error) {
	var table OrderSell
	var sellOrders []OrderSell
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("rate", "created_at").Limit(1).All(&sellOrders)
	return num, sellOrders, err
}

func SellOrders(limit int64, currencyId int64, rateCurrencyId int64) (int64, []OrderSell, error) {
	var table OrderSell
	var sellOrders []OrderSell
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("rate", "created_at").Limit(limit).All(&sellOrders)
	return num, sellOrders, err
}

func init() {
	orm.RegisterModel(new(OrderSell))
}
