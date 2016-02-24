# gotaku

[![Build Status][travis-image]][travis-url]
[![Doc][godoc-image]][godoc-url]
[![License][license-image]][license-url]

:fish: Web GYOTAKU ([ウェブ魚拓](http://megalodon.jp/)) Client for Go.
This is Go port of [gyotaku](https://github.com/moqada/gyotaku).

> Unofficial and Implemented by dirty scraping...

## Install

```
$ go get github.com/moqada/gotaku
```

## Usage

```go
import (
	"fmt"
	"github.com/moqada/gotaku"
)

// take GYOTAKU
gotaku.Take("http://example.com")
// "http://megalodon.jp/2015-1120-0000-00/example.com"

// listing GYOTAKU urls
gyotaku.List("http://example.com")
// ["http://megalodon.jp/2015-1120-0000-00/example.com", ...]

// listing GYOTAKU urls from HTML string
gyotaku.ListFromHTML("<html>...</html>">)
// ["http://megalodon.jp/2015-1120-0000-00/example.com", ...]
```

The documentation is on [GoDoc](https://godoc.org/github.com/moqada/gotaku)

## CLI

- [cmd/gotaku](https://github.com/moqada/gotaku/tree/master/cmd/gotaku).


[godoc-url]: https://godoc.org/github.com/moqada/gotaku
[godoc-image]: https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square
[travis-url]: https://travis-ci.org/moqada/gotaku
[travis-image]: https://img.shields.io/travis/moqada/gotaku.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/github/license/moqada/gotaku.svg?style=flat-square
