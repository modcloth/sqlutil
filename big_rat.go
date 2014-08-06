package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"
)

//BigRat is a wrapper for big.Rat that satisfies:
//json.Marshaler
//sql.Scanner
//driver.Valuer
type BigRat struct {
	big.Rat
	Precision int //Used when marshalling to strings via Rat.FloatString
}

//MarshalJSON marshals embedded big.Rat's FloatString
func (br *BigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(br.Rat.FloatString(br.Precision))
}

//Scan implements sql.Scanner
//
//Accepts int64, float64, and string
func (br *BigRat) Scan(value interface{}) error {
	switch value.(type) {
	case int64:
		br.Rat.SetInt64(value.(int64))
	case float64:
		br.Precision = 16 //Not sure of a better way to determine this value
		br.Rat.SetFloat64(value.(float64))
	case []uint8, string:
		str := asString(value)
		if i := strings.Index(str, "."); i != -1 {
			br.Precision = len(str) - i - 1
		}
		if _, err := fmt.Sscan(str, &br.Rat); err != nil {
			return err
		}
	default:
		return fmt.Errorf("couldn't scan %+v", reflect.TypeOf(value))
	}

	return nil
}

//Value implements driver.Valuer
//
//Returns embedded big.Rat.FloatString with the precision specified on the object
func (br BigRat) Value() (value driver.Value, err error) {
	return br.FloatString(br.Precision), nil
}

//Sub subtracts two numbers into the given struct.
//
//Analagous to big.Rat.Sub(), but also sets the precision of the resulting BigRat.
//If the two numbers have different precision, uses the less-precise one.
func (br *BigRat) Sub(x *BigRat, y *BigRat) *BigRat {
	br.Rat.Sub(&x.Rat, &y.Rat)
	br.Precision = int(math.Min(float64(x.Precision), float64(y.Precision)))
	return br
}
