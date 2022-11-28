Branch=$(shell git symbolic-ref --short -q HEAD)
Commit=$(shell git rev-parse --short HEAD)
Date=$(shell git log --pretty=format:%cd $(Commit) -1) 
Author=$(shell git log --pretty=format:%an $(Commit) -1)
shortDate=$(shell git log -1 --format="%at" | xargs -I{} date -d @{} +%Y%m%d)
Email=$(shell git log --pretty=format:%ae $(Commit) -1)
Ver=$(shell echo $(Branch)-$(Commit)-$(shortDate))
GoVersion=$(shell go version )

.PHONY: statik
statik:
	go install github.com/rakyll/statik@latest
	statik -f -src=web

.PHONY: compile
compile: statik
	GOOS=linux go build -a -installsuffix cgo \
	-ldflags "-X 'github.com/Ankr-network/uscan/cmd.Branch=$(Branch)' \
	-X 'github.com/Ankr-network/uscan/cmd.Commit=$(Commit)' \
	-X 'github.com/Ankr-network/uscan/cmd.Date=$(Date)' \
	-X 'github.com/Ankr-network/uscan/cmd.Author=$(Author)' \
	-X 'github.com/Ankr-network/uscan/cmd.Email=$(Email)' \
	-X 'github.com/Ankr-network/uscan/cmd.GoVersion=$(GoVersion)'" -o bin/uscan

.PHONY: perf
perf: compile
	bin/uscan --config .uscan.yaml 

.PHONY: race
race:
	go run -race main.go --config .uscan.yaml 

.PHONY: start 
start: compile
	bin/uscan --config .uscan.yaml 

.PHONY: init
init:
	@echo "Init data dir...."
	mkdir $(HOME)/uscan-deploy  &&  cp docker-compose.yaml $(HOME)/uscan-deploy/
	@echo "build docker image..."
	docker build -t uscan:latest .
	docker image prune -f --filter label=stage=uscan-builder
	@echo "init finish! Please go to $(HOME)/uscan-deploy"


.PHONY: build
build: 
	docker build \
	--build-arg GoVersion='$(GoVersion)' \
	--build-arg Branch='$(Branch)' \
	--build-arg Commit='$(Commit)' \
	--build-arg Date='$(Date)' \
	--build-arg Author='$(Author)' \
	--build-arg Email='$(Email)' \
	-t ankrnetwork/uscan:$(Ver) .
	docker image prune -f --filter label=stage=builder

.PHONY: release
release: build
	docker push ankrnetwork/uscan:$(Ver)
	docker tag ankrnetwork/uscan:$(Ver) ankrnetwork/uscan:latest
	docker push ankrnetwork/uscan:latest



