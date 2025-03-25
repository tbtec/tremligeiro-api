# Trem Ligeiro Application

This applications is responsible for the restaurant "Trem Ligeiro" from Company "Trem Bão Tecnologia".

## How to run

### Requirements

-   [Go 1.22](https://golang.org/doc/install)
-   [Make](https://www.gnu.org/software/make/)
-   [Docker](https://docs.docker.com/engine/install/)
-   [Docker Compose](https://docs.docker.com/compose/install/https://docs.docker.com/compose/install/)
-   [Minikube](https://minikube.sigs.k8s.io/docs/) [local]
-   [Kubectl](https://kubernetes.io/docs/tasks/tools/) [local]

### Building the Application

Use this commands to build the application.

```bash
make pre-build
make build
```

### Run locally

Before run the application, it's needed ro run docker compose in another terminal using this command.

```bash
make run-compose
```
Then, starts applications using:

```bash
make run
```

## Deploy at AWS EKS Cluster

Pre-requisite: Needed AWS Credentials on local context.

Use this commands in this order:
```bash
make kube-eks-connect
```
Then:
```bash
make kube-kube-deploy-eks
```
To destroy, use this command:
```bash
make kube-kube-deploy-eks-destroy
```

### Guides
More Guides is available at [Guides](/docs/guides/).

## Clean Architecture at Application

![Trem Ligeiro Clean Architecture](/docs/diagrams/trem-ligeiro-ca-CleanArch.jpg)

## Infrastructure

![Trem Ligeiro Infraestructure](/docs/diagrams/trem-ligeiro-ca-ArchTremLigeiro.jpg)
