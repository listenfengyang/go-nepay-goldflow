package go_nepay_goldflow

import (
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NePayInitParams{
		MerchantId:          MERCHANT_ID,
		HashKey:             HASH_KEY,
		HashIv:              HASH_IV,
		AccessKey:           ACCESS_KEY,
		DepositUrl:          DEPOSIT_URL,
		WithdrawUrl:         WITHDRAW_URL,
		DepositCallbackUrl:  NOTIFY_URL,
		WithdrawCallbackUrl: WITHDRAW_NOTIFY_URL,
		ReturnUrl:           RETURN_URL,
	})

	err := cli.WithdrawCallback(GenWdRequestDemo(), func(NePayWithdrawCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
}

func GenWdRequestDemo() NePayWithdrawCallbackReq {
	return NePayWithdrawCallbackReq{}
}
