package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.Service(
		"ni-maschine",
		"shelly-notifications",
		"syncthing",
		"wireplumber",
		"xdg-user-dirs",
	)
}
