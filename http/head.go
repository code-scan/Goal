package http
// Head 方法
func (h *Http) Head(urls string) error {
	return h.New("HEAD", urls)
}