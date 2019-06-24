func CreditWallet(wallet_id int64, amount float64, currency int64) error {
	var cwallet Wallet
	num, wallets, err := PrimaryCurrencyWallets(wallet_id, currency)
	if err == nil {
		if num > 0 {
			cwallet = wallets[0]
		} else {
			num, wallets, err := CurrencyWallets(wallet_id, currency)
			if err == nil {
				if num > 0 {
					cwallet = wallets[0]
				} else {
					//nickname := "Primary"
					//primary := "YES"
					cwallet = Wallet{WalletMasterId: wallet_id, Amount: amount, AmountLocked: 0}
					cwallet.Insert()
				}
			} else {
				return err
			}
		}
	} else {
		return err
	}
	cwallet.Amount += amount
	cwallet.Update()
	return nil
}

func DebitWallet(wallet_id int64, amount float64, currency int64) error {
	_, wallets, err := CurrencyWallets(wallet_id, currency)
	if err == nil {
		for _, dw := range wallets {
			if amount > 0 {
				if dw.Amount >= amount {
					dw.Amount -= amount
					if dw.AmountLocked >= amount {
						dw.AmountLocked -= amount
					} else {
						dw.AmountLocked = 0
					}
					amount = 0
					dw.Update()
				} else {
					amount -= dw.Amount
					dw.Amount = 0
					dw.AmountLocked = 0
					dw.Update()
				}
			}
		}
	} else {
		return err
	}
	return nil
}

func (op *OrderProcess) postNewOrderBuy(order *models.OrderBuy, amount float64) {
	if order != nil {
		totalAmount := amount * order.Rate
		//err := models.PrimaryWalletLockAmount(order.UserId, order.RateCurrencyId, totalAmount)
		//if err == nil {
			orderBuy := models.OrderBuy{UserId: order.UserId, CurrencyId: order.CurrencyId, Amount: amount, Rate: order.Rate, RateCurrencyId: order.RateCurrencyId, TotalAmount: totalAmount}
			orderBuy.Insert()
		//}
	}
}

func (op *OrderProcess) postNewOrderSell(order *models.OrderSell, amount float64) {
	if order != nil {
		totalAmount := amount * order.Rate
		//err := models.PrimaryWalletLockAmount(order.UserId, order.CurrencyId, amount)
		//if err == nil {
			orderSell := models.OrderSell{UserId: order.UserId, CurrencyId: order.CurrencyId, Amount: amount, Rate: order.Rate, RateCurrencyId: order.RateCurrencyId, TotalAmount: totalAmount}
			orderSell.Insert()
		//}
	}
}

func (op *OrderProcess) ProcessOrders(limit int64, currencyId int64, rateCurrencyId int64) {
	fmt.Println(limit, currencyId, rateCurrencyId)
	for {
		op.processSellOrders(limit, currencyId, rateCurrencyId)
		fmt.Println("ProcessOrders: ", time.Now())
		time.Sleep(3 * time.Second)
	}
}

func (op *OrderProcess) processSellOrders(limit int64, currencyId int64, rateCurrencyId int64) {
	sellNum, sellOrders, sellErr := models.SellOrders(limit, currencyId, rateCurrencyId)

	if sellErr == nil && sellNum > 0 {
		//for _, sellOrder := range sellOrders {
			op.matchBuyOrders(&sellOrders[0], limit, currencyId, rateCurrencyId)
		//}
	}
}

func (op *OrderProcess) matchBuyOrders(sellOrder *models.OrderSell, limit int64, currencyId int64, rateCurrencyId int64) {
	buyNum, buyOrders, buyErr := models.BuyOrders(limit, currencyId, rateCurrencyId)
	if buyErr == nil && buyNum > 0 {
		for _, mbuyOrder := range buyOrders {
			fmt.Println("list Buy: ", mbuyOrder.Rate, "Amount: ",  mbuyOrder.Amount)
		}//log print

		for _, buyOrder := range buyOrders {
			if sellOrder.Rate <= buyOrder.Rate {
				fmt.Println("Sell: ", sellOrder.Rate, "Amount: ", sellOrder.Amount)
				fmt.Println("Buy : ", buyOrder.Rate , "Amount: ", buyOrder.Amount)
				if sellOrder.UserId != buyOrder.UserId {
					op.processOrder(sellOrder, &buyOrder)
				}
			}
		}
	}
}

func (op *OrderProcess) matchOrder(currencyId int64, rateCurrencyId int64) {
	sellNum, sellOrders, sellErr := models.OneSellOrder(currencyId, rateCurrencyId)
	//fmt.Println(sellOrders)
	if sellErr == nil && sellNum > 0 {
		sellOrder := sellOrders[0]
		buyNum, buyOrders, buyErr := models.OneBuyOrder(currencyId, rateCurrencyId)//sellOrder.UserId, 
		//fmt.Println(buyOrders)
		if buyErr == nil && buyNum > 0 {
			buyOrder := buyOrders[0]
			if sellOrder.Rate <= buyOrder.Rate {
				//if sellOrder.UserId != buyOrder.UserId {
					//fmt.Println(sellOrder, buyOrder)
					op.processOrder(&sellOrder, &buyOrder)
				//}
			}
		}
	}
}