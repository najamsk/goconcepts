package main

import (
	"context"
	"example/graph"
	"example/graph/generated"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
// var userCtxKey = &graph.ContextKey{Name: "user"}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)

		c, err := r.Cookie("auth-cookie")

		// Allow unauthenticated users in
		if err != nil || c == nil {

			// set cookie for storing token
			cookie := http.Cookie{}
			cookie.Name = "auth-cookie"
			cookie.Value = "najam"
			cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
			cookie.Secure = false
			cookie.HttpOnly = true
			cookie.Path = "/"
			http.SetCookie(w, &cookie)

			next.ServeHTTP(w, r)
			return
		}
		if c.Value == "najam" {
			user := &graph.User{
				Name:    "najam awan",
				IsAdmin: true,
			}
			// put it in context
			ctx := context.WithValue(r.Context(), graph.UserCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return

		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Middleware decodes the share session cookie and packs the session into context
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r := mux.NewRouter()
	r.Use(authMiddleware)
	r.HandleFunc("/wow", handlerWow)
	// http.Handle("/", r)
	r.Handle("/pg", playground.Handler("grapqhl playground", "/query"))
	r.Handle("/query", srv).Methods(http.MethodPost).Name("query")

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handlerWow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from wow handler"))
}
