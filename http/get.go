package http

// Get 方法
func (h *Http) Get(url string) error {
	var err error
	err = h.New("GET", url)
	return err

}
