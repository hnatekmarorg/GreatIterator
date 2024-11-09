[![Release](https://github.com/hnatekmarorg/GreatIterator/actions/workflows/release.yaml/badge.svg?branch=main)](https://github.com/hnatekmarorg/GreatIterator/actions/workflows/release.yaml)

**GreatIterator**
================

Iterate on problems using a Language Model (LLM) until they're resolved.

**Project Description**
---------------

GreatIterator is a Go project that leverages a Language Model (LLM) to iteratively solve problems. Written primarily in Go, this project includes a Dockerfile for containerization.

**Installation**
-------------

You can install GreatIterator using one of the following methods:

### Using `go install`

1. Install the package: `go install github.com/hnatekmarorg/GreatIterator@latest`
2. Run the project: `GreatIterator --help`

### Using Binary Releases

1. Download the latest release from the [Releases page](https://github.com/hnatekmarorg/GreatIterator/releases).
2. Extract the binary and move it to your PATH: for example on linux `mv GreatIterator /usr/local/bin/`
3. Run the project: `GreatIterator --help`

### Using Docker

1. Clone the project.
2. Build the Docker image: `docker build . -t great-iterator`
3. Run the project: `docker run -ti great-iterator --help`

**Usage**
-----

To use GreatIterator, you need to provide:

* A test command that returns 0 on success and anything else on failure
* One or more files that the LLM will modify
* An access token for OpenAI or an OpenAI-compatible endpoint

Example:
```
GreatIterator fix --openai-token 'sk-...' 'gcc examples/main.c' examples/main.c
```
This command will continuously execute `gcc examples/main.c` and modify `examples/main.c` until it works.
