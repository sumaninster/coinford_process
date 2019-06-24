package models

import (
	"github.com/astaxie/beego/orm"
)

func (ob *OrderBuy) TableName() string {
    return "order_buy"
}

func (ob *OrderBuy) Insert() error {
	if _, err := orm.NewOrm().Insert(ob); err != nil {
		return err
	}
	return nil
}

func (ob *OrderBuy) Read(fields ...string) error {
	if err := orm.NewOrm().Read(ob, fields...); err != nil {
		return err
	}
	return nil
}

func (ob *OrderBuy) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(ob, field, fields...)
}

func (ob *OrderBuy) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ob, fields...); err != nil {
		return err
	}
	return nil
}

func (ob *OrderBuy) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(ob, fields ...); err != nil {
		return err
	}
	return nil
}

func OneBuyOrder(currencyId int64, rateCurrencyId int64) (int64, []OrderBuy, error) {//userId int64, .Exclude("user_id", userId)
	var table OrderBuy
	var buyOrders []OrderBuy
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-rate", "created_at").Limit(1).All(&buyOrders)	
	return num, buyOrders, err
}

func BuyOrders(limit int64, currencyId int64, rateCurrencyId int64) (int64, []OrderBuy, error) {
	var table OrderBuy
	var buyOrders []OrderBuy
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-rate", "created_at").Limit(limit).All(&buyOrders)
	return num, buyOrders, err
}

func init() {
	orm.RegisterModel(new(OrderBuy))
}
