package process

import (
	"coinford_process/models"
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type OrderProcess struct {
}

func (op *OrderProcess) ProcessOrders(currencyId int64, rateCurrencyId int64) {
	for {
		op.matchOrder(currencyId, rateCurrencyId)
		fmt.Println("ProcessOrders: ", time.Now())
		time.Sleep(3 * time.Second)
	}
}

func (op *OrderProcess) matchOrder(currencyId int64, rateCurrencyId int64) {
	sellNum, sellOrders, sellErr := models.OneSellOrder(currencyId, rateCurrencyId)
	buyNum, buyOrders, buyErr := models.OneBuyOrder(currencyId, rateCurrencyId)
	if sellErr == nil && sellNum > 0 && buyErr == nil && buyNum > 0{

			buyOrder := buyOrders[0]
			sellOrder := sellOrders[0]

		if sellOrder.Rate <= buyOrder.Rate && buyOrder.Lock == false && sellOrder.Lock == false{

			buyOrder.Lock = true
			buyOrder.Update()
			sellOrder.Lock = true
			sellOrder.Update()

			op.processOrder(&sellOrder, &buyOrder)
		} else {

			/*buyOrder.Lock = false
			sellOrder.Lock = false
			buyOrder.Update()
			sellOrder.Update()*/

		}
	} else {
		fmt.Println(sellErr, buyErr)
	}
}

func (op *OrderProcess) processOrder(sellOrder *models.OrderSell, buyOrder *models.OrderBuy) {
	fmt.Printf("Sell: %f , Amount: %f || Buy: %f , Amount: %f \n", sellOrder.Rate, sellOrder.Amount, buyOrder.Rate, buyOrder.Amount)
	if op.walletBalanceSufficient(sellOrder, buyOrder) {
		fmt.Println("walletBalanceSufficient")
		if sellOrder.Amount <= buyOrder.Amount {
			fmt.Println("sellOrder.Amount <= buyOrder.Amount")
			errB := models.CreditWallet(buyOrder.UserId, sellOrder.Amount, buyOrder.CurrencyId, buyOrder.CurrencyWalletId)
			errS := models.CreditWallet(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId, sellOrder.RateCurrencyWalletId)
			if errB == nil && errS == nil {
				errB = models.DebitWallet(buyOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.Amount * buyOrder.Rate, buyOrder.RateCurrencyId, buyOrder.RateCurrencyWalletId)
				errS = models.DebitWallet(sellOrder.UserId, sellOrder.Amount, sellOrder.Amount, sellOrder.CurrencyId, sellOrder.CurrencyWalletId)
			} else {
				//fmt.Println("CreditWallet Buy 1", errB, "CreditWallet Sell 1", errS)
				if errB == orm.ErrNoRows {
					fmt.Printf("Not row found")
					wc := WalletController{}
					wc.Add(buyOrder.UserId, sellOrder.Amount, buyOrder.CurrencyId, "CRYPTO", "", true)
				}
				if errS == orm.ErrNoRows {
					fmt.Printf("Not row found")
					wc := WalletController{}
					wc.Add(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId, "CRYPTO", "", true)
				}
			}

			if errB == nil && errS == nil {
				err := op.recordOrder(buyOrder, sellOrder, sellOrder.Amount, sellOrder.Amount * sellOrder.Rate)
				if err == nil {
					if sellOrder.Amount < buyOrder.Amount {
						op.postNewOrderBuy(buyOrder, buyOrder.Amount - sellOrder.Amount)
						sellOrder.Delete()
					} else if sellOrder.Amount == buyOrder.Amount {
						sellOrder.Delete()
						buyOrder.Delete()
					}
				}
			}
		} else if sellOrder.Amount > buyOrder.Amount {
			fmt.Println("sellOrder.Amount > buyOrder.Amount")
			errB := models.CreditWallet(buyOrder.UserId, buyOrder.Amount, sellOrder.CurrencyId, buyOrder.CurrencyWalletId)
			errS := models.CreditWallet(sellOrder.UserId, buyOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId, sellOrder.RateCurrencyWalletId)

			if errB == nil && errS == nil {
				errB = models.DebitWallet(buyOrder.UserId, buyOrder.Amount * sellOrder.Rate, buyOrder.Amount * buyOrder.Rate, buyOrder.RateCurrencyId, buyOrder.RateCurrencyWalletId)
				errS = models.DebitWallet(sellOrder.UserId, buyOrder.Amount, buyOrder.Amount, sellOrder.CurrencyId, sellOrder.CurrencyWalletId)
			} else {
				//fmt.Println("CreditWallet Buy 2", errB, "CreditWallet Sell 2", errS)
				if errB == orm.ErrNoRows {
					fmt.Printf("Not row found")
					wc := WalletController{}
					wc.Add(buyOrder.UserId, sellOrder.Amount, buyOrder.CurrencyId, "CRYPTO", "", true)
				}
				if errS == orm.ErrNoRows {
					fmt.Printf("Not row found")
					wc := WalletController{}
					wc.Add(sellOrder.UserId, sellOrder.Amount * sellOrder.Rate, sellOrder.RateCurrencyId, "CRYPTO", "", true)
				}
			}
			fmt.Println("CreditWallet Buy 2", errB, "CreditWallet Sell 2", errS)
			if errB == nil && errS == nil {
				err := op.recordOrder(buyOrder, sellOrder, buyOrder.Amount, buyOrder.Amount * sellOrder.Rate)
				if err == nil {
					op.postNewOrderSell(sellOrder, sellOrder.Amount - buyOrder.Amount)
					buyOrder.Delete()
				}
			}
		}
	}
}

func (op *OrderProcess) walletBalanceSufficient(sellOrder *models.OrderSell, buyOrder *models.OrderBuy) bool {
	sellerWallet, errS := models.CheckWallet(sellOrder.UserId, sellOrder.CurrencyId, sellOrder.CurrencyWalletId)
	buyerWallet, errB := models.CheckWallet(buyOrder.UserId, buyOrder.RateCurrencyId, buyOrder.RateCurrencyWalletId)

	fmt.Println("sellerPrimaryWallet: ", sellerWallet.Amount, "buyerPrimaryWallet: ", buyerWallet.Amount)
	if errS == nil && errB == nil {
		if sellerWallet.Amount >= sellOrder.Amount && buyerWallet.Amount >= (buyOrder.Amount * sellOrder.Rate) {
			fmt.Println("seller total: ", sellerWallet.Amount, "buyer total: ", buyerWallet.Amount)
			fmt.Println("seller amount: ", sellOrder.Amount, "buyer amount: ", (sellOrder.Amount * sellOrder.Rate))
			return true
		}
		return false
	}
	return false
}

func (op *OrderProcess) postNewOrderBuy(order *models.OrderBuy, amount float64) {
	order.Amount = amount
	order.TotalAmount = amount * order.Rate
	order.Lock = false
	order.Update()
}

func (op *OrderProcess) postNewOrderSell(order *models.OrderSell, amount float64) {
	order.Amount = amount
	order.TotalAmount = amount * order.Rate
	order.Lock = false
	order.Update()
}

func (op *OrderProcess) recordOrder(buyer *models.OrderBuy, seller *models.OrderSell, amount float64, totalAmount float64) error {
	order := models.Order{SellerUserId: seller.UserId, BuyerUserId: buyer.UserId, 
		Amount: amount, Rate: seller.Rate, TotalAmount: totalAmount, 
		CurrencyId: seller.CurrencyId, RateCurrencyId: seller.RateCurrencyId,
		SellerCurrencyWalletId: seller.CurrencyWalletId, SellerRateCurrencyWalletId: seller.RateCurrencyWalletId,
		BuyerCurrencyWalletId: buyer.CurrencyWalletId, BuyerRateCurrencyWalletId: buyer.RateCurrencyWalletId}
	err := order.Insert()
	if err == nil {
		orderGraph := OrderGraph{}
		go orderGraph.Register(order)
	}
	return err
}

func (op *OrderProcess) debugMessage(tag string, err1 error, err2 error) {
	if models.Runmode == "dev" {
		fmt.Println(tag, err1, err2)
	}
}
