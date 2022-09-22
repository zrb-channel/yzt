package yzt

import (
	"context"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/yzt/config"
	"net/http"

	"errors"

	log "github.com/zrb-channel/utils/logger"
)

// Login
// @param ctx
// @param conf
// @param id
// @param req
// @date 2022-09-22 18:32:55
func Login(ctx context.Context, conf *Config, id string, req *LoginRequest) (*LoginResponseDataResult, error) {

	body, err := NewRequest(conf, "1001100058", id, req)
	if err != nil {
		log.WithError(err).Error("登录失败")
		return nil, errors.New("操作失败")
	}

	resp, err := utils.Request(ctx).SetBody(body).SetHeader("ContentType", "application/json").Post(config.Addr)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("操作失败")
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Println(resp.StatusCode())
		return nil, errors.New("操作失败")
	}

	baseResp := &LoginResponse{}
	if err = json.Unmarshal(resp.Body(), baseResp); err != nil {
		return nil, errors.New("操作失败")
	}

	if baseResp.ResponseData.Success {
		return &baseResp.ResponseData.Result, nil
	}

	return nil, errors.New(baseResp.ResponseMessage)
}
