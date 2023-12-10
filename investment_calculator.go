package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 //初始化變數型態
	var expectedReturnRate = 5.5
	var years float64 = 10 //初始化變數型態

	//var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years) //當變數型態改變, 可省略轉換

	fmt.Println(futureValue)
}
