package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
)

type NullBigInt struct {
	BigInt BigInt
	Valid  bool
}

//Marshals nested BigInt struct or nil if invalid
func (me *NullBigInt) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if me.Valid {
		data = me.BigInt
	}

	return json.Marshal(data)
}

//Implements sql.Scanner
//
//Accepts nil, proxies everything else to nested BigInt
func (me *NullBigInt) Scan(value interface{}) (err error) {
	me.BigInt = BigInt{}

	if value == nil {
		me.Valid = false
		return nil
	}

	err = me.BigInt.Scan(value)

	if err == nil {
		me.Valid = true
	}

	return err
}

//Implements driver.Valuer
//
//Returns nil if invalid, otherwise proxies to nested BigInt
func (me *NullBigInt) Value() (value driver.Value, err error) {
	if !me.Valid {
		return nil, nil
	}
	return me.BigInt.Value()
}
