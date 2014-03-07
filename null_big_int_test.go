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

func TestNullBigIntValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.NullBigInt
		expected driver.Value
	}{
		{sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(2)}, Valid: true}, "2"},
		{sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(1844674407370955)}, Valid: true}, "1844674407370955"},
		{sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(-1)}, Valid: true}, "-1"},
		{sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(4)}, Valid: false}, nil},
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

func TestNullBigIntScan(t *testing.T) {
	var tests = []struct {
		n        driver.Value
		expected sqlutil.NullBigInt
		err      error
	}{
		{int64(2), sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(2)}, Valid: true}, nil},
		{float64(2.5), sqlutil.NullBigInt{}, errors.New("couldn't scan float64")},
		{true, sqlutil.NullBigInt{}, errors.New("couldn't scan bool")},
		{[]byte("9"), sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(9)}, Valid: true}, nil},
		{"2", sqlutil.NullBigInt{BigInt: sqlutil.BigInt{Int: *big.NewInt(2)}, Valid: true}, nil},
		{time.Now(), sqlutil.NullBigInt{}, errors.New("couldn't scan time.Time")},
		{nil, sqlutil.NullBigInt{}, nil},
	}

	for _, tt := range tests {
		actual := sqlutil.NullBigInt{}
		err := actual.Scan(tt.n)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("%+v.Scan(%v): expected error %+v, got error: %+v", actual, tt.n, tt.err, err)
		}

		if actual.Valid != tt.expected.Valid || actual.BigInt.Int.Cmp(&tt.expected.BigInt.Int) != 0 {
			t.Errorf("%+v.Scan(%v): expected %+v, actual: %+v", actual, tt.n, tt.expected, actual)
		}
	}
}
