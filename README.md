# scrman - Script Manager

A Unix Command Line Interface for bash script managing and sharing.

**Note:** This project is currently in progress. The following is the planned API usage.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Configuration](#configuration)

--------------------------------------------------------------------------------------

## Installation

### MacOS

TODO

### Linux

TODO

### Add bin to PATH

Add the following to your `.zshrc` or `.bashrc`:

```sh
PATH=$PATH:~/.scrman/bin
```

--------------------------------------------------------------------------------------

## Usage

### Creating Scripts

```sh
scr create
```

This command will start a dialogue to save commands you just ran in the terminal and
convert them into a script. You can also create a script from scratch using this command.

### Running Scripts

#### Run a script from an online repository

```sh
scr run <author-username>/<repository>
```

- This command will run a script from the web.
- Scripts are saved through **Github** repositories and fetched using git.
- Running this will save the script repo in `~/.scrman/scripts/<author-username>/<repository>`.

#### Run a local script

```sh
scr run <script-name>
```

This command will run scripts located in `~/.scrman/scripts/`

#### Installing a script

You can install a script to make it runnable as a terminal command:

```sh
scr install <script-name>
```

Run the command in your terminal by simply calling the script name

```sh
<script-name>
```

For example:

```sh
scr install helloworld
helloworld
> "Hello World!"
```

--------------------------------------------------------------------------------------

## Configuration

### Local Folder Structure

When installed, the following folder structure is created in your home directory, under `~/.scrman/`:

- **~/.scrman/**: Root Folder
  - **bin/**: Stores all installed scripts as binaries
  - **scripts/**: Stores script project files
    - **helloworld/**: Example script project
      - **config.json**: Example script Configuration
      - **scr.sh**: Example script file
    - **another-script-project/**
      - ...
  - **config.json**: General scrman configuration

### Script Configuration

In each script project, configurations are stored in a `config.json` file.

Here is an example of a config file:

```json
// config.json
{
    "location": "./",
    "arguments": [
        {
            "name": "Example Argument 1",
            "default": "Example Default Value",
        },
        {
            "name": "Example Argument 2",
            "default": 1,
        },
    ]
}
```

- `location` defines where the script will be run
- `arguments` are parameters passed down to the script.
