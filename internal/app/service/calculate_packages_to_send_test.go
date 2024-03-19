package service

import (
	"testing"

	bl "github.com/FurmanovD/postpackage/internal/pkg/businesslogic"
	"github.com/FurmanovD/postpackage/internal/pkg/db"
	"github.com/FurmanovD/postpackage/pkg/api/v1"
	"github.com/FurmanovD/postpackage/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPackages = bl.Packages{
		{
			Name:            "5000",
			ItemsPerPackage: 5000,
		},
		{
			Name:            "2000",
			ItemsPerPackage: 2000,
		},
		{
			Name:            "1000",
			ItemsPerPackage: 1000,
		},
		{
			Name:            "500",
			ItemsPerPackage: 500,
		},
		{
			Name:            "250",
			ItemsPerPackage: 250,
		},
	}
)

func Test_calculatePackagesToSend(t *testing.T) {

	pkgMap := make(map[uint]api.Package)
	for i, p := range testPackages {
		apiPkg := bl.Package(testPackages[i]).ToAPI()
		pkgMap[p.ItemsPerPackage] = *apiPkg
	}

	testCases := map[string]struct {
		amount   int64
		expected []api.PackageAmount
	}{
		"1": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  1,
					Package: pkgMap[250],
				},
			},
		},
		"250": {
			amount: 250,
			expected: []api.PackageAmount{
				{
					Number:  1,
					Package: pkgMap[250],
				},
			},
		},
		"251": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  2,
					Package: pkgMap[250],
				},
			},
		},
		"501": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  1,
					Package: pkgMap[500],
				},
				{
					Number:  1,
					Package: pkgMap[250],
				},
			},
		},
		"12001": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  2,
					Package: pkgMap[5000],
				},
				{
					Number:  1,
					Package: pkgMap[2000],
				},
				{
					Number:  1,
					Package: pkgMap[250],
				},
			},
		},
		"15000": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  3,
					Package: pkgMap[5000],
				},
			},
		},
		"15500": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  3,
					Package: pkgMap[5000],
				},
				{
					Number:  1,
					Package: pkgMap[500],
				},
			},
		},
		"18000": {
			amount: 1,
			expected: []api.PackageAmount{
				{
					Number:  3,
					Package: pkgMap[5000],
				},
				{
					Number:  1,
					Package: pkgMap[2000],
				},
				{
					Number:  1,
					Package: pkgMap[1000],
				},
			},
		},
	}

	s := NewService(db.NewStorage(nil), log.Default())
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {

			packaging, err := s.calculatePackagesToSend(testPackages, tc.amount)

			require.NoError(t, err)

			assert.ObjectsAreEqual(tc.expected, packaging)
		})
	}

}
