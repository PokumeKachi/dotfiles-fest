package main

import "github.com/emad-elsaid/fest"

func init() {
	fest.NpmPackage(
		"npm@12.2.1",
		"semver@7.8.5",
		"yarn@1.22.22",
		"acorn@8.19.0",
		"node-gyp@13.1.0",
		"nopt@10.0.1",
	)
}
