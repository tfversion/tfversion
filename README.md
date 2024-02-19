# tfversion

[![Release](https://github.com/bschaatsbergen/tfversion/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/bschaatsbergen/tfversion/actions/workflows/goreleaser.yaml) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bschaatsbergen/tfversion) ![GitHub commits since latest release (by SemVer)](https://img.shields.io/github/commits-since/bschaatsbergen/tfversion/latest) [![Go Reference](https://pkg.go.dev/badge/github.com/bschaatsbergen/tfversion.svg)](https://pkg.go.dev/github.com/bschaatsbergen/tfversion)

A simple tool to manage Terraform versions.

## Brew

To install tfversion using brew, simply run:

```sh
brew install tfversion
```

## Binaries

You can download the [latest binary](https://github.com/bschaatsbergen/tfversion/releases/latest) for Linux, MacOS, and Windows.

## Examples

Using `tfversion` is very simple.

### Install a specific version

```sh
tfversion install 1.7.0
```

### Use a specific version

```sh
tfversion use 1.7.0
```

### List versions

```sh
tfversion list
```

### List installed versions

```sh
tfversion list --installed
```

## Contributing

Contributions are highly appreciated and always welcome.
Have a look through existing [Issues](https://github.com/bschaatsbergen/tfversion/issues) and [Pull Requests](https://github.com/bschaatsbergen/tfversion/pulls) that you could help with.
