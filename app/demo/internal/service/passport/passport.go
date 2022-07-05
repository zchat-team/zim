package passport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchat-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/util/env"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
	"github.com/zmicro-team/zmicro/core/log"
	zgin "github.com/zmicro-team/zmicro/core/transport/http"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/zchat-team/zim/app/demo/internal/constant"
	"github.com/zchat-team/zim/app/demo/internal/model"
	"github.com/zchat-team/zim/errno"
	"github.com/zchat-team/zim/pkg/auth"
	"github.com/zchat-team/zim/pkg/idgen"
	"github.com/zchat-team/zim/pkg/util"
	"github.com/zchat-team/zim/pkg/zcontext"
	api "github.com/zchat-team/zim/proto/http/demo/passport"
)

type IpLimit struct {
	IP            string
	LastErrorTime int64
	ErrorTimes    int64
}

type Service struct {
	zgin.Implemented
	auth *auth.Auth
	db   *gorm.DB
}

var (
	service *Service
	once    sync.Once
)

func GetService() *Service {
	once.Do(func() {
		service = &Service{}
		service.db = runtime.GetDB()

		privKey := config.GetString("auth.privKey")
		pubKey := config.GetString("auth.pubKey")
		service.auth = auth.NewAuth(privKey, pubKey, runtime.GetRedisClient())
	})
	return service
}

func (s *Service) AuthToken(token string) (acc *auth.Account, err error) {
	// 不考虑令牌吊销问题
	return s.auth.Inspect(token)
}
func (s *Service) Login(ctx context.Context, req *api.LoginReq, rsp *api.LoginRsp) (err error) {
	v := model.User{}
	if req.Type == constant.LoginTypeMobile {
		v.Mobile = req.Account
		//v.MobileVerified = 1
	} else if req.Type == constant.LoginTypeZid {
		v.Zid = req.Account
	} else if req.Type == constant.LoginTypeEmail {
		v.Email = req.Account
		//v.EmailVerified = 1
	}

	if err = s.db.Where(&v).Take(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.ErrUserNotExists()
		}
		return
	}

	if v.Password == "" {
		err = errno.ErrUserNoPasswd()
		return
	}

	var ipLimits []*IpLimit
	var curIpLimit *IpLimit
	if v.LoginIpLimit != "" {
		_ = json.Unmarshal([]byte(v.LoginIpLimit), &ipLimits)
	}

	// 清理过期记录
	i := 0
	for _, item := range ipLimits {
		if time.Now().Unix() < item.LastErrorTime+constant.RetryCD {
			// 未过冷却时间
			ipLimits[i] = item
			if item.IP == zcontext.GetClientIp(ctx) {
				curIpLimit = item
			}
			i++
		}
	}
	ipLimits = ipLimits[:i]

	if curIpLimit != nil && curIpLimit.ErrorTimes >= 6 {
		msg := fmt.Sprintf("密码错误次数过多，请%s再试", util.HumanTime(curIpLimit.LastErrorTime+constant.RetryCD))
		_ = msg
		err = errno.ErrCustom(msg)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(req.Password)); err != nil {
		err = errno.ErrCustom("用户或密码错误")

		// 更新错误次数
		if curIpLimit == nil {
			curIpLimit = &IpLimit{
				IP:            zcontext.GetClientIp(ctx),
				LastErrorTime: time.Now().Unix(),
				ErrorTimes:    0,
			}
			ipLimits = append(ipLimits, curIpLimit)
		}
		curIpLimit.LastErrorTime = time.Now().Unix()
		curIpLimit.ErrorTimes += 1
		if b, e := json.Marshal(&ipLimits); e == nil {
			v.LoginIpLimit = string(b)
			s.db.Model(&v).Updates(&v)
		}

		// UserLoginLog
		l := model.UserLoginLog{
			Id:       idgen.Next(),
			Uid:      v.Id,
			DeviceId: "", // req.DeviceId,
			Type:     1,  // 登录
			Status:   2,  // 失败
			LoginIp:  zcontext.GetClientIp(ctx),
		}
		s.db.Create(&l)

		return
	}

	token, err := s.auth.GenerateToken(cast.ToString(v.Id))
	if err != nil {
		return err
	}

	rsp.AccessToken = token.AccessToken
	rsp.RefreshToken = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expires = token.Expires.Unix()
	rsp.Uid = v.Id

	// 清空错误次数
	if curIpLimit != nil {
		curIpLimit.LastErrorTime = 0
		curIpLimit.ErrorTimes = 0
	}

	if b, e := json.Marshal(&ipLimits); e == nil {
		v.LoginIpLimit = string(b)
		v.LastLoginIp = zcontext.GetClientIp(ctx)
		v.LastLoginTime = time.Now()
		s.db.Model(&v).Updates(&v)
	}

	// UserLoginLog
	l := model.UserLoginLog{
		Id:       idgen.Next(),
		Uid:      v.Id,
		DeviceId: "", // req.DeviceId,
		Type:     1,
		Status:   1,
		LoginIp:  zcontext.GetClientIp(ctx),
	}
	s.db.Create(&l)
	return
}

func (s *Service) SmsLogin(ctx context.Context, req *api.SmsLoginReq, rsp *api.SmsLoginRsp) (err error) {
	var has bool
	v := model.User{Mobile: req.Mobile}
	err = s.db.Where(&v).First(&v).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	} else {
		has = true
	}

	clientIp := zcontext.GetClientIp(ctx)

	if env.IsProduct() || env.IsStaging() {

	}

	// 预发或者生产环境，需要验证短信
	//if env.IsDeployRelease() {
	//	if err = GetAliyunSms().VerifyCode(req.Mobile, req.Code); err != nil {
	//		// UserLoginLog
	//		l := model.UserLoginLog{
	//			Id:       idgen.Next(),
	//			UserId:   v.Id,
	//			DeviceId: "", // req.DeviceId,
	//			Type:     1,
	//			Status:   2,
	//			LoginIp:  clientIp,
	//		}
	//		s.db.Create(&l)
	//		return
	//	}
	//}

	var token *auth.Token
	var uid int64
	if !has {
		rsp.IsSignup = true
		if uid, err = s.register(ctx, req.Mobile); err != nil {
			return err
		}
	} else {
		uid = v.Id
	}

	if token, err = s.auth.GenerateToken(cast.ToString(uid)); err != nil {
		return err
	}

	// UserLoginLog
	l := model.UserLoginLog{
		Id:       idgen.Next(),
		Uid:      uid,
		DeviceId: "", // req.DeviceId,
		Type:     1,
		Status:   1,
		LoginIp:  clientIp,
	}
	s.db.Create(&l)

	v.LastLoginIp = clientIp
	v.LastLoginTime = time.Now()
	s.db.Model(&v).Where(&model.User{Id: uid}).Updates(&v)

	rsp.AccessToken = token.AccessToken
	rsp.RefreshToken = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expires = token.Expires.Unix()
	rsp.Uid = uid

	return
}

func (s *Service) register(ctx context.Context, mobile string) (uid int64, err error) {
	clientIp := zcontext.GetClientIp(ctx)

	// 执行本地事务
	err = s.db.Transaction(func(tx *gorm.DB) error {
		v := model.User{Mobile: mobile}
		if err := tx.Where(&v).First(&v).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			return errno.ErrCustom("用户已经存在，不能注册")
		}

		uid = idgen.Next()

		// user表

		u := model.User{
			Id:             uid,
			Mobile:         mobile,
			MobileVerified: 1,
			Zid:            "ZID_" + cast.ToString(uid),
			Nickname:       util.RandomNickname(),
			//Avatar:         util.RandAvatar(),
			LastLoginTime: time.Now(),
			LastLoginIp:   clientIp,
			RegisterIp:    clientIp,
		}
		if err := tx.Create(&u).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return
}

func (s *Service) Logout(ctx context.Context, req *api.LogoutReq, rsp *api.LogoutRsp) (err error) {
	// TODO
	return
}

func (s *Service) Sms(ctx context.Context, req *api.SmsReq, rsp *api.SmsRsp) (err error) {
	//if !env.IsDeployDebug() {
	//	err = GetAliyunSms().SendCode(req.Mobile)
	//}

	if env.IsProduct() || env.IsStaging() {

	}
	rsp.Code = "123456"

	return
}

func (s *Service) ChangePwd(ctx context.Context, req *api.ChangePwdReq, rsp *api.ChangePwdRsp) (err error) {
	uid := zcontext.GetUid(ctx)

	v := model.User{Id: uid}
	if err = s.db.First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.ErrUserNotExists()
		}
		return
	}

	req.OldPassword = strings.TrimSpace(req.OldPassword)
	req.NewPassword = strings.TrimSpace(req.NewPassword)

	if err = bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(req.OldPassword)); err != nil {
		err = errno.ErrPasswordIncorrect()
		return
	}

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return
	}
	v.Password = string(passwdHash)

	token, err := s.auth.GenerateToken(
		cast.ToString(uid),
	)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Updates(&v).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Error(err)
		return
	}

	rsp.AccessToken = token.AccessToken
	rsp.RefreshToken = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expires = token.Expires.Unix()

	return
}

func (s *Service) SetPwd(ctx context.Context, req *api.SetPwdReq, rsp *api.SetPwdRsp) (err error) {
	uid := zcontext.GetUid(ctx)

	v := model.User{Id: uid}
	if err = s.db.First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.ErrUserNotExists()
		}
		return
	}

	req.Password = strings.TrimSpace(req.Password)

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return
	}
	v.Password = string(passwdHash)

	token, err := s.auth.GenerateToken(
		cast.ToString(uid),
	)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&v).Updates(&v).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.AccessToken = token.AccessToken
	rsp.RefreshToken = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expires = token.Expires.Unix()
	return
}

func (s *Service) RefreshToken(ctx context.Context, req *api.RefreshTokenReq, rsp *api.RefreshTokenRsp) (err error) {
	token, err := s.auth.RefreshToken(cast.ToString(req.Uid), req.RefreshToken)

	if err != nil {
		log.Error(err)
		return err
	}

	rsp = &api.RefreshTokenRsp{}
	rsp.AccessToken = token.AccessToken
	rsp.RefreshToken = token.RefreshToken
	rsp.Created = token.Created.Unix()
	rsp.Expires = token.Expires.Unix()

	return
}
