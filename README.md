# go-scule

This is a port of the [unjs.io](https://unjs.io/) scule package to Go.

## Install

```bash
go get github.com/givensuman/go-scule
```

Import:

```go
import (
  "github.com/givensuman/go-scule"
)
```

## Utils

### `PascalCase(str string, normalize bool) string`

Splits string and joins by PascalCase convention:

```go
PascalCase("foo-bar_baz");
// FooBarBaz
```

**Notice:** If an uppercase letter is followed by other uppercase letters (like `FooBAR`), they are preserved. You can use the `normalize` parameter for strictly following pascalCase convention.

### `CamelCase(str string, normalize bool) string`

Splits string and joins by camelCase convention:

```go
CamelCase("foo-bar_baz");
// fooBarBaz
```

### `KebabCase(str string) string`

Splits string and joins by kebab-case convention:

```go
KebabCase("fooBar_Baz");
// foo-bar-baz
```

**Notice:** It does **not** preserve case.

### `SnakeCase(str string) string`

Splits string and joins by snake_case convention:

```go
SnakeCase("foo-barBaz");
// foo_bar_baz
```

### `FlatCase(str strsing) string`

Splits string and joins by flatcase convention:

```go
FlatCase("foo-barBaz");
// foobarbaz
```

### `TrainCase(str string, normalize bool) string`

Split string and joins by Train-Case (a.k.a. HTTP-Header-Case) convention:

```go
TrainCase("FooBARb");
// Foo-Ba-Rb
```

**Notice:** If an uppercase letter is followed by other uppercase letters (like `WWWAuthenticate`), they are preserved (=> `WWW-Authenticate`). You can use `{ normalize: true }` for strictly only having the first letter uppercased.

### `TitleCase(str string, normalize bool) string`

With Title Case all words are capitalized, except for minor words.
A compact regex of common minor words (such as `a`, `for`, `to`) is used to automatically keep them lower case.

```go
TitleCase("this-IS-aTitle");
// This is a Title
```

### `UpperFirst(str string) string`

Converts first character to upper case:

```go
UpperFirst("hello world!");
// Hello world!
```

### `LowerFirst(str string) string`

Converts first character to lower case:

```go
LowerFirst("Hello world!");
// hello world!
```

### `SplitByCase(str, splitters *[]string) string`

- Splits string by the splitters provided (default: `{"-", "_", "/", ".", " "}`)
- Splits when case changes from lower to upper or upper to lower
- Ignores numbers for case changes
- Case is preserved in returned value
- Is an irreversible function since splitters are omitted

## License

[MIT](./LICENSE)

<!-- Badges -->

