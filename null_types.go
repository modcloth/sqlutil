package sqlutil

import (
	"database/sql"
	"encoding/json"
)

//NullString is a wrapper around sql.NullString that satifies json.Marshaler
type NullString struct {
	sql.NullString
}

//MarshalJSON marshals nested String or nil if invalid
func (ns *NullString) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !ns.Valid {
		data = nil
	} else {
		data = ns.String
	}

	return json.Marshal(data)
}

//NullInt64 is a wrapper around sql.NullInt64 that satifies json.Marshaler
type NullInt64 struct {
	sql.NullInt64
}

//MarshalJSON marshals nested Int64 or nil if invalid
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !ni.Valid {
		data = nil
	} else {
		data = ni.Int64
	}

	return json.Marshal(data)
}

//NullFloat64 is a wrapper around sql.NullInt64 that satifies json.Marshaler
type NullFloat64 struct {
	sql.NullFloat64
}

//MarshalJSON is a wrapper around sql.NullFloat64 that satifies json.Marshaler
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	var data interface{}

	if !nf.Valid {
		data = nil
	} else {
		data = nf.Float64
	}

	return json.Marshal(data)
}
