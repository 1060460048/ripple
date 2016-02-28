# Ripple
a MVC web framework for Go(base on [Echo](https://github.com/labstack/echo))

# Which features to include in a framework

A framework has 3 parts. A router receiving a request and directing it to a handler, a middleware system to add reusable pieces of software before and after the handler, and the handler processing the request and writing the response.
- Router
- Middlewares
- Handler processing

![architect]()
<img src="https://www.nicolasmerouze.com/img/illustrations/byof-part1@2x.png" width="400" height="275" />

Middlewares handle:
- error/panic
- logging
- security
- sessions
- cookies
- body parsing

# Features
* [x] MySQL and Foundation database support
* [x] Modular (you can choose which components to use)
* [x] Middleware support, compatible Middleware works out of the box
* [x] Lightweight. Only MVC
* [x] Multiple configuration files support (currently json)

# Overview
`ripple` is a lightweight MVC framework. It is based on the principles of simplicity, relevance and elegance.
- Simplicity. The design is simple, easy to understand and doesn't introduce many layers between you and the standard library. It is a goal of the project that users should be able to understand the whole framework in a single day.
- Relevance. `ripple` doesn't assume anything. We focus on things that matter, this way we are able to ensure easy maintenance and keep the system well-organized, well-planned and sweet.
- Elegance. `ripple` uses golang best practises. We are not afraid of heights, it's just that we need a parachute in our backpack. The source code is heavily documented, any functionality should be well explained and well tested.

# Installation

	$ go get github.com/bmbstack/ripple
	$ go get github.com/bmbstack/ripple/cmd/ripple
	$ go get github.com/bmbstack/ripple/cmd/wbs
	$ ripple new rippleApp
	$ cd $GOPATH/src/rippleApp
	$ wbs -c wbs.toml
	
Then, Open the url: http://localhost:8090

# Project structure
This is the structure of the `rippleApp` list application that will showcase how you can build web apps with `ripple`:

```shell
├── bin
│   └── forum
├── config
│   └── config.json
├── controllers
│   └── home.go
├── frontend
│   ├── static
│   │   ├── css
│   │   │   └── app.css
│   │   └── js
│   │       └── app.js
│   └── templates
│       └── home
│           ├── html.html
│           └── index.html
├── logger
│   └── logger.go
├── main.go
├── models
│   └── user.go
├── scripts
│   ├── commands.go
│   ├── const.go
│   ├── init.go
│   └── server.go
└── wbs.toml

12 directories, 15 files
```

# Screenshots

ripple new rippleApp
<img src="https://raw.githubusercontent.com/bmbstack/ripple/master/screenshots/ripple_new.png"/>
Debug mode
<img src="https://raw.githubusercontent.com/bmbstack/ripple/master/screenshots/ripple_debug.png"/>
Release mode
<img src="https://raw.githubusercontent.com/bmbstack/ripple/master/screenshots/ripple_release.png"/>
Web browser
<img src="https://raw.githubusercontent.com/bmbstack/ripple/master/screenshots/ripple_chrome.png"/>

# Acknowledgements
These amazing projects have made ripple possible:

- https://github.com/labstack/echo
- https://github.com/boj/redistore
- https://github.com/codegangsta/cli
- https://github.com/flosch/pongo2
- https://github.com/garyburd/redigo
- https://github.com/gorilla/sessions
- https://github.com/weisd/cache
- https://github.com/xyproto/cookie
- https://github.com/xyproto/pinterface
- https://github.com/xyproto/simpleredis
- https://golang.org/x/crypto
- https://gopkg.in/bluesuncorp/validator.v5
- https://github.com/achiku/wbs


