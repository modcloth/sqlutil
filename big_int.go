package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

//BigInt is a wrapper for big.Int that satisfies:
//json.Marshaler
//sql.Scanner
//driver.Valuer
type BigInt struct {
	big.Int
}

//MarshalJSON marshals embedded big.Int
func (bi *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(bi.Int)
}

//Scan implements sql.Scanner
//
//Accepts int64 and string
func (bi *BigInt) Scan(value interface{}) error {
	switch value.(type) {
	case int64:
		bi.Int = *big.NewInt(value.(int64))
	case string:
		if _, err := fmt.Sscan(value.(string), &bi.Int); err != nil {
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
//Returns embedded big.Int.String()
func (bi *BigInt) Value() (value driver.Value, err error) {
	return bi.String(), nil
}
