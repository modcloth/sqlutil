package sqlutil_test

import (
	"database/sql/driver"
	"math/big"
	"testing"

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
