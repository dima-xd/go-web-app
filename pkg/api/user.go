package api

import (
	"github.com/dimaxdqwerty/go-web-app/pkg/data"

	"github.com/gorilla/mux"
)

type userAPI struct {
	data *data.UserData
}

func ServeUserResource(r *mux.Router, data data.UserData) {
	_ = &userAPI{data: &data}

}
