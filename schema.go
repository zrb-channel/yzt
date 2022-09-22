package yzt

import (
	"encoding/json"
	"strings"
)

type (
	Config struct {
		AppId     string
		Channel   string
		PublicKey string
		AesIV     string
		AppSecret string
	}

	LoginRequest struct {
		RequestID           string `json:"requestId"`           // 唯一码
		SilentLoginFlag     string `json:"silentLoginFlag"`     // 静默登录标识符（0，非注册登录；1，已注册登录）
		ChannelCode         string `json:"channelCode"`         // 渠道号
		UserName            string `json:"userName"`            // 客户姓名
		IdCard              string `json:"certificateNum"`      // 客户身份证号
		UserPhone           string `json:"userPhone"`           //	客户手机号
		CompanyName         string `json:"companyName"`         //企业名称
		ProductCode         string `json:"productCode"`         //产品编号
		UMCode              string `json:"UMCode"`              //客户经理UM号
		MarketPersonnelCode string `json:"marketPersonnelCode"` //渠道营销人员代码
	}

	LoginResponseDataResult struct {
		OrderCode  string `json:"loanOrderCode"`
		ProductUrl string `json:"productUrl"`
	}

	LoginResponseData struct {
		Result  LoginResponseDataResult `json:"result"`
		Success bool                    `json:"success"`
	}
	LoginResponse struct {
		RequestId       string            `json:"requestId"`
		ResponseData    LoginResponseData `json:"responseData"`
		ResponseMessage string            `json:"responseMessage"`
		ResponseCode    string            `json:"responseCode"`
	}
)

func formatString(sourceStr string) string {
	if sourceStr == "" {
		return ""
	}
	return strings.ReplaceAll(strings.ReplaceAll(sourceStr, "\\r", ""), "\\n", "")
}

func (req *LoginRequest) String() string {
	v, _ := json.Marshal(req)
	return string(v)
}
