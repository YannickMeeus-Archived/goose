# goose

## Usage
> A collection of automation steps just to make the author's life a bit easier

goose

## Description

```
I would be surprised if this is going to be useful to a lot more people than just myself,
but we'll see. This CLI will help me create a bunch of ad hoc things that help me in my day to day
development tasks.'
```
## Examples

```bash
goose github pat create
goose s3 bucket create
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`goose completion`|generate the autocompletion script for the specified shell|
|`goose help`|Help about any command|
# ... completion
`goose completion`

## Usage
> generate the autocompletion script for the specified shell

goose completion

## Description

```

Generate the autocompletion script for goose for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`goose completion bash`|generate the autocompletion script for bash|
|`goose completion fish`|generate the autocompletion script for fish|
|`goose completion powershell`|generate the autocompletion script for powershell|
|`goose completion zsh`|generate the autocompletion script for zsh|
# ... completion bash
`goose completion bash`

## Usage
> generate the autocompletion script for bash

goose completion bash

## Description

```

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(goose completion bash)

To load completions for every new session, execute once:
Linux:
  $ goose completion bash > /etc/bash_completion.d/goose
MacOS:
  $ goose completion bash > /usr/local/etc/bash_completion.d/goose

You will need to start a new shell for this setup to take effect.
  
```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`goose completion fish`

## Usage
> generate the autocompletion script for fish

goose completion fish

## Description

```

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:
$ goose completion fish | source

To load completions for every new session, execute once:
$ goose completion fish > ~/.config/fish/completions/goose.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`goose completion powershell`

## Usage
> generate the autocompletion script for powershell

goose completion powershell

## Description

```

Generate the autocompletion script for powershell.

To load completions in your current shell session:
PS C:\> goose completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`goose completion zsh`

## Usage
> generate the autocompletion script for zsh

goose completion zsh

## Description

```

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:
# Linux:
$ goose completion zsh > "${fpath[1]}/_goose"
# macOS:
$ goose completion zsh > /usr/local/share/zsh/site-functions/_goose

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`goose help`

## Usage
> Help about any command

goose help [command]

## Description

```
Help provides help for any command in the application.
Simply type goose help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 13 February 2022**
