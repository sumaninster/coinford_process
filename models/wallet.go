package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"errors"
)

func (w *Wallet) TableName() string {
    return "wallet"
}

func (w *Wallet) Insert() error {
	if _, err := orm.NewOrm().Insert(w); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) Read(fields ...string) error {
	if err := orm.NewOrm().Read(w, fields...); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(w, field, fields...)
}

func (w *Wallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(w, fields...); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(w, fields ...); err != nil {
		return err
	}
	return nil
}

func Wallets(walletMasterId int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_master_id", walletMasterId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&wallets)
	return num, wallets, err
}

func WalletPrimaryFalse(walletMasterId int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_master_id", walletMasterId).OrderBy("-updated_at").All(&wallets)
	if err == nil && num > 0 {
		for _, v := range wallets {
			var update = false
			if v.Primary == true {
				v.Primary = false
				update = true
			}
			if v.Nickname == "Primary" {
				v.Nickname = "No Name"
				update = true
			}
			if update {
				v.Update()
			}
		}
	}
	return num, wallets, err
}

func CheckWallet(userId int64, currencyId int64, walletId int64) (Wallet, error) {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == currencyId {
				err = nil
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return wallet, err
}

func CreditWallet(userId int64, amount float64, currencyId int64, walletId int64) error {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		fmt.Println(walletMaster.UserId, userId, walletMaster.CurrencyId, currencyId)
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == currencyId {
				wallet.Amount += amount
				err = wallet.Update()
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return err
}

func DebitWallet(userId int64, amount float64, lockAmount float64, currencyId int64, walletId int64) error {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		fmt.Println(walletMaster.UserId, userId, walletMaster.CurrencyId, currencyId)
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == currencyId { 
				if wallet.Amount >= amount && wallet.AmountLocked >= lockAmount {
					wallet.Amount -= amount
					wallet.AmountLocked -= lockAmount
					wallet.Update()
				} else {
					errors.New("Insufficient amount locked in the wallet")
				}
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return err
}

func init() {
    orm.RegisterModel(new(Wallet))
}
