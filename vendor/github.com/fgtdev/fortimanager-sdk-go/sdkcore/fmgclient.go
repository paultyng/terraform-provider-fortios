package fmgclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	Id      uint64         `json:"id"`
	Method  string         `json:"method"`
	Session string         `json:"session"`
	Params  [1]interface{} `json:"params"`
}

type FmgSDKClient struct {
	Ipaddr string
	User   string
	Passwd string
	Debug  string
	Client *http.Client
}

func NewClient(ip, user, passwd string, client *http.Client) *FmgSDKClient {

	d := os.Getenv("TRACEDEBUG")

	return &FmgSDKClient{
		Ipaddr: ip,
		User:   user,
		Passwd: passwd,
		Client: client,
		Debug:  d,
	}
}

func (c *FmgSDKClient) Execute(req *Request) (result map[string]interface{}, err error) {

	j, _ := json.Marshal(req)

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> request: %s", j)
	}

	httpResp, err := c.Client.Post("http://"+c.Ipaddr+"/jsonrpc", "application/json", bytes.NewBuffer(j))
	if err != nil {
		err = fmt.Errorf("Login failed: %s", err)
		return
	}
	defer httpResp.Body.Close()

	// check response
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> result: %s", body)
	}

	result = map[string]interface{}{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil {
		if id := uint64(result["id"].(float64)); id != req.Id {
			err = fmt.Errorf("id not match, should be 1, but is %d", id)
			return
		}

		if result["result"] != nil {
			status := (result["result"].([]interface{}))[0].(map[string]interface{})["status"].(map[string]interface{})

			c := uint(status["code"].(float64))
			m := status["message"].(string)
			if c != 0 || m != "OK" {
				err = fmt.Errorf("status not right: code is %d, message is %s", c, m)
				return
			}
		} else {
			err = fmt.Errorf("can't get response status: %s", err)
			return
		}
	}

	return
}

func (c *FmgSDKClient) Login() (session string, err error) {
	params := map[string]interface{}{
		"data": map[string]string{
			"user":   c.User,
			"passwd": c.Passwd,
		},
		"url": "/sys/login/user",
	}

	req := &Request{
		Id:     1,
		Method: "exec",
		Params: [1]interface{}{params},
	}

	result, err := c.Execute(req)
	if err != nil {
		return "", fmt.Errorf("login failed:%s", err)
	}

	session = result["session"].(string)
	return
}

func (c *FmgSDKClient) Logout(s string) (err error) {
	params := map[string]interface{}{
		"url": "/sys/logout",
	}

	req := &Request{
		Id:      1,
		Method:  "exec",
		Params:  [1]interface{}{params},
		Session: s,
	}

	_, err = c.Execute(req)
	if err != nil {
		err = fmt.Errorf("logout failed:%s", err)
		return
	}
	return
}

func (c *FmgSDKClient) Do(method string, params map[string]interface{}) (output map[string]interface{}, err error) {
	session, err := c.Login()
	if err != nil {
		return nil, fmt.Errorf("Executing failed", err)
	}
	defer c.Logout(session)

	req := &Request{
		Id:      1,
		Method:  method,
		Params:  [1]interface{}{params},
		Session: session,
	}

	output, err = c.Execute(req)
	return
}

func (f *FmgSDKClient) Trace(s string) func() {
	if f.Debug == "ON" {
		log.Printf("[TRACEDEBUG] -> Enter %s <-", s)
		return func() { log.Printf("[TRACEDEBUG]    -> Leave %s <-", s) }
	}

	return func() {}
}
