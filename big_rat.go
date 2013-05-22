package sqlutil

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigRat struct {
	R *big.Rat
}

func (me *BigRat) MarshalJSON() ([]byte, error) {
	f, _ := me.R.Float64()
	return json.Marshal(f)
}

func (me *BigRat) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		me.R = &big.Rat{}
		if _, err := fmt.Sscan(value.(string), me.R); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return fmt.Errorf("Couldn't scan non-string %+v into BigRat", value)
	}

	return nil
}
