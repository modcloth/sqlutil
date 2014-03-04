package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"math/big"
)

type NullBigInt struct {
	BigInt *BigInt
	Valid  bool
}

func (me *NullBigInt) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if me.Valid {
		data = me.BigInt
	}

	return json.Marshal(data)
}

func (me *NullBigInt) Scan(value interface{}) error {
	me.BigInt = &BigInt{I: &big.Int{}}

	if value == nil {
		me.Valid = false
		return nil
	}

	me.Valid = true
	return me.BigInt.Scan(value)
}

func (me *BigInt) Value() (value driver.Value, err error) {
	if me == nil || me.I == nil {
		return nil, nil
	}

	return me.I.String(), nil
}
