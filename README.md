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
