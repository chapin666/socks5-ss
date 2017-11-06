package core

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

// PasswordLength const.
const PasswordLength = 256

// ErrInvalidPassword type.
var ErrInvalidPassword = errors.New("不合法的密码")

// Password byte Array.
type Password [PasswordLength]byte

func init() {
	// 更新随机种子，防止生成一样的随机密码
	rand.Seed(time.Now().Unix())
}

// 采用base64编码把密码装换为字符串
func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

// ParsePassword 方法解析采用 base64 编码的字符串获取密码
func ParsePassword(passwordString string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != PasswordLength {
		return nil, ErrInvalidPassword
	}
	password := Password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

// RandPassword method return Password
func RandPassword() *Password {
	// 随机生成一个由 0 ～ 255 组成的 byte 数组
	intArr := rand.Perm(PasswordLength)
	password := &Password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			// 索引与值相等，重新生成
			return RandPassword()
		}
	}

	return password
}
