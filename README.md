sqlutil
=======

[![Build Status](https://travis-ci.org/modcloth/sqlutil.png?branch=master)](https://travis-ci.org/modcloth/sqlutil)

[Godoc](https://godoc.org/github.com/modcloth/sqlutil)

**NOTE** API is not locked yet.

Works with Go 1.1+

Backend-agnostic SQL utilities.  Github suggested I name this
`derp-octo-ironman`.  I almost did. Almost.

This library supplies:
- Wrappers to Go's `big.Int` and `big.Rat` that
  allows them to be used with `database/sql` `Scan` and `Exec` functions
  to interact with arbitrary precision database fields. Also supplies
  Null\* analogues for each of these.
- Wrappers for `NullString`, `NullInt64`, `NullFloat64` from
  `database/sql` that satisfiy the `json.Marshaler` interface.
- An implementation of `NullTime` that satisfies the `json.Marshaler`
  interface as well as `sql.Scanner` and `driver.Valuer` for use with
  `database/sql`.
