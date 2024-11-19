package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ULID ulid.ULID

func (u ULID) Value() (driver.Value, error) {
	return ulid.ULID(u).String(), nil
}

func (u *ULID) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		id, err := ulid.Parse(src)
		if err != nil {
			return err
		}
		*u = ULID(id)
		return nil
	case []byte:
		id, err := ulid.Parse(string(src))
		if err != nil {
			return err
		}
		*u = ULID(id)
		return nil
	default:
		return fmt.Errorf("unsupported type for ULID: %T", src)
	}
}

func NewULID() ULID {
	return ULID(ulid.Make())
}
