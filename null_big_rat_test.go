package sqlutil_test

import (
	"database/sql/driver"
	"math/big"
	"testing"

	"github.com/modcloth/sqlutil"
)

func TestNullBigRatValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.NullBigRat
		expected driver.Value
	}{
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{R: *big.NewRat(2, 4), Precision: 1}, Valid: true}, "0.5"},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{R: *big.NewRat(2, 4), Precision: 1}, Valid: false}, nil},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{R: *big.NewRat(1, 3), Precision: 1}, Valid: true}, "0.3"},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{R: *big.NewRat(-2, 4), Precision: 1}, Valid: true}, "-0.5"},
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
