package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

//Implements sql.Scanner
//
//Only accepts time.Time structs (and nil)
func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}

	switch value.(type) {
	case time.Time:
		nt.Valid = true
		nt.Time = value.(time.Time)
	default:
		return fmt.Errorf("Couldn't scan %+v", reflect.TypeOf(value))
	}

	return nil
}

//Returns "" if null, otherwise time.String()
func (nt *NullTime) String() string {
	if !nt.Valid {
		return ""
	}

	return nt.Time.String()
}

//Implements driver.Valuer
//
//Returns nil if null, otherwise the nested time struct
func (nt *NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

//Marshals nested Time struct or nil if invalid
func (me *NullTime) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.Time
	}

	return json.Marshal(data)
}
