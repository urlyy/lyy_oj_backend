package util

import (
	"github.com/google/uuid"
)

func GenUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		// 处理错误
		return "", nil
	}
	return u.String(), nil
}
