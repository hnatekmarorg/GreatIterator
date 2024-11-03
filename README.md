# GreatIterator
Iterates on problems via LLM until it finishes.

# Project Description
GreatIterator is a Go project designed to iterate on problems using a Language Model (LLM) until the problem is resolved. This project is primarily written in Go and includes a Dockerfile for containerization.

# Installation
You can install GreatIterator via go install or by using the provided binary releases.

Using go install
Install the package:

go install github.com/hnatekmarorg/GreatIterator@latest
Run the project:

`GreatIterator --help`
Using Binary Releases
Download the latest release from the [Releases page](https://github.com/hnatekmarorg/GreatIterator/releases).

Extract the binary and move it to your PATH:

`mv GreatIterator /usr/local/bin/`
Run the project:

`GreatIterator --help`

# Fixing files
In order for it to work you need to provide:
- test command (that returns 0 on succes and anything else on failure)
- one or more files that llm will modify
- access token for openai or openai compatible endpoint
For example `GreatIterator fix --openai-token 'sk-...' 'gcc examples/main.c' examples/main.c` will continuously execute `gcc examples/main.c` and change `examples/main.c` until it works
