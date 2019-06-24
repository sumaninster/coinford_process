package models

import (
	"github.com/astaxie/beego/orm"
)

func (og *OrderGraph15m) TableName() string {
    return "order_graph_15m"
}

func (og *OrderGraph15m) Insert() error {
	if _, err := orm.NewOrm().Insert(og); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph15m) Read(fields ...string) error {
	if err := orm.NewOrm().Read(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph15m) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(og, field, fields...)
}

func (og *OrderGraph15m) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph15m) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(og, fields ...); err != nil {
		return err
	}
	return nil
}

func LastGraphOrder15m(currencyId int64, rateCurrencyId int64) ([]OrderGraph15m, error) {
	var table OrderGraph15m
	var orders []OrderGraph15m
	err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-date").One(&orders)
	return orders, err
}

func GraphOrders15m(lastGraphOrderId int64, currencyId int64, rateCurrencyId int64) (int64, []OrderGraph15m, error) {
	var table OrderGraph15m
	var orders []OrderGraph15m
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).Filter("id__gt", lastGraphOrderId).OrderBy("date").All(&orders)
	return num, orders, err
}

func init() {
	orm.RegisterModel(new(OrderGraph15m))
}
