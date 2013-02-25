package sqlutil

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type NullString struct {
	sql.NullString
}

func (me *NullString) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.String
	}

	return json.Marshal(data)
}

type NullInt64 struct {
	sql.NullInt64
}

func (me *NullInt64) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.Int64
	}

	return json.Marshal(data)
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (me *NullFloat64) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.Float64
	}

	return json.Marshal(data)
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	var ok bool
	if value == nil {
		nt.Time = *new(time.Time)
		nt.Valid = false
		return nil
	}

	if nt.Time, ok = value.(time.Time); !ok {
		return errors.New("Tried to scan a non-time")
	}
	nt.Valid = true
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
