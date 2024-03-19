package packages

import "github.com/gin-gonic/gin"

type PackagesHandler interface {
	GetPackages(c *gin.Context)

	AddPackage(c *gin.Context)
	UpdatePackage(c *gin.Context)
	DeletePackage(c *gin.Context)

	CalculatePackagesToSend(c *gin.Context)
}
