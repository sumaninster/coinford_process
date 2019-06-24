package models

import (
	"github.com/astaxie/beego/orm"
)

func (og *OrderGraph1h) TableName() string {
    return "order_graph_1h"
}

func (og *OrderGraph1h) Insert() error {
	if _, err := orm.NewOrm().Insert(og); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1h) Read(fields ...string) error {
	if err := orm.NewOrm().Read(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1h) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(og, field, fields...)
}

func (og *OrderGraph1h) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1h) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(og, fields ...); err != nil {
		return err
	}
	return nil
}

func LastGraphOrder1h(currencyId int64, rateCurrencyId int64) ([]OrderGraph1h, error) {
	var table OrderGraph1h
	var orders []OrderGraph1h
	err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-date").One(&orders)
	return orders, err
}

func GraphOrders1h(lastGraphOrderId int64, currencyId int64, rateCurrencyId int64) (int64, []OrderGraph1h, error) {
	var table OrderGraph1h
	var orders []OrderGraph1h
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).Filter("id__gt", lastGraphOrderId).OrderBy("date").All(&orders)
	return num, orders, err
}

func init() {
	orm.RegisterModel(new(OrderGraph1h))
}
