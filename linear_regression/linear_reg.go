package linear_regression

import (
	"math/rand"
	"time"
)

type Regression struct {
	X []float64
	Y []float64
	yHat []float64

	dataLength int
	b0, b1, yBar, xBar,
	xSqaure, ySquare,
	meanSquaredError float64
}

// Generates a random data set
func (linReg *Regression) CreateDataSet(n int, xMin float64, xMax float64, yMin float64, yMax float64) {

	// set the seed
	rand.Seed(time.Now().UnixNano())

	linReg.X = make([]float64, n)
	linReg.Y = make([]float64, n)
	linReg.dataLength = n

	// load in the randomly generated data
	for i := 0; i < linReg.dataLength; i++ {
		linReg.X[i] = xMin + rand.Float64()*(xMax-xMin)
		linReg.Y[i] = yMin + rand.Float64()*(yMax-yMin)

	}

}

// Calculates the X mean
func (linReg *Regression) GetXMean() float64 {
	for i := range linReg.X {
		linReg.xBar += linReg.X[i]
	}

	return linReg.xBar / float64(linReg.dataLength)
}

// Calculates the Y mean
func (linReg *Regression) GetYMean() float64 {
	for i := range linReg.Y {
		linReg.yBar += linReg.Y[i]
	}

	return linReg.yBar / float64(linReg.dataLength)
}

// Calculates b1
func (linReg *Regression) Getb1() float64 {

	var diff float64 = 0.0
	var xSquaredDiff float64 = 0.0
	for i := 0; i < linReg.dataLength; i++ {
		diff += (linReg.X[i] - linReg.xBar)*(linReg.Y[i] - linReg.yBar)
		xSquaredDiff += (linReg.X[i] - linReg.xBar)*(linReg.X[i] - linReg.xBar)
	}

	// finally calculate the value of b1
	linReg.b1 = diff / xSquaredDiff
	return linReg.b1
}

// Calculates b0
func (linReg *Regression) Getb0() float64 {
	linReg.b0 = linReg.yBar - linReg.b1*linReg.xBar
	return linReg.b0
}


// Calculates yHat(fitted model)
func (linReg *Regression) FitModel() {

	for i := 0; i < linReg.dataLength; i++ {
		linReg.yHat[i] = linReg.b0 + linReg.X[i]*linReg.b1
	}
}


// Calculates the MSE(mean squared error)
func (linReg *Regression) MSE() float64 {
	linReg.meanSquaredError = 0.0

	for i := 0; i < linReg.dataLength; i++ {
		linReg.meanSquaredError += (linReg.Y[i] - linReg.yHat[i])*(linReg.Y[i] - linReg.yHat[i])
	}

	return linReg.meanSquaredError
}
