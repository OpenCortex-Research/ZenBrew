package zb_types

type Package struct {
	Name       string         `json:"name"`
	Format     string         `json:"format"`
	Maintainer string         `json:"maintainer"`
	Versions   []PackageVersion `json:"versions"`
	Latest     string         `json:"latest"`
}

type PackageVersion struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type PackageLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type InstalledPackage struct {
	Name       string         `json:"name"`
	Format     string         `json:"format"`
	Maintainer string         `json:"maintainer"`
	Version   PackageVersion `json:"version"`
	Status	string `json:"status"`
	Repository	string `json:"repository"`
}