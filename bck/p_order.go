package process

import (
	"coinford_api/models"
	"coinford_api/configs"
	"fmt"
	"time"
)

func ProcessOrders() {
	for {
		processSellOrders()
		fmt.Println("ProcessOrders: ", time.Now())
		time.Sleep(5 * time.Second)
	}
}

func processSellOrders() {
	var buyOrders, sellOrders []models.Order
	var sellOrder models.Order
	sellNum, err := models.GetSellOrders(&sellOrders)
	if err == nil && sellNum > 0 {
		for _, sellOrder = range sellOrders {
			isPairSupported, exchangeType := models.IsCurrencyPairSupported(sellOrder.CurrencyId, sellOrder.RateCurrencyId)
			if isPairSupported {
				buyNum, err := models.GetBuyOrders(&sellOrder, &buyOrders)
				if err == nil && buyNum > 0 {
					bestDeal(&sellOrder, &buyOrders, exchangeType)
				}
			}
		}
	}
}

func bestDeal(sellOrder *models.Order, buyOrders *[]models.Order, exchangeType string) {
	var buyOrder models.Order
	for _, mbuyOrder := range *buyOrders {
		fmt.Println("list Buy: ", mbuyOrder.Rate, "Amount: ",  mbuyOrder.Amount)
	}

	first := true
	for _, mbuyOrder := range *buyOrders {
		if sellOrder.Amount <= mbuyOrder.Amount {
			buyOrder = mbuyOrder
			first = false
		} else {
			if first {
				buyOrder = mbuyOrder
			}
			break
		}
	}
	fmt.Println("Sell: ", sellOrder.Rate, "Amount: ", sellOrder.Amount)
	fmt.Println("Buy : ", buyOrder.Rate , "Amount: ", buyOrder.Amount)

	switch exchangeType {
	case "Digital":
		processDigital(sellOrder, &buyOrder)
	case "Fiat":
		processDigital(sellOrder, &buyOrder)
		//processFiat(sellOrder, &buyOrder)
	}
}

func processDigital(sellOrder *models.Order, buyOrder *models.Order) {
	fmt.Printf("Sell: %f , Amount: %f || Buy: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount, buyOrder.Rate, buyOrder.Amount)
	if walletBalanceSufficient(sellOrder, buyOrder) {
		fmt.Println("walletBalanceSufficient")
		if sellOrder.Amount <= buyOrder.Amount {
			fmt.Println("sellOrder.Amount <= buyOrder.Amount")
			models.CreditWallet(buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyId)
			models.DebitWallet(sellOrder.UserId, sellOrder.Amount, sellOrder.CurrencyId)
			recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, sellOrder.Amount, sellOrder.CurrencyId)

			models.CreditWallet(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)
			models.DebitWallet(buyOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)
			recordTransaction(buyOrder, buyOrder.UserId, sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)

			if sellOrder.Amount < buyOrder.Amount {
				postNewOrder(buyOrder, buyOrder.Amount - sellOrder.Amount)
			}
		} else if sellOrder.Amount > buyOrder.Amount {
			fmt.Println("sellOrder.Amount > buyOrder.Amount")
			models.CreditWallet(buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyId)
			models.DebitWallet(sellOrder.UserId, buyOrder.Amount, sellOrder.CurrencyId)
			recordTransaction(sellOrder, sellOrder.UserId, buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyId)

			models.CreditWallet(sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)
			models.DebitWallet(buyOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)
			recordTransaction(buyOrder, buyOrder.UserId, sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId)

			postNewOrder(sellOrder, sellOrder.Amount - buyOrder.Amount)
		}
	}
}

func walletBalanceSufficient(sellOrder *models.Order, buyOrder *models.Order) bool {
	seller_num, seller_wallets, err1 := models.CurrencyWallets(sellOrder.UserId, sellOrder.CurrencyId)
	buyer_num, buyer_wallets, err2 := models.CurrencyWallets(buyOrder.UserId, buyOrder.RateCurrencyId)

	if seller_num > 0 && buyer_num > 0 && err1 == nil && err2 == nil {
		var seller_total, buyer_total float64
		for _, sw := range seller_wallets {
			seller_total += sw.Amount
		}
		for _, bw := range buyer_wallets {
			buyer_total += bw.Amount
		}
		if seller_total >= sellOrder.Amount && buyer_total >= (sellOrder.Amount * sellOrder.Rate) {
			fmt.Println("seller total: ", seller_total, "buyer total: ", buyer_total)
			fmt.Println("seller amount: ", sellOrder.Amount, "buyer amount: ", (sellOrder.Amount * sellOrder.Rate))
			return true
		}
		return false
	}
	return false
}

func recordTransaction(order *models.Order, from int64, to int64, amount float64, currency int64) {
	if order != nil {
		transaction := models.Transaction{OrderId: order.Id, FromUser: from, ToUser: to, Amount: amount, CurrencyId: currency, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		transaction.Insert()
		order.ProcessedAt = time.Now()
		order.UpdatedAt = time.Now()
		order.ProcessedType = "FULL"
		order.Update()
	}
}

func postNewOrder(order *models.Order, amount float64) {
	if order != nil {
		order.ProcessedType = "PARTIAL"
		order.Update()
		orderchild := models.Order{UserId: order.UserId, CurrencyId: order.CurrencyId, Amount: amount, Rate: order.Rate, RateCurrencyId: order.RateCurrencyId, OrderType: order.OrderType, ProcessedType: "NOT_PROCESSED", ProcessedAt: *configs.NullTime, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
		orderchild.Insert()
		mapPartialOrder(order.Id, orderchild.Id)
	}
}

func mapPartialOrder(parentOrderId int64, childOrderId int64) {
	orderpartial := models.OrderPartial{ParentOrderId: parentOrderId, ChildOrderId: childOrderId, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *configs.NullTime}
	orderpartial.Insert()
}

func processFiat(sellOrder *models.Order, buyOrder *models.Order) {
	//Not Supported
}

func debugMessage(tag string, err1 error, err2 error) {
	if models.Runmode == "dev" {
		fmt.Println(tag, err1, err2)
	}
}
