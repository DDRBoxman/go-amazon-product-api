Simple library to simplify grabbing and posting data from the Amazon Affiliate API

[![Build Status](https://travis-ci.org/DDRBoxman/go-amazon-product-api.svg?branch=master)](https://travis-ci.org/DDRBoxman/go-amazon-product-api)
[![MIT](http://img.shields.io/badge/license-MIT-green.svg)](LICENSE) [![GODOC](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/DDRBoxman/go-amazon-product-api)


Have a look at the go docs here:
http://godoc.org/github.com/DDRBoxman/go-amazon-product-api

Example
_______

	package main

	import (
		"fmt"
		"github.com/DDRBoxman/go-amazon-product-api"
	)

	func main() {
		var api amazonproduct.AmazonProductAPI

		api.AccessKey = ""
		api.SecretKey = ""
		api.Host = "webservices.amazon.com"
		api.AssociateTag = ""

		result,err := api.ItemSearchByKeyword("sgt+frog")
		if (err != nil) {
			fmt.Println(err)
		}

		fmt.Println(result)
	}
