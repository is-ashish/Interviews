package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main_() {
	router := mux.NewRouter()
	router.Use(guidMiddleware)
	router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}

func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	uuid := r.Context().Value("uuid")
	log.Printf("[%v] Returning 200 - Healthy", uuid)
	w.Write([]byte("Healthy"))
}

func guidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
		next.ServeHTTP(w, r)
	})
}

// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main_() {
// 	// example of data sharing
// 	ctx := context.Background()
// 	ctx = addValue(ctx)
// 	readValue(ctx)
// }
// func addValue(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, "key", "ashish-jain")
// }
// func readValue(ctx context.Context) {
// 	fmt.Println(ctx.Value("key")) // Print context value
// }

// Applications in golang use Contexts for controlling and managing very important
// aspects of reliable applications, such as cancellation and data sharing in concurrent programming.

// it’s good way to share scoped values between those concurrent functions
// (meaning that each context will keep their own values on its scope).
// That’s exactly what net/http package does when handling HTTP requests.

// Let’s say you have a function where you start tens of goroutines. That main function waits for all goroutines to finish or a cancellation signal before proceeding.
// If you receive the cancellation signal you may want to propagate it to all your goroutines, so you don’t waste compute resources.
// If you share the same context among all goroutines you can easily do that.

//A Context carries deadlines, cancellation signals, and other request-scoped
// values across API boundaries and goroutines.
