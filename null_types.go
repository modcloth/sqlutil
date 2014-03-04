package sqlutil

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

//Marshals nested String or nil if invalid
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

//Marshals nested Int64 or nil if invalid
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

//Marshals nested Float64 or nil if invalid
func (me *NullFloat64) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !me.Valid {
		data = nil
	} else {
		data = me.Float64
	}

	return json.Marshal(data)
}
