package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zmicro-team/zim/pkg/constant"
	"github.com/zmicro-team/zim/pkg/util"
	"github.com/zmicro-team/zmicro/core/log"
)

// 给登录设备增加TAG标记
// 同一帐号不同TAG当成独立设备处理
// 当设备冲突时，可以有两种策略，1：踢掉较早登录的设备（主动登录情况） 2：提示当前设备登录冲突（重连情况）

// 离线推送通知
// 推送服务的设备数据与用户UIN进行关联
// 云端自动将即时通讯消息转成特定的推送通知发送至客户端
// 根据UIN找到关联设备

// 离线消息同步
// 云端主动推送方案
// 客户端从云端拉取方案，用户登录上线后，计算出用户离线期间产生的未读消息的对话列表及对应的未读消息数，以未读消息更新事件通知到客户端

type DeviceInfo struct {
	DeviceId       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	Tag            string `json:"tag"`
	Platform       string `json:"platform"`
	Server         string `json:"server"`
	LoginTime      int64  `json:"login_time"`
	DisconnectTime int64  `json:"disconnect_time"`
	Status         int    `json:"status"` // 获取状态，请调用GetStatus()方法
}

func (d *DeviceInfo) GetRealStatus() int {
	status := d.Status
	if d.DisconnectTime != 0 && d.Status == constant.PushOnline {
		if time.Since(time.Unix(d.DisconnectTime, 0)) > time.Duration(constant.PushOnlineKeepDays*24)*time.Hour {
			d.Status = constant.Offline
		}
	}
	return status
}

func (s *Service) addConn(ctx context.Context, uin string, info *DeviceInfo) (err error) {
	b, err := json.Marshal(info)
	if err != nil {
		return
	}
	_, err = s.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		key := util.KeyOnline(uin, info.DeviceId)
		pipe.Set(ctx, key, string(b), time.Minute*30)
		key = util.KeyDevice(uin)
		pipe.HSet(ctx, key, info.DeviceId, string(b))
		return nil
	})
	return
}

func (s *Service) delConn(ctx context.Context, uin string, info *DeviceInfo) (err error) {
	b, err := json.Marshal(info)
	if err != nil {
		return
	}
	_, err = s.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, util.KeyOnline(uin, info.DeviceId))
		pipe.HSet(ctx, util.KeyDevice(uin), info.DeviceId, string(b))
		return nil
	})

	return
}

func (s *Service) getDevice(ctx context.Context, uin, deviceId string) *DeviceInfo {
	key := util.KeyDevice(uin)

	if b, err := s.client.HGet(ctx, key, deviceId).Bytes(); err != nil {
		return nil
	} else {
		info := &DeviceInfo{}
		if err := json.Unmarshal(b, info); err != nil {
			return nil
		}
		return info
	}
}

//func (s *Service) GetAllDevice(ctx context.Context, uin string) (devices map[string]*DeviceInfo, err error) {
//	devices = make(map[string]*DeviceInfo)
//	result, err := s.client.HGetAll(ctx, util.KeyDevice(uin)).Result()
//	if err != nil {
//		return
//	}
//	for k, v := range result {
//		info := DeviceInfo{}
//		if err := json.Unmarshal([]byte(v), &info); err != nil {
//			continue
//		}
//		devices[k] = &info
//	}
//	return
//}
//
//func (s *Service) GetOnlineDevice(ctx context.Context, uin, deviceId string) (device *DeviceInfo, err error) {
//	key := util.KeyOnline(uin, deviceId)
//	result, err := s.client.Get(ctx, key).Result()
//	if err != nil {
//		if err == redis.Nil {
//			err = nil
//		}
//		return
//	}
//
//	if result == "" {
//		return
//	}
//
//	device = &DeviceInfo{}
//	if err = json.Unmarshal([]byte(result), device); err != nil {
//		return nil, err
//	}
//
//	return
//}

func (s *Service) getOnline(ctx context.Context, uin string) (devices map[string][]*DeviceInfo, err error) {
	devices = make(map[string][]*DeviceInfo)
	keys, err := s.client.Keys(ctx, fmt.Sprintf("online:%s:*", uin)).Result()
	if err != nil {
		return
	}
	if len(keys) == 0 {
		return
	}
	log.Infof("online keys=%v", keys)
	result, err := s.client.MGet(ctx, keys...).Result()
	if err != nil {
		return
	}

	for _, v := range result {
		info := DeviceInfo{}
		if err := json.Unmarshal([]byte(v.(string)), &info); err != nil {
			continue
		}
		devices[info.Server] = append(devices[info.Server], &info)
	}

	return
}

func (s *Service) getOnlineOfTag(ctx context.Context, uin string, tag string) (devices []*DeviceInfo, err error) {
	devices = make([]*DeviceInfo, 0)
	keys, err := s.client.Keys(ctx, fmt.Sprintf("online:%s:*", uin)).Result()
	if err != nil {
		return
	}
	if len(keys) == 0 {
		return
	}

	log.Infof("online keys=%v", keys)
	result, err := s.client.MGet(ctx, keys...).Result()
	if err != nil {
		return
	}
	for _, v := range result {
		info := DeviceInfo{}
		if err := json.Unmarshal([]byte(v.(string)), &info); err != nil {
			continue
		}

		if info.Status != constant.Online {
			continue
		}
		if info.Tag == tag {
			devices = append(devices, &info)
			return
		}
	}

	if len(devices) > 1 {
		sort.Slice(devices, func(i, j int) bool { return devices[i].LoginTime < devices[j].LoginTime })
	}
	return
}
