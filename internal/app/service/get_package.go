package service

import (
	"context"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
)

func (s *serviceImpl) GetPackage(ctx context.Context, id int64) (*api.Package, error) {
	log := s.logger.WithField("method", "GetPackage")

	pkg, err := s.db.Packages.Get(ctx, id)
	if err != nil {
		log.Errorf("error getting package[ID:%d]: %v", id, err)
		return nil, ErrDBError
	}

	if pkg == nil {
		return nil, ErrNotFound
	}

	return bl.Package(*pkg).ToAPI(), nil
}
