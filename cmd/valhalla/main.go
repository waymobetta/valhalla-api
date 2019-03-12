//go:generate goagen bootstrap -d github.com/waymobetta/endorse/design

package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/valhalla-api/app"
	controllers "github.com/waymobetta/valhalla-api/controllers"
	mw "github.com/waymobetta/valhalla-api/middleware"
)

func main() {
	log.SetReportCaller(true)

	rtr := mux.NewRouter()

	rootHandler := handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type"},
		),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "OPTIONS"},
		),
		handlers.AllowedOrigins([]string{"*"}),
	)(rtr)

	port := "5000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	host := fmt.Sprintf("0.0.0.0:%s", port)

	// Create service
	service := goa.New("valhalla")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// mount controllers
	helseCtrlr := controllers.NewHelseController(service)
	app.MountHelseController(service, helseCtrlr)

	godkjentCtrlr := controllers.NewGodkjentController(service)
	app.MountGodkjentController(service, godkjentCtrlr)

	// goa handler
	goaHandler := mw.RateLimitHandler(service.Server.Handler)

	rootMux := http.NewServeMux()

	// NOTE: merging current routes with goa routes
	rootMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// limit payload sizes
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		// Update regex to include any base goa routes in order to properly forward to goa handler
		goaRoutesRegex := regexp.MustCompile(`v1/(helse|godkjent)`)
		isGoaRoute := goaRoutesRegex.Match([]byte(strings.ToLower(r.URL.Path)))

		if isGoaRoute {
			goaHandler.ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/documentation") {
			http.FileServer(http.Dir("./web")).ServeHTTP(w, r)
		} else {
			rootHandler.ServeHTTP(w, r)
		}
	})

	log.Printf("[cmd] listening on %s\n", host)

	// Start service
	panic(http.ListenAndServe(host, cors.Default().Handler(rootMux)))

}
