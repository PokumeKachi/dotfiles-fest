package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.Service(
		"gnome-keyring-daemon",
		"shelly-notifications",
		"syncthing",
		"wireplumber",
		"xdg-user-dirs",
	)
}
