package sqlutil_test

import (
	"database/sql/driver"
	"errors"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/modcloth/sqlutil"
)

func TestNullBigRatValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.NullBigRat
		expected driver.Value
	}{
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(2, 4), Precision: 1}, Valid: true}, "0.5"},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(2, 4), Precision: 1}, Valid: false}, nil},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(1, 3), Precision: 1}, Valid: true}, "0.3"},
		{sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(-2, 4), Precision: 1}, Valid: true}, "-0.5"},
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

func TestNullBigRatScan(t *testing.T) {
	var tests = []struct {
		n        driver.Value
		expected sqlutil.NullBigRat
		err      error
	}{
		{int64(2), sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(2, 1), Precision: 0}, Valid: true}, nil},
		{float64(2.5), sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(5, 2), Precision: 16}, Valid: true}, nil},
		{true, sqlutil.NullBigRat{}, errors.New("couldn't scan bool")},
		{[]byte("9.55"), sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(955, 100), Precision: 2}, Valid: true}, nil},
		{"2.2", sqlutil.NullBigRat{BigRat: sqlutil.BigRat{Rat: *big.NewRat(11, 5), Precision: 1}, Valid: true}, nil},
		{time.Now(), sqlutil.NullBigRat{}, errors.New("couldn't scan time.Time")},
		{nil, sqlutil.NullBigRat{}, nil},
	}

	for _, tt := range tests {
		actual := sqlutil.NullBigRat{}
		err := actual.Scan(tt.n)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("%+v.Scan(%v): expected error %+v, got error: %+v", actual, tt.n, tt.err, err)
		}

		if actual.Valid != tt.expected.Valid ||
			actual.BigRat.Rat.Cmp(&tt.expected.BigRat.Rat) != 0 ||
			actual.BigRat.Precision != tt.expected.BigRat.Precision {
			t.Errorf("%+v.Scan(%v): expected %+v, actual: %+v", actual, tt.n, tt.expected, actual)
		}
	}
}
