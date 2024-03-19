package service

import (
	"context"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
)

func (s *serviceImpl) ListPackages(ctx context.Context) (*api.ListPackagesResponse, error) {
	log := s.logger.WithField("method", "ListPackages")

	pkgs, err := s.db.Packages.List(ctx)
	if err != nil {
		log.Errorf("error listing packages: %v", err)
		return nil, ErrDBError
	}

	return bl.Packages(pkgs).ToAPI(), nil
}
