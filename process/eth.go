package process

import (
	//"coinford_process/models"
	"coinford_process/configs"
	"fmt"
	//"time"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	//"strconv"
	//"errors"
)

type ETHProcess struct {
	currencyid 	int64
}

type ETHNewAccount struct {
	Jsonrpc 	string
	Id 			int64
	Result		string
}

func (ethp *ETHProcess) GetNewAddress(passphrase string) (string, error) {
	jsonstr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"personal_newAccount","params":["%s"],"id":1}`, passphrase)
	body, err := ethp.query(jsonstr)
	var rqd ETHNewAccount
	json.Unmarshal(body, &rqd)
	return rqd.Result, err
}

func (ethp *ETHProcess) query(jsonstr string) ([]byte, error) {
	fmt.Println("Request Body:", string(jsonstr))
	jsonbytes := []byte(jsonstr)
	resp, _ := http.Post(configs.ETHAuthUrl, "application/json", bytes.NewBuffer(jsonbytes))
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
    return body, nil
}