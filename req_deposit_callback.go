package go_nepay_goldflow

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-nepay-goldflow/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值-成功回调
func (cli *Client) DepositCallback(req NePayDepositCallbackReq, processor func(NePayDepositCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	flag, err := utils.Verify(params, cli.Params.AccessKey, cli.Params.HashKey, cli.Params.HashIv)
	if !flag || err != nil {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nowPay deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
