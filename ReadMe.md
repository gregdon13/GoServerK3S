# Go Server k3s

## Do yo have Go installed?

No? 

> https://golang.org/doc/install

The one you want is go1.16.3.linux-armv6l

when finished, make sure you're careful to checkout the following:

> The package installs the Go distribution to /usr/local/go. 
> The package should put the /usr/local/go/bin directory in your PATH environment variable. 
> You may need to restart any open Terminal sessions for the change to take effect.

Do a `go version` at a shell prompt, to see what version of Go you just installed.

following command 

`go run server.go`

 You should be able to see “Starting server at port 8080”

Now the next step is to build our go project into an executable. 
Run the following command at the root directory of the project(where server.go occurs) 

`GOOS=linux GOARCH=arm CGO_ENABLED=0 go build .`

(notice the period at the end of that line ^^^^)

You should be able to see a newly created file with no extension.

build a docker image

https://hub.docker.com get an account here.

```
sudo docker build . -t xt0fer/go-service
sudo docker run xt0fer/go-service
sudo docker login -u xt0fer
sudo docker push xt0fer/go-service
```

```
sudo kubectl create -f microsvc.yaml

sudo kubectl get pods
sudo kubectl expose deployment go-server-deployment --type=NodePort 

# WAIT! should you need to re-run create or expose, you have to delete first
sudo kubectl delete deployment go-server-deployment 
sudo kubectl delete service go-server-deployment 
```

This command will expose our deployment to the outer world on a random port

```
sudo kubectl get svc
```
