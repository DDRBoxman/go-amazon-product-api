package amazonproduct 

import (
	"net/url"
	"sort"
	"fmt"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type AmazonProductAPI struct {
	AccessKey string
	SecretKey string
	AssociateTag string
}

func (api AmazonProductAPI) ItemSearchByKeyword(Keywords string) {
	params := map[string] string {
		"Keywords": Keywords,
	}
	api.ItemSearch("All", params)
}

func (api AmazonProductAPI) ItemSearch(SearchIndex string, Parameters map[string] string) {

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
