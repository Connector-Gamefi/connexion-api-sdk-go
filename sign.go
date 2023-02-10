package connexionapisdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func Sign(params map[string]interface{}, secret string) string {
	data := JoinParams(params)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func JoinParams(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedParams := make([]string, 0)
	for _, k := range keys {
		sortedParams = append(sortedParams, k+"="+fmt.Sprintf("%v", params[k]))
	}
	return strings.Join(sortedParams, "&")
}
