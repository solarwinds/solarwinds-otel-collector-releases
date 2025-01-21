package domain

type Domain struct {
	Domain     string // name of domain
	FQDN       string // Fully Qualified Domain Name
	DomainRole int    // domain role
	Workgroup  string // workgroup
}
