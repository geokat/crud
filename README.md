Description
=======

A simple CRUD API demo server written in Go using only the standard library.

How To Run
=======
If you have the Go toolchain installed on your system, you can run:

```bash
go get github.com/geokat/crud
crud
```

This will download, build and start the demo server.

To start the demo server using docker instead, `cd` to the source directory and run:

```bash
docker build -t crud .
docker run -it -p 8080:8080 crud
```

In both cases the running app will be available on localhost:8080.

How To Test
=======

You can test the demo server by making API requests using `curl`. The server provides four RESTful APIs for creating, fetching, updating and deleting users. Examples:

```bash
curl -X POST http://localhost:8080/users/ -d 'name=george' -d 'email=george@example.com'
curl -X POST http://localhost:8080/users/ -d 'name=bob' -d 'email=bob@example.com'
curl -X POST http://localhost:8080/users/ -d 'name=alice' -d 'email=alice@example.com'
```

```bash
curl -X GET http://localhost:8080/users/ | json_pp
```

```bash
curl -X PUT http://localhost:8080/users/40f8d096a3777232204cb3f796c577b7 -d 'name=james' -d 'email=george@example.com'
curl -X GET http://localhost:8080/users/ | json_pp
```


```bash
curl -X DELETE http://localhost:8080/users/40f8d096a3777232204cb3f796c577b7
curl -X GET http://localhost:8080/users/ | json_pp
```

Implementation notes
=======
The goal was to write an API server using only the Go standard library; therefore the endpoint handler code is a bit messy.  For a real project, using a 3rd party routing library (e.g. [gorilla/mux](https://github.com/gorilla/mux)) should make the code more maintainable.

The server's internal data APIs are packaged separately (package `model`). This allows us to easily import them into other codebases, which is useful in a service-oriented backend architecture (in this example, though, it doesn't matter because we're using an in-memory data store which can't be shared between services; in a real project we'd use an external store).

The server's in-memory data store is implemented using a Go map. To make it thread safe, access is synchronized using mutexes.

Possible improvements
=======

* Use a networked data store (e.g. Redis or an RDS)
* Implement basic unit tests
* Implement a hypermedia format (e.g. Siren) for the API JSON responses to make it easier to consume them using a standard parser
* Implement authentication using [JSON Web Tokens](https://jwt.io/) or similar
