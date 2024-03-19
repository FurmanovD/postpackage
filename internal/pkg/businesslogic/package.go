package businesslogic

import (
	"github.com/FurmanovD/postpackage/internal/pkg/db/model"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
)

type Package model.Package

func NewPackage(p *api.Package) *model.Package {
	if p == nil {
		return nil
	}

	return &model.Package{
		ID:              p.ID,
		Name:            p.Name,
		ItemsPerPackage: p.ItemsPerPackage,
	}
}

func (p Package) ToAPI() *api.Package {
	return &api.Package{
		ID:              p.ID,
		Name:            p.Name,
		ItemsPerPackage: p.ItemsPerPackage,
	}
}
