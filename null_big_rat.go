package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
)

//NullBigRat is a wrapper for BigRat that allows nil and satisfies:
//json.Marshaler
//sql.Scanner
//driver.Valuer
type NullBigRat struct {
	BigRat BigRat
	Valid  bool
}

//MarshalJSON marshals nested BigRat struct or nil if invalid
func (nbr *NullBigRat) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if nbr.Valid {
		data = nbr.BigRat
	}

	return json.Marshal(data)
}

//Scan implements sql.Scanner
//
//Accepts nil, proxies everything else to nested BigRat
func (nbr *NullBigRat) Scan(value interface{}) (err error) {
	nbr.BigRat = BigRat{}

	if value == nil {
		nbr.Valid = false
		return nil
	}

	err = nbr.BigRat.Scan(value)

	if err == nil {
		nbr.Valid = true
	}

	return err
}

//Value implements driver.Valuer
//
//Returns nil if invalid, otherwise proxies to nested BigRat
func (nbr *NullBigRat) Value() (value driver.Value, err error) {
	if !nbr.Valid {
		return nil, nil
	}
	return nbr.BigRat.Value()
}
