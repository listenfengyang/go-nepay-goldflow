package go_nepay_goldflow

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
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

	req := GenCallbackRequestDemo()
	err := cli.DepositCallback(req, func(NePayDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", err)
}

func GenCallbackRequestDemo() NePayDepositCallbackReq {
	return NePayDepositCallbackReq{
		Scode:        "CPT",
		OrderId:      "202604281013170136",
		OrderNo:      "PAY202604280713171892",
		PayType:      "card_to_card",
		Amount:       "6844",
		ProductName:  "",
		Currency:     "CNY",
		Memo:         "",
		RespTime:     "2026-04-28T07:28:37Z",
		Status:       1,
		RespCode:     "00",
		TxId:         "",
		CreditAmount: "6844",
		Sign:         "3C3E6C4ECD07AB7585009B49A252A692FE4240CBFA7F2570BB321212C3B10CEA",
	}
}
