package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func (o *Order) TableName() string {
    return "order"
}

func (o *Order) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *Order) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Order) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(o, field, fields...)
}

func (o *Order) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Order) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(o, fields ...); err != nil {
		return err
	}
	return nil
}

func Orders(limit int64, currencyId int64, rateCurrencyId int64) (int64, []Order, error) {
	var table Order
	var orders []Order
	num, err := orm.NewOrm().QueryTable(table).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-created_at").Limit(limit).All(&orders)
	return num, orders, err
}

func TimeOrders(lastOrderId int64, duration time.Duration, day int, timeType string, currencyId int64, rateCurrencyId int64) (int64, []Order, error) {
	var table Order
	var orders []Order
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var cond1 *orm.Condition

	cond1 = cond.And("currency_id", currencyId).And("rate_currency_id", rateCurrencyId).And("id__gt", lastOrderId)

	start := time.Now()

	switch timeType {
		case "m":
			start = start.Add(-duration * time.Minute)
		case "h":
			start = start.Add(-duration * time.Hour)
		case "d":
			start = start.AddDate(0, 0, -day)
	}
	
	cond1 = cond.And("created_at__gte", start).AndCond(cond1)
	
	num, err := qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("id").All(&orders)
	return num, orders, err
}

func MyOrders(user *User, limit int64, currencyId int64, rateCurrencyId int64) (int64, []Order, error) {
	var table Order
	var orders []Order
	num, err := orm.NewOrm().QueryTable(table).Filter("from_user_id", user.Id).Filter("currency_id", currencyId).Filter("rate_currency_id", rateCurrencyId).OrderBy("-created_at").Limit(limit).All(&orders)
	return num, orders, err
}

func init() {
	orm.RegisterModel(new(Order))
}
