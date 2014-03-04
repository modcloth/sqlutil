package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

type BigInt struct {
	I big.Int
}

func (me *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.I)
}

func (me *BigInt) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		me.I = big.Int{}
		if _, err := fmt.Sscan(value.(string), &me.I); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return fmt.Errorf("Couldn't scan non-string %+v into BigInt", value)
	}

	return nil
}

func (me *BigInt) Value() (value driver.Value, err error) {
	return me.I.String(), nil
}
