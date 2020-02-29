# URL Shortener

Brief: 

```
Write an API for a URL shortener that satisfies the following behaviour:

Accepts a URL to be shortened.
Generates a short URL for the original URL.
Accepts a short URL and redirects the caller to the original URL.

Bonus points
Comes with a CLI that can be used to call your service.

Things we'd like to see
Sound design approach that's not overly complicated or over-engineered.
Code that's easy to read and not too "clever".
Sensible tests in place
```

Server

```
curl -X POST -H "Content-Type: application/json" \
 -d '{"URL": "https://example.com/foo"}' \
 localhost:8080/shorten

curl localhost:8080/Ti0-MV4cifgD
```

CLI Example

Currently I haven't implemented an actual _persistent_ store.. which is a bit lazy but seems to work for the time being.
So the CLI only half-works; you have to both add and retrieve at the same time.

```
go run main.go -help
go run main.go -shorten https://example.com/yolo -retrieve  LpnX-1cqUSsL -redirect
```

TODO: Maybe remove external non-stdlib dependencies?


Some shortcomings
We don't cleanly exit - it'd be better to be able to clean up database connections, etc, and have a nicer way to turn the server off.
