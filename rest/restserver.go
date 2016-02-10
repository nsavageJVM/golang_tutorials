package rest


import (
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	GET    = "GET"
	POST   = "POST"
)

// GetSupported is the interface that provides the Get
// method a resource must support to receive HTTP GETs.
type GetSupported interface {
	Get(url.Values, http.Header) (int, interface{}, http.Header)
}

// PostSupported is the interface that provides the Post
// method a resource must support to receive HTTP POSTs.
type PostSupported interface {
	Post(url.Values, http.Header) (int, interface{}, http.Header)
}


// here we are binding this function to a struct 'API'
func  (api *API)  BootStrapServer(port int) error {

	if !api.muxInitialized {
		return errors.New("You must add at least one resource to this API.")
	}
	portString := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(portString, api.Mux())
}

// NewAPI allocates and returns a new API.
func NewAPI() *API {
	return &API{}
}


type API struct {
	mux     *http.ServeMux
	muxInitialized bool
}


// AddResource adds a new resource to an API. The API will route
// requests that match one of the given paths to the matching HTTP
// method on the resource.
func (api *API) AddResource(resource interface{}, paths ...string) {
	for _, path := range paths {
		api.Mux().HandleFunc(path, api.requestHandler(resource))
	}
}

// Mux returns the http.ServeMux used by an API. If a ServeMux has
// does not yet exist, a new one will be created and returned.
func (api *API) Mux() *http.ServeMux {
	if api.muxInitialized {
		return api.mux
	} else {
		api.mux = http.NewServeMux()
		api.muxInitialized = true
		return api.mux
	}
}

func (api *API) requestHandler(resource interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		var handler func(url.Values, http.Header) (int, interface{}, http.Header)

		switch request.Method {
		case GET:
			if resource, ok := resource.(GetSupported); ok {
				handler = resource.Get  }

		case POST:
			if resource, ok := resource.(PostSupported); ok {
				handler = resource.Post    }
		}

		if handler == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		code, data, header := handler(request.Form, request.Header)
		content, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for name, values := range header {
			for _, value := range values {
				rw.Header().Add(name, value)
			}
		}
		rw.WriteHeader(code)
		rw.Write(content)


	}
}


