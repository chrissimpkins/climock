package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version = "1.0.0"

	usage = `Usage: $ climock [option] [arg]`

	help = `==================================================
 climock
 Copyright 2018 Christopher Simpkins
 MIT License

 Source: https://github.com/chrissimpkins/climock
==================================================
 Usage:
  $ climock [option] [arg]
 Options:
      --exit           Exit status code
  -h, --help           Application help
      --stderr         Standard error string
      --stdout         Standard output string
      --usage          Application usage
  -v, --version        Application version
`
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool
var stdOut, stdErr *string
var exitCode *int

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
	stdOut = flag.String("stdout", "", "Standard output string")
	stdErr = flag.String("stderr", "", "Standard error string")
	exitCode = flag.Int("exit", 0, "Exit status code")
}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("climock v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	// write requested standard output
	if len(*stdOut) > 0 {
		os.Stdout.WriteString(*stdOut)
	}

	// write requested standard error
	if len(*stdErr) > 0 {
		os.Stderr.WriteString(*stdErr)
	}

	// exit with default exit status code 0 or value defined on the command line
	os.Exit(*exitCode)

}
