package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.SystemService(
		"ananicy-cpp",
		"avahi-daemon",
		"bluetooth",
		"chronyd",
		"cups",
		"docker",
		"getty@",
		"libvirtd",
		"limine-snapper-sync",
		"NetworkManager-dispatcher",
		"NetworkManager-wait-online",
		"NetworkManager",
		"systemd-resolved",
		"tor",
		"ufw",
	)
}
