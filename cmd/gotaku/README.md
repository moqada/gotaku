# gotaku

:fish: Web GYOTAKU ([ウェブ魚拓](http://megalodon.jp/)) command.
This is Go port of [gyotaku-cli](https://github.com/moqada/gyotaku-cli).

This command can take new GYOTAKU and list GYOTAKU urls.

## Install

```
$ go get github.com/moqada/gotaku/cmd/gotaku
```

## Usage

```
Usage: gotaku [options] <URL>

Options:
  -l    Listing GYOTAKU urls

Examples:
  gotaku http://example.com/     Take GYOTAKU of specific url
  gotaku -l http://example.com/  List GYOTAKU urls for specific url
```
