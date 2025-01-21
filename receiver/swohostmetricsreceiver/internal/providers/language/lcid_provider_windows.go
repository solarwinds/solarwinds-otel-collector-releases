package language

import "github.com/go-ole/go-ole"

type LCIDProvider interface {
	GetUserDefaultLCID() uint32
}

type windowsLCIDProvider struct{}

func NewWindowsLCIDProvider() LCIDProvider {
	return &windowsLCIDProvider{}
}

func (*windowsLCIDProvider) GetUserDefaultLCID() uint32 {
	return ole.GetUserDefaultLCID()
}
