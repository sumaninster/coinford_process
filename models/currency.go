package models

import (
	"github.com/astaxie/beego/orm"
)

func (e *Currency) TableName() string {
    return "currency"
}

func (e *Currency) Read(fields ...string) error {
	if err := orm.NewOrm().Read(e, fields...); err != nil {
		return err
	}
	return nil
}

func Currencies(currencyType string) (int64, []Currency, error) {
	var table Currency
	var currencies []Currency
	var num int64
	var err error

	if currencyType == "FIAT" || currencyType == "CRYPTO" {
		num, err = orm.NewOrm().QueryTable(table).Filter("type", currencyType).Filter("deleted_at__isnull", true).OrderBy("id").All(&currencies)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("id").All(&currencies)	
	}
	return num, currencies, err
}

func GetCurrencyId(currency string) int64 {
	e := Currency{Code: currency}
	err := e.Read("code")
	if err == nil {
		return e.Id
	}
	return 0
}

func init() {
    orm.RegisterModel(new(Currency))
}