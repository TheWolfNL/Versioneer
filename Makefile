REPO=$(shell git remote get-url origin | sed 's/https:\/\///;s/.*@//;s/.git$$//;s/:/\//' | tr '[:upper:]' '[:lower:]')
APP=$(shell basename -s .git `git remote get-url origin`)
PHONY: build run test vendor
build:
	docker run --rm \
	-e Repo="${REPO}" \
	-v ${PWD}:/go/src/${REPO} \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/${REPO} \
	goreleaser/goreleaser release --snapshot --skip-publish --rm-dist

run:
	./dist/${APP}_$(shell uname -s | tr '[:upper:]' '[:lower:]')_amd64/${APP}

test:
	go test -v ${REPO}/cmd

vendor:
	go mod vendor