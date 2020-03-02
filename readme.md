# URL Shortener

## Running locally
This project uses `go mod`, so you'll have to make sure you've got `EXPORT GO111MODULE=ON` defined in your shell or `~/.bashrc`.
`make install`

## Architecture

#### shortener
`shortener` is the central domain. Most of the other packages rely on either the `shortener` interface or `URL` / `ShortCode` types. I tried to make this vaguely [hexagonal](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) (in other words, these central business rules don't rely on any other packages & connect to other packages via interfaces).

#### app
`app` contains a concrete implementation of the `shortener` interface. It connects to the `hasher` and `stores` interfaces.

#### stores
`stores` are where the persistance layer lives. As there's a `stores` interface, you can easily create new layers and actually create a useful database backend. I've only implemented an in-memory cache as it was easiest to get working.

#### hasher
`hasher` contains a way to generate random codes from the input URL. At the moment, it just hashes the string, but if you wanted you coudl change the algorithm. I was thinking you may also want to change the interface to allow a `store` to be passed in... so that you could check whether the url already exists in the database if you were to use a different algorithm.

#### drivers
There are currently two main ways to run the app:
1. API (`web` package)
2. CLI (`cli` package')
These both use the same `shortener` interface, so there is no duplication of buisness logic. Almost all of the code in the `web` and `cli` packages just deals with the specifics of their interfaces (eg. accessing data and returning error messages or results).

### Server Example

```
$ make run_server_redirect # or go run main.go -server true -redirect true


$ curl -X POST -H "Content-Type: application/json" \
 -d '{"URL": "https://example.com/foo"}' \
 localhost:8080/shorten
# => {"success":true,"error":"","redirectURL":"localhost:8080/Ti0-MV4cifgD"}


$ curl localhost:8080/Ti0-MV4cifgD
# => <a href="https://example.com/foo">Moved Permanently</a>.

```

### CLI Example

Currently I haven't implemented an actual _persistent_ store.. which is a bit lazy but seems to work for the time being.
So the CLI only half-works; you have to both add and retrieve at the same time.
If we were to add a store that *actually stores data persistently*, this would work a bit more smoothly.
```
$ go run main.go -shorten https://example.com/yolo 
# => Successfully shortened https://example.com/yolo to LpnX-1cqUSsL

$ go run main.go -shorten https://example.com/yolo -retrieve  LpnX-1cqUSsL -redirect
# => Successfully shortened https://example.com/yolo to LpnX-1cqUSsL
# => Successfully retrieved https://example.com/yolo from LpnX-1cqUSsL
# => ***OPENS BROWSER & OPENS https://example.com/yolo***

```

### Other things to explore
We've got a pretty minimal makefile
```
$ go run main.go -help
```

### Potential improvements
- It would have been good to know a bit more about the service's requirements; the amount of uptime needed, concurrent connections, peak load, etc.. all would have helped influence the design of this.
- We don't cleanly exit - it'd be better to be able to clean up database connections, etc, and have a nicer way to turn the server or database connection off.
- We don't actually properly use the `shortener` interface; I reckon that `cli` and `web` should rely on `shortener.Shortener` rather than `app.URLShortener` (the actual interface, rather than just an implementation of it).
Having said that, I'm unconvinced that we actually do *need* an interface here. If we're trying to keep things simple (as we should be), perhaps I should just remove the pointless `shortener.Shortener` interface alltogether as we're unlikely to make different implementations of the interface. Something to think about.
- the `shortener` / `app` distinction is perhaps a bit vague. Maybe I should totally remove the `shortener` interface and just move `URL` / `ShortCode` to a generic `models` package??
- This is probably over-engineered for a simple url-shortening service. Something with less complexity would definitely work (eg. only actual important interface is the `Store` one).
- The CLI only really works locally. This is fine if you're happy to SSH into the server and run the CLI there - but that's not a great solution. You obviously could also connect to the prod database (you'd have to have the connection / credentials stored locally), but that's also not a great idea. Ideally the CLI would just be a wrapper to the server.
- I'm not sure `hasher` is a good name. Maybe `tokenisation` would be better?
