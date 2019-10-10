package code

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	codekey = "e5fcdg3hqa4b1nopij2rstuv67mwx89klyz"
)

// CreateInviterCode 根据用户id创建邀请码
func CreateInviterCode(userID string) string {
	code := ""
	bytekey := []byte(codekey)

	num, err := strconv.Atoi(userID)
	if err != nil {
		return ""
	}

	for num > 0 {
		mod := num % 35
		num = (num - mod) / 35
		code += string(bytekey[mod])
	}

	// 不足5位补上5位
	padnum := 5 - len(code)
	for padnum > 0 {
		code = "0" + code
		padnum--
	}

	return code
}

// DecodeInviterCode 根据邀请码解析出用户id
func DecodeInviterCode(code string) string {
	// 先把0去掉
	code = strings.Replace(code, "0", "", -1)
	len := len(code)
	num := 0
	for i := 0; i < len; i++ {
		index := strings.Index(codekey, string(code[i]))
		num += index * int(math.Pow(35, float64(i)))
	}

	return fmt.Sprintf("%d", num)
}
