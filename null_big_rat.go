package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
)

type NullBigRat struct {
	BigRat BigRat
	Valid  bool
}

func (me *NullBigRat) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if me.Valid {
		data = me.BigRat
	}

	return json.Marshal(data)
}

func (me *NullBigRat) Scan(value interface{}) error {
	me.BigRat = BigRat{}

	if value == nil {
		me.Valid = false
		return nil
	}

	me.Valid = true
	return me.BigRat.Scan(value)
}

func (me *NullBigRat) Value() (value driver.Value, err error) {
	if !me.Valid {
		return nil, nil
	}
	return me.BigRat.Value()
}
