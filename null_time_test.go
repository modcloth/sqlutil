package sqlutil_test

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/modcloth/sqlutil"
)

func TestNullTimeValue(t *testing.T) {
	var now = time.Now()
	var valueTests = []struct {
		n        sqlutil.NullTime
		expected driver.Value
	}{
		{sqlutil.NullTime{Time: now, Valid: false}, nil},
		{sqlutil.NullTime{Time: now, Valid: true}, now},
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

func TestNullTimeScan(t *testing.T) {
	var now = time.Now()
	var tests = []struct {
		n        driver.Value
		expected sqlutil.NullTime
		err      error
	}{
		{int64(2), sqlutil.NullTime{}, errors.New("couldn't scan int64")},
		{float64(2.5), sqlutil.NullTime{}, errors.New("couldn't scan float64")},
		{true, sqlutil.NullTime{}, errors.New("couldn't scan bool")},
		{[]byte{5}, sqlutil.NullTime{}, errors.New("couldn't scan []uint8")},
		{"2", sqlutil.NullTime{}, errors.New("couldn't scan string")},
		{now, sqlutil.NullTime{Time: now, Valid: true}, nil},
		{nil, sqlutil.NullTime{}, nil},
	}

	for _, tt := range tests {
		actual := sqlutil.NullTime{}
		err := actual.Scan(tt.n)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("%+v.Scan(%v): expected error %+v, got error: %+v", actual, tt.n, tt.err, err)
		}

		if actual.Valid != tt.expected.Valid || !actual.Time.Equal(tt.expected.Time) {
			t.Errorf("%+v.Scan(%v): expected %+v, actual: %+v", actual, tt.n, tt.expected, actual)
		}
	}
}
