# ReGet
[![Go Report Card](https://goreportcard.com/badge/github.com/januznl/reget)](https://goreportcard.com/report/github.com/januznl/reget) 
![build](https://github.com/Januznl/reget/actions/workflows/goreleaser.yml/badge.svg)
![ci](https://github.com/Januznl/reget/actions/workflows/ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview
Download the latest release/tag from online SVN tools like github. It determines based on the API which is the latest release or tag and downloads the package based on your architecture. This app is inspired by the python app [lastversion](https://github.com/dvershinin/lastversion).

## Build for automated use
This tool is ideally to use in build scripts, like docker containers or in CI environments.

## Features
* Determine latest release/tag for given project
* Pin to a minor, to download always the latest patch
* Pin to a major, to download always the latest minor

## Usage

### Github usage
```
# Download latest release
reget github <owner>/<repo>

# Download latest tag
reget github <owner>/<repo> -t/--use-tag

# Download latest release and save it with custom filename
reget github <owner>/<repo> -o/--output local_filename.tar.gz

# Download and pin on Major release 1
reget github <owner>/<repo> -p/--pinned-release 1 

# Download and pin on Minor release 1.10 (will download 1.10.5)
reget github <owner>/<repo> -p/--pinned-release 1.10

# Download and pin on full release 1.10.3 (do not download any other version)
reget github <owner>/<repo> -r/--release 1.10.3

```

### Pecl usage
```
# Download latest release
reget pecl <package>

# Download latest release and save it with custom filename
reget pecl <package> -o/--output local_filename.tar.gz

# Download and pin on Major release 1
reget pecl <package> -p/--pinned-release 1 

# Download and pin on Minor release 1.10 (will download 1.10.5)
reget pecl <package> -p/--pinned-release 1.10

# Download and pin on full release 1.10.3 (do not download any other version)
reget pecl <package> -r/--release 1.10.3

```

## Installation

### Use MacOS / Homebrew

```
brew tap januznl/januznl
brew install reget
```

## Use in Docker

```
FROM example
...
COPY --from=januznl/reget:latest /reget /reget
...

```

## Todo
* Add other SVN's than github

