# GopherLibs => Humanize [![Go Reference][gopkg]](https://pkg.go.dev/github.com/gopherlibs/humanize) [![Go Report Card](https://goreportcard.com/badge/github.com/gopherlibs/humanize)](https://goreportcard.com/report/github.com/gopherlibs/humanize) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/humanize/trunk/LICENSE)

*This project is a hard fork of `github.com/dustin/go-humanize`. The project continues here, with maintenance and potentially new features.*

`Humanize` is a Go package to format various units and words into human friendly terms.


## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Migrating](#migrating)
- [Credits](#credits)


## Requirements

- The minimum Go version supported is v1.24.


## Installation

Run `go get github.com/gopherlibs/humanize/humanize` in your project directory.
Then you can import the package in whatever files you use it in.


## Usage

This module is fairly simple. Most formatters are functions of the package.
Simply determine the conversion you need and find the corresponding function.

Here are some examples.

### Sizes

This lets you take numbers like `82854982` and convert them to useful strings like, `83 MB` or `79 MiB` (whichever you prefer).

```go
fmt.Printf("That file is %s.", humanize.Bytes(82854982)) // That file is 83 MB.
```

### Times

This lets you take a `time.Time` and spit it out in relative terms.
For example, `12 seconds ago` or `3 days from now`.

```go
fmt.Printf("This was touched %s.", humanize.Time(someTimeInstance)) // This was touched 7 hours ago.
```

### Ordinals

This lets you label ordinals.

    0 -> 0th
    1 -> 1st
    2 -> 2nd
    3 -> 3rd
    4 -> 4th
    [...]

```go
fmt.Printf("You're my %s best friend.", humanize.Ordinal(193)) // You are my 193rd best friend.
```

### Commas

Want to shove commas into numbers? Be my guest.

    0 -> 0
    100 -> 100
    1000 -> 1,000
    1000000000 -> 1,000,000,000
    -100000 -> -100,000

```go
fmt.Printf("You owe $%s.\n", humanize.Comma(6582491)) // You owe $6,582,491.
```

### Ftoa

Nicer float64 formatter that removes trailing zeros.

```go
fmt.Printf("%f", 2.24)                // 2.240000
fmt.Printf("%s", humanize.Ftoa(2.24)) // 2.24
fmt.Printf("%f", 2.0)                 // 2.000000
fmt.Printf("%s", humanize.Ftoa(2.0))  // 2
```

### SI notation

Format numbers with [SI notation](http://en.wikipedia.org/wiki/Metric_prefix).

```go
humanize.SI(0.00000000223, "M") // 2.23 nM
```

### English-specific functions

The following functions are in the `github.com/gopherlibs/humanize/humanize/english` package.

#### Plurals
Simple English pluralization

```go
english.PluralWord(1, "object", "") // object
english.PluralWord(42, "object", "") // objects
english.PluralWord(2, "bus", "") // buses
english.PluralWord(99, "locus", "loci") // loci

english.Plural(1, "object", "") // 1 object
english.Plural(42, "object", "") // 42 objects
english.Plural(2, "bus", "") // 2 buses
english.Plural(99, "locus", "loci") // 99 loci
```

#### Word series
Format comma-separated words lists with conjuctions:

```go
english.WordSeries([]string{"foo"}, "and") // foo
english.WordSeries([]string{"foo", "bar"}, "and") // foo and bar
english.WordSeries([]string{"foo", "bar", "baz"}, "and") // foo, bar and baz

english.OxfordWordSeries([]string{"foo", "bar", "baz"}, "and") // foo, bar, and baz
```

See the [Go Pkg Docs][gopkg] for complete documentation.


## Development

This library is written and tested with Go v1.24+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## Migrating

Are you migrating from `github.com/dustin/go-humanize`?
It isn't complicated.
Run through the [installation](#installation) steps above.
In the import statements for each file, replace `github.com/dustin/go-humanize` with `github.com/gopherlibs/humanize/humanize`.
That should be it.


## Credits

This module is maintained by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).

This module is a hard fork of [dustin/go-humaize](https://github.com/dustin/go-humanize).
That project seems to no longer be maintained.
This project will continue where that one left off, including the original git history.





[gopkg]: https://pkg.go.dev/badge/github.com/gopherlibs/humanize.svg
