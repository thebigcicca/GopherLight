# GopherLight ![gopher1](./img/typing-furiously.gif)

[![GitHub License](https://img.shields.io/github/license/BrunoCiccarino/express-go?style=for-the-badge&color=blue&link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino%2Fexpress-go%2Fblob%2Fmain%2FLICENSE)](https://github.com/BrunoCiccarino/GopherLight/blob/main/LICENSE) 
![Go Reference](https://img.shields.io/badge/reference-grey?style=for-the-badge&logo=go&link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino%2Fexpress-go) 
![pr's welcome](https://img.shields.io/badge/PR'S-WELCOME-green?style=for-the-badge) 
[![discord](https://img.shields.io/badge/discord-grey?style=for-the-badge&logo=discord)](https://discord.gg/fJ2gvdvCpk) 
![GitHub Repo stars](https://img.shields.io/github/stars/BrunoCiccarino/express-go) 
[![GitHub followers](https://img.shields.io/github/followers/BrunoCiccarino?link=https%3A%2F%2Fgithub.com%2FBrunoCiccarino)](https://github.com/BrunoCiccarino) 
![GitHub forks](https://img.shields.io/github/forks/BrunoCiccarino/express-go) 

<img src="./img/image.png" align="right">

### What is GopherLight?
Hey there! So, you know how building web applications can sometimes feel like climbing a mountain? Well, GopherLight is like that cool hiking buddy who helps you navigate the trail, making things way easier and way more fun!

GopherLight is a micro framework for Go (Golang) that brings a bit of the simplicity and flexibility of the popular Express.js framework from the Node.js world right to your Go projects. It’s perfect for those times when you want to whip up a web server or an API without getting bogged down in all the nitty-gritty details.

Imagine you want to handle HTTP requests and create endpoints to manage users—just like in a classic CRUD (Create, Read, Update, Delete) app. GopherLight, you can define your routes and handlers in a snap. No need to wrestle with the standard net/http package; instead, you get a clean and straightforward way to manage your routes and responses.

The cool part? You get to focus on writing your application logic while the framework handles the heavy lifting under the hood. Need to add a new route? Just call a method and pass in your handler. Want to send a JSON response? Easy peasy!

Plus, it’s lightweight, so it won’t weigh down your application. You get all the goodies of a modern web framework while keeping things simple and fast. Whether you're a seasoned pro or just dipping your toes into web development, express-go makes it a breeze to get your ideas off the ground.

So, if you’re looking for a friendly and efficient way to build web apps in Go, GopherLight is your new best friend. Grab your backpack, and let’s hit the trail!

> [!WARNING]
> We are in an initial beta version, so it is likely that the framework will change a lot, always stay up to date, with an updated version of your code using the framework.

### Tasks

- [x] router 100%
- [x] http requests 100%
- [x] manipulation of the methods (get, post, put, delete ...) 100%
- [x] plugin support 100%
- [x] more detailed error logs 100%
- [x] basical middlewares (~~authentication~~, ~~timeout~~, ~~CORS~~, ~~csrf~~, ~~logging~~, etc...) 100%
- [x] More complete documentation 100%
- [ ] next func 0 %

### Installation

```bash
go get github.com/BrunoCiccarino/GopherLight/router
go get github.com/BrunoCiccarino/GopherLight/req
go get github.com/BrunoCiccarino/GopherLight/middleware
go get github.com/BrunoCiccarino/GopherLight/plugins
```

### basic usage example

```go
package main

import (
	"fmt"
	"github.com/BrunoCiccarino/GopherLight/router"
	"github.com/BrunoCiccarino/GopherLight/req"
)


func main() {
	app := router.NewApp()

	// Define a route that responds to a GET request at "/hello".
	app.Get("/hello", func(r *req.Request, w *req.Response) {
		w.Send("Hello, World!")
	})

	fmt.Println("Server listening on port 3333")
	app.Listen(":3333")
}
```

Do you want to learn how to create APIs like a professional in a simple, fast and efficient way using our framework? Follow this link to the documentation: [link](./docs/docs.md)

### Contribute

That said, there's a bunch of ways you can contribute to this project, like by:

* ⭐ Giving a star on this repository (this is very important and costs nothing)
* 🪲 Reporting a bug
* 📄 Improving this [documentation](./docs/)
* 🚨 Sharing this project and recommending it to your friends
* 💻 Submitting a pull request to the official repository
* ⚠️ Before making a pull request, it is important that you read our [doc](.github/CONTRIBUTING)


### Contributors

This project exists thanks to all the people who contribute. 

<a href="https://github.com/BrunoCiccarino/GopherLight/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=BrunoCiccarino/GopherLight&max=24" />
</a>

