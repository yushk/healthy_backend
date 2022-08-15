
SERVICES = apiserver manager
# Docker informations
# docker push docker.pkg.gitee.com/healthy_backend
DOCKER_REPO := 192.168.1.118:5000/
DOCKER_NAMESPACE := healthy
DOCKER_TAG       := latest

##### =====> build <===== #####
.PHONY: build
build: $(SERVICES)
	@echo "building completed!"

.PHONY: $(SERVICES)
$(SERVICES):
	@echo building $@ ...
	@SERVICE_NAME=$@ make .build;

.build:
	make -C $(SERVICE_NAME)

##### =====> tidy <===== #####
.PHONY: tidy
tidy: $(addprefix tidy-, $(SERVICES))
	@echo "go mod tidy completed!"

$(addprefix tidy-, $(SERVICES)):
	@SERVICE_NAME=$(subst tidy-,,$@) make .tidy

.tidy:
	make tidy -C ./$(SERVICE_NAME)

##### =====> gen <===== #####
.PHONY: gen
gen: $(addprefix gen-, $(SERVICES)) gen-pkg
	@echo "generate code completed!"

$(addprefix gen-, $(SERVICES)):
	@SERVICE_NAME=$(subst gen-,,$@) make .gen

gen-pkg:
	@SERVICE_NAME=$(subst gen-,,$@) make .gen

.gen:
	make gen -C ./$(SERVICE_NAME)

##### =====> docker <===== #####
.PHONY: docker
docker: $(addprefix docker-build-, $(SERVICES))
	@echo "docker building completed!"

$(addprefix docker-build-, $(SERVICES)):
	@SERVICE_NAME=$(subst docker-build-,,$@) \
		make .docker-build;

.docker-build:
	docker build \
		-f $(SERVICE_NAME)/Dockerfile.dev \
		--label=$(DATE) \
		-t $(DOCKER_REPO)$(DOCKER_NAMESPACE)/$(SERVICE_NAME):$(DOCKER_TAG) .
	docker push $(DOCKER_REPO)$(DOCKER_NAMESPACE)/$(SERVICE_NAME):$(DOCKER_TAG)
	docker rmi $(DOCKER_REPO)$(DOCKER_NAMESPACE)/$(SERVICE_NAME):$(DOCKER_TAG)

##### =====> clean <===== #####
clean: docker-clean
	for i in $(SERVICES); do rm -rf ./$${i}/bin; done
	make clean -C ui

docker-clean:
	docker rmi $$(docker images | grep "none" | awk '{print $$3}')


##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Internals <===== #####

# 版本号 v1.0.3-6-g0c2b1cf-dev
# 1、6:表示自打tag v1.0.3以来有6次提交（commit）
# 2、g0c2b1cf：g 为git的缩写，在多种管理工具并存的环境中很有用处
# 3、0c2b1cf：7位字符表示为最新提交的commit id 前7位
# 4、如果本地仓库有修改，则认为是dirty的，则追加-dev，表示是开发版：v1.0.3-6-g0c2b1cf-dev
VERSION          := $(shell git describe --tags --always --dirty="-dev")

# 时间
DATE             := $(shell date -u '+%Y-%m-%d-%H%M')
