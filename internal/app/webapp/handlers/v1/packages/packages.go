package packages

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FurmanovD/postpackage/internal/app/service"
	"github.com/FurmanovD/postpackage/internal/app/webapp/weberror"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
)

// packagesHandlerImpl an EndpointHandler interface holder
type packagesHandlerImpl struct {
	service service.PostService
}

// NewPackagesHandler instantiates a packages handler
func NewPackagesHandler(svc service.PostService) PackagesHandler {
	return &packagesHandlerImpl{
		service: svc,
	}
}

// GetPackage is a GET /api/v1/packages/{:packageID} handler
func (h *packagesHandlerImpl) GetPackages(c *gin.Context) {
	pkgIDStr := c.Param(api.KeyPackageID)
	if pkgIDStr == "" {
		h.ListPackages(c)
		return
	}

	id, err := strconv.ParseInt(pkgIDStr, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				api.KeyPackageID+" expected to be a natural number",
			),
		)
		return
	}

	pkg, err := h.service.GetPackage(c.Request.Context(), id)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	// return package object
	c.JSON(
		http.StatusOK,
		pkg,
	)
}

// ListPackages is a GET /api/v1/packages/ handler
func (h *packagesHandlerImpl) ListPackages(c *gin.Context) {
	packages, err := h.service.ListPackages(c.Request.Context())
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, "errir listing packages"),
		)
		return
	}

	// return packages list
	c.JSON(
		http.StatusOK,
		packages,
	)
}

// AddPackage is a POST /api/v1/packages handler
func (h *packagesHandlerImpl) AddPackage(c *gin.Context) {
	// construct the request object:
	req := api.Package{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				err.Error(),
			),
		)
		return
	}

	pkg, err := h.service.AddPackage(c.Request.Context(), req)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, "error creating new package object"),
		)
		return
	}

	// return package
	c.JSON(
		http.StatusOK,
		pkg,
	)
}

// UpdatePackage is a PATCH /api/v1/packages handler
func (h *packagesHandlerImpl) UpdatePackage(c *gin.Context) {
	// construct the request object:
	req := api.Package{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				err.Error(),
			),
		)
		return
	}

	pkg, err := h.service.UpdatePackage(c.Request.Context(), req)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, "error updating package object"),
		)
		return
	}

	// return package
	c.JSON(
		http.StatusOK,
		pkg,
	)
}

// UpdatePackage is a DELETE /api/v1/packages/:packageID handler
func (h *packagesHandlerImpl) DeletePackage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param(api.KeyPackageID), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				api.KeyPackageID+" expected to be a natural number",
			),
		)
		return
	}

	err = h.service.DeletePackage(c.Request.Context(), id)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, "error deleting package object"),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		api.GetCommonResponseOk(),
	)
}

// CalculatePackagesToSend is a GET /api/v1/packages/post_packaging handler
func (h *packagesHandlerImpl) CalculatePackagesToSend(c *gin.Context) {
	itemsAmount, err := strconv.ParseInt(c.Query(api.ParamItemsAmountToSend), 10, 64)
	if err != nil || itemsAmount <= 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				api.ParamItemsAmountToSend+" expected to be a natural number",
			),
		)
		return
	}

	packaging, err := h.service.CalculatePackagesToSend(c.Request.Context(), itemsAmount)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, "error deleting package object"),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		packaging,
	)
}
