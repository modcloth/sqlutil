package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

type BigInt struct {
	big.Int
}

//Marshals embedded big.Int
func (me *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.Int)
}

//Implements sql.Scanner
//
//Accepts int64 and string
func (me *BigInt) Scan(value interface{}) error {
	switch value.(type) {
	case int64:
		me.Int = *big.NewInt(value.(int64))
	case string:
		if _, err := fmt.Sscan(value.(string), &me.Int); err != nil {
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
//Returns embedded big.Int.String()
func (me *BigInt) Value() (value driver.Value, err error) {
	return me.String(), nil
}
