//Package amazonproduct provides methods for interacting with the Amazon Product Advertising API
package amazonproduct

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
ItemLookupWithResponseGroup takes a product ID (ASIN) and a ResponseGroup and returns the result
*/
func (api AmazonProductAPI) ItemLookupWithResponseGroup(ItemId string, ResponseGroup string) (string, error) {
	params := map[string]string{
		"ItemId":        ItemId,
		"ResponseGroup": ResponseGroup,
	}

	return api.genSignAndFetch("ItemLookup", params)
}

/*
ItemLookupWithParams takes the params for ItemLookup and returns the result
*/
func (api AmazonProductAPI) ItemLookupWithParams(params map[string]string) (string, error) {
	_, present := params["ItemId"]
	if !present {
		return "", errors.New("ItemId property is required in the params map")
	}

	return api.genSignAndFetch("ItemLookup", params)
}

/*
MultipleItemLookup takes an array of product IDs (ASIN) and returns the result
*/
func (api AmazonProductAPI) MultipleItemLookup(ItemIds []string) (string, error) {
	params := map[string]string{
		"ItemId":        strings.Join(ItemIds, ","),
		"ResponseGroup": "Images,ItemAttributes,Small,EditorialReview",
	}

	return api.genSignAndFetch("ItemLookup", params)
}

/*
MultipleItemLookupWithResponseGroup takes an array of product IDs (ASIN) as well as a ResponseGroup and returns the result
*/
func (api AmazonProductAPI) MultipleItemLookupWithResponseGroup(ItemIds []string, ResponseGroup string) (string, error) {
	params := map[string]string{
		"ItemId":        strings.Join(ItemIds, ","),
		"ResponseGroup": ResponseGroup,
	}

	return api.genSignAndFetch("ItemLookup", params)
}

/*
ItemSearchByKeyword takes a string containing keywords and returns the search results
*/
func (api AmazonProductAPI) ItemSearchByKeyword(Keywords string, page int) (string, error) {
	params := map[string]string{
		"Keywords":      Keywords,
		"ResponseGroup": "Images,ItemAttributes,Small,EditorialReview",
		"ItemPage":      strconv.FormatInt(int64(page), 10),
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

/*
CartAdd takes a map containing ASINs and quantities and adds them to the given cart.
Up to 10 items are allowed
*/
func (api AmazonProductAPI) CartAdd(items map[string]int, cartid, HMAC string) (string, error) {

	params := map[string]string{
		"CartId": cartid,
		"HMAC":   HMAC,
	}

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
	return api.genSignAndFetch("CartAdd", params)
}

/*
CartClear takes a CartId and HMAC that were returned when generating a cart
It then removes the contents of the cart
*/
func (api AmazonProductAPI) CartClear(CartId, HMAC string) (string, error) {

	params := map[string]string{
		"CartId": CartId,
		"HMAC":   HMAC,
	}

	return api.genSignAndFetch("CartClear", params)
}

/*
Cart get takes a CartID and HMAC that were returned when generating a cart
Returns the contents of the specified cart
*/
func (api AmazonProductAPI) CartGet(CartId, HMAC string) (string, error) {

	params := map[string]string{
		"CartId": CartId,
		"HMAC":   HMAC,
	}

	return api.genSignAndFetch("CartGet", params)
}

/*
BrowseNodeLookup takes a BrowseNodeId and returns the result.
*/
func (api AmazonProductAPI) BrowseNodeLookup(nodeId string) (string, error) {
	params := map[string]string{
		"BrowseNodeId": nodeId,
	}
	return api.genSignAndFetch("BrowseNodeLookup", params)
}

func (api AmazonProductAPI) BrowseNodeLookupWithResponseGroup(nodeId string, responseGroup string) (string, error) {
	params := map[string]string{
		"BrowseNodeId":  nodeId,
		"ResponseGroup": responseGroup,
	}
	return api.genSignAndFetch("BrowseNodeLookup", params)
}
