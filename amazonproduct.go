package amazonproduct 

import (
	"net/url"
	"sort"
	"fmt"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
	"net/http"
	"io/ioutil"
)

type AmazonProductAPI struct {
	AccessKey string
	SecretKey string
	AssociateTag string
	Host string
}

func (api AmazonProductAPI) ItemSearchByKeyword(Keywords string) (string, error) {
	params := map[string] string {
		"Keywords": Keywords,
		"ResponseGroup" : "Images,ItemAttributes,Small,EditorialReview",
	}
	return api.ItemSearch("All", params)
}

func (api AmazonProductAPI) ItemSearchByKeywordWithResponseGroup(Keywords string, ResponseGroup string) (string, error) {	
	params := map[string] string {
		"Keywords": Keywords,
		"ResponseGroup" : ResponseGroup,
	}
	return api.ItemSearch("All", params)
}

func (api AmazonProductAPI) ItemSearch(SearchIndex string, Parameters map[string] string) (string,error){
	Parameters["SearchIndex"] = SearchIndex
	genUrl, err := GenerateAmazonUrl(api, "ItemSearch", Parameters)
	if (err != nil) {
		return "", err
	}

	SetTimestamp(genUrl)

	signedurl,err := SignAmazonUrl(genUrl, api)
	if (err != nil) {
		return "", err
	}

	resp, err := http.Get(signedurl)
	if (err != nil) {
		return "", err
	}	

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		return "", err
	}

	return string(body), nil
}

func GenerateAmazonUrl(api AmazonProductAPI, Operation string, Parameters map[string] string) (finalUrl *url.URL, err error) {

	result,err := url.Parse(api.Host)
	if (err != nil) {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "http"
	result.Path = "/onca/xml"

	values := url.Values{}
	values.Add("Operation", Operation)
	values.Add("Service", "AWSECommerceService")
	values.Add("AWSAccessKeyId", api.AccessKey)
	values.Add("Version", "2009-01-01")
	values.Add("AssociateTag", api.AssociateTag)

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if (err != nil) {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func SignAmazonUrl(origUrl *url.URL, api AmazonProductAPI) (signedUrl string , err error){

	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if (err != nil) {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}
