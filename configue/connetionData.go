package configue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//커넥션 정보 가져오기
func GetConnectionData(connecter string) interface{} {

	var jsonData map[string]interface{}
	data, err := os.Open("./configue/connection.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(data)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &jsonData)

	if err != nil {
		fmt.Println(err)
	}
	return jsonData[connecter]
}
