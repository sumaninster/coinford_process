package configs

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"time"
	"math/rand"
	)

var (
	SignBytes []byte
	NullTime	*time.Time
	NullString	*string

	BTCAuthUrl 	string
	BTC_IN_ACCOUNT string

	ETHAuthUrl 	string
	ETHPasswordExpiryHour int64
	err error
)

func Init() {
	SignBytes, err = ioutil.ReadFile("conf/app.rsa")
	fatal("SignBytes: ", err)

	NullTime = new(time.Time)
	NullString = new(string)

	BTCAuthUrl = "127.0.0.1:18332"
	BTC_IN_ACCOUNT = beego.AppConfig.String("BTC_IN_ACCOUNT")

	ETHAuthUrl = "http://localhost:8545"
	ETHPasswordExpiryHour = 1000000
}

func GetSha512(s string) string {
    h := sha512.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum([]byte(s)))
}

func RandString(str_size int) string {
    alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_-+=<>?:;"
    var bytes = make([]byte, str_size)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}

func Int64ToInterface(t []int64) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
	    s[i] = v
	}
	return s
}

func fatal(tag string, err error) {
	if err != nil {
		log.Fatal(tag, err)
	}
}
