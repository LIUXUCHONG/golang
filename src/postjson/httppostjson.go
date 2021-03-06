package postjson

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPostJson(getip string , gethostname string) {
	url := "https://"
	method := "POST"

	payload := strings.NewReader("{\"hostName\": \""+gethostname+"\",\"ip\": \""+getip+"\",\"env\":\"开发环境\",\"machineRoom\":\"DZ\",\"info\":\"\",\"label\":\"k8s\",\"user\":\"zhangsan\"}")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
