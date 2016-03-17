# cardamomo

**Warning!** This project is under development and at this time can be unstable!

### Installation

Import next package in your project

```sh
import (
    "github.com/zombcode/cardamomo"
)
```

Or if you want, can clone the repository with

```sh
$ git clone https://github.com/zombcode/cardamomo.git
````

### First steps

##### Basics

For instanciate the cardamomo framework in your project, you must write

```sh
c := cardamomo.Instance()
```

at this moment your cardamomo is instanciated. When you are ready, you can do

```sh
c.Run()
```

for run the cardamomo http server

##### GET

For generate GET patterns you can do

```sh
c.Get("/", func(res cardamomo.Response) {
    res.Send("Hello world!");
})
```

you can see, the variable **res** in callback function is for send data to the client.

##### POST

For generate POST patterns you can do

```sh
c.Get("/", func(res cardamomo.Response) {
    res.Send("Hello world!");
})
```

you can see, the variable **res** in callback function is for send data to the client.

##### BASE

The **base** is used for GROUP various routes below the same base path.

```sh
c.Base("/base", func(router cardamomo.Router) {
    router.Get("/route", func(res cardamomo.Response) {
        res.Send("Hello route base/route!");
    })
})
```

##### Future

At this moment the framework is very simple, in the future we want to implement:

- JSON responses
- GET and POST request data
- REQUEST body parser
- Cookies
- Layout manager

### Version
**0.0.1**
