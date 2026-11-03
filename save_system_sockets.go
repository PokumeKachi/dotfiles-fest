package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.SystemSocket(
		"avahi-daemon",
		"cups",
		"libvirtd-admin",
		"libvirtd-ro",
		"libvirtd",
		"nix-daemon",
		"systemd-resolved-monitor",
		"systemd-resolved-varlink",
		"systemd-userdbd",
		"virtlockd-admin",
		"virtlockd",
		"virtlogd-admin",
		"virtlogd",
	)
}
