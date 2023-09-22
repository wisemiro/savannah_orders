package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (repo *HandlerRepo) InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*", "Strict-Transport-Security"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(repo.securityHeaders)

	router.Route("/", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/git-auth", GitAuth())
			r.Get("/callback", repo.HandleGitHubCallback())
		})
		//
		r.Route("/products", func(r chi.Router) {
			r.Post("/", repo.CreateProduct())
			r.Put("/{id}", repo.UpdateProduct())
			r.Delete("/{id}", repo.DeleteProduct())
			r.Get("/", repo.ListProducts())
			r.Get("/{id}", repo.GetProduct())

		})
		//
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", repo.CreateOrder())
			r.Put("/{id}", repo.UpdateOrder())
			r.Get("/{id}", repo.GetOrder())
			r.Get("/customer/{cutomer_id}", repo.ListOrdersByCustomer())
			r.Get("/", repo.ListOrders())
		})
	})
	return router
}
func (repo *HandlerRepo) securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

// func (repo *HandlerRepo) AuthenticationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authorizationHeader := r.Header.Get("authorization")

// 		if len(authorizationHeader) == 0 {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest)
// 			render.Respond(w, r, errors.New("authorization header is not provided"))
// 			return
// 		}

// 		fields := strings.Fields(authorizationHeader)
// 		authorizationType := strings.ToLower(fields[0])

// 		if authorizationType != "bearer" {
// 			err := fmt.Sprintf("unsupported authorization type %s", authorizationType)
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest)
// 			render.Respond(w, r, errors.New(err))
// 			return
// 		}
// 		err = repo.store.Verify(r.Context(), authorizationHeader, )

// 	})
// }
