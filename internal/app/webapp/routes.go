package webapp

import (
	"github.com/FurmanovD/postpackage/internal/app/webapp/handlers/v1/packages"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
)

// RegisterRoutes registers all the routes
func (s *webServer) RegisterRoutes() {
	s.registerAPIv1Routes()
}

func (s *webServer) registerAPIv1Routes() {
	s.registerAPIv1PackagesRoutes()
	// add other API v1 routes here
}

func (s *webServer) registerAPIv1PackagesRoutes() {
	handler := packages.NewPackagesHandler(s.serviceFacade)

	s.routerGroupAPIV1.GET(api.PathPackagesID, handler.GetPackages)

	s.routerGroupAPIV1.POST(api.PathPackages, handler.AddPackage)

	s.routerGroupAPIV1.PATCH(api.PathPackages, handler.UpdatePackage)

	s.routerGroupAPIV1.DELETE(api.PathPackagesID, handler.DeletePackage)

	s.routerGroupAPIV1.GET(api.PathPackagesCalculateAmount, handler.CalculatePackagesToSend)
}
