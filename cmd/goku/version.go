package main

import "fmt"

type versionCmd struct{}

func (v versionCmd) name() string { return "version" }

func (v versionCmd) usage() string { return "" }

func (v versionCmd) short() string { return fmt.Sprintf("Print version (version: %s)", version) }

func (v versionCmd) long() string { return v.short() }

func (v versionCmd) run(_ argSlice) error {
	fmt.Println(version)
	return nil
}
