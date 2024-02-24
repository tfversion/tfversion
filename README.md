# tfversion

[![Release](https://github.com/tfversion/tfversion/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/tfversion/tfversion/actions/workflows/goreleaser.yaml) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tfversion/tfversion) ![GitHub commits since latest release (by SemVer)](https://img.shields.io/github/commits-since/tfversion/tfversion/latest) [![Go Reference](https://pkg.go.dev/badge/github.com/tfversion/tfversion.svg)](https://pkg.go.dev/github.com/tfversion/tfversion)

A simple tool to manage Terraform versions.

## Brew

To install tfversion using brew, simply run:

```sh
brew install tfversion
```

## Binaries

You can download the [latest binary](https://github.com/tfversion/tfversion/releases/latest) for Linux, MacOS, and Windows.

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
Have a look through existing [Issues](https://github.com/tfversion/tfversion/issues) and [Pull Requests](https://github.com/tfversion/tfversion/pulls) that you could help with.
