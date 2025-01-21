package firewall

type Profile struct {
	Name    string
	Enabled int
}

type Container struct {
	FirewallProfiles []Profile
	Error            error
}
