package main

import (
	"fmt"
	"math"
	"github.com/themoriarty/yfinance"
	"github.com/themoriarty/fanalysis"
	"github.com/grd/stat"
	)

func main(){
	yf := yfinance.Interface{}
	allPrices, _ := yf.GetPrices([]string{"MSFT"}, yfinance.Date(2009, 1, 1), yfinance.Date(2009, 12, 31))
	prices, _ := allPrices.Prices("MSFT")
	events := fanalysis.FindEvents(prices, func(today yfinance.Price, history fanalysis.History) bool{
		if len(history.Prices) < 20{
			return false
		}
		yesterday, _ := history.Yesterday()
		diff := today.AdjustedClose - yesterday.AdjustedClose
		dev := stat.Absdev(history.LastNDays(20))
		if math.Abs(float64(diff)) > 1.5 * dev{
			fmt.Println(today, diff, 1.5 * dev)
			return true
		}
		return false
		
	});	
	fmt.Println(events)
}