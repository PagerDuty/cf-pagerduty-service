package auth

import "net/http"

// Wrapper auth wrapper
type Wrapper struct {
	username string
	password string
}

// NewWrapper creates new wrapper
func NewWrapper(username, password string) *Wrapper {
	return &Wrapper{
		username: username,
		password: password,
	}
}

const notAuthorized = "Not Authorized"

// WrapFunc wraps HandlerFunc
func (wrapper *Wrapper) WrapFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authorized(wrapper, r) {
			http.Error(w, notAuthorized, http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
	})
}

func authorized(wrapper *Wrapper, r *http.Request) bool {
	username, password, isOk := r.BasicAuth()
	return isOk && username == wrapper.username && password == wrapper.password
}
