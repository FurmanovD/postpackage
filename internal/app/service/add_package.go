package service

import (
	"context"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
)

func (s *serviceImpl) AddPackage(ctx context.Context, pkg api.Package) (*api.Package, error) {
	log := s.logger.WithField("method", "AddPackage")

	pkgDB := bl.NewPackage(&pkg)
	pkgDB, err := s.db.Packages.Insert(ctx, *pkgDB)
	if err != nil {
		log.Errorf(
			"error adding new object[%+v] to DB: %v",
			pkgDB,
			err,
		)
		return nil, ErrDBError
	}

	return bl.Package(*pkgDB).ToAPI(), nil
}
