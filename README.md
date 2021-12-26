# traindown-cli

traindown-cli is a command line tool for fitness data using
[Traindown](https://traindown.com/).

## Installation

```bash
go install github.com/OliverCardoza/traindown-cli@latest
```

## Usage

To see usage documentation:

```bash
traindown-cli --help
```

This is still a big work in progress. The initial functionality focuses on
validating file content to ensure it follows the
[traindown specifiction](https://traindown.com/spec/).

```bash
traindown-cli validate -i $FILE_OR_DIR
```

## Development

To run the local code during development:

```bash
go run main.go
```

[cobra](https://github.com/spf13/cobra) is used to build the CLI scaffolding
which has some specific guides:

* [user guide](https://github.com/spf13/cobra/blob/master/user_guide.md)
* [code generation](https://github.com/spf13/cobra/blob/master/cobra/README.md): how to add more commands
