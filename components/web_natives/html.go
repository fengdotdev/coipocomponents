package display

import "github.com/fengdotdev/coipoblog/pkg/coipocom/components"



type Html struct {
	Code string
}

func (h *Html) Render() string {
	return h.Code
}

func (h *Html) String() string {
	return h.Render()
}

func (h *Html) iam() (components.Component, {
	return h
}