## 微信开发平台 message encrypt & decrpt 

### Install
```bash
go get github.com/bowin/wcrypto
``` 

### Usage
```go
 import "github.com/bowin/wcrypto"
 wcp := wcrypto.New("your token", "your aes key", "your open appid")
 // encrypt
 wcp.Encrypt("your messge")
 // decrypt
 wcp.Decrypt("messge from wechat")
```
