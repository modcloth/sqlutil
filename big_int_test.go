package sqlutil_test

import (
	"database/sql/driver"
	"math/big"
	"testing"

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
