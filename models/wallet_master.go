package models

import (
	"github.com/astaxie/beego/orm"
)

func (wm *WalletMaster) TableName() string {
    return "wallet_master"
}

func (wm *WalletMaster) Insert() error {
	if _, err := orm.NewOrm().Insert(wm); err != nil {
		return err
	}
	return nil
}

func (wm *WalletMaster) Read(fields ...string) error {
	if err := orm.NewOrm().Read(wm, fields...); err != nil {
		return err
	}
	return nil
}

func (wm *WalletMaster) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(wm, field, fields...)
}

func (wm *WalletMaster) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(wm, fields...); err != nil {
		return err
	}
	return nil
}

func (wm *WalletMaster) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(wm, fields ...); err != nil {
		return err
	}
	return nil
}

func WalletMasters(userId int64) (int64, []WalletMaster, error) {
	var table WalletMaster
	var walletMasters []WalletMaster
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).OrderBy("-created_at").All(&walletMasters)
	return num, walletMasters, err
}


func CurrencyWallets(walletId int64, currency int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).Filter("currency_id", currency).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&wallets)
	return num, wallets, err
}

func PrimaryCurrencyWallets(walletId int64, currency int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).Filter("currency_id", currency).Filter("primary", true).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&wallets)
	return num, wallets, err
}

func init() {
    orm.RegisterModel(new(WalletMaster))
}