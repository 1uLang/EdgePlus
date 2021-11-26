// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	teaconst "github.com/TeaOSLab/EdgePlus/pkg/const"
	"github.com/TeaOSLab/EdgePlus/pkg/encrypt"
	"github.com/iwind/TeaGo/maps"
	"time"
)

// Encode 加密
func Encode(data []byte) (string, error) {
	instance, err := encrypt.NewMethodInstance("aes-256-cfb", teaconst.PlusKey, teaconst.PlusIV)
	if err != nil {
		return "", errors.New("不支持选择的加密方式")
	}
	dist, err := instance.Encrypt(data)
	if err != nil {
		return "", errors.New("加密失败：" + err.Error())
	}
	return base64.StdEncoding.EncodeToString(dist), nil
}

// EncodeMap 加密Map
func EncodeMap(m maps.Map) (string, error) {
	m["updatedAt"] = time.Now().Unix() // 用来校验Authority服务是否已经更新

	data, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return Encode(data)
}

// Decode 解密
func Decode(data []byte) (maps.Map, error) {
	instance, err := encrypt.NewMethodInstance("aes-256-cfb", teaconst.PlusKey, teaconst.PlusIV)
	if err != nil {
		return nil, errors.New("encrypt method not supported")
	}
	source, err := base64.StdEncoding.DecodeString(string(bytes.TrimSpace(data)))
	if err != nil {
		return nil, errors.New("decode key failed: base64 decode failed: " + err.Error())
	}
	dist, err := instance.Decrypt(source)
	if err != nil {
		return nil, errors.New("decode key failed: decrypt failed: " + err.Error())
	}
	m := maps.Map{}
	err = json.Unmarshal(dist, &m)
	if err != nil {
		return nil, errors.New("decode key failed: decode json failed: " + err.Error())
	}
	return m, nil
}

// DecodeKey 解密Key
func DecodeKey(data []byte) (*Key, error) {
	instance, err := encrypt.NewMethodInstance("aes-256-cfb", teaconst.PlusKey, teaconst.PlusIV)
	if err != nil {
		return nil, errors.New("encrypt method not supported")
	}
	source, err := base64.StdEncoding.DecodeString(string(bytes.TrimSpace(data)))
	if err != nil {
		return nil, errors.New("decode key failed: base64 decode failed: " + err.Error())
	}
	dist, err := instance.Decrypt(source)
	if err != nil {
		return nil, errors.New("decode key failed: decrypt failed: " + err.Error())
	}

	var result = &Key{}
	err = json.Unmarshal(dist, result)
	if err != nil {
		return nil, errors.New("decode key failed: " + err.Error())
	}

	return result, nil
}
