# optionGen

## Functional Options
Functional options are an idiomatic way of creating APIs with options on types. The initial idea for this design pattern can be found in an article published by Rob Pike called [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html).

optionGen is a fork of [XSAM/optionGen](https://github.com/XSAM/optionGen), a tool to generate go Struct option for test, mock or more flexible. The purpose of this fork is to provide more powerful and flexible option generation. 

## Install
Install using go get, and this will build the optionGen binary in $GOPATH/bin.
```bash
go get github.com/timestee/optionGen/...
```

optionGen require [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) to format code which is generated. So you may confirm that `goimports` has been installed

```bash
go get golang.org/x/tools/cmd/goimports
```

## Using optionGen
To generate struct option, you need write a function declaration to tell optionGen how to generate.struct name and `OptionDeclareWithDefault` suffix. In this function, just return a variable which type is `map[string]interface{}`.

The key of the map means option name, and the value of the map should consist of two parts, one for option type(except func type), and the other option default value.

Here is an example.
```go
//go:generate optionGen --option_with_struct_name=false
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Sounds": string("Meow"),
		"Food":   (*string)(nil),
		"Walk": func() {
			log.Println("Walking")
		},
	}
}

```

To use a optionGen, you must tell optionGen that you want to use it using a special comment in your code. For example
```
//go:generate optionGen --option_with_struct_name=false
```
This tells go generate to run optionGen and that you want to ignore the struct name for option function.

Here is the sample result generate by `optionGen`

```go
// Code generated by optionGen. DO NOT EDIT.
package example

import "log"

type Config struct {
	Sounds string
	Food   (*string)
	Walk   func()
}

type ConfigOption func(oo *Config)

func WithSounds(v string) ConfigOption { return func(oo *Config) { oo.Sounds = v } }
func WithFood(v *string) ConfigOption  { return func(oo *Config) { oo.Food = v } }
func WithWalk(v func()) ConfigOption   { return func(oo *Config) { oo.Walk = v } }

func NewConfig(opts ...ConfigOption) *Config {
	ret := newDefaultConfig()
	for _, o := range opts {
		o(ret)
	}
	return ret
}

var defaultConfigOptions = [...]ConfigOption{
	WithSounds("Meow"),
	WithFood(nil),
	WithWalk(func() {
		log.Println("Walking")
	}),
}

func newDefaultConfig() *Config {
	ret := &Config{}
	for _, o := range defaultConfigOptions {
		o(ret)
	}
	return ret
}


```

See a complete example in the [example](https://github.com/timestee/optionGen/blob/master/example/cat.go) directory.
