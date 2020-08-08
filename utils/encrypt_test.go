package utils

import "testing"

var str = "Base64编码要求把3个8位字节(3*8=24)转化为4个6位的字节(4*6=24),之后在6位的前面补两个0,形成8位一个字节的形式。"
var bstr = "QmFzZTY057yW56CB6KaB5rGC5oqKM+S4qjjkvY3lrZfoioIoMyo4PTI0Kei9rOWMluS4ujTkuKo25L2N55qE5a2X6IqCKDQqNj0yNCks5LmL5ZCO5ZyoNuS9jeeahOWJjemdouihpeS4pOS4qjAs5b2i5oiQOOS9jeS4gOS4quWtl+iKgueahOW9ouW8j+OAgg=="

func TestBase64Decode(t *testing.T) {

	str1 := Base64Encode(str)
	str2, err := Base64Decode(str1)

	t.Log(str1)
	t.Log(str2)
	t.Log(err)
}

func TestBase64Encode(t *testing.T) {

	str2, err := Base64Decode(bstr)
	str1 := Base64Encode(str2)

	t.Log(str1)
	t.Log(str2)
	t.Log(err)
}

func TestBase64EncodeBytes(t *testing.T) {
	bb := []byte(str)
	cc := Base64EncodeBytes(bb)

	t.Log(string(cc))
}

func TestBase64DecodeBytes(t *testing.T) {
	bb := []byte(bstr)
	cc, err := Base64DecodeBytes(bb)

	t.Log(string(cc))
	t.Log(err)
}

func TestMd5(t *testing.T) {
	t.Log(Md5(str))
}

func TestSha1(t *testing.T) {
	t.Log(Sha1(str))
}

func TestSha256(t *testing.T) {
	t.Log(Sha256(str))
}

func TestSha512(t *testing.T) {
	t.Log(Sha512(str))
}
