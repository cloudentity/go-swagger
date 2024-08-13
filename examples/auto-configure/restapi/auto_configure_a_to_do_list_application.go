// Code generated by go-swagger; DO NOT EDIT.
// Auto configures api handlers Implementations.

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cloudentity/go-swagger/examples/auto-configure/implementation"
	"github.com/cloudentity/go-swagger/examples/auto-configure/restapi/operations"
	"github.com/cloudentity/go-swagger/examples/auto-configure/restapi/operations/todos"
)

//go:generate swagger generate server --target ../../auto-configure --name AToDoListApplication --spec ../swagger.yml --implementation-package github.com/cloudentity/go-swagger/examples/auto-configure/implementation --principal interface{}

// This file auto configures the api backend implementation.
// implementation package must already exist.
// implementation.New() is implemented by user, and must return an object
// or interface that implements Handler interface defined below.
var Impl Handler = implementation.New()

// Handler handles all api server backend configurations and requests
type Handler interface {
	Authable
	Configurable
	TodosHandler
}

// Configurable handles all server configurations
type Configurable interface {
	ConfigureFlags(api *operations.AToDoListApplicationAPI)
	ConfigureTLS(tlsConfig *tls.Config)
	ConfigureServer(s *http.Server, scheme, addr string)
	CustomConfigure(api *operations.AToDoListApplicationAPI)
	SetupMiddlewares(handler http.Handler) http.Handler
	SetupGlobalMiddleware(handler http.Handler) http.Handler
}

// Authable handles server authentication
type Authable interface {
	// Applies when the "x-todolist-token" header is set
	KeyAuth(token string) (interface{}, error)
}

/* TodosHandler  */
type TodosHandler interface {
	AddOne(params todos.AddOneParams, principal interface{}) middleware.Responder
	DestroyOne(params todos.DestroyOneParams, principal interface{}) middleware.Responder
	FindTodos(params todos.FindTodosParams, principal interface{}) middleware.Responder
	UpdateOne(params todos.UpdateOneParams, principal interface{}) middleware.Responder
}

func configureFlags(api *operations.AToDoListApplicationAPI) {
	Impl.ConfigureFlags(api)
}

func configureAPI(api *operations.AToDoListApplicationAPI) http.Handler {

	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-todolist-token" header is set
	api.KeyAuth = func(token string) (interface{}, error) {
		return Impl.KeyAuth(token)
	}

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams, principal interface{}) middleware.Responder {
		return Impl.AddOne(params, principal)
	})
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams, principal interface{}) middleware.Responder {
		return Impl.DestroyOne(params, principal)
	})
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams, principal interface{}) middleware.Responder {
		return Impl.FindTodos(params, principal)
	})
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams, principal interface{}) middleware.Responder {
		return Impl.UpdateOne(params, principal)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	// CustomConfigure can override or add to configurations set above
	Impl.CustomConfigure(api)

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	Impl.ConfigureTLS(tlsConfig)
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	Impl.ConfigureServer(s, scheme, addr)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return Impl.SetupMiddlewares(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return Impl.SetupGlobalMiddleware(handler)
}
