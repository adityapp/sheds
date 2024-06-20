package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adityapp/sheds/configs"
	"github.com/adityapp/sheds/internal/controllers"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type sheds struct {
}

func Init() (*sheds, error) {
	err := os.MkdirAll(configs.Get().WorkDir, 0755)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	fmt.Printf("\nCreated directory %+v\n", configs.Get().WorkDir)

	return &sheds{}, nil
}

func (app *sheds) Start() {
	route := chi.NewRouter()
	route.Route("/api", func(r chi.Router) {
		r.Post("/file/upload", controllers.Upload)
		r.Post("/file/delete", controllers.Delete)
		r.Post(("/file/move"), controllers.Move)
		r.Post(("/file/copy"), controllers.Copy)
	})

	fmt.Printf("Stating server in :8080\n\n")
	http.ListenAndServe(":8080", route)
}
