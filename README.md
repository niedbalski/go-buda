# Go-Buda

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/niedbalski/go-buda)
[![Build Status](https://travis-ci.org/niedbalski/go-buda.svg?branch=master)](https://travis-ci.org/niedbalski/go-buda)

Golang implementation of a API client for the [Buda](https://api.buda.com) crypto exchange.

### Usage

Please check the buda_test.go for additional examples.

```go
package main

import (
	"github.com/niedbalski/go-buda"
	"fmt"
)


func main() (){
	buda, err := buda.NewAPIClient("key", "secret")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(buda.GetMarkets())
	
}

```
