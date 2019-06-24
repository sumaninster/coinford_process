package process

import (
	"coinford_process/models"
	"coinford_process/configs"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
)

type WalletController struct {
	beego.Controller
}

func (w *WalletController) Add(userId int64, currencyId int64, currencyType string, nickname string, primary bool) error {
	walletMaster := models.WalletMaster{UserId: userId, CurrencyId: currencyId, CurrencyType: currencyType}
	_, _, err := walletMaster.ReadOrCreate("user_id", "currency_id")
	if err == nil {
		notDeleted := true
		if walletMaster.DeletedAt != *configs.NullTime {
			walletMaster.DeletedAt = *configs.NullTime
			err = walletMaster.Update()
			notDeleted = false
		}
		if err == nil {
			if notDeleted {
				w.updateStatus(walletMaster.Id, &nickname, &primary)
				wallet := models.Wallet{WalletMasterId: walletMaster.Id, Amount: 0, AmountLocked: 0, Nickname: nickname, Primary: primary}
				err := wallet.Insert()
				if err == nil {
					err = w.addWalletCrypto(wallet.Id, walletMaster.CurrencyId)
				}
			}
		}
	}
	return err
}

func (w *WalletController) updateStatus(walletMasterId int64, nickname *string, primary *bool) {
	var num int64
	var err error
	if *primary {
		num, _, err = models.WalletPrimaryFalse(walletMasterId)
	} else {
		num, _, err = models.Wallets(walletMasterId)
	}

	if num == 0 && err == nil {
		if *nickname == *configs.NullString {
			*nickname = "Primary"
		}
		*primary = true
	} else {
		if *nickname == *configs.NullString {
			*nickname = "No Name"
		}
	}
}

func (w *WalletController) addWalletCrypto(walletId int64, CurrencyId int64) error {
	switch CurrencyId {
		case models.GetCurrencyId("BTC"):
			return w.addBTCWalletCrypto(walletId)
		case models.GetCurrencyId("ETH"):
			return w.addETHWalletCrypto(walletId)
		default:
			return nil
	}
	return nil
}

func (w *WalletController) addBTCWalletCrypto(walletId int64) error {
	btcprocess := BTCProcess{}
	address, err := btcprocess.GetNewAddress()
	fmt.Println(address)
	if err == nil {
		walletaddress := models.WalletCrypto{WalletId: walletId, Address: address}
		err = walletaddress.Insert()
		fmt.Println(err)
		return err
	} else {
		return err
	}
	return nil
}

func (w *WalletController) addETHWalletCrypto(walletId int64) error {
	expirationTime := time.Now().Add(time.Hour * time.Duration(configs.ETHPasswordExpiryHour))
	passphrase, _ := w.token(expirationTime.Unix())
	//passphrase := configs.RandString(100)
	//fmt.Println(passphrase)
	walletPassphrase := models.WalletPassphrase{WalletId: walletId, Passphrase: passphrase}
	err := walletPassphrase.Insert()
	if err == nil {
		ethprocess := ETHProcess{}
		address, err := ethprocess.GetNewAddress(passphrase)
		if err == nil {
			walletAddress := models.WalletCrypto{WalletId: walletId, Address: address}
			err = walletAddress.Insert()
			fmt.Println(err)
			return err
		} else {
			return err
		}
	}
	return nil
}

func (w *WalletController) token(expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, 
    jwt.MapClaims{
        "exp": expirationTime,
        "iat": time.Now().Unix() })

    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(configs.SignBytes)
	w.debugMessage("ParseRSAPrivateKeyFromPEM Error: ", err)
	if err != nil {return "Invalid Token", err}

    tokenString, err := token.SignedString(privateKey)
	w.debugMessage("SignedString Error: ", err)
	if err != nil {return "Invalid Token", err}

	return tokenString, nil
}

func (w *WalletController) debugMessage(tag string, err error) {
	if err != nil && models.Runmode == "dev" {
		fmt.Println(tag, err)
	}
}