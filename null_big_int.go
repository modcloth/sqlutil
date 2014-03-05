package sqlutil

import (
	"database/sql/driver"
	"encoding/json"
)

//NullBigInt is a wrapper for BigInt that allows nil and satisfies:
//json.Marshaler
//sql.Scanner
//driver.Valuer
type NullBigInt struct {
	BigInt BigInt
	Valid  bool
}

//MarshalJSON marshals nested BigInt struct or nil if invalid
func (nbi *NullBigInt) MarshalJSON() ([]byte, error) {
	var data interface{}

	data = nil
	if nbi.Valid {
		data = nbi.BigInt
	}

	return json.Marshal(data)
}

//Scan implements sql.Scanner
//
//Accepts nil, proxies everything else to nested BigInt
func (nbi *NullBigInt) Scan(value interface{}) (err error) {
	nbi.BigInt = BigInt{}

	if value == nil {
		nbi.Valid = false
		return nil
	}

	err = nbi.BigInt.Scan(value)

	if err == nil {
		nbi.Valid = true
	}

	return err
}

//Value implements driver.Valuer
//
//Returns nil if invalid, otherwise proxies to nested BigInt
func (nbi *NullBigInt) Value() (value driver.Value, err error) {
	if !nbi.Valid {
		return nil, nil
	}
	return nbi.BigInt.Value()
}
