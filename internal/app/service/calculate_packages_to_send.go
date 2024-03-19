package service

import (
	"context"
	"errors"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
)

func (s *serviceImpl) CalculatePackagesToSend(ctx context.Context, minAmount int64) (*api.CalculatePackagingResponse, error) {
	log := s.logger.WithField("method", "CalculatePackagesToSend")

	pkgs, err := s.db.Packages.List(ctx)
	if err != nil {
		log.Errorf("error listing packages: %v", err)
		return nil, ErrDBError
	}

	if len(pkgs) == 0 {
		return &api.CalculatePackagingResponse{}, nil
	}

	pkgAmounts, err := s.calculatePackagesToSend(pkgs, minAmount)
	if err != nil {
		log.Errorf("error calculating packages to send: %v", err)
		return nil, err
	}

	return &api.CalculatePackagingResponse{Packages: pkgAmounts}, nil
}

func (s *serviceImpl) calculatePackagesToSend(pkgs bl.Packages, minAmount int64) ([]api.PackageAmount, error) {
	if minAmount < 0 {
		return []api.PackageAmount{}, errors.New("invalid amount")
	}

	pkgs = pkgs.SortByAmount(false)

	if minAmount <= 0 || len(pkgs) == 0 {
		return []api.PackageAmount{}, nil
	}

	biggestPackageAmount := minAmount / int64(pkgs[0].ItemsPerPackage)
	uncoveredRest := minAmount % int64(pkgs[0].ItemsPerPackage)
	var smallerPkgs []api.PackageAmount
	var err error

	if uncoveredRest > 0 { // the biggest packages number don't exactly cover the whole amount required
		smallerPkgs, err = s.calculatePackagesToSend(pkgs[1:], uncoveredRest)
		if err != nil {
			return []api.PackageAmount{}, err
		}

		if len(smallerPkgs) == 0 {
			biggestPackageAmount++ // add one big package to cover the rest
		}
	}

	result := make([]api.PackageAmount, 0, 1+len(smallerPkgs))
	// set the number of the biggest packages
	if biggestPackageAmount > 0 {
		result = append(result, api.PackageAmount{
			Number:  biggestPackageAmount,
			Package: *bl.Package(pkgs[0]).ToAPI(),
		})
	}

	// and append the smaller packaging
	result = append(result, smallerPkgs...)

	return result, nil
}
