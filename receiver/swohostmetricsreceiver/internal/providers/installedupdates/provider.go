package installedupdates

type InstalledUpdate struct {
	Caption     string
	HotFixID    string
	InstalledOn string
	InstalledBy string
	Description string
}

type Provider interface {
	GetUpdates() ([]InstalledUpdate, error)
}

// Null Object pattern provider, returns nothing, but does not fail.
type noUpdatesProvider struct{}

var _ (Provider) = (*noUpdatesProvider)(nil)

// Create new instance of NullProvider, whitch returns no updates.
func createNoUpdatesProvider() Provider {
	return new(noUpdatesProvider)
}

func (noUpdatesProvider) GetUpdates() ([]InstalledUpdate, error) {
	return []InstalledUpdate{}, nil
}
