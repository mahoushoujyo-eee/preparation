package work

import (
	"fmt"
	"io"
)

type login interface {
	login()
	logout()
}

type oauth2Login struct {
	source string
}

func (o *oauth2Login) login() {
	fmt.Println("oauth2 login")
}

func (o *oauth2Login) logout() {
	fmt.Println("oauth2 logout")
}

var myOauth2Login login = &oauth2Login{}

type HookReader struct {
	r io.Reader
}

func (h *HookReader) Read(p []byte) (int, error) {
	// do something before read ...
	n, err := h.r.Read(p)
	// do something after read ...
	return n, err
}