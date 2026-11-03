package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.Socket(
		"gnome-keyring-daemon",
		"p11-kit-server",
		"pipewire-pulse",
		"pipewire",
	)
}
