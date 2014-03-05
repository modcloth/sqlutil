package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
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
		br.Rat.SetFloat64(value.(float64))
	case string:
		if _, err := fmt.Sscan(value.(string), &br.Rat); err != nil {
			fmt.Println(err)
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
func (br *BigRat) Value() (value driver.Value, err error) {
	return br.FloatString(br.Precision), nil
}
