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

func TestBigRatValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.BigRat
		expected driver.Value
	}{
		{sqlutil.BigRat{Rat: *big.NewRat(2, 4), Precision: 1}, "0.5"},
		{sqlutil.BigRat{Rat: *big.NewRat(-2, 4), Precision: 1}, "-0.5"},
		{sqlutil.BigRat{Rat: *big.NewRat(1, 3), Precision: 3}, "0.333"},
		{sqlutil.BigRat{Rat: *big.NewRat(1, 3), Precision: 1}, "0.3"},
		{sqlutil.BigRat{Rat: *big.NewRat(1, 3), Precision: 0}, "0"},
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

func TestBigRatScan(t *testing.T) {
	var tests = []struct {
		n        driver.Value
		expected sqlutil.BigRat
		err      error
	}{
		{int64(2), sqlutil.BigRat{Rat: *big.NewRat(2, 1), Precision: 0}, nil},
		{float64(2.5), sqlutil.BigRat{Rat: *big.NewRat(5, 2), Precision: 16}, nil},
		{true, sqlutil.BigRat{}, errors.New("couldn't scan bool")},
		{[]byte("9.55"), sqlutil.BigRat{Rat: *big.NewRat(955, 100), Precision: 2}, nil},
		{"2.2", sqlutil.BigRat{Rat: *big.NewRat(11, 5), Precision: 1}, nil},
		{time.Now(), sqlutil.BigRat{}, errors.New("couldn't scan time.Time")},
		{nil, sqlutil.BigRat{}, errors.New("couldn't scan <nil>")},
	}

	for _, tt := range tests {
		actual := sqlutil.BigRat{}
		err := actual.Scan(tt.n)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("%+v.Scan(%v): expected error %+v, got error: %+v", actual, tt.n, tt.err, err)
		}

		if actual.Rat.Cmp(&tt.expected.Rat) != 0 || actual.Precision != tt.expected.Precision {
			t.Errorf("%+v.Scan(%v): expected %+v, actual: %+v", actual, tt.n, tt.expected, actual)
		}
	}
}

func TestBigRatSub(t *testing.T) {
	var valueTests = []struct {
		x        sqlutil.BigRat
		y        sqlutil.BigRat
		expected sqlutil.BigRat
	}{
		{
			sqlutil.BigRat{Rat: *big.NewRat(5, 2), Precision: 16},
			sqlutil.BigRat{Rat: *big.NewRat(3, 4), Precision: 16},
			sqlutil.BigRat{Rat: *big.NewRat(7, 4), Precision: 16},
		},
		{
			sqlutil.BigRat{Rat: *big.NewRat(5, 2), Precision: 16},
			sqlutil.BigRat{Rat: *big.NewRat(3, 4), Precision: 2},
			sqlutil.BigRat{Rat: *big.NewRat(7, 4), Precision: 2},
		},
		{
			sqlutil.BigRat{Rat: *big.NewRat(5, 2), Precision: 0},
			sqlutil.BigRat{Rat: *big.NewRat(3, 4), Precision: 2},
			sqlutil.BigRat{Rat: *big.NewRat(7, 4), Precision: 0},
		},
	}

	for _, tt := range valueTests {

		actual := sqlutil.BigRat{}
		actual.Sub(&tt.x, &tt.y)

		if actual.Rat.Cmp(&tt.expected.Rat) != 0 || actual.Precision != tt.expected.Precision {
			t.Errorf("br.Sub(%+v, %+v): expected %+v, actual %+v", tt.x, tt.y, tt.expected, actual)
		}
	}
}
