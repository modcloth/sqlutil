package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

type BigRat struct {
	big.Rat
	Precision int
}

func (me *BigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.Rat.String())
}

func (me *BigRat) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		if _, err := fmt.Sscan(value.(string), &me.Rat); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return fmt.Errorf("Couldn't scan non-string %+v into BigRat", value)
	}

	return nil
}

func (me *BigRat) Value() (value driver.Value, err error) {
	return me.FloatString(me.Precision), nil
}
