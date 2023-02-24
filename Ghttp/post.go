package Ghttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

// POST 方法
func (h *Http) Post(urls string, params interface{}) *Http {
	h.New("POST", urls)
	if h.err != nil {
		return h
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
func (h *Http) SetPostJson(values map[string]interface{}) *Http {
	bytesData, err := json.Marshal(values)
	if err != nil {
		log.Println("[!] SetPostJson Error: ", err)

	}
	h.HttpBody = bytes.NewReader(bytesData)
	h.SetContentType("json")
	h.setParams()
	return h

}

// 字符串格式的参数 a=1&b=2
func (h *Http) SetPostString(values string) *Http {
	h.HttpBody = strings.NewReader(values)
	h.setParams()
	return h
}
func (h *Http) setParams() *Http {
	if h.HttpRequest != nil && h.HttpBody != nil {
		h.HttpRequest.Body = ioutil.NopCloser(h.HttpBody)
		switch v := h.HttpBody.(type) {
		case *bytes.Buffer:
			h.HttpRequest.ContentLength = int64(v.Len())
			buf := v.Bytes()
			h.HttpRequest.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
		case *bytes.Reader:
			h.HttpRequest.ContentLength = int64(v.Len())
			snapshot := *v
			h.HttpRequest.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *strings.Reader:
			h.HttpRequest.ContentLength = int64(v.Len())
			snapshot := *v
			h.HttpRequest.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		default:

		}
	}
	return h
}
