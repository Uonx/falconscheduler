scheduler:
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=$(OS_NAME) GOARCH=$(OS_ARCH) GOPROXY="https://goproxy.cn,direct" go build -mod=mod \
		--ldflags "-s -w \
        	-X console/cmd.BuildTime=$(NOW) \
        	-X console/cmd.PWD=$(PWD) \
        	-X console/cmd.BRANCH=$(CURRENT_BRANCH) \
        " \
	-o build/dist/falcon ./cmd/falcon