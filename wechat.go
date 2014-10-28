package wxutils

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"sort"
)

type WXAuth struct {
	Token string
}

// 用户发送的消息
type WXUserMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	PicUrl       string
	Location_X   float64
	Location_Y   float64
	Scale        float64
	Label        string
	Title        string
	Description  string
	Url          string
	MsgId        int64
}

// 回复消息
type WXReplyMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	FuncFlag     int
}

// http://mp.weixin.qq.com/wiki/index.php?title=%E6%B6%88%E6%81%AF%E6%8E%A5%E5%8F%A3%E6%8C%87%E5%8D%97
// 加密/校验流程：
//   1. 将token、timestamp、nonce三个参数进行字典序排序
//   2. 将三个参数字符串拼接成一个字符串进行sha1加密
//   3. 开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
// param:
//   s: signature
//   t: timestamp
//   n: nonce
func (wx WXAuth) CheckSignature(s, t, n string) (err error) {
	params := []string{t, n, wx.Token}
	sort.Sort(sort.StringSlice(params))

	res := string(sha1.New().Sum([]byte(params[0] + params[1] + params[2])))
	if res == s {
		return nil
	}

	return fmt.Errorf("signature not equal with sha1(timestamp+token+nonce): signature=%s sha1=%s!", s, res)
}

func DecodeWXMsg(b []byte) (msg WXUserMsg, err error) {
	err = xml.Unmarshal(b, &msg)
	if err != nil {
		return
	}

	return
}