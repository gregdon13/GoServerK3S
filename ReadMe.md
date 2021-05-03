# Go Server k3s

## Do yo have Go installed?

No?

> https://golang.org/doc/install

when finished, make sure you're careful to checkout the following:

> The package installs the Go distribution to /usr/local/go. 
> The package should put the /usr/local/go/bin directory in your PATH environment variable. 
> You may need to restart any open Terminal sessions for the change to take effect.

Do a `go version` at a shell prompt, to see what version of Go you just installed.

following command 

`$ go run server.go`

 You should be able to see “Starting server at port 8080”

Now the next step is to build our go project into an executable. 
Run the following command at the root directory of the project(where server.go occurs) 

`$ GOOS=linux GOARCH=arm CGO_ENABLED=0 go build .`

(notice the period at the end of that line ^^^^)

You should be able to see a newly created file with no extension.

> $ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

build a docker image

https://hub.docker.com get an account here.


> $ kubectl create -f microsvc.yaml
