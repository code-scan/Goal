package Ghttp

func (h *Http) Session() *Http {
	h.isSession = true
	return h
}
