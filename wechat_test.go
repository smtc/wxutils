package wxutils

import (
	"testing"
)

func TestWXAuth(t *testing.T) {
	// signature=0f70745f80fdb37664c51f862c51b221f6529616&echostr=4405949986800074302&timestamp=1414684216&nonce=966454691
	signature := "0f70745f80fdb37664c51f862c51b221f6529616"
	//echostr := "4405949986800074302"
	timestamp := "1414684216"
	nonce := "966454691"

	tk := "KHOpv0C0mFt9xfb3AlReUzGb"
	wa := CreateWXAuth(tk)
	err := wa.CheckSignature(signature, timestamp, nonce)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDecodeWXMsg(t *testing.T) {

}
