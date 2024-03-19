package service

import (
	"context"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
)

func (s *serviceImpl) UpdatePackage(ctx context.Context, pkg api.Package) (*api.Package, error) {
	log := s.logger.WithField("method", "UpdatePackage")

	pkgDB := bl.NewPackage(&pkg)
	pkgDB, err := s.db.Packages.Update(ctx, *pkgDB)
	if err != nil {
		log.Errorf(
			"error updating object[%+v] in a DB: %v",
			pkgDB,
			err,
		)
		return nil, ErrDBError
	}

	return bl.Package(*pkgDB).ToAPI(), nil
}
