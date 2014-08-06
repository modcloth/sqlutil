package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//NullTime is a wrapper for time.Time that allows for 'null' values from databases
//
//Satisfies:
//json.Marshaler
//sql.Scanner
//driver.Valuer
type NullTime struct {
	Time  time.Time
	Valid bool
}

//Scan implements sql.Scanner
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
		return fmt.Errorf("couldn't scan %+v", reflect.TypeOf(value))
	}

	return nil
}

//String returns "" if null, otherwise time.String()
func (nt *NullTime) String() string {
	if !nt.Valid {
		return ""
	}

	return nt.Time.String()
}

//Value implements driver.Valuer
//
//Returns nil if null, otherwise the nested time struct
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

//MarshalJSON marshals nested Time struct or nil if invalid
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !nt.Valid {
		data = nil
	} else {
		data = nt.Time
	}

	return json.Marshal(data)
}
