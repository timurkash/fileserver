# Makefile for releasing images
#
# The release version is controlled from pkg/version

NAME:=images
GOPRIVATE:=gitlab.mcsolutions.ru
GROUP:=find-psy
REGISTRY:=gitlab.mcsolutions.ru:4567
REPOSITORY:=find-psy/back/users/images
DOCKER_IMAGE_NAME:=$(REGISTRY)/$(REPOSITORY)
DOCKER_PUBLIC_IMAGE:=docker.io/timurkash/find-psy-back-users-images-v1
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"' | tr -d '[:space:]')
GIT_COMMIT:=$(shell git describe --dirty --always)
PORT:=3027
HELMREPO_URL:=https://timurkash.github.com/helm-example/
CHART:=sfpg
RELEASE_NAME:=back-users-images
STAGING:=$(GROUP)-staging
PRODUCTION:=$(GROUP)-production

version:
	@echo $(VERSION)

v:
	@echo $(VERSION)

vendor:
	export GOPRIVATE
	go mod init gitlab.mcsolutions.ru/find-psy/back/users/images
	go mod vendor

revendor:
	rm -rf vendor go.mod go.sum
	make vendor

init:
	git init
	git add .
#	git remote add origin git@gitlab.mcsolutions.ru:find-psy/back/users/images.git
	git remote add origin https://gitlab.mcsolutions.ru/find-psy/back/users/images.git
	git commit -m "Init"
	git push origin master

commit:
	git add .
	git commit -m "autocommit"
	git push -u origin master

next:
	git add .
	git commit -m "next"
	git push -u origin master

pull:
	git pull origin master

push:
	git push -u origin master

list:
	go list all -mod=mod -m

run:
	GO111MODULE=on go run -ldflags "-s -w -X gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/version.REVISION=$(GIT_COMMIT)" cmd/images/* --level=debug

test:
	GO111MODULE=on go test -v -race ./...

build:
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor  -ldflags "-s -w -X gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/images ./cmd/images/*

#build-charts:
#	helm lint charts/*
#	helm package charts/*
#
build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

run-container:
	@docker run -dp $(PORT):$(PORT) --name=images $(DOCKER_IMAGE_NAME):$(VERSION)

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest

push-container-public:
	docker tag  $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_PUBLIC_IMAGE):$(VERSION)
	docker tag  $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_PUBLIC_IMAGE):latest
	docker push $(DOCKER_PUBLIC_IMAGE):$(VERSION)

	docker push $(DOCKER_PUBLIC_IMAGE):latest

run-compose:
	docker-compose up -d

stop-container:
	@docker stop $(NAME)
	@docker ps

clear-container:
	@docker rm -f $(NAME)

net-compose:
	docker network create -d=bridge find-psy.net

screen:
	@echo Running $(NAME) binary
	screen -S $(NAME) ./bin/$(NAME)

#version-set:
#	@next="latest" && \
#	current="$(VERSION)" && \
#	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
#	sed -i '' "s/tag: $$current/tag: $$next/g" charts/images/values.yaml && \
#	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/images/Chart.yaml && \
#	sed -i '' "s/version: $$current/version: $$next/g" charts/images/Chart.yaml && \
#	sed -i '' "s/images:$$current/images:$$next/g" kustomize/deployment.yaml && \
#	echo "Version $$next set in code, deployment, chart and kustomize"
#
release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	GO111MODULE=on go get github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init -g server.go

rmi-win:
	@echo In Windows PowerShell
	#	docker rmi $(docker images --format "{{.Repository}}:{{.Tag}}" | findstr "$(DOCKER_IMAGE_NAME)")

rmi:
	docker rmi $(docker images | grep $(DOCKER_IMAGE_NAME))

images:
	docker images $(DOCKER_IMAGE_NAME)

logs:
	docker logs $(NAME)

in:
	docker exec -it $(NAME) sh

port-forward-staging:
	kubectl -n $(STAGING) port-forward $(kubectl get pods -n $(STAGING) -l app=$(RELEASE_NAME) -o name) 1$(PORT):$(PORT)

port-forward-production:
	kubectl -n $(PRODUCTION) port-forward $(kubectl get pods -n $(PRODUCTION) -l app=$(RELEASE_NAME) -o name) 1$(PORT):$(PORT)

template:
	helm repo add $(GROUP) $(HELMREPO_URL)
	helm repo update
	helm template $(RELEASE_NAME) $(GROUP)/$(CHART) -f=values.yaml > template.yaml

uninstall:
	helm uninstall $(RELEASE_NAME) -n $(STAGING)

install-istio:
	kubectl apply -f istio.yaml
