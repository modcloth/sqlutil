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

func TestBigIntValue(t *testing.T) {
	var valueTests = []struct {
		n        sqlutil.BigInt
		expected driver.Value
	}{
		{sqlutil.BigInt{Int: *big.NewInt(2)}, "2"},
		{sqlutil.BigInt{Int: *big.NewInt(1844674407370955)}, "1844674407370955"},
		{sqlutil.BigInt{Int: *big.NewInt(-1)}, "-1"},
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
func TestBigIntScan(t *testing.T) {
	var tests = []struct {
		n        driver.Value
		expected sqlutil.BigInt
		err      error
	}{
		{int64(2), sqlutil.BigInt{Int: *big.NewInt(2)}, nil},
		{float64(2.2), sqlutil.BigInt{}, errors.New("couldn't scan float64")},
		{true, sqlutil.BigInt{}, errors.New("couldn't scan bool")},
		{[]byte{5}, sqlutil.BigInt{}, errors.New("couldn't scan []uint8")},
		{"2", sqlutil.BigInt{Int: *big.NewInt(2)}, nil},
		{time.Now(), sqlutil.BigInt{}, errors.New("couldn't scan time.Time")},
		{nil, sqlutil.BigInt{}, errors.New("couldn't scan <nil>")},
	}

	for _, tt := range tests {
		actual := sqlutil.BigInt{}
		err := actual.Scan(tt.n)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("%+v.Scan(%v): expected error %+v, got error: %+v", actual, tt.n, tt.err, err)
		}

		if actual.Int.Cmp(&tt.expected.Int) != 0 {
			t.Errorf("%+v.Scan(%v): expected %+v, actual: %+v", actual, tt.n, tt.expected, actual)
		}
	}
}
