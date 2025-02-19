# Makefile

# 检测操作系统类型
ifeq ($(OS),Windows_NT)
	DETECTED_OS := Windows
	# Windows specific commands
	SHELL := cmd.exe
	RM_CMD := rd /s /q
	MKDIR_CMD := mkdir
	NULL_DEV := 2>nul
	PATH_SEP := \\
	PROTO_FILES := $(shell powershell -Command "Get-ChildItem -Path './proto' -Filter '*.proto' -Recurse | Select-Object -ExpandProperty FullName | ForEach-Object { $$_.Replace('\', '/')}")
	WHICH_CMD := where
	SCRIPT_EXT := .bat

else
	DETECTED_OS := $(shell uname -s)
	# Unix-like systems commands
	RM_CMD := rm -rf
	MKDIR_CMD := mkdir -p
	NULL_DEV := 2>/dev/null
	PATH_SEP := /
	PROTO_FILES := $(shell find ./proto -name "*.proto")
	WHICH_CMD := which
	SCRIPT_EXT := .sh
endif

# 项目配置
PROJECT_ROOT ?= ..$(PATH_SEP)farms-rpc-server$(PATH_SEP)app
SERVICES ?= agent farm manage user
#SERVICES ?= mqtt
PROTO_ROOT ?= .$(PATH_SEP)proto
PROTO_THIRD_PARTY ?= $(PROTO_ROOT)$(PATH_SEP)third_party
GENERATED_DIR ?= .$(PATH_SEP)proto-gen-go

# 从 .proto 文件路径中提取服务目录
PROTO_SERVICES := $(sort $(dir $(PROTO_FILES)))

# RPC 服务生成配置
ifeq ($(DETECTED_OS),Windows)
define generate_rpc
	echo "Generating RPC service: $(1)" && \
	cmd /c scripts\genrpc.bat $(1)
endef
else
define generate_rpc
	echo "Generating RPC service: $(1)" && \
	./scripts/genrpc.sh $(1)
endef
endif

# 生成所有 RPC 服务
.PHONY: genrpc
genrpc: check-tools
	@echo "Starting RPC service generation..."
	@$(foreach service,$(SERVICES),$(call generate_rpc,$(service));)
	@echo "RPC service generation completed."

# 检查必要工具是否安装
.PHONY: check-tools
check-tools:
	@$(WHICH_CMD) protoc $(NULL_DEV) || (echo "Error: protoc not found" && exit 1)
	@$(WHICH_CMD) protoc-gen-go $(NULL_DEV) || (echo "Error: protoc-gen-go not found" && exit 1)
	@$(WHICH_CMD) protoc-gen-go-grpc $(NULL_DEV) || (echo "Error: protoc-gen-go-grpc not found" && exit 1)
	@$(WHICH_CMD) goctl $(NULL_DEV) || (echo "Error: goctl not found" && exit 1)

# 动态创建目录结构
.PHONY: prepare
prepare:
ifeq ($(DETECTED_OS),Windows)
	@for %%d in ($(PROTO_SERVICES)) do ( \
		$(MKDIR_CMD) "$(GENERATED_DIR)$(PATH_SEP)%%~pd" $(NULL_DEV) || exit 0 \
	)
else
	@for service in $(PROTO_SERVICES); do \
		$(MKDIR_CMD) "$(GENERATED_DIR)/$(subst $(PROTO_ROOT)/,,$(patsubst %/,%,$$service))"; \
	done
endif

# 生成所有服务的 Go 代码
.PHONY: gengo
gengo: prepare
ifeq ($(DETECTED_OS),Windows)
	@echo "Generating Go code from proto files..."
	@for /r proto %%i in (*.proto) do (
			echo Processing file: %%i
			protoc --proto_path=proto \
			--go_out=paths=source_relative:proto-gen-go \
			--go-grpc_out=paths=source_relative:proto-gen-go %%i
		)
	@echo "Generation complete."
else
	@if [ -z "$(PROTO_FILES)" ]; then \
		echo "No .proto files found"; \
	else \
		for service in $(PROTO_SERVICES); do \
			echo "Generating code for: $$(basename $$service)" && \
			protoc \
			--proto_path=$(PROTO_ROOT) \
			--proto_path=$(PROTO_THIRD_PARTY) \
			--go_out=$(GENERATED_DIR) \
			--go_opt=paths=source_relative \
			--go-grpc_out=$(GENERATED_DIR) \
			--go-grpc_opt=paths=source_relative \
			$$service/*.proto; \
		done \
	fi
	@rm -rf $(GENERATED_DIR)/$(PROTO_ROOT)
endif

# 帮助信息
.PHONY: help
help:
	@echo "使用说明:"
	@echo "  make genrpc     - 生成所有配置的 RPC 服务"
	@echo "  make gengo      - 生成所有配置的 proto 文件"
	@echo "  make install-tools - 安装所需工具"
	@echo "  make clean      - 清理生成的文件"
	@echo ""
	@echo "可以通过环境变量自定义:"
	@echo "  SERVICES=\"service1 service2 ...\"  - 指定要生成的服务列表 (默认: game user)"
	@echo "  PROJECT_ROOT    - RPC 服务的基础目录 (默认: ../live-server-rpc/app)"

# 清理生成的文件
.PHONY: clean
clean:
	@echo "清理生成的文件..."
ifeq ($(DETECTED_OS),Windows)
	@powershell -Command "if (Test-Path '$(PROJECT_ROOT)') { Get-ChildItem -Path '$(PROJECT_ROOT)' -Filter 'pb' -Directory -Recurse | Remove-Item -Recurse -Force }"
else
	@find $(PROJECT_ROOT) -type d -name "pb" -exec rm -rf {} +
endif

# 安装依赖工具
.PHONY: install-tools
install-tools:
	@echo "Installing required tools..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/zeromicro/go-zero/tools/goctl@latest
