package block

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//JSON RPC RPC 客户端 (over HTTP(s))
type RpcClient struct {
	serverAddr string
	user       string
	passwd     string
	httpClient *http.Client
}

//请求参数配置
type Config struct {
	Key   string
	Value interface{}
}

//  RCP 请求
type rpcRequest struct {
	Method  string        `json:"method"`
	Configs []interface{} `json:"configs"`
	Id      int64         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
}

// RCP error(错误)
type rpcError struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

//RPC 响应
type rpcResponse struct {
	Id     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
	Err    interface{}     `json:"error"`
}

//创建一个客户端链接
func newClient(host string, port int, user, passwd string, useSSL bool) (c *RpcClient, err error) {
	if len(host) == 0 {
		err = errors.New("请求主机失败")
	}
	var serverAddr string
	var httpClient *http.Client

	if useSSL {
		serverAddr = "https://"
		t := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		httpClient = &http.Client{Transport: t}
	} else {
		serverAddr = "http://"
		httpClient = &http.Client{}
	}
	c = &RpcClient{
		serverAddr: fmt.Sprintf("%s%s:%d", serverAddr, host, port),
		user:       user,
		passwd:     passwd,
		httpClient: httpClient}

	return
}

//http 超时请求
func (c *RpcClient) doTimeoutRequest(timer *time.Timer, req *http.Request) (re *http.Response, err error) {
	type reslut struct {
		resp *http.Response
		err  error
	}
	done := make(chan reslut, 1)

	go func() {
		resp, err := c.httpClient.Do(req)
		done <- reslut{resp: resp, err: err}
	}()

	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("链接服务器超时")
	}

}

//请求回复
func (c *RpcClient) call(method string, configs []interface{}) (rr rpcResponse, err error) {
	connectionTimer := time.NewTimer(RPCCLIENT_TIMEOUT * time.Second)
	c.httpClient.CloseIdleConnections()

	rpsR := rpcRequest{
		method,
		configs,
		time.Now().UnixNano(),
		"2.0",
	}

	payloadBuffer := &bytes.Buffer{}
	jsonEncode := json.NewEncoder(payloadBuffer) //json 序列化

	err = jsonEncode.Encode(rpsR) //序列化请求
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	//使用提供的用户名和密码进行基本身份验证。
	if len(c.user) > 0 || len(c.passwd) > 0 {
		req.SetBasicAuth(c.user, c.passwd)
	}

	resp, err := c.doTimeoutRequest(connectionTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	//log.Printf("DATA:%s",data)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("HTTP error: " + resp.Status)
		return
	}
	err = json.Unmarshal(data, &rr)

	return
}
