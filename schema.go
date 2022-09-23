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

		// 唯一码
		RequestID string `json:"requestId"`

		// 静默登录标识符（0，非注册登录；1，已注册登录）
		SilentLoginFlag string `json:"silentLoginFlag"`

		// 渠道号
		ChannelCode string `json:"channelCode"`

		// 客户姓名
		UserName string `json:"userName"`

		// 客户身份证号
		IdCard string `json:"certificateNum"`

		//	客户手机号
		UserPhone string `json:"userPhone"`

		//企业名称
		CompanyName string `json:"companyName"`

		//产品编号
		ProductCode string `json:"productCode"`

		//客户经理UM号
		UMCode string `json:"UMCode"`

		//渠道营销人员代码
		MarketPersonnelCode string `json:"marketPersonnelCode"`
	}

	LoginResponseDataResult struct {
		OrderCode string `json:"loanOrderCode"`

		ProductUrl string `json:"productUrl"`
	}

	LoginResponseData struct {
		Result LoginResponseDataResult `json:"result"`

		Success bool `json:"success"`
	}
	LoginResponse struct {
		RequestId string `json:"requestId"`

		ResponseData LoginResponseData `json:"responseData"`

		ResponseMessage string `json:"responseMessage"`

		ResponseCode string `json:"responseCode"`
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
