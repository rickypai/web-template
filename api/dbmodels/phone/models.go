// Code generated by sqlc. DO NOT EDIT.

package phone

import (
	"time"
)

type Phone struct {
	ID                int64
	Name              string
	ManufacturerID    int32
	OperatingSystemID int32
	CreatedAt         time.Time
	ModifiedAt        time.Time
}
