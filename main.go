package main

type OxPayHead struct {
	Version       string `json:"version"`       //OxPayAPI version
	AppType       string `json:"appType"`       //app类型 I: iPhone; A: Android; W: Web
	AppVersion    string `json:"appVersion"`    //For example: "AppName.01.20.0" or "WebName.0002.00012.1"
	Status        Status `json:"status"`        //json字符串 只存在于response 的head中
	McpTerminalId string `json:"mcpTerminalId"` // 登录后相应体中获取Mcp Terminal Id mandatory for sale, void, refund, reverse...
	Signature     string `json:"signature"`     //签名 登录后获取 ，并且在每次发送请求时都需要带上
	Uuid          string `json:"uuid"`          // 在pos机上需要设置
}
type Status struct {
	ResponseCode string `json:"responseCode"`
	Message      string `json:"message"`
}

type OxPayReq struct {
	Currency             string `json:"currency"`
	CardNo               string `json:"cardno"`
	CardExpiryDate       string `json:"cardExpiryDate"`
	CardHolderName       string `json:"cardHolderName"`
	Threedsecure         string `json:"threedsecure"` //0或1 是否启用线程安全
	TotalAmount          string `json:"totalAmount"`
	SalesAmount          string `json:"salesAmount"`
	Cav                  string `json:"cav"`
	Eci                  string `json:"eci"`                  //电子商务指示 VISA - 05 or 06  MASTER - 01 or 02 JCB - 05  AMEX - 05 or 06
	Cvv                  string `json:"cvv"`                  // 银行卡cvc Visa - CVV2; Master - CVC; Amex - CID (Reference Document)cav
	ReferenceNo          string `json:"referenceNo"`          //我们这边的订单号
	CustomerEmailAddress string `json:"customerEmailAddress"` //用户邮箱 用于收收据
	CustomerName         string `json:"customerName"`         //用户姓名 用于收据开头
	ServiceAmount        string `json:"serviceAmount"`        // 通过支付网关增加销售金额的价值,不带十进制的服务金额。要求1.23新币等于123新币。日元12000将是12000。
	ServiceRate          string `json:"serviceRate"`          // 服务费加在销售额之上
	GstAmount            string `json:"gstAmount"`            //在销售金额上添加的实际消费税金额不带小数点的总额。要求1.23新币等于123新币。日元12000将是12000。
	GstRate              string `json:"gstRate"`              // 消费税加在销售额之上
	TipAmount            string `json:"tipAmount"`            // 提示数量
	BrandName            string `json:"brandName"`            // 发卡行名称
	MasterpassTxnId      string `json:"masterpassTxnId"`      // Masterpass标准C/O, Express C/O(商家托管)交易标识符
	Tokenize             string `json:"tokenize"`             // Y or No
	Cardtoken            string `json:"cardtoken"`            // 与tokenize成对出现,卡片相关信息Token注:如果请求有cardToken，则不需要提供卡片信息
	MasterpassParingId   string `json:"masterpassParingId"`   //Masterpass Express C/O(网关托管)交易标识符仅为快速结帐请求返回  PairingIds是一次性使用的。
	Status               string `json:"status"`               // 3DS验证(tx_status= Y/N/U) Y——是的, N -不, U -不可查
	Is3DSEnabled         string `json:"is3DSEnabled"`         // 是否能3d验证取值为true(1)/false(0)。
	AdditionalData       string `json:"additionalData"`       // json字符串  "{\"otpRedirectUrl\":\"http://example.com/api/returnUrl\"}"

}

const ()

func main() {

}
