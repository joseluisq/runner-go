// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

// +build windows

package shell

import "github.com/joseluisq/runner-go/shell/powershell"

// Suffix provides the shell script suffix.
const Suffix = ".ps1"

// Command returns the powershell command and arguments.
func Command() (string, []string) {
	return powershell.Command()
}

// Script converts a slice of individual shell commands to
// a powershell script.
func Script(commands []string) string {
	return powershell.Script(commands)
}
