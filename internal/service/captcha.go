package service

import (
	"shippo-server/utils"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

type CaptchaService struct {
	*Service
}

func NewCaptchaService(s *Service) *CaptchaService {
	return &CaptchaService{s}
}

func (s *CaptchaService) CaptchaSmsSend(phone string, token string) (err error) {

	if !check.CheckPhone(phone) {
		err = ecode.ServerErr
		return
	}

	// 过期所有验证码
	err = s.dao.Captcha.CaptchaDel(phone)
	if err != nil {
		return
	}
	// 生成新的验证码
	r, err := s.dao.Captcha.CaptchaSmsInsert(phone, token)
	if err != nil {
		return
	}

	// 发送验证码
	utils.SendSms(r.Target, r.Code)
	return
}

func (s *CaptchaService) CaptchaEmailSend(email string, token string) (err error) {

	if !check.CheckQQEmail(email) {
		err = ecode.ServerErr
		return
	}

	// 过期所有验证码
	err = s.dao.Captcha.CaptchaDel(email)
	if err != nil {
		return
	}
	// 生成新的验证码
	r, err := s.dao.Captcha.CaptchaEmailInsert(email, token)
	if err != nil {
		return
	}

	// 发送验证码
	utils.SendEmail(r.Target, r.Code)
	return
}
