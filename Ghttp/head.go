package Ghttp

// Head 方法
func (h *Http) Head(urls string) *Http {
	return h.New("HEAD", urls)
}
