package service

import (
	"context"

	"github.com/FurmanovD/postpackage/internal/pkg/db"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
	"github.com/FurmanovD/postpackage/pkg/log"
)

// confirms PostService interface is implemented
var _ PostService = &serviceImpl{}

type PostService interface {
	GetPackage(ctx context.Context, id int64) (*api.Package, error)
	ListPackages(ctx context.Context) (*api.ListPackagesResponse, error)

	AddPackage(ctx context.Context, pkg api.Package) (*api.Package, error)
	UpdatePackage(ctx context.Context, pkg api.Package) (*api.Package, error)
	DeletePackage(ctx context.Context, id int64) error

	CalculatePackagesToSend(ctx context.Context, minAmount int64) (*api.CalculatePackagingResponse, error)
}

type serviceImpl struct {
	logger log.Logger
	db     *db.Storage
}

func NewService(
	db *db.Storage,
	logger log.Logger,
) *serviceImpl {
	return &serviceImpl{
		db:     db,
		logger: logger.WithField("scope", "service"),
	}
}
