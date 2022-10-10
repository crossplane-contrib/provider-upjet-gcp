# ====================================================================================
# Setup Project

PROJECT_NAME := provider-gcp
PROJECT_REPO := github.com/upbound/$(PROJECT_NAME)
PROJECT_VERSION_TAG_GROUP := gcp

GOLANGCILINT_VERSION ?= 1.50.0
GO_REQUIRED_VERSION ?= 1.19

export TERRAFORM_VERSION := 1.2.1
export TERRAFORM_PROVIDER_SOURCE := hashicorp/google
export TERRAFORM_PROVIDER_VERSION := 4.22.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-google
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX := https://releases.hashicorp.com/terraform-provider-google/$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_PROVIDER_REPO ?= https://github.com/hashicorp/terraform-provider-google
export TERRAFORM_DOCS_PATH ?= website/docs/r

PLATFORMS ?= linux_amd64 linux_arm64

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

#GO_REQUIRED_VERSION ?= 1.19
#GOLANGCILINT_VERSION ?= 1.50.0
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis
GO111MODULE = on
-include build/makelib/golang.mk

# ====================================================================================
# Setup Kubernetes tools

UP_VERSION = v0.13.0
UP_CHANNEL = stable
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

DOCKER_REGISTRY ?= upbound
IMAGES = provider-gcp
-include build/makelib/image.mk

# ====================================================================================
# Setup XPKG

XPKG_REGISTRY ?= xpkg.upbound.io
XPKG_ORG ?= upbound
XPKG_REPO ?= $(PROJECT_NAME)-staging

# ====================================================================================
# Targets

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_generated.deepcopy | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	UPBOUND_CONTEXT="local" $(GO_OUT_DIR)/provider --debug

# NOTE(hasheddan): we ensure up is installed prior to running platform-specific
# build steps in parallel to avoid encountering an installation race condition.
build.init: $(UP)

xpkg.build: $(UP) do.build.images
	@$(INFO) Building package $(PROJECT_NAME)-$(VERSION).xpkg for $(PLATFORM)
	@mkdir -p $(OUTPUT_DIR)/xpkg/$(PLATFORM)
	@$(UP) xpkg build \
		--controller $(BUILD_REGISTRY)/$(PROJECT_NAME)-$(ARCH) \
		--package-root ./package \
		--examples-root ./examples \
		--output ./_output/xpkg/$(PLATFORM)/$(PROJECT_NAME)-$(VERSION).xpkg || $(FAIL)
	@$(OK) Built package $(PROJECT_NAME)-$(VERSION).xpkg for $(PLATFORM)

build.artifacts.platform: xpkg.build

xpkg.push: $(UP)
	@$(INFO) Pushing package $(PROJECT_NAME)-$(VERSION).xpkg
	@$(UP) xpkg push \
		--package $(OUTPUT_DIR)/xpkg/linux_amd64/$(PROJECT_NAME)-$(VERSION).xpkg \
		--package $(OUTPUT_DIR)/xpkg/linux_arm64/$(PROJECT_NAME)-$(VERSION).xpkg \
		$(XPKG_REGISTRY)/$(XPKG_ORG)/$(XPKG_REPO):$(VERSION) || $(FAIL)
	@$(OK) Pushed package $(PROJECT_NAME)-$(VERSION).xpkg

xpkg.load: $(UP)
	@$(INFO) Loading package $(PROJECT_NAME)-$(VERSION).xpkg for $(PLATFORM) into Docker daemon
	@docker load -i $(OUTPUT_DIR)/xpkg/$(PLATFORM)/$(PROJECT_NAME)-$(VERSION).xpkg
	@$(OK) Loaded package $(PROJECT_NAME)-$(VERSION).xpkg for $(PLATFORM) into Docker daemon

.PHONY: cobertura submodules fallthrough run crds.clean

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json

$(TERRAFORM):
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init -upgrade > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)

generate.init: $(TERRAFORM_PROVIDER_SCHEMA)

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA)

# ====================================================================================
# Extract Terraform registry metadata
pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))" ]; then \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))"; \
	fi
	@git -C "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: pull-docs

.PHONY: pull-docs

# ====================================================================================

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

endef

# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

# NOTE(hasheddan): the build submodule currently overrides XDG_CACHE_HOME in
# order to force the Helm 3 to use the .work/helm directory. This causes Go on
# Linux machines to use that directory as the build cache as well. We should
# adjust this behavior in the build submodule because it is also causing Linux
# users to duplicate their build cache, but for now we just make it easier to
# identify its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

go.mod.cachedir:
	@go env GOMODCACHE

uptest: $(KIND) $(KUBECTL) $(HELM3) $(UP) $(KUTTL)
	@$(INFO) running uptest using kind $(KIND_VERSION)
	@./cluster/install_provider.sh || $(FAIL)
	@echo "$${UPTEST_EXAMPLE_VALUE_REPLACEMENTS}" > $(WORK_DIR)/replacements.yaml
	@KIND=$(KIND) KUBECTL=$(KUBECTL) KUTTL=$(KUTTL) go run github.com/upbound/official-providers/testing/cmd --data-source "$(WORK_DIR)/replacements.yaml" || $(FAIL)

uptest-local: $(KUBECTL) $(KUTTL)
	@$(INFO) running automated tests with uptest using current kubeconfig $(KIND_VERSION)
	@KUBECTL=$(KUBECTL) KUTTL=$(KUTTL) go run github.com/upbound/official-providers/testing/cmd || $(FAIL)

cluster_dump: $(KUBECTL)
	@mkdir -p ${DUMP_DIRECTORY}
	@$(KUBECTL) cluster-info dump --output-directory ${DUMP_DIRECTORY} --all-namespaces || true
	@$(KUBECTL) get managed -o yaml > ${DUMP_DIRECTORY}/managed.yaml || true
	@cat /tmp/automated-tests/case/*.yaml > ${DUMP_DIRECTORY}/kuttl-inputs.yaml

.PHONY: cobertura reviewable submodules fallthrough go.mod.cachedir go.cachedir
