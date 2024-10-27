## Middleware Documentation
Hey there, fellow dev! We’ve got a batch of middlewares ready for you to add some serious functionality to your Go web app. Each of these middlewares brings its own magic—security, logging, timeouts, and more! Let’s break them down one by one. 👇

### Authentication Middleware (JWT)
Our AuthMiddleware helps protect your routes with JSON Web Tokens (JWT). It’s flexible, letting you customize the secret key, error handling, and token extraction method.

Setup
To get started, configure your JWT settings using JWTConfig:

* SecretKey: The secret key for signing JWTs.
* SigningMethod: The JWT signing algorithm.
* ErrorHandler: Custom error handler for handling auth errors (optional).
* TokenExtractor: Extracts the token from the request header (optional).

### Example

```go
import (
	"github.com/BrunoCiccarino/GopherLight/middleware"
)

config := middleware.JWTConfig{
	SecretKey: []byte("your_secret_key"),
}
app.Use(middleware.NewAuthMiddleware(config))
```

### CORS Middleware
Need to allow cross-origin requests? No problem! Our CORSMiddleware configures the Cross-Origin Resource Sharing (CORS) settings to make your API accessible from other domains.

Config Options
* AllowOrigin: Set to "*" to allow any origin or specify a domain (e.g., "http://example.com").
* AllowMethods: Which HTTP methods are allowed? Common choices include "GET", "POST", etc.
* AllowHeaders: Specify which headers clients can use.
* AllowCredentials: Set to true if you want cookies or HTTP auth to be included.
* ExposeHeaders: Let the client read specific headers from the response.
* MaxAge: Cache time (in seconds) for preflight requests.

### Example
```go
corsOptions := middleware.CORSOptions{
	AllowOrigin: "*",
	AllowMethods: []string{"GET", "POST"},
}
app.Use(middleware.CORSMiddleware(corsOptions))
```

### CSRF Middleware
Our CSRFMiddleware protects against Cross-Site Request Forgery by validating a CSRF token sent with each request. Use GenerateCSRFToken() to create a secure token, then validate it with your own isValidToken function.

### Example

```go
app.Use(middleware.CSRFMiddleware(func(token string) bool {
	return token == "your_valid_token"
}))
```

And don’t forget to generate tokens with:
```go
csrfToken := middleware.GenerateCSRFToken()
```

### Logging Middleware
Want to keep track of what’s happening on your server? LoggingMiddleware logs each request, including the method, path, and time taken. It’s a great way to stay informed on app performance and any unusual activity.

### Example
```go
app.Use(middleware.LoggingMiddleware)
```
Each request will be logged like this:

* Started: Logs the request start time.
* Completed: Logs when the request finishes, including the duration.

### Timeout Middleware
Avoid those endless waits by setting time limits on request processing with TimeoutMiddleware. This middleware will cancel the request if it doesn’t complete in time, sending a 504 Gateway Timeout status to the client.

### Example
```go
import (
	"time"
	"github.com/BrunoCiccarino/GopherLight/middleware"
)

timeout := 2 * time.Second
app.Use(middleware.TimeoutMiddleware(timeout))
```

Putting It All Together
Alright, now that you’re equipped with JWT auth, CORS controls, CSRF protection, request logging, and request timeouts, you’re ready to make your app secure, flexible, and robust! Mix and match these middlewares as needed, and build a resilient API like a pro. Go forth and code!