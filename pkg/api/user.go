package api

import (
	"github.com/dimaxdqwerty/go-web-app/go-web-app/pkg/data"

	"github.com/dimaxdqwerty/go-web-app/github.com/gorilla/mux"
)

type userAPI struct {
	data *data.UserData
}

func ServeUserResource(r *mux.Router, data data.UserData) {
	_ = &userAPI{data: &data}

}
