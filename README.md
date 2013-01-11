Simple library to simplify grabbing and posting data from the Amazon Affiliate API

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
