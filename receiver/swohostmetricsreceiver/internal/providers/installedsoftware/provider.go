package installedsoftware

type InstalledSoftware struct {
	Name        string
	Publisher   string
	Version     string
	InstallDate string
}

type Provider interface {
	GetSoftware() ([]InstalledSoftware, error)
}
