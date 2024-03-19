package api

const (
	KeyPackageID           = "package_id"
	ParamItemsAmountToSend = "items_amount_to_send"

	PathPackages   = "packages"
	PathPackagesID = "packages/:" + KeyPackageID

	PathPackagesCalculateAmount = PathPackages + "/post_packaging"
)

type Packages []Package

type Package struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	ItemsPerPackage uint   `json:"items_per_package,omitempty"`
}

// TODO(DF): add pagination
type ListPackagesResponse = Packages

type PackageAmount struct {
	Number  int64   `json:"number"`
	Package Package `json:"package"`
}

type CalculatePackagingResponse struct {
	Packages []PackageAmount `json:"packages"`
}
