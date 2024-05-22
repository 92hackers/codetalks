# Codetalks

A tool that helps analyze your codebase, make it self-talk!

## Quick Start

Please make sure you have installed Go 1.18+.

### Installation

```bash
go install github.com/92hackers/codetalks@latest
```

### Usage

```bash
cd /path/to/your/project

# Analyze whole codebase
codetalks

# Also, you can analyze a specific directory
codetalks ./path/to/your/project
```

## Features overview

Alias: `cloc` means `count number of lines of code, comments and blanks`.

1. Show `cloc` info grouped by language in the project.
2. Show `cloc` info of every directory and it's subdirectories.
3. Sort files or directories by `cloc` info, such as set `lines of code` as the sort criteria.
4. Blazing fast analyzing.


## Roadmap

- [ ] Show `cloc` info of every file in the project.
- [x] Show `cloc` info of every file in the directory and it's subdirectories.
- [ ] Deep analysis on Golang codebase. (Currently only Golang codebase is supported.)
