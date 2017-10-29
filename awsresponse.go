package amazonproduct

// Response describes the generic API Response
type AWSResponse struct {
	OperationRequest struct {
		RequestID             string     `xml:"RequestId"`
		Arguments             []Argument `xml:"Arguments>Argument"`
		RequestProcessingTime float64
	}
}

// Argument todo
type Argument struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

// Image todo
type Image struct {
	URL    string
	Height uint16
	Width  uint16
}

// Price describes the product price as
// Amount of cents in CurrencyCode
type Price struct {
	Amount         uint
	CurrencyCode   string
	FormattedPrice string
}

type TopSeller struct {
	ASIN  string
	Title string
}

// Item represents a product returned by the API
type Item struct {
	ASIN             string
	URL              string
	DetailPageURL    string
	ItemAttributes   *ItemAttributes
	OfferSummary     OfferSummary
	Offers           Offers
	SalesRank        int
	SmallImage       *Image
	MediumImage      *Image
	LargeImage       *Image
	ImageSets        *ImageSets
	EditorialReviews EditorialReviews
	BrowseNodes      struct {
		BrowseNode []BrowseNode
	}
}

// BrowseNode represents a browse node returned by API
type BrowseNode struct {
	BrowseNodeID string `xml:"BrowseNodeId"`
	Name         string
	TopSellers   struct {
		TopSeller []TopSeller
	}
	Ancestors struct {
		BrowseNode []BrowseNode
	}
}

// ItemAttributes response group
type ItemAttributes struct {
	Author          string
	Binding         string
	Brand           string
	Color           string
	EAN             string
	Creator         string
	Title           string
	ListPrice       Price
	Manufacturer    string
	Publisher       string
	NumberOfItems   int
	PackageQuantity int
	Feature         string
	Model           string
	ProductGroup    string
	ReleaseDate     string
	Studio          string
	Warranty        string
	Size            string
	UPC             string
}

// Offer response attribute
type Offer struct {
	Condition       string `xml:"OfferAttributes>Condition"`
	ID              string `xml:"OfferListing>OfferListingId"`
	Price           Price  `xml:"OfferListing>Price"`
	PercentageSaved uint   `xml:"OfferListing>PercentageSaved"`
	Availability    string `xml:"OfferListing>Availability"`
}

// Offers response group
type Offers struct {
	TotalOffers     int
	TotalOfferPages int
	MoreOffersURL   string  `xml:"MoreOffersUrl"`
	Offers          []Offer `xml:"Offer"`
}

// OfferSummary response group
type OfferSummary struct {
	LowestNewPrice   Price
	LowestUsedPrice  Price
	TotalNew         int
	TotalUsed        int
	TotalCollectible int
	TotalRefurbished int
}

// EditorialReview response attribute
type EditorialReview struct {
	Source  string
	Content string
}

// EditorialReviews response group
type EditorialReviews struct {
	EditorialReview EditorialReview
}

// BrowseNodeLookupRequest is the confirmation of a BrowseNodeInfo request
type BrowseNodeLookupRequest struct {
	BrowseNodeId  string
	ResponseGroup string
}

// ItemLookupRequest is the confirmation of a ItemLookup request
type ItemLookupRequest struct {
	IDType        string `xml:"IdType"`
	ItemID        string `xml:"ItemId"`
	ResponseGroup string `xml:"ResponseGroup"`
	VariationPage string
}

// ItemLookupResponse describes the API response for the ItemLookup operation
type ItemLookupResponse struct {
	AWSResponse
	Items struct {
		Request struct {
			IsValid           bool
			ItemLookupRequest ItemLookupRequest
		}
		Item Item `xml:"Item"`
	}
}

// ItemSearchRequest is the confirmation of a ItemSearch request
type ItemSearchRequest struct {
	Keywords      string `xml:"Keywords"`
	SearchIndex   string `xml:"SearchIndex"`
	ResponseGroup string `xml:"ResponseGroup"`
}

type ItemSearchResponse struct {
	AWSResponse
	Items struct {
		Request struct {
			IsValid           bool
			ItemSearchRequest ItemSearchRequest
		}
		Items                []Item `xml:"Item"`
		TotalResults         int
		TotalPages           int
		MoreSearchResultsUrl string
	}
}

type BrowseNodeLookupResponse struct {
	AWSResponse
	BrowseNodes struct {
		Request struct {
			IsValid                 bool
			BrowseNodeLookupRequest BrowseNodeLookupRequest
		}
		BrowseNode BrowseNode
	}
}

type ImageSets struct {
	ImageSet []ImageSet
}

type ImageSet struct {
	//Category string `xml:"Category,attr"`
	Category       string `xml:",attr"`
	SwatchImage    *Image
	SmallImage     *Image
	ThumbnailImage *Image
	TinyImage      *Image
	MediumImage    *Image
	LargeImage     *Image
}
