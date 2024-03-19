package businesslogic

import (
	"sort"

	"github.com/FurmanovD/postpackage/internal/pkg/db/model"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
)

type Packages []model.Package

func (p Packages) ToAPI() *api.ListPackagesResponse {
	if len(p) == 0 {
		return nil
	}

	pkgs := make([]api.Package, len(p))
	for i, v := range p {
		pkgs[i] = *Package(v).ToAPI()
	}

	list := api.ListPackagesResponse(pkgs)
	return &list
}

func (p Packages) SortByAmount(asc bool) Packages {
	sort.Slice(p, func(i, j int) bool {
		if asc {
			return p[i].ItemsPerPackage < p[j].ItemsPerPackage
		} else {
			return p[i].ItemsPerPackage > p[j].ItemsPerPackage
		}
	})
	return p
}
