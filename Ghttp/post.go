package Ghttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

// POST 方法
func (h *Http) Post(urls string, params interface{}) error {
	err := h.New("POST", urls)
	if err != nil {
		return err
	}
	switch params.(type) {
	case string:
		h.SetPostString(params.(string))
	case map[string]interface{}:
		h.SetPostJson(params.(map[string]interface{}))
	case url.Values:
		h.SetPostValues(params.(url.Values))
	}
	return nil
}

//urlpkg.Values 格式的参数
func (h *Http) SetPostValues(values url.Values) {
	h.HttpBody = strings.NewReader(values.Encode())
	h.setParams()

}

//json 格式的参数
func (h *Http) SetPostJson(values map[string]interface{}) {
	bytesData, err := json.Marshal(values)
	if err != nil {
		log.Println("[!] SetPostJson Error: ", err)

	}
	h.HttpBody = bytes.NewReader(bytesData)
	h.SetContentType("json")
	h.setParams()

}

// 字符串格式的参数 a=1&b=2
func (h *Http) SetPostString(values string) {
	h.HttpBody = strings.NewReader(values)
	h.setParams()
}
func (h *Http) setParams() {
	if h.HttpRequest != nil && h.HttpBody != nil {
		h.HttpRequest.Body = ioutil.NopCloser(h.HttpBody)
	}
}
