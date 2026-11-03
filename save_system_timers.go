package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.SystemTimer(
		"cachyos-rate-mirrors",
		"fstrim",
		"snapper-cleanup",
	)
}
