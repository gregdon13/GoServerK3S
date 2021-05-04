clean:
	go clean .

server:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build .

drop:
	sudo kubectl delete deployment go-server-deployment 
	sudo kubectl delete service go-server-deployment 

docker:
	sudo docker build . -t xt0fer/go-service
	sudo docker login -u xt0fer
	sudo docker push xt0fer/go-service

k3s:
	sudo kubectl create -f microsvc.yaml
	sudo kubectl expose deployment go-server-deployment --type=NodePort 
	sudo kubectl get all
