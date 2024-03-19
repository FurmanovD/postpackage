package service

import (
	"context"
)

func (s *serviceImpl) DeletePackage(ctx context.Context, id int64) error {
	log := s.logger.WithField("method", "DeletePackage")

	err := s.db.Packages.Delete(ctx, id)
	if err != nil {
		log.Errorf("error deleting package[ID:%d]: %v", id, err)
		return ErrDBError
	}

	return nil
}
