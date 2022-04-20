package router

import (
	"encoding/json"
	"net/http"

	"github.com/deFarro/fastpoke_backend.git/app/config"
	"github.com/deFarro/fastpoke_backend.git/app/database"
)

type Router struct {
	Settings config.Config
	Database database.Database
}

func NewRouter(settings config.Config) (Router, error) {
	db, err := database.NewDatabase(settings)
	if err != nil {
		return Router{}, err
	}

	return Router{
		Settings: settings,
		Database: db,
	}, nil
}

type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewError(m string) Error {
	return Error{
		Error:   true,
		Message: m,
	}
}

func SendError(m string, w http.ResponseWriter) {
	err := NewError(m)

	payload, _ := json.Marshal(err)

	w.Write(payload)
}
