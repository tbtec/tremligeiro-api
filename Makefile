BINARY_NAME=tremligeiro
AWS_EKS_CLUSTER_NAME=tremligeiro-eks-cluster

run:
	go run cmd/main.go

test:
	go test -cover ./internal/... -coverpkg ./...

pre-build:
	go mod download
	go mod verify
	go mod tidy

build:
	go build -o bin/${BINARY_NAME} -ldflags="-s -w" -tags appsec cmd/main.go

build-ci:
	go build -o bin/${BINARY_NAME} -ldflags="-s -w" -tags appsec cmd/main.go

build-docker:
	docker build -t tbtec/tremligeiro:1.0.0 .	

run-docker:
	docker run -p 8080:8080 tbtec/tremligeiro:1.0.0 --env-file .env

run-compose:
	docker compose up

run-compose-enviroment:
	docker compose -f docker-compose-enviroment.yaml up

docker-push:
	docker push tbtec/tremligeiro:1.0.0

kube-config:
#	eval $(minikube docker-env)
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/postgresql/postgresql-base.yaml 
	kubectl apply -f k8s/postgresql/postgresql.yaml 

kube-deploy:
	kubectl apply -f k8s/configmap.yaml
	kubectl apply -f k8s/secret.yaml
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/ingress.yaml
	kubectl apply -f k8s/hpa.yaml

kube-deploy-eks:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/configmap.yaml
	kubectl apply -f k8s/secret.yaml
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/ingress.yaml
	kubectl apply -f k8s/hpa.yaml


kube-eks-connect:
	aws eks update-kubeconfig --name ${AWS_EKS_CLUSTER_NAME} --region us-east-1
