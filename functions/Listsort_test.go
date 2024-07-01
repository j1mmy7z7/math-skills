package skills_test

import (
	"math/big"
	"reflect"
	skills "skills/functions"
	"testing"
)

var testcases = []struct {
	name  string
	data  []*big.Float
	expected  []*big.Float
}{
	{
	"test 1",
	[]*big.Float{big.NewFloat(-2),big.NewFloat(-44),big.NewFloat(5),big.NewFloat(99),big.NewFloat(-56),big.NewFloat(56),big.NewFloat(77),big.NewFloat(-100)},
	[]*big.Float{big.NewFloat(-100),big.NewFloat(-56),big.NewFloat(-44),big.NewFloat(-2),big.NewFloat(5),big.NewFloat(56),big.NewFloat(77),big.NewFloat(99)},
	},
	{
	"test 2",
	[]*big.Float{},
	[]*big.Float{},
	},
	{
	"test 3",
	[]*big.Float{big.NewFloat(5)},
	[]*big.Float{big.NewFloat(5)},
	},
	{
	"test 4",
	[]*big.Float{big.NewFloat(-2),big.NewFloat(2),big.NewFloat(5),big.NewFloat(-5),big.NewFloat(-10),big.NewFloat(10),big.NewFloat(100),big.NewFloat(-100)},
	[]*big.Float{big.NewFloat(-100),big.NewFloat(-10),big.NewFloat(-5),big.NewFloat(-2),big.NewFloat(2),big.NewFloat(5),big.NewFloat(10),big.NewFloat(100)},
	
	},
}

func TestQuicksort(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			skills.QuickSort(tc.data, 0, len(tc.data)-1)
			if !reflect.DeepEqual(tc.data, tc.expected) {
				t.Errorf("got %v wnated %v", tc.data, tc.expected)
			}

		})
	}
}