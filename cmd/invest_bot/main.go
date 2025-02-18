package main

import (
	"fmt"
	"os"

	"github.com/parchinskiy/TinkoffInvestBot.git/internal/configuration"
)

func main() {
	file, err := os.Open("configs/main.yaml")
	if err != nil {
		fmt.Println("Error when open configs/main.yaml:", err)
		return
	}
	settings, err := configuration.Load(os.ReadFile(file))
	if err != nil {
		fmt.Println("Error when open configs/main.yaml:", err)
		return
	}

	current_stocks := tinkoff.RequestCurrentStocks()
	// update database of current actives
	for {
		// если не время торговать, то спать до открытия
		// calculation best for buy(many variants) по отдельным типам активов
		// выбор случайного типа активов(но в соответствии с процентами, то есть в соответствии с распределением вероятностей) и случайны конкретный тикет(тоже с распределением)
		// buy through tinkoff api
		// update database
		// sleep until next wakeup
	}
	fmt.Println("line")
}
