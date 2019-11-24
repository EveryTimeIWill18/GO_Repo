package main

import (
	"ProjOne/linear_regression"
	"fmt"
)

func main() {
	fmt.Println("Running app ...")
	// create an instance of the linear regression model
	linearReg := linear_regression.Regression{}
	linearReg.CreateDataSet(50, 0.0, 150.123, 0.0, 1000.599)
	for i := range linearReg.X {
		fmt.Println("x[", i, "] = ", linearReg.X[i], "|", "y[", i, "] = ", linearReg.Y[i])
		fmt.Println()
	}

	fmt.Println("xMean: ", linearReg.GetXMean())
	fmt.Println("yMean", linearReg.GetYMean())
}
