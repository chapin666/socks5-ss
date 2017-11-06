package core

import (
	"math/rand"
	"time"
)

// PasswordLength const.
const PasswordLength = 256

// Password byte Array.
type Password [PasswordLength]byte

func init() {
	// 更新随机种子，防止生成一样的随机密码
	rand.Seed(time.Now().Unix())
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
