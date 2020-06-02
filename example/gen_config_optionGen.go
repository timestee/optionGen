// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optionGen

package example

import "log"

type Config struct {
	TestNil             interface{}
	TestBool            bool
	TestInt             int
	TestInt64           int64
	TestSliceInt        []int
	TestSliceInt64      []int64
	TestSliceString     []string
	TestSliceBool       []bool
	TestSliceIntNil     []int
	TestSliceIntEmpty   []int
	TestMapIntInt       map[int]int
	TestMapIntString    map[int]string
	TestMapStringInt    map[string]int
	TestMapStringString map[string]string
	TestString          string
	Food                *string
	Walk                func()
	TestNilFunc         func()
	TestReserved1_      []byte
	TestReserved2Inner  int
}

func (cc *Config) SetOption(opt ConfigOption) {
	_ = opt(cc)
}

func (cc *Config) GetSetOption(opt ConfigOption) ConfigOption {
	return opt(cc)
}

type ConfigOption func(cc *Config) ConfigOption

func WithTestNil(v interface{}) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestNil
		cc.TestNil = v
		return WithTestNil(previous)
	}
}
func WithTestBool(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestBool
		cc.TestBool = v
		return WithTestBool(previous)
	}
}
func WithTestInt(v int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestInt
		cc.TestInt = v
		return WithTestInt(previous)
	}
}
func WithTestInt64(v int64) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestInt64
		cc.TestInt64 = v
		return WithTestInt64(previous)
	}
}
func WithTestSliceInt(v []int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceInt
		cc.TestSliceInt = v
		return WithTestSliceInt(previous)
	}
}
func WithTestSliceInt64(v []int64) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceInt64
		cc.TestSliceInt64 = v
		return WithTestSliceInt64(previous)
	}
}
func WithTestSliceString(v []string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceString
		cc.TestSliceString = v
		return WithTestSliceString(previous)
	}
}
func WithTestSliceBool(v []bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceBool
		cc.TestSliceBool = v
		return WithTestSliceBool(previous)
	}
}
func WithTestSliceIntNil(v []int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceIntNil
		cc.TestSliceIntNil = v
		return WithTestSliceIntNil(previous)
	}
}
func WithTestSliceIntEmpty(v []int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceIntEmpty
		cc.TestSliceIntEmpty = v
		return WithTestSliceIntEmpty(previous)
	}
}
func WithTestMapIntInt(v map[int]int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapIntInt
		cc.TestMapIntInt = v
		return WithTestMapIntInt(previous)
	}
}
func WithTestMapIntString(v map[int]string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapIntString
		cc.TestMapIntString = v
		return WithTestMapIntString(previous)
	}
}
func WithTestMapStringInt(v map[string]int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapStringInt
		cc.TestMapStringInt = v
		return WithTestMapStringInt(previous)
	}
}
func WithTestMapStringString(v map[string]string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapStringString
		cc.TestMapStringString = v
		return WithTestMapStringString(previous)
	}
}
func WithTestString(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestString
		cc.TestString = v
		return WithTestString(previous)
	}
}
func WithFood(v *string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Food
		cc.Food = v
		return WithFood(previous)
	}
}
func WithWalk(v func()) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Walk
		cc.Walk = v
		return WithWalk(previous)
	}
}
func WithTestNilFunc(v func()) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestNilFunc
		cc.TestNilFunc = v
		return WithTestNilFunc(previous)
	}
}

func NewConfig(opts ...ConfigOption) *Config {
	cc := newDefaultConfig()
	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogConfig != nil {
		watchDogConfig(cc)
	}
	return cc
}

func InstallConfigWatchDog(dog func(cc *Config)) {
	watchDogConfig = dog
}

var watchDogConfig func(cc *Config)

var defaultConfigOptions = [...]ConfigOption{
	WithTestNil(nil),
	WithTestBool(false),
	WithTestInt(32),
	WithTestInt64(32),
	WithTestSliceInt([]int{1, 2, 3}),
	WithTestSliceInt64([]int64{1, 2, 3}),
	WithTestSliceString([]string{"test1", "test2"}),
	WithTestSliceBool([]bool{false, true}),
	WithTestSliceIntNil(nil),
	WithTestSliceIntEmpty(nil),
	WithTestMapIntInt(map[int]int{1: 1, 2: 2, 3: 3}),
	WithTestMapIntString(map[int]string{1: "test"}),
	WithTestMapStringInt(map[string]int{"test": 1}),
	WithTestMapStringString(map[string]string{"test": "test"}),
	WithTestString("Meow"),
	WithFood(nil),
	WithWalk(func() {
		log.Println("Walking")
	}),
	WithTestNilFunc(nil),
}

func newDefaultConfig() *Config {
	cc := &Config{
		TestReserved1_:     nil,
		TestReserved2Inner: 1,
	}

	for _, opt := range defaultConfigOptions {
		_ = opt(cc)
	}

	return cc
}

type Config2 struct {
	TestNil             interface{}
	TestBool            bool
	TestInt             int
	TestInt64           int64
	TestSliceInt        []int
	TestSliceInt64      []int64
	TestSliceString     []string
	TestSliceBool       []bool
	TestSliceIntNil     []int
	TestSliceIntEmpty   []int
	TestMapIntInt       map[int]int
	TestMapIntString    map[int]string
	TestMapStringInt    map[string]int
	TestMapStringString map[string]string
	TestString          string
	Food                *string
	Walk                func()
	TestNilFunc         func()
	TestReserved1_      []byte
	TestReserved2Inner  int
}

func (cc *Config2) SetOption(opt Config2Option) {
	_ = opt(cc)
}

func (cc *Config2) GetSetOption(opt Config2Option) Config2Option {
	return opt(cc)
}

type Config2Option func(cc *Config2) Config2Option

func WithTestNil(v interface{}) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestNil
		cc.TestNil = v
		return WithTestNil(previous)
	}
}
func WithTestBool(v bool) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestBool
		cc.TestBool = v
		return WithTestBool(previous)
	}
}
func WithTestInt(v int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestInt
		cc.TestInt = v
		return WithTestInt(previous)
	}
}
func WithTestInt64(v int64) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestInt64
		cc.TestInt64 = v
		return WithTestInt64(previous)
	}
}
func WithTestSliceInt(v []int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceInt
		cc.TestSliceInt = v
		return WithTestSliceInt(previous)
	}
}
func WithTestSliceInt64(v []int64) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceInt64
		cc.TestSliceInt64 = v
		return WithTestSliceInt64(previous)
	}
}
func WithTestSliceString(v []string) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceString
		cc.TestSliceString = v
		return WithTestSliceString(previous)
	}
}
func WithTestSliceBool(v []bool) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceBool
		cc.TestSliceBool = v
		return WithTestSliceBool(previous)
	}
}
func WithTestSliceIntNil(v []int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceIntNil
		cc.TestSliceIntNil = v
		return WithTestSliceIntNil(previous)
	}
}
func WithTestSliceIntEmpty(v []int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestSliceIntEmpty
		cc.TestSliceIntEmpty = v
		return WithTestSliceIntEmpty(previous)
	}
}
func WithTestMapIntInt(v map[int]int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestMapIntInt
		cc.TestMapIntInt = v
		return WithTestMapIntInt(previous)
	}
}
func WithTestMapIntString(v map[int]string) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestMapIntString
		cc.TestMapIntString = v
		return WithTestMapIntString(previous)
	}
}
func WithTestMapStringInt(v map[string]int) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestMapStringInt
		cc.TestMapStringInt = v
		return WithTestMapStringInt(previous)
	}
}
func WithTestMapStringString(v map[string]string) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestMapStringString
		cc.TestMapStringString = v
		return WithTestMapStringString(previous)
	}
}
func WithTestString(v string) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestString
		cc.TestString = v
		return WithTestString(previous)
	}
}
func WithFood(v *string) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.Food
		cc.Food = v
		return WithFood(previous)
	}
}
func WithWalk(v func()) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.Walk
		cc.Walk = v
		return WithWalk(previous)
	}
}
func WithTestNilFunc(v func()) Config2Option {
	return func(cc *Config2) Config2Option {
		previous := cc.TestNilFunc
		cc.TestNilFunc = v
		return WithTestNilFunc(previous)
	}
}

func NewConfig2(opts ...Config2Option) *Config2 {
	cc := newDefaultConfig2()
	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogConfig2 != nil {
		watchDogConfig2(cc)
	}
	return cc
}

func InstallConfig2WatchDog(dog func(cc *Config2)) {
	watchDogConfig2 = dog
}

var watchDogConfig2 func(cc *Config2)

var defaultConfig2Options = [...]Config2Option{
	WithTestNil(nil),
	WithTestBool(false),
	WithTestInt(32),
	WithTestInt64(32),
	WithTestSliceInt([]int{1, 2, 3}),
	WithTestSliceInt64([]int64{1, 2, 3}),
	WithTestSliceString([]string{"test1", "test2"}),
	WithTestSliceBool([]bool{false, true}),
	WithTestSliceIntNil(nil),
	WithTestSliceIntEmpty(nil),
	WithTestMapIntInt(map[int]int{1: 1, 2: 2, 3: 3}),
	WithTestMapIntString(map[int]string{1: "test"}),
	WithTestMapStringInt(map[string]int{"test": 1}),
	WithTestMapStringString(map[string]string{"test": "test"}),
	WithTestString("Meow"),
	WithFood(nil),
	WithWalk(func() {
		log.Println("Walking")
	}),
	WithTestNilFunc(nil),
}

func newDefaultConfig2() *Config2 {
	cc := &Config2{
		TestReserved1_:     nil,
		TestReserved2Inner: 1,
	}

	for _, opt := range defaultConfig2Options {
		_ = opt(cc)
	}

	return cc
}
