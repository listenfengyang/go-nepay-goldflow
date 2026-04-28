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

// { "memo": "", "sign": "A35168AB1C8FF9CEDFCACF49EADF68817991BE890314F96F1034A7B67704F8BC", "txid": "", "scode": "CPT", "amount": "2000", "status": 1, "orderid": "202612345678", "orderno": "PAY202604240844157956", "paytype": "card_to_card", "currency": "CNY", "respcode": "00", "resptime": "2026-04-24T08:58:37Z", "productname": "", "credit_amount": "2000" }
func GenCallbackRequestDemo() NePayDepositCallbackReq {
	return NePayDepositCallbackReq{
		Scode:        "CPT",
		OrderId:      "202612345678",
		OrderNo:      "PAY202604240844157956",
		PayType:      "card_to_card",
		Amount:       "2000",
		ProductName:  "",
		Currency:     "CNY",
		Memo:         "",
		RespTime:     "2026-04-24T08:58:37Z",
		Status:       1,
		RespCode:     "00",
		TxId:         "",
		CreditAmount: "2000",
		Sign:         "A35168AB1C8FF9CEDFCACF49EADF68817991BE890314F96F1034A7B67704F8BC",
	}
}
