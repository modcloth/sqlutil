package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
)

type NullBigRat struct {
	BigRat BigRat
	Valid  bool
}

//Marshals nested BigRat struct or nil if invalid
func (me *NullBigRat) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if me.Valid {
		data = me.BigRat
	}

	return json.Marshal(data)
}

//Implements sql.Scanner
//
//Accepts nil, proxies everything else to nested BigRat
func (me *NullBigRat) Scan(value interface{}) (err error) {
	me.BigRat = BigRat{}

	if value == nil {
		me.Valid = false
		return nil
	}

	err = me.BigRat.Scan(value)

	if err == nil {
		me.Valid = true
	}

	return err
}

//Implements driver.Valuer
//
//Returns nil if invalid, otherwise proxies to nested BigRat
func (me *NullBigRat) Value() (value driver.Value, err error) {
	if !me.Valid {
		return nil, nil
	}
	return me.BigRat.Value()
}
