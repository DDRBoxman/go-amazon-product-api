//Package amazonproduct provides methods for interacting with the Amazon Product Advertising API
package amazonproduct

import (
	"fmt"
	"strconv"
)

/*
ItemLookup takes a product ID (ASIN) and returns the result
*/
func (api AmazonProductAPI) ItemLookup(ItemId string) (string, error) {
	params := map[string]string{
		"ItemId":        ItemId,
		"ResponseGroup": "Images,ItemAttributes,Small,EditorialReview",
	}

	return api.genSignAndFetch("ItemLookup", params)
}

/*
ItemSearchByKeyword takes a string containg keywords and returns the search results
*/
func (api AmazonProductAPI) ItemSearchByKeyword(Keywords string) (string, error) {
	params := map[string]string{
		"Keywords":      Keywords,
		"ResponseGroup": "Images,ItemAttributes,Small,EditorialReview",
	}
	return api.ItemSearch("All", params)
}

func (api AmazonProductAPI) ItemSearchByKeywordWithResponseGroup(Keywords string, ResponseGroup string) (string, error) {
	params := map[string]string{
		"Keywords":      Keywords,
		"ResponseGroup": ResponseGroup,
	}
	return api.ItemSearch("All", params)
}

func (api AmazonProductAPI) ItemSearch(SearchIndex string, Parameters map[string]string) (string, error) {
	Parameters["SearchIndex"] = SearchIndex
	return api.genSignAndFetch("ItemSearch", Parameters)
}

/*
CartCreate takes a map containing ASINs and quantities. Up to 10 items are allowed
*/
func (api AmazonProductAPI) CartCreate(items map[string]int) (string, error) {

	params := make(map[string]string)

	i := 1
	for k, v := range items {
		if i < 11 {
			key := fmt.Sprintf("Item.%d.ASIN", i)
			params[key] = string(k)

			key = fmt.Sprintf("Item.%d.Quantity", i)
			params[key] = strconv.Itoa(v)

			i++
		} else {
			break
		}
	}

	return api.genSignAndFetch("CartCreate", params)
}

func (api AmazonProductAPI) CartClear(CartId, HAMC string) (string, error) {

	params := map[string]string{
		"CartId": CartId,
		"HMAC":   HMAC,
	}

	return api.genSignAndFetch("CartClear", params)
}
