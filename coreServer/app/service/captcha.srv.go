package service

import "C"
import (
	"bytes"
	"context"
	"core-server/config"
	"core-server/pb"
	"errors"
	"github.com/LyricTian/captcha"
	"github.com/google/wire"
)

// CaptchaSet 注入Captcha
var CaptchaSet = wire.NewSet(wire.NewSet(new(Captcha)))

// Captcha 图形验证码
type Captcha struct {
	pb.UnimplementedCaptchaGreeterServer
}

// Get 获取CaptchaID
func (c *Captcha) Get(ctx context.Context, in *pb.CaptchaGetRequest) (*pb.CaptchaGetResult, error) {
	captchaID := captcha.NewLen(config.C.Captcha.Length)
	return &pb.CaptchaGetResult{CaptchaID: captchaID}, nil
}

// Generate 生成图形验证码
func (c *Captcha) Generate(ctx context.Context, in *pb.CaptchaGenerateRequest) (*pb.CaptchaGenerateResult, error) {
	captchaID := in.GetCaptchaID()
	if captchaID == "" {
		return nil, errors.New("请提供验证码ID")
	}

	if in.GetReload() {
		if !captcha.Reload(captchaID) {
			return nil, errors.New("未找到验证码ID")
		}
	}

	buffer := new(bytes.Buffer)
	err := captcha.WriteImage(buffer, captchaID, config.C.Captcha.Width, config.C.Captcha.Height)
	if err != nil {
		if err == captcha.ErrNotFound {
			return nil, errors.New("未找到指定资源")
		}
		return nil, err
	}
	return &pb.CaptchaGenerateResult{Buffer: buffer.Bytes()}, nil
}

// Verify 验证图形验证码
func (c *Captcha) Verify(ctx context.Context, in *pb.CaptchaVerifyRequest) (*pb.StatusResult, error) {
	captchaID := in.GetCaptchaID()
	captchaCode := in.GetCaptchaCode()
	if captchaID == "" || captchaCode == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("无效的验证码")
	}

	if !captcha.VerifyString(captchaID, captchaCode) {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("无效的验证码")
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}
