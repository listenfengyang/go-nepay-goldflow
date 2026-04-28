package go_nepay_goldflow

import (
	"testing"
)

func TestDeposit(t *testing.T) {

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
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() NePayDepositReq {
	return NePayDepositReq{
		Orderid:      "2026123456782",
		Amount:       "2000",
		Currency:     "CNY",
		Userid:       "1234",
		Accountname:  "简",
		Payeraccount: "346236236",
		Paytype:      "card_to_card", // 银联：card_to_card 支付宝：qr_pay
	}
}
