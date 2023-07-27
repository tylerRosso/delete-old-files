package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var (
	args struct {
		cutoffTime time.Time
		path       string
	}

	canSkipDir bool
)

func deleteFileIfOlder(path string, info fs.FileInfo, err error) error {
	if info.IsDir() {
		if canSkipDir {
			return filepath.SkipDir
		}

		return nil
	}

	canSkipDir = true

	creationTime := time.Unix(info.Sys().(*syscall.Stat_t).Ctim.Sec, 0)
	if creationTime.Before(args.cutoffTime) {
		removeErr := os.Remove(path)
		if removeErr != nil {
			printErrorMessageAndExit("The following error ocurred when trying to delete the filepath \"" + path + "\": " + removeErr.Error())
		}
	}

	return nil
}

func parseArgs() {
	switch len(os.Args) {
	case 1:
		printErrorMessageAndExit("You must give folder path and file age (in days).")
	case 2:
		printErrorMessageAndExit("You must inform file age (in days).")
	}

	parsePath()

	parseCutoffTime()
}

func parseCutoffTime() {
	age, err := strconv.Atoi(os.Args[2])
	if err != nil || age < 0 {
		printErrorMessageAndExit("Given file age is not a valid integer.")
	}

	args.cutoffTime = time.Now().AddDate(0, 0, -age)
}

func parsePath() {
	path := os.Args[1]

	pathInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			printErrorMessageAndExit("Given path does not exist.")
		} else {
			printErrorMessageAndExit("The following error ocurred when getting statistics on the given path: " + err.Error())
		}
	}

	if !pathInfo.IsDir() {
		printErrorMessageAndExit("Given path is not a directory.")
	}

	args.path = path
}

func printErrorMessageAndExit(errorMessage string) {
	fmt.Fprintln(os.Stderr, errorMessage)

	os.Exit(1)
}

func main() {
	parseArgs()

	filepath.Walk(args.path, deleteFileIfOlder)
}
