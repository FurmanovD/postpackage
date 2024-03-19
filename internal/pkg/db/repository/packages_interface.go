package repository

import (
	"context"

	"github.com/FurmanovD/postpackage/internal/pkg/db/model"
)

// PackagesRepository contains all functions required to manage Packages objects and their state
type PackagesRepository interface {
	Insert(ctx context.Context, pkg model.Package) (*model.Package, error)
	Get(ctx context.Context, id int64) (*model.Package, error)
	List(ctx context.Context) ([]model.Package, error)
	Update(ctx context.Context, pkg model.Package) (*model.Package, error)
	Delete(ctx context.Context, id int64) error
}
