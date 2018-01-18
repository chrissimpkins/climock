# climock


climock is a command line executable that returns user-defined mock standard output stream text, standard error stream text, and exit status code testing data.

## Installation

climock is developed in Go and compiled to the command line executable `climock` (`climock.exe` on Windows). A variety of cross-compiled binaries are available for use on Linux, macOS, and Windows systems, or you can download the source and compile the application yourself. Instructions for both approaches follow.

## Installation Approaches

### Approach 1: Install the pre-compiled binary executable file

Download the latest compiled release file for your operating system and architecture from [the Releases page](https://github.com/chrissimpkins/climock/releases/latest).

#### Linux / macOS

Unpack the tar.gz archive and move the `climock` executable file to a directory on your system PATH (e.g. `/usr/local/bin`).  This can be performed by executing the following command in the root of the unpacked archive:

```
$ mv climock /usr/local/bin/climock
```

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

#### Windows

Unpack the zip archive and move the `climock.exe` executable file to a directory on your system PATH. See [details here](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows) for more information about how to do this.

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

### Approach 2: Compile from the source code and install

You must install the Go programming language (which includes the `go` tool) to compile the project from source.  Follow the [instructions on the Go download page](https://golang.org/dl/) for your platform. 

Once you have installed Go and configured your settings so that Go executables are installed on your system PATH, use the following command to (1) pull the master branch of the climock repository; (2) compile the `climock` executable from source for your platform/architecture configuration; (3) install the executable on your system:

```
$ go get -u github.com/chrissimpkins/climock
```

## Uninstall

The installation includes a single executable binary file.  If you installed with `go get` or added one of the pre-compiled binaries on your system `$PATH` on *.nix systems, you can uninstall with:

```
$ rm $(which climock)
```

## Usage

### Options

```
      --exit           Exit status code
  -h, --help           Application help
      --stderr         Standard error string
      --stdout         Standard output string
      --usage          Application usage
  -v, --version        Application version
```

### Define standard output stream

Use the `--stdout` option to define standard output text:

```
$ climock --stdout "This is the standard output"
This is the standard output
```

### Define standard error stream

Use the `--stderr` option to define standard error text:

```
$ climock --stderr "This is the standard error"
This is the standard error
```

### Define the exit status code

Use the `--exit` option to define the exit status code integer value:

```
$ climock --exit 1
```
The exit status code option can be combined with the standard output and standard error options to modify the exit status code returned when there is text present in these streams.

## Issues

Please [file an issue report](https://github.com/chrissimpkins/climock/issues/new) on the repository for any problems that arise with use.

## Contributing

Contributions to the project are welcomed. Please submit your changes as a pull request on this repository.

### License

climock is licensed under the [MIT License](LICENSE)