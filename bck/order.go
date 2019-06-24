package models

import (
	"github.com/astaxie/beego/orm"
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

func UserOrders(userId int64, orderType string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error

	if orderType == "BUY" || orderType == "SELL" {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("order_type", orderType).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
	}
	return num, orders, err
}

func UserOrdersIsProcessed(userId int64, orderType string, isProcessed string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error

	if orderType == "BUY" || orderType == "SELL" {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("order_type", orderType).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("order_type", orderType).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		}
	} else {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		}
	}
	return num, orders, err
}

func Orders(userId int64, orderType string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error
	icurrency_ids := UserCurrencyIds(userId)

	if orderType == "BUY" || orderType == "SELL" {
		num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
	}
	return num, orders, err
}

func OrdersIsProcessed(userId int64, orderType string, isProcessed string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error
	icurrency_ids := UserCurrencyIds(userId)

	if orderType == "BUY" || orderType == "SELL" {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		}
	} else {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
		}
	}
	return num, orders, err
}

func OrdersExcludeMine(userId int64, orderType string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error
	icurrency_ids := UserCurrencyIds(userId)

	if orderType == "BUY" || orderType == "SELL" {
		num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
	}
	return num, orders, err
}

func OrdersIsProcessedExcludeMine(userId int64, orderType string, isProcessed string) (int64, []Order, error) {
	var table Order
	var orders []Order
	var num int64
	var err error
	icurrency_ids := UserCurrencyIds(userId)

	if orderType == "BUY" || orderType == "SELL" {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("order_type", orderType).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
		}
	} else {
		if isProcessed == "YES" {
			num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("processed_at__isnull", true).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
		} else {
			num, err = orm.NewOrm().QueryTable(table).Filter("currency_id__in", icurrency_ids...).Filter("rate_currency_id__in", icurrency_ids...).Filter("deleted_at__isnull", true).Exclude("user_id", userId).OrderBy("-updated_at").All(&orders)
		}
	}
	return num, orders, err
}

func GetSellOrders(sellOrders *[]Order) (int64, error) {
	var table Order
	return orm.NewOrm().QueryTable(table).Filter("order_type", "SELL").Filter("deleted_at__isnull", true).Filter("processed_at__isnull", true).OrderBy("rate", "created_at").All(sellOrders)
}

func GetBuyOrders(sellOrder *Order, buyOrders *[]Order) (int64, error) {
	var table Order
	return orm.NewOrm().QueryTable(table).Filter("order_type", "BUY").Filter("currency_id", sellOrder.CurrencyId).Filter("rate__gte", sellOrder.Rate).Filter("rate_currency_id", sellOrder.RateCurrencyId).Exclude("user_id", sellOrder.UserId).Filter("deleted_at__isnull", true).Filter("processed_at__isnull", true).OrderBy("-amount", "rate", "created_at").All(buyOrders)
}

func init() {
    orm.RegisterModel(new(Order))
}