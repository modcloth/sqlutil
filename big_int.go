package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

func (me *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.Int)
}

func (me *BigInt) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		me.Int = big.Int{}
		if _, err := fmt.Sscan(value.(string), &me.Int); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return fmt.Errorf("Couldn't scan non-string %+v into BigInt", value)
	}

	return nil
}

func (me *BigInt) Value() (value driver.Value, err error) {
	return me.String(), nil
}
