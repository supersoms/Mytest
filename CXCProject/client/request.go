package client

import (
	"errors"
	"log"
	"net/http"
)

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go GET URL : %s \n", req.URL.String())
	return client.Do(req)
}

// example: http://host:port/uri/?param1=1&param2=2
/*func Get(reqUrl string, jsonParams interface{}, toWho string) ([]byte, error) {
	var params url.Values = url.Values{}
	var jsonObj map[string]interface{}
	jsonB := StringifyJsonToBytes(jsonParams)
	//logger.L.Info("["+toWho+"]请求参数:"+string(jsonB), "url", reqUrl)
	if err := ParseJsonFromBytes(jsonB, &jsonObj); err != nil {
		logger.L.Trace(err)
		return nil, err
	}
	for k, v := range jsonObj {
		params.Set(k, fmt.Sprintf("%v", v))
	}
	logger.L.Debug("["+toWho+"]get请求url:", "url", reqUrl+"?"+params.Encode())
	req, rErr := http.NewRequest("GET", reqUrl+"?"+params.Encode(), nil)
	if rErr != nil {
		return nil, rErr
	}
	//req.Header.Set("sign", sig)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.L.Error("["+toWho+"]请求不成功:", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		logger.L.Error("["+toWho+"]读取响应体报错:", "err", err2.Error())
		return nil, err2
	}
	if len(string(b)) <= 200 {
		logger.L.Info("[" + toWho + "]请求结果:" + string(b))
	} else {
		logger.L.Info("[" + toWho + "]请求结果太长不打印")
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("[" + toWho + "]请求不成功:" + string(b))
	}
	return b, nil
}*/
