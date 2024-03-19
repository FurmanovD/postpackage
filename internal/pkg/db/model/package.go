package model

import "time"

type Package struct {
	ID              int64      `db:"id" json:"id,omitempty"`
	Name            string     `db:"name" json:"name,omitempty"`
	ItemsPerPackage uint       `db:"items_per_package" json:"items_per_package,omitempty"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
