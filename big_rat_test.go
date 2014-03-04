package sqlutil_test

import (
	"database/sql/driver"
	"math/big"
	"testing"

	"github.com/modcloth/sqlutil"
)

func TestBigRatValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.BigRat
		expected driver.Value
	}{
		{sqlutil.BigRat{R: *big.NewRat(2, 4), Precision: 1}, "0.5"},
		{sqlutil.BigRat{R: *big.NewRat(-2, 4), Precision: 1}, "-0.5"},
		{sqlutil.BigRat{R: *big.NewRat(1, 3), Precision: 3}, "0.333"},
		{sqlutil.BigRat{R: *big.NewRat(1, 3), Precision: 1}, "0.3"},
		{sqlutil.BigRat{R: *big.NewRat(1, 3), Precision: 0}, "0"},
	}

	for _, tt := range valueTests {
		actual, err := tt.n.Value()

		if err != nil {
			t.Errorf("%+v.Value(): got error: %+v", tt.n, err)
		}

		if actual != tt.expected {
			t.Errorf("%+v.Value(): expected %s, actual %s", tt.n, tt.expected, actual)
		}
	}
}
