package main

import (
	compendit "compendit/internal"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"

	// limiter "github.com/ulule/limiter/v3"
	// mhttp "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	// sredis "github.com/ulule/limiter/v3/drivers/store/redis"

	"log"
	"net/http"
)

func main() {
	if cfg, err := compendit.RetrieveConfig(); err != nil {
		log.Fatal(err)
		// } else if app, err := compendit.NewApp(cfg); err != nil {
		// 	log.Fatal(err)
		// } else if store, err := sredis.NewStoreWithOptions(app.RedisClient, limiter.StoreOptions{
		// 	Prefix: fmt.Sprintf("%v:compenditApi:rateLimit", cfg.App.Env),
		// }); err != nil {
		// log.Fatal(fmt.Errorf("error while making redis store for rate limiting: %w", err))
		// } else if server, err := NewServer(app); err != nil {
		// 	log.Fatal(fmt.Errorf("error while making server: %w", err))
	} else {
		// if key := cfg.BugsnagApiKey(); key != nil {
		// 	bugsnag.Configure(bugsnag.Configuration{
		// 		APIKey:          *key,
		// 		ReleaseStage:    cfg.App.Env,
		// 		ProjectPackages: []string{"main", "compendit/cmd/core"},
		// 	})
		// }

		// if env := cfg.AppEnv(); env != "development" {
		// 	tracer.Start(
		// 		tracer.WithEnv(env),
		// 	)
		// 	defer tracer.Stop()
		// }

		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		// router.Use(mhttp.NewMiddleware(limiter.New(store, limiter.Rate{
		// 	Period: time.Minute,
		// 	Limit:  int64(cfg.Api.RateLimit),
		// }, limiter.WithTrustForwardHeader(true))).Handler)

		corsInstance := cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{
				"X-Pagination-Offset",
				"X-Pagination-Limit",
				"X-Pagination-Total-Count",
			},
			AllowCredentials: true,
		})
		router.Use(corsInstance.Handler)

		RegisterRoutes(router) //routes register

		// Host spec
		// router.Use(func(next http.Handler) http.Handler {
		// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 		if r.URL.Path == cfg.App.coreBaseURL+"/spec.yaml" {
		// 			if _, err := w.Write(coreApiSpecFileContent); err != nil {
		// 				panic(err)
		// 			}
		// 		} else {
		// 			next.ServeHTTP(w, r)
		// 		}
		// 	})
		// })

		// router.Use(middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		// 	Options: openapi3filter.Options{
		// 		AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		// 			return nil // Do nothing for now, since we already have another middleware handling the user authentication
		// 		},
		// 	},
		// }))

		// oapi.HandlerFromMuxWithBaseURL(strictHandler, router, cfg.App.coreBaseURL)

		httpHandler := http.Handler(router)
		// if cfg.BugsnagApiKey() != nil {
		// 	httpHandler = bugsnag.Handler(httpHandler)
		// }

		addr := fmt.Sprintf(":%d", cfg.App.Port)
		fmt.Printf("Listening at %q\n", addr)

		err := http.ListenAndServe(addr, httpHandler)
		log.Fatal(err)
	}
}
