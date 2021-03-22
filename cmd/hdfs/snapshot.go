package main

import (
	"fmt"
)

func snapshotadd(args []string) {
	if len(args) != 2 {
		printHelp()
	}

	sources, nn, err := normalizePaths(args[0:1])
	if err != nil {
		fatal(err)
	}

	source := sources[0]
	name := args[1]
	client, err := getClient(nn)
	if err != nil {
		fatal(err)
	}
	// TODO. On-demand snapshot enable?
	err = client.AllowSnapshots(source)
	if err != nil {
		fatal(err)
	}
	snapshotPath, err := client.CreateSnapshot(source, name)
	if err != nil {
		fatal(err)
	}
	fmt.Println(snapshotPath)
}

func snapshotrm(args []string) {
	if len(args) != 2 {
		printHelp()
	}

	sources, nn, err := normalizePaths(args[0:1])
	if err != nil {
		fatal(err)
	}

	source := sources[0]
	name := args[1]
	client, err := getClient(nn)
	if err != nil {
		fatal(err)
	}
	err = client.DeleteSnapshot(source, name)
	if err != nil {
		fatal(err)
	}
}
