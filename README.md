# Go Problem Sheet 2
My name is Cian Hatton, I'm currently in my 3rd year studying Software Development at [GMIT](http://www.gmit.ie/).

I created this repository to hold my solutions to this [problem sheet](https://github.com/data-representation/data-representation.github.io/blob/master/problems/go-web-applications.md) on creating web applications in Go.

This is problem sheet 2 for the [Data Representation and Querying](https://data-representation.github.io/) module.

All solutions are written in the [Go](https://golang.org/) programming language.

In order to run the files, first clone the repository.

```bash
> git clone https://github.com/CHatton/Go-Problem-Sheet-2.git
```

Make sure that you have the [Go compiler](https://golang.org/dl/)  installed.

Navigate to the folder of the file you want to run.

```bash
> cd src
```

You can run the program by either first building it.

```bash
> go build <file-name>
```

And then running the executable.

```bash
> ./<file-name>
```

Or you can simply use the command

```bash
> go run <file-name>
```

In this case, you can provide an optional port that the server will listen on. The deafult is 7777.

```bash
> go run main.go 9999
```

The server should now be running. In order to make a http request using curl, first make sure you have [curl](https://curl.haxx.se/download.html) downloaded and installed.

When you do, you can use the command

```bash
> curl -v http://localhost:9999
```

And you should see an output like this

```bash
 Rebuilt URL to: http://localhost:9999/
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9999 (#0)
> GET / HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Length: 1234
< Content-Type: text/html; charset=utf-8
< Last-Modified: Sat, 21 Oct 2017 18:02:45 GMT
< Date: Sat, 21 Oct 2017 19:13:32 GMT
< 
<!DOCTYPE html>
<html lang="en">
<!-- html will appear here -->
* Connection #0 to host localhost left intact

```
The 3 different endpoints in this application are

```bash
/
/guess/
/newgame/
```
