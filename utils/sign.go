package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Sign(params map[string]interface{}, key, hashKey, hashIv string) (string, error) {
	// 1. 依照 ASCII 顺序由小到大做排序
	//  key1=value1&key2=value2...的方式组出字串，最后再加上&hash_key={密钥}&hash_iv={密钥}
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" && value != "" { // && value != ""
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}
	signStr := url.QueryEscape(sb.String())
	signStr += fmt.Sprintf("hash_key=%s&hash_iv=%s", hashKey, hashIv)
	signStr, err := url.QueryUnescape(signStr)
	if err != nil {
		fmt.Println("QueryUnescape error:", err)
		return "", err
	}

	fmt.Printf("[rawString]%s\n", signStr)

	// 第2步骤产生签名字串做 md5 加签得到sign
	// 第2步骤产生签名字串做 sha256 加签得到sign
	signResult := sha256HexBytes([]byte(signStr))

	signResult = strings.ToUpper(signResult)

	fmt.Printf("SHA256签名str: %s\n\n", signResult)
	return signResult, nil
}

// 入金&出金回调-成功-验签
func Verify(params map[string]interface{}, signKey, hashKey, hashIv string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := Sign(params, signKey, hashKey, hashIv)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}

func sha256HexBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
