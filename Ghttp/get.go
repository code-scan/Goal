package Ghttp

// Get 方法
func (h *Http) Get(url string) *Http {
	return h.New("GET", url)
}
