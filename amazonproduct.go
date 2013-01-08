//Package amazonproduct provides methods for interacting with the Amazon Product Advertising API
package amazonproduct

import ()

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
