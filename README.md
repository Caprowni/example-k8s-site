# example-k8s-site

## Overview

This repo installs a Golang site into a Kubernetes cluster, the following tools are required in order to get this working:

- Helm installed
- Kubectl
- Minikube(For local stuff)
- Docker installed

If you have the tools installed then simply do the following to get it up and running:

- Clone the repository.
- minikube start
- Build the image like so:
    `docker build -t "go-site" .`
- helm upgrade --install go-site my\_web\_service/  --values my\_web\_service/values.yaml 
- eval $(minikube -p minikube docker-env) 
- minikube service go-site --url
