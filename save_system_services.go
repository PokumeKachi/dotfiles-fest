package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.SystemService(
		"ananicy-cpp",
		"bluetooth",
		"chronyd",
		"docker",
		"getty@",
		"limine-snapper-sync",
		"NetworkManager-dispatcher",
		"NetworkManager-wait-online",
		"NetworkManager",
		"ollama",
		"tor",
		"ufw",
	)
}
