package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

var (
	authInstance *AuthHelper
	authOnce     sync.Once
)

type AuthHelper struct {
	name   string
	secret string
}

func GetAuthInstance() *AuthHelper {
	authOnce.Do(func() {
		authInstance = &AuthHelper{}
		authInstance.name = "admin"
		authInstance.GeneratePassword()
	})
	return authInstance
}

// 生成秘钥
func (h *AuthHelper) GeneratePassword() {
	key := make([]byte, 24)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	h.secret = hex.EncodeToString(key)
}

// 获取秘钥
func (h *AuthHelper) GetSecret() string {
	return h.secret
}

func (h *AuthHelper) GetName() string {
	return h.name
}
