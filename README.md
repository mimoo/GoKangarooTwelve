# Kangaroo Twelve Implementation in Go

This is an implementation of [KangarooTwelve](http://keccak.noekeon.org/kangarootwelve.html) in Go.

It is heavily based on the official Go's [x/crypto/sha3](https://godoc.org/golang.org/x/crypto/sha3) library. But because of minor implementation details the relevant files have been copied and modified here so you do not need Go's SHA-3 implementation to run this package. Hopefuly one day Go's SHA-3 library will be more flexible to allow other keccak construction to rely on it.

I have tested this implementation with different test vectors and it works fine. Note that it has not received proper peer review. If you look at the code and find issues (or not) please [let me know](https://www.cryptologie.net/contact/)!

See here [why you should use KangarooTwelve instead of SHA-3](https://www.cryptologie.net/article/393/kangarootwelve/). But [see here first why you should still not skip SHA-3](https://www.cryptologie.net/article/400/maybe-you-shouldnt-skip-sha-3/).

## Installation

```sh
go get github.com/mimoo/GoKangarooTwelve/K12
```

## Usage

```go
package main

import(
    "fmt"
    "github.com/mimoo/GoKangarooTwelve/K12"
    "encoding/hex"
)

func main(){
    // custom string allows you to customize your hash function 🙃
    customString = []byte("myWebsite.com")
    hash := K12.NewK12(customString)
    payload := []byte("salut!")
    hash.Write(payload)
    out := make([]byte, 30)
    a.Read(out)
    fmt.Println(hex.EncodeToString(out))

    // or simpler
    fmt.Println(hex.EncodeToString(K12Sum(customString, payload, out)
}
```

## Other Keccak-based things

I've implemented a few other constructions that might be helpful:

* [STROBE](https://github.com/mimoo/StrobeGo/blob/master/golang.org/x/crypto/sha3/strobe.go) 
* [cSHAKE](https://github.com/mimoo/StrobeGo/blob/master/golang.org/x/crypto/sha3/sp800-185.go)


