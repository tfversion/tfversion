# tfversion

<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/31e433b6b4ea0e11257fffebe26893d2259f34c6.png" width="125" height="125"> [![Release](https://github.com/tfversion/tfversion/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/tfversion/tfversion/actions/workflows/goreleaser.yaml) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tfversion/tfversion) [![Go Reference](https://pkg.go.dev/badge/github.com/tfversion/tfversion.svg)](https://pkg.go.dev/github.com/tfversion/tfversion)

A simple tool to manage Terraform versions.

## Brew

To install tfversion using brew, simply run:

```sh
brew install tfversion/tap/tfversion
```

## Binaries

You can download the [latest binary](https://github.com/tfversion/tfversion/releases/latest) for Linux, MacOS, and Windows.

## Configuring your shell

To ensure that `tfversion` is always available in your shell, append the following line to your shell profile (e.g., .bashrc, .zshrc or fish config):

```sh
export PATH="$HOME/.tfversion/bin:$PATH"
```

We prefer to not automagically modify your shell profile, and therefore leave this step to you.

## Examples

`tfversion` is user-friendly and easy to use. For more information, please visit the [documentation](https://tfversion.xyz/use.html).

### Install a specific version

```sh
tfversion install 1.7.0
```

### Install the latest stable version

```sh
tfversion install --latest
```

### Install the latest pre-release version

```sh
tfversion install --latest --pre-release
```

### Install the required version for your current directory

```sh
tfversion install --required
```

### Use a specific version

```sh
tfversion use 1.7.0
```

### Use the latest stable version

```sh
tfversion use --latest
```

### Use the latest pre-release version

```sh
tfversion use --latest --pre-release
```

### Use the required version for your current directory

```sh
tfversion use --required
```

### List versions

```sh
tfversion list
```

### List installed versions

```sh
tfversion list --installed
```

### Uninstall a specific version

```sh
tfversion uninstall 1.7.4
```

### Create and use an alias

```sh
tfversion alias default 1.7.4
tfversion use default
```

## Contributing

Contributions are highly appreciated and always welcome.
Have a look through existing [Issues](https://github.com/tfversion/tfversion/issues) and [Pull Requests](https://github.com/tfversion/tfversion/pulls) that you could help with.
