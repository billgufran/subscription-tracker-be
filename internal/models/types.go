package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type ULID ulid.ULID

// Used behind the scenes by GORM. Value implements the driver.Valuer interface.
func (u ULID) Value() (driver.Value, error) {
	return ulid.ULID(u).String(), nil
}

// Used behind the scenes by GORM. Scan implements the sql.Scanner interface.
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

func (u *ULID) BeforeCreate(tx *gorm.DB) error {
	if *u == (ULID{}) { // Check if ULID is empty
		*u = NewULID()
	}
	return nil
}

func (u ULID) String() string {
	return ulid.ULID(u).String()
}

func (u ULID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}

func (u *ULID) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid ULID format")
	}
	id, err := ulid.Parse(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*u = ULID(id)
	return nil
}
