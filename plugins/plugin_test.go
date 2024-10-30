package plugins_test

import (
	"testing"

	"github.com/BrunoCiccarino/GopherLight/req"
)

type MockResponse struct {
	Body string
}

func (res *MockResponse) Send(body string) {
	res.Body = body
}

type MockRouter struct {
	registeredRoutes map[string]func(req *req.Request, res *MockResponse)
}

func (m *MockRouter) Route(method, path string, handler func(req *req.Request, res *MockResponse)) {
	if m.registeredRoutes == nil {
		m.registeredRoutes = make(map[string]func(req *req.Request, res *MockResponse))
	}
	m.registeredRoutes[path] = handler
}

type MockPlugin struct{}

func (p *MockPlugin) Register(route func(method, path string, handler func(req *req.Request, res *MockResponse))) {
	route("GET", "/mock-plugin-route", func(req *req.Request, res *MockResponse) {
		res.Send("Mock Plugin Route")
	})
}

func TestPluginRegistration(t *testing.T) {
	mockRouter := &MockRouter{}
	mockPlugin := &MockPlugin{}

	mockPlugin.Register(mockRouter.Route)

	if len(mockRouter.registeredRoutes) == 0 {
		t.Fatal("No routes were registered by the plugin")
	}

	handler, exists := mockRouter.registeredRoutes["/mock-plugin-route"]
	if !exists {
		t.Fatalf("Route '/mock-plugin-route' was not registered correctly")
	}

	req := &req.Request{}
	res := &MockResponse{}

	handler(req, res)

	if res.Body != "Mock Plugin Route" {
		t.Fatalf("Expected 'Mock Plugin Route' response, but got '%s'", res.Body)
	}
}
