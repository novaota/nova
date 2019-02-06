package rest

import (
	"fmt"
	"log"
	"net/http"
	"nova/shared"

	"github.com/gorilla/mux"
)

type Api struct {
	Port        int
	Controllers []Controller
	useTls      bool
	shared.CertificateSettings
	router      *mux.Router
	name        string
}

func NewRestApi(port int) *Api {
	return &Api{
		Port:     port,
		router:   mux.NewRouter(),
	}
}

func (api *Api) SetName(name string) {
	api.name = name
}

func (api *Api) AddController(controllers ...Controller) {
	for _, controller := range controllers {
		api.Controllers = append(api.Controllers, controller)
	}
}

func (api *Api) setupRoutes() {
	for _, controller := range api.Controllers {
		for _, route := range controller.GetRoutes() {
			log.Printf(" > Found Route %v %v\n", route.Method, route.Path)
			api.router.NewRoute().Path(route.Path).Methods(route.Method).Handler(route)
		}
	}
}

func (api *Api) UseTls(cert string, key string) {
	api.useTls = true
	api.CertificateSettings = shared.CertificateSettings{
		CACertificate: cert,
		CAKey:         key,
	}
}

func (api *Api) adress() string {
	return fmt.Sprintf(":%v", api.Port)
}

func (api *Api) StartServing() {
	api.setupRoutes()
	log.Printf("Starting %v on %v", api.name, api.adress())


	if api.useTls {
		http.ListenAndServeTLS(api.adress(), api.CertificateSettings.CACertificate, api.CertificateSettings.CAKey, api.router)
	} else {
		http.ListenAndServe(api.adress(), api.router)
	}
}
