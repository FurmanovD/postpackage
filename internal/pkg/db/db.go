package db

import (
	"github.com/FurmanovD/postpackage/internal/pkg/db/repository"
	"github.com/jmoiron/sqlx"
)

// Storage is a DB dependency locator.
type Storage struct {
	Packages repository.PackagesRepository
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Packages: repository.NewPackagesRepository(db),
	}
}
