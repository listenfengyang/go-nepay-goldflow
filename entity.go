package go_nepay_goldflow

type NePayInitParams struct {
	MerchantId          string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`                                     // 商户代号
	HashKey             string `json:"hashKey" mapstructure:"hashKey" config:"hashKey"  yaml:"hashKey"`                                                 // hashKey
	HashIv              string `json:"hashIv" mapstructure:"hashIv" config:"hashIv"  yaml:"hashIv"`                                                     // hashIv
	AccessKey           string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`                                         // accessKey
	DepositUrl          string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`                                     // 入金地址
	WithdrawUrl         string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`                                 // 出金地址
	DepositCallbackUrl  string `json:"depositCallbackUrl" mapstructure:"depositCallbackUrl" config:"depositCallbackUrl"  yaml:"depositCallbackUrl"`     // 入金回调地址
	WithdrawCallbackUrl string `json:"withdrawCallbackUrl" mapstructure:"withdrawCallbackUrl" config:"withdrawCallbackUrl"  yaml:"withdrawCallbackUrl"` // 出金回调地址
	ReturnUrl           string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl"  yaml:"returnUrl"`                                         // 重定向地址
}

//============================================================

// nepay入金
type NePayDepositReq struct {
	Scode        string `json:"scode" mapstructure:"scode" config:"scode"  yaml:"scode"`                             // 商户代号
	Orderid      string `json:"orderid" mapstructure:"orderid" config:"orderid"  yaml:"orderid"`                     // 商户交易序号
	Paytype      string `json:"paytype" mapstructure:"paytype" config:"paytype"  yaml:"paytype"`                     // 支付方式
	Amount       string `json:"amount" mapstructure:"amount" config:"amount"  yaml:"amount"`                         // 支付金额
	Currency     string `json:"currency" mapstructure:"currency" config:"currency"  yaml:"currency"`                 // 支付币种
	Userid       string `json:"userid" mapstructure:"userid" config:"userid"  yaml:"userid"`                         // 用户ID
	Redirectpage string `json:"redirectpage" mapstructure:"redirectpage" config:"redirectpage"  yaml:"redirectpage"` // 同步回覆模式（固定值："0"）
	Accountname  string `json:"accountname" mapstructure:"accountname" config:"accountname"  yaml:"accountname"`     // 付款人姓名
	Payeraccount string `json:"payeraccount" mapstructure:"payeraccount" config:"payeraccount"  yaml:"payeraccount"` // 付款人账号

	// Productname string `json:"productname" mapstructure:"productname" config:"productname"  yaml:"productname"` // 商品名称
	// Memo        string `json:"memo" mapstructure:"memo" config:"memo"  yaml:"memo"`                             // 备注
	// Callbackurl string `json:"callbackurl" mapstructure:"callbackurl" config:"callbackurl"  yaml:"callbackurl"` // 重定向地址（部分通道支援完成订单后可以跳转至该重定向地址）
	// Noticeurl    string `json:"noticeurl" mapstructure:"noticeurl" config:"noticeurl"  yaml:"noticeurl"`             // 商家接收交易结果网址（接收异步交易结果）
	// Sign         string `json:"sign" mapstructure:"sign" config:"sign"  yaml:"sign"`                                 // 签名
}

type NePayDepositRsp struct {
	Status      interface{} `json:"status" mapstructure:"status"`           // 状态
	Respcode    string      `json:"respcode" mapstructure:"respcode"`       // 请求结果码
	Respmsg     string      `json:"respmsg" mapstructure:"respmsg"`         // 请求结果消息
	Scode       string      `json:"scode" mapstructure:"scode"`             // 商户代号
	Orderid     string      `json:"orderid" mapstructure:"orderid"`         // 商户交易序号
	Orderno     string      `json:"orderno" mapstructure:"orderno"`         // 系统交易序号
	Paytype     string      `json:"paytype" mapstructure:"paytype"`         // 支付方式
	Amount      string      `json:"amount" mapstructure:"amount"`           // 支付金额
	Productname string      `json:"productname" mapstructure:"productname"` // 商品名称
	Currency    string      `json:"currency" mapstructure:"currency"`       // 支付币种
	Memo        string      `json:"memo" mapstructure:"memo"`               // 备注
	Url         string      `json:"url" mapstructure:"url"`                 // 跳转地址
	// 其他字段...
	Bankname       string `json:"bankname" mapstructure:"bankname"`             // 收款银行名称
	Bankno         string `json:"bankno" mapstructure:"bankno"`                 // 收款银行账号
	Acctname       string `json:"acctname" mapstructure:"acctname"`             // 收款户名
	Qr_url         string `json:"qr_url" mapstructure:"qr_url"`                 // 收款二维码
	Floatingamount int32  `json:"floatingamount" mapstructure:"floatingamount"` // 下浮金额
}

// nepay出金
type NePayWithdrawReq struct {
	Scode    string `json:"scode" mapstructure:"scode"`       // 厂商編號
	OrderId  string `json:"orderid" mapstructure:"orderid"`   // 商户交易序号
	Money    string `json:"money" mapstructure:"money"`       // 支付金额
	Currency string `json:"currency" mapstructure:"currency"` // 支付币种
	// BankName    string `json:"bankname" mapstructure:"bankname"`       // 收款银行名称
	BankNo      string `json:"bankno" mapstructure:"bankno"`           // 收款银行账号
	AccountNo   string `json:"accountno" mapstructure:"accountno"`     // 付款人账号
	AccountName string `json:"accountname" mapstructure:"accountname"` // 付款人姓名
	PayType     string `json:"paytype" mapstructure:"paytype"`         // 支付方式
	// NotifyUrl   string `json:"notifyurl" mapstructure:"notifyurl"`     // 通知步地址
	// Sign        string `json:"sign" mapstructure:"sign"`               // 驗證碼
}

type NePayWithdrawRsp struct {
	Prc     interface{} `json:"prc" mapstructure:"prc"`         // 状态
	Errcode string      `json:"errcode" mapstructure:"errcode"` // 请求结果码
	Msg     string      `json:"msg" mapstructure:"msg"`         // 请求结果消息
	Orderno string      `json:"orderno" mapstructure:"orderno"` // 系统交易序号
}

// 入金回调
type NePayDepositCallbackReq struct {
	Scode        string `json:"scode" mapstructure:"scode"`                 // 厂商編號
	OrderId      string `json:"orderid" mapstructure:"orderid"`             // 商家交易序號
	OrderNo      string `json:"orderno" mapstructure:"orderno"`             // 本系統交易序號（具唯一性）
	PayType      string `json:"paytype" mapstructure:"paytype"`             // 支付方式
	Amount       string `json:"amount" mapstructure:"amount"`               // 支付金額（格式：00.00）
	ProductName  string `json:"productname" mapstructure:"productname"`     // 商品名稱
	Currency     string `json:"currency" mapstructure:"currency"`           // 支付幣別
	Memo         string `json:"memo" mapstructure:"memo"`                   // 備註（回傳原廠商傳入的memo欄位值）
	RespTime     string `json:"resptime" mapstructure:"resptime"`           // 交易完成時間（UTC 時間，ISO 8601 格式：yyyy-MM-ddTHH:mm:ssZ，例：2025-03-01T12:05:30Z）
	Status       int32  `json:"status" mapstructure:"status"`               // 交易結果：1=交易成功，-1=交易失敗
	RespCode     string `json:"respcode" mapstructure:"respcode"`           // 交易訊息代碼（請參閱錯誤代碼說明）
	TxId         string `json:"txid" mapstructure:"txid"`                   // hash
	CreditAmount string `json:"credit_amount" mapstructure:"credit_amount"` // 實際付款金額（如有回傳此欄位，手續費以此欄位計算）
	Sign         string `json:"sign" mapstructure:"sign"`                   // 驗證碼（詳見簽名規則章節）
}

// 出金回调
type NePayWithdrawCallbackReq struct {
	Scode    string `json:"scode" mapstructure:"scode"`       // 厂商編號
	OrderId  string `json:"orderid" mapstructure:"orderid"`   // 商家交易序號
	OrderNo  string `json:"orderno" mapstructure:"orderno"`   // 本系統交易序號（具唯一性）
	Money    string `json:"money" mapstructure:"money"`       // 代付金額（以元為單位，支援整數或最多兩位小數）
	Status   string `json:"status" mapstructure:"status"`     // 代付結果：S=代付成功，F=代付失敗（交易結果以此欄位的狀態為主）
	RespCode string `json:"respcode" mapstructure:"respcode"` // 交易訊息代碼（請參閱錯誤代碼說明）
	ResTime  string `json:"resptime" mapstructure:"resptime"` // 出金完成時間（UTC 時間，ISO 8601 格式：yyyy-MM-ddTHH:mm:ssZ；處理成功則回傳，未成功則為空）
	Sign     string `json:"sign" mapstructure:"sign"`         // 驗證碼（詳見簽名規則章節）
}
