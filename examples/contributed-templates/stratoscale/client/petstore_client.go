// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/contributed-templates/stratoscale/client/pet"
	"github.com/cloudentity/go-swagger/examples/contributed-templates/stratoscale/client/store"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "petstore.service.strato"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new petstore HTTP client.
func New(c Config) *Petstore {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(Petstore)
	cli.Transport = transport
	cli.Pet = pet.New(transport, strfmt.Default, c.AuthInfo)
	cli.Store = store.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// Petstore is a client for petstore
type Petstore struct {
	Pet       *pet.Client
	Store     *store.Client
	Transport runtime.ClientTransport
}
