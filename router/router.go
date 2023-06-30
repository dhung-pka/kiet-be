package router

import (
	"log"
	"net/http"
	"service/config"
	"service/controller"

	_ "service/docs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)
	r.Use(jwtauth.Verifier(config.GetJWT()))

	authController := controller.NewAuthController()
	basicQuery := controller.NewBasicQueryController()

	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:18888/api/v1/swagger/doc.json"), //The url pointing to API definition
	))

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"value": "Test",
		})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(authRouter chi.Router) {
			authRouter.Post("/login", authController.Login)
		})
		r.Route("/basic-query", func(basicQueryRouter chi.Router) {
			basicQueryRouter.Post("/", basicQuery.BasicQuery)
		})
	})

	log.Println("Run server: http://localhost:18888/api/v1")
	log.Println("Run server swagger: http://localhost:18888/api/v1/swagger/index.html")

	return r
}
