# 


## Description

```
A collection of automation steps just to make the author's life a bit easier.
I would be surprised if this is going to be useful to a lot more people than just myself,
but we'll see. This CLI will help me create a bunch of ad hoc things that help me in my day to day
development tasks.
```
## Examples

```bash
goose github personal-access-token create
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
|` completion`|Generate the autocompletion script for the specified shell|
|` fun`|A collection of useless commands.|
|` help`|Help about any command|
# ... completion
` completion`

## Usage
> Generate the autocompletion script for the specified shell

 completion

## Description

```
Generate the autocompletion script for  for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|` completion bash`|Generate the autocompletion script for bash|
|` completion fish`|Generate the autocompletion script for fish|
|` completion powershell`|Generate the autocompletion script for powershell|
|` completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
` completion bash`

## Usage
> Generate the autocompletion script for bash

 completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <( completion bash)

To load completions for every new session, execute once:

#### Linux:

	 completion bash > /etc/bash_completion.d/

#### macOS:

	 completion bash > /usr/local/etc/bash_completion.d/

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
` completion fish`

## Usage
> Generate the autocompletion script for fish

 completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	 completion fish | source

To load completions for every new session, execute once:

	 completion fish > ~/.config/fish/completions/.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
` completion powershell`

## Usage
> Generate the autocompletion script for powershell

 completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	 completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
` completion zsh`

## Usage
> Generate the autocompletion script for zsh

 completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	 completion zsh > "${fpath[1]}/_"

#### macOS:

	 completion zsh > /usr/local/share/zsh/site-functions/_

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... fun
` fun`

## Usage
> A collection of useless commands.

 fun

## Description

```
Sometimes also contains sandboxes useful commands that need a bit more polish.
```

## Commands
|Command|Usage|
|-------|-----|
|` fun mock`|Prints out a string AlL FuNNy LiKe|
# ... fun mock
` fun mock`

## Usage
> Prints out a string AlL FuNNy LiKe

 fun mock [string]

## Description

```
Don't judge me, I do this for fun, alright?'
```
# ... help
` help`

## Usage
> Help about any command

 help [command]

## Description

```
Help provides help for any command in the application.
Simply type  help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 05 April 2022**
