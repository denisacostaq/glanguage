# glanguage
Translate English to Gophers's language (simple demo)

[![Build Status](https://travis-ci.com/denisacostaq/glanguage.svg?branch=develop)](https://travis-ci.com/denisacostaq/glanguage)
[![Coverage Status](https://coveralls.io/repos/github/denisacostaq/glanguage/badge.svg?branch=develop)](https://coveralls.io/github/denisacostaq/glanguage?branch=develop)

# CLI Server Documentation

Gopher's language command line interface server documentation.

<!-- MarkdownTOC autolink="true" bracket="round" levels="1,2,3" -->

# Documentation
- [CLI Documentation](#cli-documentation)
  - [System Requirements](#system-requirements)
  - [Install](#install)
  - [Usage](#usage)
    - [Run the server](#run-the-server)
      - [Examples](#examples)
        - [Show Help](#show-help)
        - [Change Default Port](#change-default-port)

<!-- /MarkdownTOC -->

## System Requirements

You need to have golang (version >= 1.12) installed on your system. To install go, just follow the [official instructions](https://golang.org/doc/install)

## Install

The instruction for install the server from source is the following:
```bash
GO111MODULE=on go get github.com/denisacostaq/glanguage/cmd/glanguage@develop
```

## Usage

### Run the server

Start listening in the port 8080 (default). 

```bash
$GOPATH/bin/glanguage
```

#### Examples
##### Show Help

```bash
$GOPATH/bin/glanguage
```
<details>
 <summary>View Output</summary>

```
NAME:
   glanguage - Translate English to the Gophers's language

USAGE:
   glanguage [global options] command [command options] [arguments...]

AUTHOR:
   Alvaro Denis <denisacostaq@gmail.co>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  Listen port [$GLANGUAGE_PORT]
   --help, -h              show help (default: false)
```
</details>

##### Change Default Port

```bash
$GOPATH/bin/glanguage --port 8081
```
<details>
 <summary>View Output</summary>

```
INFO[0000] Starting server...                            addr=":8081"
```
</details>