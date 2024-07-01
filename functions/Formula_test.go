package skills_test

import (
	"testing"
	skills "skills/functions"
	"math/big"
)

var testmedian = []struct{
	name string
	data []*big.Float
	expected *big.Int
}{
	{
		"test 1",
		[]*big.Float{big.NewFloat(5)},
		big.NewInt(5),
	},
	{
		"test 2",
		[]*big.Float{big.NewFloat(4),big.NewFloat(6),big.NewFloat(8),big.NewFloat(10),big.NewFloat(12),big.NewFloat(14),big.NewFloat(16),big.NewFloat(18)},
		big.NewInt(11),
	},
	{
		"test 3",
		[]*big.Float{big.NewFloat(4),big.NewFloat(6),big.NewFloat(8),big.NewFloat(10),big.NewFloat(12),big.NewFloat(14),big.NewFloat(16)},
		big.NewInt(10),
	},
	{
		"test 4",
		[]*big.Float{big.NewFloat(4),big.NewFloat(6),big.NewFloat(8),big.NewFloat(11),big.NewFloat(12),big.NewFloat(14),big.NewFloat(16), big.NewFloat(18)},
		big.NewInt(12),
	},
}

var TestDeviation = []struct{
	name string
	data *big.Float
	expected *big.Int
}{
	{
		"test 1",
		big.NewFloat(1),
		big.NewInt(1),

	},
	{
		"test 2",
		big.NewFloat(32478601),
		big.NewInt(5699),
	},
	{
		"test 3",
		big.NewFloat(15649996.556),
		big.NewInt(3956),

	},
	{
		"test 4",
		big.NewFloat(15649999678),
		big.NewInt(125100),

	},

}

var testVariance = []struct {
	name string
	array []*big.Float
	times *big.Float
	average *big.Float
	expected *big.Float
}{
	{
		"test 1",
		[]*big.Float{big.NewFloat(4)},
		big.NewFloat(1),
		big.NewFloat(4),
		big.NewFloat(0),
	},
	{
		"test 2",
		[]*big.Float{big.NewFloat(6), big.NewFloat(8), big.NewFloat(10), big.NewFloat(12), big.NewFloat(14)},
		big.NewFloat(5),
		big.NewFloat(10),
		big.NewFloat(8),
	},
	{
		"test 3",
		[]*big.Float{big.NewFloat(4), big.NewFloat(8), big.NewFloat(10), big.NewFloat(12)},
		big.NewFloat(4),
		big.NewFloat(8),
		big.NewFloat(9),
	},
}

func TestMedian(t *testing.T) {
	for _ , test := range testmedian {
		t.Run(test.name, func(t *testing.T) {
			got, err := skills.Median(test.data)
			if err != nil {
				t.Errorf("function returned an error")
			}
			want := test.expected
			if want.Cmp(got) == -1 || want.Cmp(got) == 1 {
				t.Errorf("wanted %v but got %v ", want, got)
			}
		})
	} 
}

func TestSatndardDeviation(t *testing.T) {
	for _ , test := range TestDeviation {
		t.Run(test.name, func(t *testing.T) {
			got := skills.StandardDeviation(test.data)
			want := test.expected
			if want.Cmp(got) == -1 || want.Cmp(got) == 1 {
				t.Errorf("wanted %v but got %v ", want, got)
			}
		})
	} 
}

func TestVariance(t *testing.T) {
	for _ , test := range testVariance {
		t.Run(test.name, func(t *testing.T) {
			got := skills.Variance(test.array, test.times, test.average)
			want := test.expected
			if want.Cmp(got) == -1 || want.Cmp(got) == 1 {
				t.Errorf("wanted %v but got %v ", want, got)
			}
		})
	} 
}