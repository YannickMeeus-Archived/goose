# Quick Start - Install 

> [!TIP]
>  is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/Silly-Goose-Software/goose/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -sSL instl.sh/Silly-Goose-Software/goose/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -sSL instl.sh/Silly-Goose-Software/goose/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile  from source, you have to have [Go](https://golang.org/) installed.

Compiling  from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of .

```command
go install github.com/Silly-Goose-Software/goose@latest
```

<!-- tabs:end -->
