package util

import (
	"encoding/json"
	"log"
)

// 其他转json
func JsonEncode(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("【JsonEncode fail】：%s\n", err)
	}

	return b
}

// json 转其他
func JsonDecode(data []byte, v interface{}) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Printf("【JsonDecode fail】：%s\n", err)
	}
}
