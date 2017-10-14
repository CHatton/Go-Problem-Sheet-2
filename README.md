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
> curl http://localhost:9999
```

And you should see an output like this

```bash
StatusCode        : 200
StatusDescription : OK
Content           : Guessing Game
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 13
                    Content-Type: text/plain; charset=utf-8
                    Date: Sat, 14 Oct 2017 15:00:07 GMT

                    Guessing Game
Forms             : {}
Headers           : {[Content-Length, 13], [Content-Type, text/plain; charset=utf-8], [Date, Sat, 14 Oct 2017 15:00:07 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 13
```

