define prepare_build_vars
	$(eval DATE_FLAG := -X 'main.date=$(shell date)')
    $(eval VERSION_FLAG=-X 'main.version=$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))')
    $(eval COMMIT_FLAG=-X 'main.commit=$(shell git rev-parse HEAD)')
endef

build/docker:
	$(call prepare_build_vars)
	go build -a -tags "osusergo netgo" --ldflags "${DATE_FLAG} ${VERSION_FLAG} ${COMMIT_FLAG} -extldflags '-static'" -o build/jwt-server ./main.go

build/local:
	$(call prepare_build_vars)
	go build -a --ldflags "${DATE_FLAG} ${VERSION_FLAG} ${COMMIT_FLAG}" -o build/jwt-server ./main.go