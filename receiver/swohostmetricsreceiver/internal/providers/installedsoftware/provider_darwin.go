package installedsoftware

type darwinProvider struct{}

var _ Provider = (*darwinProvider)(nil)

func NewInstalledSoftwareProvider() Provider {
	return createInstalledSoftwareProvider()
}

func createInstalledSoftwareProvider() Provider {
	return new(darwinProvider)
}

// GetSoftware implements Provider.
func (p *darwinProvider) GetSoftware() ([]InstalledSoftware, error) {
	return make([]InstalledSoftware, 0), nil
}
