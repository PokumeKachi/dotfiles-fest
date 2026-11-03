package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.Group(
		"sys",
		"network",
		"scanner",
		"wheel",
		"audio",
		"lp",
		"storage",
		"video",
		"users",
		"rfkill",
		"nopasswdlogin",
		"cups",
		"adbusers",
		"libvirt",
		"docker",
	)
}
