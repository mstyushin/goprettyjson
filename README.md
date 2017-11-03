# GoPrettyJson

Trivial go program for pretty-printing json data read from stdin.

## Installation:

1. [Install Go](https://golang.org/doc/install)
2. [Setup your GOPATH](https://golang.org/doc/code.html#GOPATH)
3. Run `go get github.com/mstyushin/goprettyjson`
4. Run `cd $GOPATH/src/github.com/mstyushin/goprettyjson`
5. Run `make`

## How to use it:

Say you have some some http service that responses with some json string. You can get it pretty-printed like this:

```
curl -XGET https://my.awesome.service.com/api/givemejson | pretty
```

## License
The MIT License (MIT)