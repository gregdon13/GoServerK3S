# Go Server k3s

## Do you have Go installed?

No? 

> https://golang.org/doc/install

The one you want is go1.16.3.linux-armv6l

when finished, make sure you're careful to checkout the following:

> The package installs the Go distribution to /usr/local/go. 
> The package should put the /usr/local/go/bin directory in your PATH environment variable. 
> You may need to restart any open Terminal sessions for the change to take effect.

Do a `go version` at a shell prompt, to see what version of Go you just installed.


## Testing the server 

`go run go-server.go`

 You should be able to see “Starting server at port 8080”, then Control-C to break.

Now the next step is to build our go project into an executable for inside of a docker container. 
Run the following command at the root directory of the project(where go-server.go occurs).

`GOOS=linux GOARCH=arm CGO_ENABLED=0 go build .`

(notice the period at the end of that line ^^^^)

You should be able to see a newly created file with no extension.

## Build a docker image

https://hub.docker.com get an account here.
I am using the xt0fer account for these examples. 
You want to change xt0fer everywhere to your
dockerhub account name.

```
sudo docker build . -t xt0fer/go-service
sudo docker run xt0fer/go-service
sudo docker login -u xt0fer
sudo docker push xt0fer/go-service
```

## Wrangle the Docker Image into Kubernetes

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
sudo kubectl get all
```

The key is to find the PORT of the new web service. After the command above, look in the results for somethng like:

```
service/go-server-deployment   NodePort    10.43.5.148     <none>        8080:31807/TCP   2s
```

See the 2nd port in that line?? The `31807` one?
Thats the port that K3s is running our service on.

So, we should be able to 

```
http://raspberrypi.local:31807
-or-
http://raspberrypi.local:31807/app
```

and see the results.

# Where I'm at
* go is functioning on the pi
* now to set up the docker on it
