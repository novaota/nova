package rest

type Controller interface {
	GetRoutes() []Route
}

