package OxApi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

const (
	ServerURL = "https://gw2.mcpayment.net/api/v5/"
	AuthApi   = "auth/terminal"
)

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
	Header OxPayHead `json:"header"`
	Data   any       `json:"data"`
}

type TerminalLoginReq struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
type TerminalLoginResp struct {
	McpTerminalId     string `json:"mcpTerminalId"`
	SignKey           string `json:"signKey"`
	Brands            string `json:"brands"`
	WalletTypes       string `json:"walletTypes"`
	TripleAModuleData string `json:"tripleAModuleData"`
}

var client = http.Client{}

func TerminalLogin(userID, password string) error {
	req := OxPayReq{
		Header: OxPayHead{
			Version:    "5",
			AppType:    "W",
			AppVersion: "spoonxtest.0002.00012.1",
		},
		Data: TerminalLoginReq{
			userID,
			password,
		},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(data)
	request, err := http.NewRequest(http.MethodPost, ServerURL+AuthApi, reader)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("version", "5")
	//request.Header.Set("appType", "W")
	//request.Header.Set("appVersion", "spoonxtest.0002.00012.1")
	//for s, strings := range request.Header {
	//	fmt.Println(s, strings)
	//
	//}

	//err = setHead(request, OxPayHead{
	//	Version:    "5",
	//	AppType:    "W",
	//	AppVersion: "spoonxtest.0002.00012.1",
	//})
	//if err != nil {
	//	return err
	//}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	var resp TerminalLoginResp
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return err
	}
	return nil
}

func setHead(req *http.Request, head OxPayHead) error {
	types := reflect.TypeOf(head)
	values := reflect.ValueOf(head)

	if types.Kind() != reflect.Struct {
		return errors.New("the head not struct ")
	}
	for i := 0; i < types.NumField(); i++ {
		k := types.Field(i).Tag.Get("json")

		v := values.Field(i).String()
		if v != "" && values.Field(i).Kind() == reflect.String {
			fmt.Println(k + ":" + v)
			req.Header.Set(k, v)
		}
	}
	fmt.Println(req.Header)
	return nil

}
