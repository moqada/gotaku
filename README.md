# gotaku

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

## CLI

- [cmd/gotaku](https://github.com/moqada/gotaku/tree/master/cmd/gotaku).
