## Changelog

### GopherLight v0.1 Release Notes
🚀 GopherLight v0.1 - The First Release!
We are excited to announce the first version of express-go, a micro-framework inspired by Express.js, now for the Go universe! This initial release brings the simplicity and flexibility you already love from Express, but with the robust performance and reliability of Go.

🎯 Main Features
* Simple Routing: Set up routes quickly and easily. Just define the URL and a handler, and that's it!

```go
app.Route("/hello", func(r *req.Request, w *req.Response) {
    w.Send("Hello, World!")
})
```

* Easy Request Handling: The framework encapsulates HTTP request details to facilitate access to query parameters, headers and request body.
```go
name := r.QueryParam("name")
token := r.Header("Authorization")
```

* Flexible Responses: Send responses in plain text or JSON, all with convenient methods.
```go 
w.Send("Hello, Go!")
w.JSON(map[string]string{"message": "Hello, JSON"})
```

🛠️ What's Coming
This is only version 0.1, so there's still a lot to go! We are working on:

* Middleware for handling authentication and validations.
* Support dynamic routes and route parameters.
* Improvements in error management and customized responses.

📝 Contributions
This is an early version, and we are open to suggestions, improvements and contributions from the community. Feel free to explore the code, open issues, or submit PRs to make express-go even better!

### GopherLight v0.2 Release Notes
🚀 GopherLight v0.2 - Enhanced Routing and Middleware Support!

We are excited to announce the second version of GopherLight, bringing more flexibility and features to our lightweight Go framework. With this update, you’ll experience a more powerful routing system and the addition of convenient middleware to streamline your development process.

### 🎯 Additions

* Support for Multiple HTTP Methods: GopherLight now allows routing for various HTTP methods such as GET, POST, PUT, DELETE, and more. You can define routes for each specific method, increasing flexibility in handling requests.

```go
app.Route("POST", "/submit", func(r *req.Request, w *req.Response) {
    w.Send("Data received!")
})
```

* Logging Middleware: Automatically log the start and completion time of each request, making it easier to debug and monitor the behavior of your app.

* Timeout Middleware: Prevent long-running requests from blocking your application with the newly introduced timeout middleware. You can set time limits for your request handling.

```go
app.Use(middleware.LoggingMiddleware)
app.Use(middleware.TimeoutMiddleware(2 * time.Second))
```

### 🛠️ Improvements

* HTTP Method Validation: We've implemented stricter validation for HTTP methods. Now, your routes will only respond to the defined methods, ensuring security and better management of requests.
### 🔄 Changes

* Route Function Refactoring: The route function has been restructured to handle multiple HTTP methods efficiently, improving performance and code clarity.
* Documentation Update: The documentation has been updated to reflect all the new features, including detailed examples of how to utilize multiple HTTP methods and middleware in your application.
### 🚀 What’s Next?

* Enhanced error handling and better integration with third-party services.
Middleware for authentication and CSRF protection.

* 📝 Contributions GopherLight continues to grow with your support! We welcome contributions, suggestions, and improvements from the community. Feel free to explore, submit issues, or open PRs to make the framework even better.