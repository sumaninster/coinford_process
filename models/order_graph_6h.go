package models

import (
	"github.com/astaxie/beego/orm"
)

func (og *OrderGraph6h) TableName() string {
    return "order_graph_6h"
}

func (og *OrderGraph6h) Insert() error {
	if _, err := orm.NewOrm().Insert(og); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph6h) Read(fields ...string) error {
	if err := orm.NewOrm().Read(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph6h) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(og, field, fields...)
}

func (og *OrderGraph6h) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph6h) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(og, fields ...); err != nil {
		return err
	}
	return nil
}

func LastGraphOrder6h(currencyId int64, rateCurrencyId int64) ([]OrderGraph6h, error) {
	var table OrderGraph6h
	var orders []OrderGraph6h
	err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-date").One(&orders)
	return orders, err
}

func GraphOrders6h(lastGraphOrderId int64, currencyId int64, rateCurrencyId int64) (int64, []OrderGraph6h, error) {
	var table OrderGraph6h
	var orders []OrderGraph6h
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).Filter("id__gt", lastGraphOrderId).OrderBy("date").All(&orders)
	return num, orders, err
}

func init() {
	orm.RegisterModel(new(OrderGraph6h))
}
