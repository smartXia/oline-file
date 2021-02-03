package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"log"
	"path"
	"strconv"
	"time"
)

func ByteEncoder(s interface{}) []byte {
	var enc_result bytes.Buffer
	enc := gob.NewEncoder(&enc_result)
	if err := enc.Encode(s); err != nil {
		log.Fatal("encode error:", err)
	}

	return enc_result.Bytes()
}
func GetIndex(array []interface{}, val interface{}) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func GetStringIndex(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func StringToUInt(str string) (val uint, err error) {
	valInt, err := strconv.Atoi(str)
	if valInt < 0 {
		valInt = 0
	}
	val = uint(valInt)
	return val, err
}
func StringToInt(str string) (val int, err error) {
	valInt, err := strconv.Atoi(str)
	if valInt < 0 {
		valInt = 0
	}
	val = int(valInt)
	return val, err
}

func StringToUInt8(str string) (val uint8, err error) {
	valInt, err := strconv.Atoi(str)
	if valInt < 0 {
		valInt = 0
	}
	val = uint8(valInt)
	return val, err
}

func StringToUInt16(str string) (val uint16, err error) {
	valInt, err := strconv.Atoi(str)
	if valInt < 0 {
		valInt = 0
	}
	val = uint16(valInt)
	return val, err
}

func StringToUInt32(str string) (val uint32, err error) {
	valInt, err := strconv.Atoi(str)
	if valInt < 0 {
		valInt = 0
	}
	val = uint32(valInt)
	return val, err
}

func FormatIntBoll(val string) string {
	if val != "1" {
		return "0"
	}
	return val
}

func StructToMap(i interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	j, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &m)
	return m, err
}

func GenerateFileKey(appkey string, channel uint, fileName string) string {
	if fileName == "" {
		return ""
	}
	m := md5.New()
	timeStampStr := strconv.Itoa(int(time.Now().Unix()))
	m.Write([]byte(fileName + timeStampStr))
	keyStr := hex.EncodeToString(m.Sum(nil))
	keyRune := []rune(keyStr)
	keyStr = string(keyRune[8:len(keyRune)-8]) + path.Ext(fileName)
	if appkey != "" {
		keyStr = appkey + strconv.Itoa(int(channel)) + "/" + keyStr
	}
	return keyStr
}
