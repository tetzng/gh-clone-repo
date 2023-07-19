# gh-clone-repo

This GitHub CLI extension is a tool for cloning GitHub repositories in a specified manner.

## Features

- When executing `gh clone-repo [repo]`, a directory named after **the currently logged-in username** is created, and the repository is cloned into it. If the directory already exists, the repository is cloned into that directory.
- When executing `gh clone-repo [owner/repo]`, a directory named after `owner` is created, and the repository is cloned into it. If the directory already exists, the repository is cloned into that directory.
- When executing `gh clone-repo [repo url]`, `owner/repo` is extracted from the URL, a directory named after owner is created, and the repository is cloned into it. If the directory already exists, the repository is cloned into that directory.
- When specifying `-u <string>` or `--upstream-remote-name <string>` option, it changes the remote name of the upstream to the specified name.

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation)

2. Install this extension:

   ```sh
   gh extension install tetzng/gh-clone-repo
   ```

## Usage

```sh
gh clone-repo [repo | owner/repo | repo url] [-u, --upstream-remote-name <string>]
```
