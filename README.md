Example
_______

	package main

	import (
		"fmt"
		"github.com/ddrboxman/go-amazon-product-api"
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

		products, err := amazonproduct.ParseItemResponse(result)
		if (err != nil) {
			fmt.Println(err)
		} else {
			fmt.Println(products)
		}
	}
