## Changelog

### express-go v0.1 Release Notes
🚀 express-go v0.1 - The First Release!
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

## [0.2] - 2024-10-21

### Additions
- **Support for multiple HTTP methods**: The framework now allows the routing of requests to HTTP methods such as GET, POST, PUT, DELETE and others, providing greater flexibility in defining routes.

### Improvements
- **HTTP method validation**: Implementation of validation to ensure that routes respond only to appropriate methods, increasing security and clarity when handling requests.

### Changes
- **Route function refactoring**: The function has been adjusted to accept and manage different HTTP methods efficiently.
- **Documentation update**: The documentation has been revised and updated to reflect the new capabilities of the framework, including examples of use of the different HTTP methods.