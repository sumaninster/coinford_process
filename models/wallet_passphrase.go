package models

import (
	"github.com/astaxie/beego/orm"
)

func (wp *WalletPassphrase) TableName() string {
    return "wallet_passphrase"
}

func (wp *WalletPassphrase) Insert() error {
	if _, err := orm.NewOrm().Insert(wp); err != nil {
		return err
	}
	return nil
}

func (wp *WalletPassphrase) Read(fields ...string) error {
	if err := orm.NewOrm().Read(wp, fields...); err != nil {
		return err
	}
	return nil
}

func (wp *WalletPassphrase) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(wp, field, fields...)
}

func (wp *WalletPassphrase) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(wp, fields...); err != nil {
		return err
	}
	return nil
}

func (wp *WalletPassphrase) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(wp, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
    orm.RegisterModel(new(WalletPassphrase))
}