package go_nepay_goldflow

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
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

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() NePayWithdrawReq {
	return NePayWithdrawReq{
		OrderId:     "23236326789",
		Money:       "1000",
		Currency:    "CNY",
		AccountNo:   "42623612",
		AccountName: "test",
		BankNo:      "TH0001",
	}
}
