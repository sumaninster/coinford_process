
package process

import (
	"coinford_process/models"
	"fmt"
	"time"
	//"github.com/astaxie/beego/orm"
)

type OrderGraph struct {
}

func (og *OrderGraph) Register(order *models.Order) {
	og.OrderGraph1m(order)
	og.OrderGraph5m(order)
	og.OrderGraph15m(order)
	og.OrderGraph30m(order)

	og.OrderGraph1h(order)
	og.OrderGraph6h(order)
	og.OrderGraph12h(order)

	og.OrderGraph1d(order)
	og.OrderGraph7d(order)
}

func (og *OrderGraph) OrderGraph1m(order *models.Order) {
	buyOrders, err := models.LastGraphOrder1m(order.currencyId, order.rateCurrencyId)
	fmt.Println(buyOrders, err)
	if err == orm.ErrNoRows {
		lastOrderId = 0
		og.recordDetailsAll(duration, day, timeType, currencyId, rateCurrencyId)
	} else if err == nil {
		lastOrderGraph := buyOrders[0]

		currentDate = time.Now()
		lastOrderGraph.Date - currentDate
		lastOrderGraphId = lastOrderGraph.Id

		lastOrderId = lastOrderGraph.LastOrderId

		open = lastOrderGraph.Close
		high = lastOrderGraph.Close
		low = lastOrderGraph.Close
		closeValue = lastOrderGraph.Close

		absoluteChange = 0
		percentChange = 0
		volume = 0
	}
}

func (og *OrderGraph) OrderGraph5m(order *models.Order) {
}

func (og *OrderGraph) OrderGraph15m(order *models.Order) {
}

func (og *OrderGraph) OrderGraph30m(order *models.Order) {
}

func (og *OrderGraph) OrderGraph1h(order *models.Order) {
}

func (og *OrderGraph) OrderGraph6h(order *models.Order) {
}

func (og *OrderGraph) OrderGraph12h(order *models.Order) {
}

func (og *OrderGraph) OrderGraph1d(order *models.Order) {
}

func (og *OrderGraph) OrderGraph7d(order *models.Order) {
}

func (og *OrderGraph) recordDetails(duration time.Duration, day int, timeType string, currencyId int64, rateCurrencyId int64) {
	var lastOrderGraphId, lastOrderId int64
	var open, high, low, closeValue, volume, split, dividend, absoluteChange, percentChange float64

	lastOrderGraphId = 0

	switch timeType {
		case "m":
			switch(duration) {
				case 1:
					buyOrders, err := models.LastGraphOrder1m(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderId = 0
						og.recordDetailsAll(duration, day, timeType, currencyId, rateCurrencyId)
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 5:
					buyOrders, err := models.LastGraphOrder5m(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderGraphId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 15:
					buyOrders, err := models.LastGraphOrder15m(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderGraphId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 30:
					buyOrders, err := models.LastGraphOrder30m(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderGraphId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
			}
		case "h":
			switch(duration) {
				case 1:
					buyOrders, err := models.LastGraphOrder1h(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 6:
					buyOrders, err := models.LastGraphOrder6h(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 12:
					buyOrders, err := models.LastGraphOrder12h(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderGraphId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
			}
		case "d":
			switch(day) {
				case 1:
					buyOrders, err := models.LastGraphOrder1d(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
				case 7:
					buyOrders, err := models.LastGraphOrder7d(currencyId, rateCurrencyId)
					fmt.Println(buyOrders, err)
					if err == orm.ErrNoRows {
						lastOrderId = 0
					} else if err == nil {
						lastOrderGraph := buyOrders[0]
						lastOrderGraphId = lastOrderGraph.Id

						lastOrderId = lastOrderGraph.LastOrderId

						open = lastOrderGraph.Close
						high = lastOrderGraph.Close
						low = lastOrderGraph.Close
						closeValue = lastOrderGraph.Close

						absoluteChange = 0
						percentChange = 0
						volume = 0
					}
			}
	}

	num, orders, err := models.TimeOrders(lastOrderGraphId, duration, day, timeType, currencyId, rateCurrencyId)
	fmt.Println(num, orders, err)

	if num > 0 && err == nil {

		firstOrder := orders[0]
		lastOrder := orders[0]
		highOrder := orders[0]
		lowOrder := orders[0]

		volume = firstOrder.Amount

		for _, order := range orders {
			if order.Rate > highOrder.Rate  {
				highOrder = order
			} else if order.Rate < lowOrder.Rate  {
				lowOrder = order
			}
			volume += order.Amount
			lastOrder = order
		}

		lastOrderId = lastOrder.Id

		open = firstOrder.Rate
		high = highOrder.Rate
		low = lowOrder.Rate
		closeValue = lastOrder.Rate

		absoluteChange = open - closeValue
		percentChange = (absoluteChange / open) * 100
	}

	switch timeType {
		case "m":
			switch(duration) {
				case 1:
					newOrderGraph1m := models.OrderGraph1m{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph1m.Insert()
				case 5:
					newOrderGraph5m := models.OrderGraph5m{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph5m.Insert()
				case 15:
					newOrderGraph15m := models.OrderGraph15m{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph15m.Insert()
				case 30:
					newOrderGraph30m := models.OrderGraph30m{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph30m.Insert()
			}
		case "h":
			switch(duration) {
				case 1:
					newOrderGraph1h := models.OrderGraph1h{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph1h.Insert()
				case 6:
					newOrderGraph6h := models.OrderGraph6h{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph6h.Insert()
				case 12:
					newOrderGraph12h := models.OrderGraph12h{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph12h.Insert()
			}
		case "d":
			switch(day) {
				case 1:
					newOrderGraph1d := models.OrderGraph1d{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph1d.Insert()
				case 7:
					newOrderGraph7d := models.OrderGraph7d{
					    LastOrderId:        lastOrderId,
					    CurrencyId:         currencyId,
					    RateCurrencyId:     rateCurrencyId,
					    Open:               open,
					    High:               high,
					    Low:                low,
					    Close:              closeValue,
					    Volume:             volume,
					    Split:              split,
					    Dividend:           dividend,
					    AbsoluteChange:     absoluteChange,
					    PercentChange:      percentChange}

					newOrderGraph7d.Insert()
			}
	}
}