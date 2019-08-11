
Hello world app in golang that calls to other hello world apps

```
go build .
```

Open 3 terminals, so you can run 3 copies of the server on different ports, where #0 calls #1, #1 calls #2, and #2 calls #0.

```
./recurse-world -port 8080 -upstream localhost:8081 -server-name zero
./recurse-world -port 8081 -upstream localhost:8082 -server-name one
./recurse-world -port 8082 -upstream localhost:8080 -server-name two
```

When you do a curl, you'll see:

```
% curl 'localhost:8080/hello/world/hello/world/meow/echo?foo=bar'
2018-07-19T01:28:07.851351-07:00 zero /hello/world/hello/world/meow/echo
2018-07-19T01:28:07.854602-07:00 one /world/hello/world/meow/echo
2018-07-19T01:28:07.856630-07:00 two /hello/world/meow/echo
2018-07-19T01:28:07.858659-07:00 zero /world/meow/echo
2018-07-19T01:28:07.859904-07:00 one /meow/echo
echo says: foo=bar
```
