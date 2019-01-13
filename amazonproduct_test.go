package amazonproduct

import (
	"net/url"
	"testing"
)

func Test_ItemSearchByKeyword_1(t *testing.T) {

}

func Test_SignAmazonUrl_1(t *testing.T) {

	urlString := `http://webservices.amazon.com/onca/xml?Service=AWSECommerceService&AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&Operation=ItemLookup&ItemId=0679722769&ResponseGroup=ItemAttributes,Offers,Images,Reviews&Version=2009-01-06&Timestamp=2009-01-01T12:00:00Z`

	signedUrl := `http://webservices.amazon.com/onca/xml?AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=ItemAttributes%2COffers%2CImages%2CReviews&Service=AWSECommerceService&Timestamp=2009-01-01T12%3A00%3A00Z&Version=2009-01-06&Signature=M%2Fy0%2BEAFFGaUAp4bWv%2FWEuXYah99pVsxvqtAuC8YN7I%3D`

	var api AmazonProductAPI

	api.SecretKey = "1234567890"

	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Could not parse urlstring")
	}

	resultUrl, err := SignAmazonUrl(url, api)
	if err != nil {
		t.Errorf("Signing failure: %v", err)
	}

	if signedUrl != resultUrl {
		t.Error("Signed url does not match")
	}
}

func Test_GenerateAmazonUrl_ItemSearch(t *testing.T) {

	urlString := `http://ecs.amazonaws.co.uk/onca/xml?Service=AWSECommerceService&AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&Operation=ItemSearch&Keywords=sgt+frog&Version=2013-08-01&SearchIndex=All&AssociateTag=mytag-20`

	var api AmazonProductAPI
	api.Host = "ecs.amazonaws.co.uk"
	api.AccessKey = "AKIAIOSFODNN7EXAMPLE"
	api.AssociateTag = "mytag-20"
	api.SecretKey = "1234567890"

	params := map[string]string{
		"Keywords":    "sgt frog",
		"SearchIndex": "All",
	}

	genurl, err := GenerateAmazonUrl(api, "ItemSearch", params)
	if err != nil {
		t.Error(err)
	}

	if genurl != nil {
		resultUrl, err := SignAmazonUrl(genurl, api)
		if err != nil {
			t.Errorf("Signing failure: %v", err)
		}

		parsedurl, err := url.Parse(urlString)
		if err != nil {
			t.Error("Could not parse urlstring")
		}

		urlString, err = SignAmazonUrl(parsedurl, api)
		if err != nil {
			t.Errorf("Signing failure! %v", err)
		}

		t.Log(urlString)
		t.Log(resultUrl)

		if urlString != resultUrl {
			t.Error("Signed url does not match")
		}
	} else {
		t.Error("Returned url was null")
	}
}
