# Code-talks

A tool that helps analyze your codebase, make it self-talk!

## Quick Start

Please make sure you have installed Go 1.16+.

### Installation

```bash
go install github.com/92hackers/code-talks@latest
```

### Usage

```bash
cd /path/to/your/project

# Analyze the whole project
codetalks
```

## Features overview

Alias: `cloc` means `count number of lines of code, comments and blanks`.

1. Show `cloc` info grouped by language in the project.
2. Show `cloc` info of every directory and it's subdirectories.
3. Sort files or directories by `cloc` info, such as set `lines of code` as the sort criteria.
4. Blazing fast analyze.


## Roadmap

- [ ] Show `cloc` info of every file in the project.
- [x] Show `cloc` info of every file in the directory and it's subdirectories.
- [ ] Deep analysis on Golang codebase. (Currently only Golang codebase is supported.)
