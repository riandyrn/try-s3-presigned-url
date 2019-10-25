# Go Env

[![Build Status](https://travis-ci.org/riandyrn/go-env.svg?branch=master)](https://travis-ci.org/riandyrn/go-env)
[![Go Report Card](https://goreportcard.com/badge/github.com/riandyrn/go-env)](https://goreportcard.com/report/github.com/riandyrn/go-env)

Fetching primitive values from env without hassle in Go!

## Why

Have you been annoyed when your code become ugly because you need to fetch value other than string from env but you also need to take care of its parsing error yet the only thing you need is just the value?

Here is an example of such situation when you are trying to fetch `int` value:

```go
...
strVal := os.Getenv("IntKey")
intVal, _ := strconv.Atoi(strVal) // too much hassle!
...
```

This module is offering you solution to that!

```go
intVal := env.GetInt("IntKey") // so simple!
```

## Installation

```bash
go get github.com/riandyrn/go-env
```

## Code Usage

```go
...

import "github.com/riandyrn/go-env"

...
strVal := env.GetString("StrKey") // fetch string value
intVal := env.GetInt("IntKey") // fetch int value
boolVal := env.GetBool("BoolKey") // fetch bool value
...
```

No need to care that troublesome errors!

All errors in the process would simply make the functions return their returned type zero value (e.g `env.GetInt("abc")` would returns `0`).

Check `env_test.go` to see possible scenarios handled by this module.

For full working example, check out `/example` directory.

---
