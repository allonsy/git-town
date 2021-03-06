package drivers

import "github.com/Originate/git-town/src/git"

// Core provides the public API for the drivers subsystem

var registry = Registry{}

var activeDriver CodeHostingDriver

// GetActiveDriver returns the code hosting driver to use based on the git config
func GetActiveDriver() CodeHostingDriver {
	if activeDriver == nil {
		activeDriver = GetDriver(git.GetRemoteOriginURL())
		if activeDriver != nil {
			activeDriver.SetAPIToken(git.GetConfigurationValue(activeDriver.GetAPITokenKey()))
		}
	}
	return activeDriver
}

// GetDriver returns the code hosting driver to use based on given origin url
func GetDriver(originURL string) CodeHostingDriver {
	return registry.DetermineActiveDriver(originURL)
}

// ValidateHasDriver returns an error if there is no code hosting driver
func ValidateHasDriver() error {
	driver := GetActiveDriver()
	if driver == nil {
		return UnsupportedHostingServiceError{&registry}
	}
	return nil
}
