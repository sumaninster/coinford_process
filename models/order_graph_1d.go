package models

import (
	"github.com/astaxie/beego/orm"
)

func (og *OrderGraph1d) TableName() string {
    return "order_graph_1d"
}

func (og *OrderGraph1d) Insert() error {
	if _, err := orm.NewOrm().Insert(og); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1d) Read(fields ...string) error {
	if err := orm.NewOrm().Read(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1d) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(og, field, fields...)
}

func (og *OrderGraph1d) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(og, fields...); err != nil {
		return err
	}
	return nil
}

func (og *OrderGraph1d) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(og, fields ...); err != nil {
		return err
	}
	return nil
}

func LastGraphOrder1d(currencyId int64, rateCurrencyId int64) ([]OrderGraph1d, error) {
	var table OrderGraph1d
	var orders []OrderGraph1d
	err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-date").One(&orders)
	return orders, err
}

func GraphOrders1d(lastGraphOrderId int64, currencyId int64, rateCurrencyId int64) (int64, []OrderGraph1d, error) {
	var table OrderGraph1d
	var orders []OrderGraph1d
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).Filter("id__gt", lastGraphOrderId).OrderBy("date").All(&orders)
	return num, orders, err
}

func init() {
	orm.RegisterModel(new(OrderGraph1d))
}
