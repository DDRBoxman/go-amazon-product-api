package amazonproduct

import (
	"encoding/xml"
)

type ItemResponse struct {
	XMLName xml.Name `xml:"ItemSearchResponse"`
	Products []Product `xml:"Items>Item"`
}

type Product struct {
    Id string `xml:"ASIN"`
    Name string `xml:"ItemAttributes>Title"`
    Cost string `xml:"ItemAttributes>ListPrice>FormattedPrice"`
    AffilateProgramId int
    Images string `xml:"LargeImage>URL"`
    Description string `xml:"EditorialReviews>EditorialReview>Content"`
    Rating int
}

func ParseItemResponse(Response string) ([]Product, error) {
	var q ItemResponse
	err := xml.Unmarshal([]byte(Response), &q)
	if (err != nil) {
		return nil, err
	}

	return q.Products, nil
}
