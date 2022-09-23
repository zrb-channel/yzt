package yzt

import (
	"encoding/json"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/aesutil"
	"github.com/zrb-channel/utils/hash"
	log "github.com/zrb-channel/utils/logger"
	"github.com/zrb-channel/utils/rsautil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BaseRequest struct {

	// 服务ID
	ServiceId string `json:"serviceId"`

	// 应用ID    文本文件内容获取（key: appId）
	AppId string `json:"appId"`

	// 请求ID    长度最长为64位，数字和字符串的组合，客户端请求唯一标识（建议用UUID）, 客户端生成，每次接口请求的requestId都是不同的
	RequestId string `json:"requestId"`

	// 时间戳    long毫秒数（当前时间）, 也用于加签
	Timestamp string `json:"timestamp"`

	// 渠道    应用所属渠道
	Channel string `json:"channel"`

	// 加签内容    视接口需要，若接口需要加签则必填，若接口不需要加签则不需要。
	Signture string `json:"signture"`

	// Base64编码后的加密的AES秘钥    视接口需要，若接口需要加密则必填，若接口不需要加密则不需要
	Ak string `json:"ak"`

	// 业务参数    AES秘钥加密的业务入参，业务参数为appId的值
	Message string `json:"message"`
}

// SetMessage
// @param message
// @date 2022-09-24 01:28:54
func (req *BaseRequest) SetMessage(message string) {
	req.Message = message
}

// NewRequest
// @param conf
// @param serviceID
// @param id
// @param msg
// @date 2022-09-24 01:28:52
func NewRequest(conf *Config, serviceID string, id string, msg interface{}) (*BaseRequest, error) {
	base := &BaseRequest{
		AppId:     conf.AppId,
		RequestId: id,
		Timestamp: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		Channel:   conf.Channel,
	}

	base.SetServiceID(serviceID)

	if err := base.Sign(conf, msg); err != nil {
		log.WithError(err).Error("消息签名失败")
		return nil, err
	}

	return base, nil
}

// Sign
// @param conf
// @param msg
// @date 2022-09-24 01:28:51
func (req *BaseRequest) Sign(conf *Config, msg interface{}) error {

	publicKey, err := utils.NewPublicKey(conf.PublicKey)
	if err != nil {
		return err
	}

	aesKey := utils.RandString(16)
	ak, err := rsautil.PublicEncryptToBase64(publicKey, []byte(aesKey))
	if err != nil {
		return err
	}

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	var encryptMessage string
	encryptMessage, err = aesutil.EncryptToBase64(message, []byte(aesKey), []byte(conf.AesIV))
	if err != nil {
		return err
	}

	req.SetAk(ak)

	sign := hash.SHA256String(encryptMessage + conf.AppSecret + req.Timestamp + req.ServiceId + conf.Channel)

	encryptMessage = url.QueryEscape(formatString(encryptMessage))
	req.SetMessage(encryptMessage)
	req.SetSignture(strings.ToUpper(sign))
	return nil
}

// SetAk
// @param v
// @date 2022-09-24 01:28:50
func (req *BaseRequest) SetAk(v string) {
	req.Ak = v
}

// SetServiceID
// @param id
// @date 2022-09-24 01:28:50
func (req *BaseRequest) SetServiceID(id string) {
	req.ServiceId = id
}

// SetSignture
// @param v
// @date 2022-09-24 01:28:49
func (req *BaseRequest) SetSignture(v string) {
	req.Signture = v
}

// String
// @date 2022-09-24 01:28:47
func (req *BaseRequest) String() string {
	v, _ := json.Marshal(req)
	return string(v)
}

type BaseResponse struct {
	// 返回信息码	成功：000000
	ResponseCode string `json:"responseCode"`

	// 返回信息内容	成功：000000
	ResponseMessage string `json:"responseMessage"`

	// 响应数据	具体见另一个excel文档
	ResponseData json.RawMessage `json:"responseData"`
}
