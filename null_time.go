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

func (nt *NullTime) String() string {
	if !nt.Valid {
		return ""
	}

	return nt.Time.String()
}

func (nt *NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

func (me *NullTime) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.Time
	}

	return json.Marshal(data)
}
