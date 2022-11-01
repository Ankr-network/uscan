SHELL = /bin/bash

REPO ?= ankrnetwork/coqscan

.PHONY: docker-build
docker-build:
	@echo "build docker image"
	@BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD); \
	if [[ $$BRANCH_NAME == "develop" ]]; then \
		ENV="stage"; \
	elif [[ $$BRANCH_NAME == "main" ]]; then \
		ENV="prod"; \
	else \
		ENV="feat"; \
	fi; \
	docker build -t $(REPO):$$ENV .

.PHONY: docker-push
docker-push: docker-build
	@echo "tag & push image"
	@BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD); SHA1_SHORT=$(shell git rev-parse --short HEAD); \
	if [[ $$BRANCH_NAME == "develop" ]]; then \
		ENV="stage"; \
	elif [[ $$BRANCH_NAME == "master" ]]; then \
		ENV="prod"; \
	else \
		ENV="feat"; \
	fi;  \
	docker tag $(REPO):$$ENV  $(REPO):$$SHA1_SHORT; \
	docker push $(REPO):$$SHA1_SHORT; \
	docker push $(REPO):$$ENV;