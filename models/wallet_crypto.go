package models

import (
	"github.com/astaxie/beego/orm"
)

func (wc *WalletCrypto) TableName() string {
    return "wallet_crypto"
}

func (wc *WalletCrypto) Insert() error {
	if _, err := orm.NewOrm().Insert(wc); err != nil {
		return err
	}
	return nil
}

func (wc *WalletCrypto) Read(fields ...string) error {
	if err := orm.NewOrm().Read(wc, fields...); err != nil {
		return err
	}
	return nil
}

func (wc *WalletCrypto) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(wc, field, fields...)
}

func (wc *WalletCrypto) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(wc, fields...); err != nil {
		return err
	}
	return nil
}

func (wc *WalletCrypto) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(wc, fields ...); err != nil {
		return err
	}
	return nil
}

func WalletCryptos(walletId int64) (int64, []WalletCrypto, error) {
	var table WalletCrypto
	var walletCryptos []WalletCrypto
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).OrderBy("-created_at").All(&walletCryptos)
	return num, walletCryptos, err
}

func init() {
    orm.RegisterModel(new(WalletCrypto))
}