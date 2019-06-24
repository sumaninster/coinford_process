
/*

func PrimaryWallet(userId int64, currencyId int64) (Wallet, error) {
	walletMaster := WalletMaster{UserId: userId, CurrencyId: currencyId}
	err := walletMaster.Read("user_id", "currency_id")
	var wallet Wallet
	if err == nil {
		wallet = Wallet{WalletMasterId: walletMaster.Id, Primary: true}
		err = wallet.Read("wallet_master_id", "primary")
	}
	return wallet, err
}

func CheckWallettLockAmount(userId int64, currencyId int64, walletId int64, amount float64) (Wallet, error) {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == currencyId {
				if (wallet.Amount - wallet.AmountLocked) >= amount {
					wallet.AmountLocked += amount
					wallet.Update()
				} else {
					err = errors.New("Insufficient balance in primary wallet")
				}
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return wallet, err
}

func CheckWallettUnlockAmount(userId int64, currencyId int64, walletId int64, amount float64) (Wallet, error) {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == currencyId {
				if wallet.AmountLocked >= amount {
					wallet.AmountLocked -= amount
					wallet.Update()
				} else {
					errors.New("Insufficient amount locked in primary wallet")
				}
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return wallet, err
}

func PrimaryWalletLockAmount(userId int64, currencyId int64, amount float64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == nil {
		if (wallet.Amount - wallet.AmountLocked) >= amount {
			wallet.AmountLocked += amount
			wallet.Update()
		} else {
			errors.New("Insufficient balance in primary wallet")
		}
	} else {
		return err
	}
	return nil
}

func CreditWallet(userId int64, amount float64, currencyId int64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	//fmt.Println("CreditWallet", userId, amount, currencyId, wallet, "\n")
	if err == nil {
		wallet.Amount += amount
		err = wallet.Update()
		return err
	}
	return err
}

func DebitWallet(userId int64, amount float64, lockAmount float64, currencyId int64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == orm.ErrNoRows {
		fmt.Printf("Not row found")
	}
	//fmt.Println("DebitWallet", userId, wallet, "\n")
	if err == nil {
		wallet.Amount -= amount
		wallet.AmountLocked -= lockAmount
		err = wallet.Update()
		return err
	}
	return err
}*/

//sellerUserId int64, buyerUserId int64, rate float64, currencyId int64, rateCurrencyId int64

//sellOrder.UserId, buyOrder.UserId, sellOrder.Rate, sellOrder.CurrencyId, sellOrder.RateCurrencyId

//sellOrder.UserId, buyOrder.UserId, sellOrder.Rate, sellOrder.CurrencyId, sellOrder.RateCurrencyId