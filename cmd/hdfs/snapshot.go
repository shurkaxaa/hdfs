package main

import (
	"fmt"
)

func snapshotallow(args []string) {
	if len(args) != 1 {
		printHelp()
	}

	sources, nn, err := normalizePaths(args[0:1])
	if err != nil {
		fatal(err)
	}

	source := sources[0]
	client, err := getClient(nn)
	if err != nil {
		fatal(err)
	}
	err = client.AllowSnapshots(source)
	if err != nil {
		fatal(err)
	}
}

func snapshotdisallow(args []string) {
	if len(args) != 1 {
		printHelp()
	}

	sources, nn, err := normalizePaths(args[0:1])
	if err != nil {
		fatal(err)
	}

	source := sources[0]
	client, err := getClient(nn)
	if err != nil {
		fatal(err)
	}
	err = client.DisallowSnapshots(source)
	if err != nil {
		fatal(err)
	}
}

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

func snapshotdirs() {
	_, nn, err := normalizePaths([]string{})
	if err != nil {
		fatal(err)
	}
	client, err := getClient(nn)
	if err != nil {
		fatal(err)
	}
	folders, err := client.ListSnapshottableFolders()
	if err != nil {
		fatal(err)
	}
	for _, f := range folders {
		fmt.Println(f)
	}
}
