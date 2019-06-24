package main

import (
	"coinford_process/process"
	"coinford_process/models"
	//"fmt"
)

func Process() {
	numC, currencies, errC := models.Currencies("CRYPTO")
	numRC, rateCurrencies, errRC := models.Currencies("ALL")

	if errC == nil && numC > 0 && errRC == nil && numRC > 0 {
		for _, currency := range currencies {
			for _, rateCurrency := range rateCurrencies {
				if (currency.Id != rateCurrency.Id) {
					
					orderProcess := process.OrderProcess{}
					orderProcess.ProcessOrders(currency.Id, rateCurrency.Id)

					orderGraph := process.OrderGraph{}
					go orderGraph.RecordStats(1, 0, "m", currency.Id, rateCurrency.Id)
					go orderGraph.RecordStats(5, 0, "m", currency.Id, rateCurrency.Id)
					go orderGraph.RecordStats(15, 0, "m", currency.Id, rateCurrency.Id)
					go orderGraph.RecordStats(30, 0, "m", currency.Id, rateCurrency.Id)

					go orderGraph.RecordStats(1, 0, "h", currency.Id, rateCurrency.Id)
					go orderGraph.RecordStats(6, 0, "h", currency.Id, rateCurrency.Id)
					go orderGraph.RecordStats(12, 0, "h", currency.Id, rateCurrency.Id)
					
					go orderGraph.RecordStats(1, 1, "d", currency.Id, rateCurrency.Id)
					   orderGraph.RecordStats(7, 7, "d", currency.Id, rateCurrency.Id)
				}
			}
		}
	}
}

func main() {
	//Process()
	//orderProcess := process.OrderProcess{}
	//go orderProcess.ProcessOrders(1, 8)

	orderGraph := process.OrderGraph{}
	orderGraph.RecordStats(1, 0, "m", 1, 8)
	/*go orderGraph.RecordStats(5, 0, "m", 1, 8)
	go orderGraph.RecordStats(15, 0, "m", 1, 8)
	go orderGraph.RecordStats(30, 0, "m", 1, 8)

	go orderGraph.RecordStats(1, 0, "h", 1, 8)
	go orderGraph.RecordStats(6, 0, "h", 1, 8)
	go orderGraph.RecordStats(12, 0, "h", 1, 8)

	go orderGraph.RecordStats(1, 1, "d", 1, 8)
	   orderGraph.RecordStats(7, 7, "d", 1, 8)*/
}
