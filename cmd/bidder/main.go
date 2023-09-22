package main

import (
	"github.com/adgear/go-commons/pkg/buildinfo"
	"github.com/adgear/rtb-bidder/internal/bidder/app"
)

func main() {
	flags, err := buildinfo.ParseFlags()
	if err == nil {
		// -version flag
		if flags.Version {
			buildinfo.Print()
			return
		}
	}

	// Initialize the bidder application and all application services and components.
	// Setup the dependency injection.
	bidderApp := app.Wire()

	// Start the bidder application.
	bidderApp.Start()
}
