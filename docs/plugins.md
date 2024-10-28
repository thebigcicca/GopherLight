## Plugin Documentation

Hey, plugin creators! 🚀 Ready to extend GopherLight with some cool custom functionalities? Our Plugin interface is here to make that happen! With just a bit of setup, you can add your own routes, features, and enhancements to any GopherLight app.

### The Plugin Interface
The Plugin interface is super simple but super powerful. It gives you a single method: Register. This lets you hook into the app’s routing system to add any routes you need—whether it’s a new API endpoint, a webhook, or anything else you can imagine.

The Register Method
Here’s the magic part of the Plugin interface:

```go
type Plugin interface {
	Register(route func(method, path string, handler func(req *req.Request, res *req.Response)))
}
```

The Register method accepts a route function that lets you define new routes in your plugin by specifying:

- method: HTTP method (e.g., "GET", "POST", etc.)
- path: The route path (e.g., "/my-plugin-route")
- handler: The function to execute when the route is hit. This function receives:
    - req: The request object with access to query parameters, headers, and body.
    - res: The response object to send data back to the client.

### Example Plugin

Let’s say you want to create a plugin that adds a simple endpoint at /hello-plugin to greet users. Here’s what the plugin would look like:

```go
package main

import (
	"github.com/BrunoCiccarino/GopherLight/plugins"
	"github.com/BrunoCiccarino/GopherLight/req"
)

type HelloPlugin struct{}

// Register adds a new route for the HelloPlugin.
func (p *HelloPlugin) Register(route func(method, path string, handler func(req *req.Request, res *req.Response))) {
	route("GET", "/hello-plugin", func(req *req.Request, res *req.Response) {
		res.Send("Hello from the HelloPlugin!")
	})
}
```

### Adding the Plugin to Your App
To load a plugin, simply create an instance and call Register in your main app setup:

```go
package main

import (
	"github.com/BrunoCiccarino/GopherLight/router"
)

func main() {
	app := router.NewApp()
	helloPlugin := &HelloPlugin{}
	helloPlugin.Register(app.Route)

	app.Listen(":3333")
}
```

### Customizing Your Plugins
Each plugin can add as many routes as needed. Just call route multiple times in your Register function to define additional endpoints. Use different HTTP methods, paths, and handlers to shape your plugin’s functionality however you want.

So there you have it! With this flexible Plugin interface, you can easily add new features to GopherLight and make your app even more powerful. Happy coding, and enjoy plugging into the framework! 🔌