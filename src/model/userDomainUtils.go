package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
