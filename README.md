# mus-gen-go: Code Generator for MUS

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-gen-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-gen-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-gen-go)](https://goreportcard.com/report/github.com/mus-format/mus-gen-go)
[![codecov](https://codecov.io/gh/mus-format/mus-gen-go/graph/badge.svg?token=J1JN0CEY9S)](https://codecov.io/gh/mus-format/mus-gen-go)

**mus-gen** is a Go code generator for the [mus](https://github.com/mus-format/mus-go) and [mus-stream](https://github.com/mus-format/mus-stream-go) serializers.

## Capabilities

- Generates high-performance serialization code with optional unsafe
  optimizations.
- Supports both in-memory (`mus-go`) and streaming (`mus-stream-go`) data 
  processing models.
- Can generate code for parameterized types, interfaces and types with multiple 
  versions.
- Provides multi-package support.
- Enables cross-package code generation.
- Can be extended to support additional binary serialization formats beyond MUS.

## Contents

- [mus-gen-go: Code Generator for MUS](#mus-gen-go-code-generator-for-mus)
  - [Capabilities](#capabilities)
  - [Contents](#contents)
  - [Getting Started Example](#getting-started-example)
  - [Generator](#generator)
    - [Configuration](#configuration)
      - [Required Options](#required-options)
      - [Streaming](#streaming)
      - [Modes](#modes)
        - [Safe Mode](#safe-mode)
        - [Unsafe Mode](#unsafe-mode)
        - [Not Unsafe Mode](#not-unsafe-mode)
      - [Imports](#imports)
      - [Serializer Name](#serializer-name)
    - [Methods](#methods)
      - [AddDefinedType()](#adddefinedtype)
      - [AddStruct()](#addstruct)
      - [AddTyped()](#addtyped)
      - [AddInterface()](#addinterface)
      - [RegisterInterface()](#registerinterface)
      - [AddVersioned()](#addversioned)
      - [RegisterVersioned()](#registerversioned)
  - [Multi-package support](#multi-package-support)
  - [Cross-Package Code Generation](#cross-package-code-generation)
  - [Serialization Options](#serialization-options)
    - [Numbers](#numbers)
    - [String](#string)
    - [Array](#array)
    - [Slice](#slice)
    - [Map](#map)
    - [time.Time](#timetime)
  - [MUS Format Defaults](#mus-format-defaults)
    - [Safe Mode](#safe-mode-1)
    - [Unsafe Mode](#unsafe-mode-1)
    - [Not Unsafe Mode](#not-unsafe-mode-1)

## Getting Started Example

Here, we will generate a MUS serializer for the `Foo` type.

First, download and install Go (version 1.24 or later). Then, create a `foo`
folder with the following structure:

```
foo/
 |‒‒‒gen/
 |    |‒‒‒main.go
 |‒‒‒foo.g

**foo.go**

```go
//go:generate go run gen/main.go
package foo

type Int int

type Foo[T any] struct {
  s string
  t T
}
```

**gen/main.go**

```go
package main

import (
  "os"
  "reflect"

  "example.com/foo"

  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

func main() {
  g, err := musgen.NewGenerator(
    genopts.WithPkgPath("example.com/foo"),
    // genopts.WithPackage("bar"), // Can be used to specify the package name for
    // the generated file.
  )
  if err != nil {
    panic(err)
  }
  err = g.AddDefinedType(reflect.TypeFor[foo.Int]())
  if err != nil {
    panic(err)
  }
  err = g.AddStruct(reflect.TypeFor[foo.Foo[foo.Int]]())
  if err != nil {
    panic(err)
  }
  bs, err := g.Generate()
  if err != nil {
    // In case of an error (e.g., if you forget to specify an import path using 
    // genopts.WithImport), the generated code can be inspected for additional 
    // details.
    log.Println(err)
  }
  err = os.WriteFile("./mus.gen.go", bs, 0644)
  if err != nil {
    panic(err)
  }
}
```

Run from the command line:

```bash
cd ~/foo
go mod init example.com/foo
go mod tidy
go generate
go mod tidy
```

Now you can see `mus.gen.go` file in the `foo` folder with `IntMUS`
and `FooMUS` serializers (see full [example](https://github.com/mus-format/examples-go/tree/main/mus-gen)). 

## Generator

The `Generator` is responsible for generating serialization code.

### Configuration

#### Required Options

There is only one required configuration option:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g, err := musgen.NewGenerator(
  genopts.WithPkgPath("pkg path"),  // Sets the package path where the generated 
  // file will be placed. The path must match the standard Go package path 
  // format (e.g., github.com/user/project/pkg) and can be obtained using:
  //
  //   pkgPath := reflect.TypeFor[YourType]().PkgPath()
  //
)
```

#### Streaming

To generate streaming code:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...
  genopts.WithStream(),
)
```

In this case [mus-stream](https://github.com/mus-format/mus-stream-go)
library will be used instead of `mus`.

#### Modes

##### Safe Mode

By default, the generator generates safe code (`unsafe` package is not used).

##### Unsafe Mode

To generate unsafe code:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...
  genopts.WithUnsafe(),
)
```

##### Not Unsafe Mode

In this mode, the `unsafe` package is used for all types except `string`:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...  
  genopts.WithNotUnsafe(),
)
```

It produces the fastest serialization code without unsafe side effects.

#### Imports

In some cases import statement of the generated file can miss one or more
packages. To fix this:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...
  genopts.WithImport("import path"),
  genopts.WithImportAlias("import path", "alias"),
)
```

Also, `genopts.WithImportAlias` helps prevent name conflicts when multiple
packages are imported with the same alias.

#### Serializer Name

Generated serializers follow the following naming convention:

```
pkg.YourType[T,V] -> YourTypeMUS  // Serialization format is appended to the type name.
```

To override this behavior, use `genopts.WithSerName()`:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...
  genopts.WithSerName(reflect.TypeFor[pkg.YourType](), "CustomSerName"),
)
```

### Methods

#### AddDefinedType()

Supports types defined with the following underlying types:

- Number (`uint`, `int`, `float64`, `float32`, ...)
- String
- Array
- Slice
- Map
- Pointer

For example:

```go
type Int int
type StringSlice []string
type UintPtr *uint
```

It can be used as follows:

```go
import tpopts "github.com/mus-format/mus-gen-go/options/type"

type Int int // Where int is the underlying type of Int.

err := g.AddDefinedType(reflect.TypeFor[Int]())
```

Or with serialization options, for example:

```go
err := g.AddDefinedType(reflect.TypeFor[Int](),
  tpopts.WithNumEncoding(tpopts.NumEncodingRaw), // The raw.Int serializer will be used
  // to serialize the underlying int type.
  tpopts.WithValidator("ValidateInt")) // After unmarshalling, the Int 
  // value will be validated using the ValidateInt function.
  // Validator functions in general should have the following signature:
  //
  //   func(value Type) error
  //
  // where Type denotes the type the validator is applied to.
```

#### AddStruct()

Supports types defined with the `struct` underlying type, such as:

```go
type Struct struct { ... }
type AnotherStruct Struct
```

It can be used as follows:

```go
import (
  "reflect"
 
  genopts "github.com/mus-format/mus-gen-go/options/gen"
  fldopts "github.com/mus-format/mus-gen-go/options/field"
  stopts "github.com/mus-format/mus-gen-go/options/struct"
  tpopts "github.com/mus-format/mus-gen-go/options/type"
)

type Struct struct {
  Str string
  Ignore int
  Slice []int
  // Interface Interface  // Interface fields are supported as well.
  // Any any              // But not the `any` type.
}

// ...

err := g.AddStruct(reflect.TypeFor[Struct]())
```

Or with serialization options, for example:

```go
// The number of options should be equal to the number of fields. If you don't
// want to specify options for some field, use stopts.WithField() without
// any parameters.
err := g.AddStruct(reflect.TypeFor[Struct](),
  // No options for the first field.
  stopts.WithField(), 

  // The second field will not be serialized.
  stopts.WithField(fldopts.WithIgnore()), 

  // Options for the third field.
  stopts.WithField( 
    fldopts.WithType(
      // The length of the slice field will be validated using the ValidateLength 
      // function before the rest of the slice is unmarshalled.
      tpopts.WithLenValidator("ValidateLength"), 
      // Each slice element, after unmarshalling, will be validated using the 
      // ValidateSliceElem function.
      tpopts.WithElemValidator("ValidateSliceElem"), 
    ),
  ),
)
```

A special case for the `time.Time` underlying type:

```go
type Time time.Time

err = g.AddStruct(reflect.TypeFor[Time](),
  stopts.WithUnderlyingTime(
    // By default TimeUnitSecUTC is used, but you can change it:
    // stopts.WithUnderlyingTimeUnit(tpopts.TimeUnitMilli),
  ),
)
```

#### AddTyped()

Supports all types acceptable by the `AddDefinedType`, `AddStruct`, and
`AddInterface` methods.

It can be used as follows:

```go
import (
  "reflect"
)

type Int int

t := reflect.TypeFor[Int]()
err := g.AddDefinedType(t)
// ...
err = g.AddTyped(t)
```

The typed serializer definition will be generated for the specified type.

#### AddInterface()

Supports types defined with the `interface` underlying type, such as:

```go
type Interface interface { ... }
type AnyInterface any
type AnotherInterface Interface
```

It can be used as follows:

```go
import (
  com "github.com/mus-format/common-go"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
)

// 1. Define DTM values for each implementation type.
const (
  Impl1DTM com.DTM = iota + 1
  Impl2DTM
)

type Interface interface {...}
type Impl1 struct {...}
type Impl2 int

var (
  t1 = reflect.TypeFor[Impl1]()
  t2 = reflect.TypeFor[Impl2]()
)

// 2. Add interface implementations.
err := g.AddStruct(t1)
// ...
err = g.AddDefinedType(t2)
// ...

// 3. Add typed serializers for each implementation type.
err = g.AddTyped(t1)
// ...
err = g.AddTyped(t2)
// ...

// 4. Add interface with implementations.
err = g.AddInterface(reflect.TypeFor[Interface](),
  intropts.WithImplType(t1),
  intropts.WithImplType(t2),
  // intropts.WithMarshaller(), // Enables serialization using the 
  // mus.MarshallerTyped interface, that should be satisfied by all 
  // implementation types. Disabled by default.
)
```

#### RegisterInterface()

A convenience method that performs the full registration flow for an interface
and all of its implementations.

Unlike `AddInterface`, `RegisterInterface` does not require you to define DTM
values or call `AddStruct`,  `AddDefinedType`, `AddTyped` manually:

```go
import (
  "reflect"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
)

type Interface interface { ... }
type Impl1 struct { ... }
type Impl2 int

// ...

err := g.RegisterInterface(reflect.TypeFor[Interface](),
  intropts.WithStructImpl(reflect.TypeFor[Impl1]()),
  intropts.WithDefinedTypeImpl(reflect.TypeFor[Impl2]()),
  // intropts.WithRegisterMarshaller() // optional
)
```

#### AddVersioned()

Produce a versioned serializer for the specified type.

```go
import (
  com "github.com/mus-format/common-go"
  veropts "github.com/mus-format/mus-gen-go/options/versioned"
)

// 1. Define DTM values for each type version.
const (
  FooV1DTM com.DTM = iota + 1
  FooV2DTM
)

type Foo FooV2    // target type Foo
type FooV2 string // current type version
type FooV1 int    // old type version

var (
  t1 = reflect.TypeFor[FooV1]()
  t2 = reflect.TypeFor[FooV2]()
)

// 2. Add type versions.
err := g.AddDefinedType(t1)
// ...
err = g.AddDefinedType(t2)
// ...

// 3. Add typed serializers for each version.
err = g.AddTyped(t1)
// ...
err = g.AddTyped(t2)
// ...

// 4. Add target type with versions.
err = g.AddVersioned(reflect.TypeFor[Foo](),
  veropts.WithVersion(t1, "MigrateFooV1"),
  veropts.WithCurrentVersion(t2), // Do not need migrate function.
)
```

The migration function must accept an old type version and return the target type:

```go
func MigrateFooV1(v FooV1) Foo { ... }
```

#### RegisterVersioned()

A convenience method that performs the full registration flow for a specified
type and all of its versions.

Unlike `AddVersioned`, `RegisterVersioned` does not require you to define DTM
values or call `AddStruct`/`AddDefinedType`, `AddTyped` for each type version
manually:

```go
import (
  "reflect"
  veropts "github.com/mus-format/mus-gen-go/options/versioned"
)

type Foo FooV2    // target type Foo
type FooV2 string // current type version
type FooV1 int    // old type version

err := g.RegisterVersioned(reflect.TypeFor[Foo](),
  veropts.WithVersion(reflect.TypeFor[FooV1](), "MigrateFooV1"),
  veropts.WithCurrentVersion(reflect.TypeFor[FooV2]()),
)
```

## Multi-package support

By default, `mus-gen` expects a type’s serializer to reside in the same package
as the type itself. For example, generating a serializer for the `Foo` type:

```go
package foo

type Foo struct{
  Bar bar.Bar
}
```

will result in:

```go
package foo

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  return bar.BarMUS(v.Bar) // mus-gen assumes the Bar serializer is located
  // in the bar package and follows the default naming convention.
}
```

To reference a `Bar` serializer defined in a different package or with a
non-standard name, use the `genopts.WithSerName` option:

```go
import (
  musgen "github.com/mus-format/mus-gen-go/mus"
  genopts "github.com/mus-format/mus-gen-go/options/gen"
)

g := musgen.NewGenerator(
  // ...
  genopts.WithSerName(reflect.TypeFor[bar.Bar](), "another.AwesomeBar"),
  // If only the name is non-standard:
  // genopts.WithSerName(reflect.TypeFor[bar.Bar](), "AwesomeBar"),
)
```

The string `another.AwesomeBar` will be used as-is, with the serialization
format appended:

```go
func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  return another.AwesomeBarMUS(v.Bar)
}
```

## Cross-Package Code Generation

`mus-gen` allows to generate a serializer for a type from an external package,
for example, `foo.BarMUS` for the `bar.Bar` type.

## Serialization Options

Different types support different serialization options. If an unsupported
option is specified for a type, it will simply be ignored.

```go
import tpopts "github.com/mus-format/mus-gen-go/options/type"
```

### Numbers

- `tpopts.WithNumEncoding`
- `tpopts.WithValidator`

### String

- `tpopts.WithLenEncoding`
- `tpopts.WithLenValidator`
- `tpopts.WithValidator`

### Array

- `tpopts.WithLenEncoding`
- `tpopts.WithElemValidator`
- `tpopts.WithValidator`

### Slice

- `tpopts.WithLenEncoding`
- `tpopts.WithLenValidator`
- `tpopts.WithElemValidator`
- `tpopts.WithValidator`

### Map

- `tpopts.WithLenEncoding`
- `tpopts.WithLenValidator`
- `tpopts.WithKeyValidator`
- `tpopts.WithElemValidator`
- `tpopts.WithValidator`

### time.Time

- `tpopts.WithTimeUnit`
- `tpopts.WithValidator`

## MUS Format Defaults

### Safe Mode

```text
int:        varint
uint:       varint
float:      raw
byte:       raw
bool:       ord
string:     ord
byte_slice: ord
time:       raw
time_ser:   TimeUnixUTC // Default time serializer.
```

### Unsafe Mode

```text
int:        unsafe
uint:       unsafe
float:      unsafe
byte:       unsafe
bool:       unsafe
string:     unsafe
byte_slice: unsafe // ord for mus-stream
time:       unsafe
time_ser:   TimeUnixUTC // Default time serializer.
```

### Not Unsafe Mode

```text
int:        unsafe
uint:       unsafe
float:      unsafe
byte:       unsafe
bool:       unsafe
string:     ord
byte_slice: unsafe // ord for mus-stream
time:       unsafe
time_ser:   TimeUnixUTC // Default time serializer.
```
