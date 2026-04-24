package go_nepay_goldflow

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-nepay-goldflow/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) WithdrawReq(req NePayWithdrawReq) (*NePayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl
	// 2. Convert struct to map for signing
	var params map[string]string
	mapstructure.Decode(req, &params)

	params["scode"] = cli.Params.MerchantId
	params["paytype"] = "card_to_card" // 卡对卡支付
	params["notifyurl"] = cli.Params.MerchantInfo.WithdrawCallbackUrl

	//params转换map[string]interface{}格式
	paramsMap := map[string]interface{}{}
	mapstructure.Decode(params, &paramsMap)

	// Generate signature
	signStr, _ := utils.Sign(paramsMap, cli.Params.AccessKey, cli.Params.HashKey, cli.Params.HashIv)
	params["sign"] = signStr
	var result NePayWithdrawRsp
	fmt.Println(params)

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#nepay#goldflow#withdraw->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
