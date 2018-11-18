package main

import "github.com/Urethramancer/slog"

const (
	programName = "Revolut"
)

// Version is filled in by the build script.
var Version = "undefined"

// VersionCmd is empty.
type VersionCmd struct{}

// Execute shows the program name and version.
func (cmd *VersionCmd) Execute(args []string) error {
	slog.Msg("%s %s", programName, Version)
	return nil
}
