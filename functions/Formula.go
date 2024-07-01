package skills

import (
	"errors"
	"math/big"
)

func Median(array []*big.Float) (*big.Int, error) {
	final := new(big.Int)
	n := len(array)
    // if the array is empty
	if n == 0 {
		return nil, errors.New("this array is empty")
	}
    // if the length of the array is not even
	if n%2 != 0 {
		array[n/2].Int(final)
		return final, nil
	} else {
		// For even length, return the average of the two middle elements
		median := new(big.Float).Add(array[n/2-1], array[n/2])
		median.Quo(median, big.NewFloat(2))
		median.Add(median, big.NewFloat(0.5))
		median.Int(final)
		return final, nil
	}
}

func Variance(array []*big.Float, times *big.Float, average *big.Float) *big.Float {
	sum := new(big.Float)
	for _, v := range array {
		initial := new(big.Float).Sub(v, average)
		initial.Mul(initial, initial)
		sum.Add(sum, initial)
	}
	sum.Quo(sum, times)
	return sum
}

func StandardDeviation(Variance *big.Float) *big.Int {
	final := new(big.Int)
    Variance.Sqrt(Variance)
    Variance.Add(Variance, big.NewFloat(0.5))
    Variance.Int(final)
    return final
}
