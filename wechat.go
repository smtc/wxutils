package wxutils

import (
	"crypto/sha1"
	//"encoding/base64"
	"encoding/xml"
	"fmt"
	//"io"
	"sort"
)

type WXAuth struct {
	token string
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

type WXArticle struct {
	Title       string
	Description string
	PicUrl      string
	Url         string
}

// 回复消息
type WXReplyMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MusicUrl     string
	HQMusicUrl   string
	ArticleCount int
	Articles     []WXArticle `xml:Articles>item`
	FuncFlag     int
}

func CreateWXAuth(token string) *WXAuth {
	return &WXAuth{token: token}
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
func (wx *WXAuth) CheckSignature(s, t, n string) (err error) {
	params := []string{t, n, wx.token}
	sort.Sort(sort.StringSlice(params))
	param := params[0] + params[1] + params[2]

	sha := sha1.New()
	sha.Write([]byte(param))
	res := fmt.Sprintf("%x", string(sha.Sum(nil)[0:]))

	if res == s {
		return nil
	}

	return fmt.Errorf("signature not equal with sha1(timestamp+token+nonce): signature=%s sha1=%s!", s, res)
}

func DecodeWXUserMsg(b []byte) (msg WXUserMsg, err error) {
	err = xml.Unmarshal(b, &msg)
	if err != nil {
		return
	}

	return
}
