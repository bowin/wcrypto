package eglass
import (
	b64 "encoding/base64"
	"math/rand"
	"encoding/binary"
	"bytes"
	"crypto/cipher"
	"crypto/aes"
)

// WechatCrypto 微信加密
type WechatCrypto struct {
	token string
	id []byte
	key []byte
	iv []byte
}

// Encrypt 加密
func (wc *WechatCrypto) Encrypt(text string) string {
	token := make([]byte, 16)
	rand.Read(token)
	b := make([]byte, 4) 
	binary.BigEndian.PutUint32(b, uint32(len(text)))
	msgBytes := bytes.Join([][]byte{token, b, []byte(text), wc.id}, []byte(""))
	// aes
	block, _ := aes.NewCipher(wc.key)
	B := cipher.NewCBCEncrypter(block, wc.iv)
	encoded := encode(msgBytes)
	encrypted := make([]byte, len(encoded))
	B.CryptBlocks(encrypted, encoded)
	return b64.StdEncoding.EncodeToString(encrypted)
}

// Decrypt 解密
func (wc *WechatCrypto) Decrypt(text string) string {
	block, error := aes.NewCipher(wc.key)
	if error != nil {
		panic(error)
	}
	B := cipher.NewCBCDecrypter(block, wc.iv)
	dst, error := b64.StdEncoding.DecodeString(text)
	// var s []byte
	s:= make([]byte, len(dst))
	B.CryptBlocks(s, dst)
	deciphered := decode(s)
	msg:= deciphered[16:]
	length := binary.BigEndian.Uint32(msg[0:4])
	return string(msg[4: 4 + length])
}

// New init
func New(token, key, appid string) *WechatCrypto {
	r, _ := b64.StdEncoding.DecodeString(key + "=")
	return &WechatCrypto{
		token: token,
		id: []byte(appid),
		key: []byte(r),
		iv: ([]byte(r))[0:16],
	}
}

func encode(text []byte) []byte {
	blockSize := 32
	textLength := len(text)
	amountToPad := blockSize - (textLength % blockSize)
	fillBytes := make([]byte, amountToPad)
	for i:=0; i < amountToPad; i++ {
		fillBytes[i] = byte(amountToPad)
	}
	return bytes.Join([][]byte{[]byte(text), fillBytes}, []byte(""))
}

func decode(text []byte) []byte {
	pad := int(text[len(text) - 1])
	if (pad < 1 || pad > 32) {
		pad = 0
	}
	return text[0: len(text) - pad]
}
