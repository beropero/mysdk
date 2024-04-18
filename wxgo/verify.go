package wxgo

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"
	"sort"
)


// 签名验证
func (w *Wechat) VerifySignature(vp VerifyParams) (res bool, err error) {
	// 在配置中获取 token 字段
	token, err := w.Cfg.GetToken()
	if err != nil {
		return false, err
	}
	// 构造匹配字段
	strs := []string{vp.Timestamp, vp.Nonce, token}
	// 按字典序排列后拼接成一个字符串
	sort.Strings(strs)
	str := strings.Join(strs, "")
	// 对拼接后的字符串进行 SHA1 加密
	hash := sha1.New()
	hash.Write([]byte(str))
	hashed := fmt.Sprintf("%x", hash.Sum(nil))
	// 加密结果与 signature 比较
	if hashed != vp.Signature {
		return false, errors.New("error: Signature mismatch")
	}
	return true, nil
}


