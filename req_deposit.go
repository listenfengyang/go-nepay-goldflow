package go_nepay_goldflow

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-nepay-goldflow/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req NePayDepositReq) (*NePayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	params["scode"] = cli.Params.MerchantId
	params["redirectpage"] = "0"
	params["noticeurl"] = cli.Params.ReturnUrl
	params["callbackurl"] = cli.Params.DepositCallbackUrl

	//params转换map[string]interface{}格式
	paramsMap := map[string]interface{}{}
	mapstructure.Decode(params, &paramsMap)

	// Generate signature
	signStr, _ := utils.Sign(paramsMap, cli.Params.AccessKey, cli.Params.HashKey, cli.Params.HashIv)
	params["sign"] = signStr
	// params["sign"] = "b920c43e6f8411045152532fe29371ff" //signStr
	fmt.Println(params)
	var result NePayDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#nepay#goldflow#deposit->%s", string(restLog))

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
