package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

type BigRat struct {
	big.Rat
	Precision int //Used when marshalling to strings via Rat.FloatString
}

//Marshals embedded big.Rat.FloatString()
func (me *BigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.Rat.FloatString(me.Precision))
}

//Implements sql.Scanner
//
//Accepts int64, float64, and string
func (me *BigRat) Scan(value interface{}) error {
	switch value.(type) {
	case int64:
		me.Rat.SetInt64(value.(int64))
	case float64:
		me.Rat.SetFloat64(value.(float64))
	case string:
		if _, err := fmt.Sscan(value.(string), &me.Rat); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return fmt.Errorf("Couldn't scan %+v", reflect.TypeOf(value))
	}

	return nil
}

//Implements driver.Valuer
//
//Returns embedded big.Rat.FloatString with the precision specified on the object
func (me *BigRat) Value() (value driver.Value, err error) {
	return me.FloatString(me.Precision), nil
}
