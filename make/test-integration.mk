BUILD_TAGS_K8S_INTEGRATION := k8s_integration
BUILD_TAGS_E2E := e2e integration

BUILD_TAGS_ALL := "$(BUILD_TAGS_K8S_INTEGRATION) $(BUILD_TAGS_E2E)"

# This is statically specified in the vagrant configuration
# todo @troian check it still necessary
KUBE_NODE_IP ?= 172.18.8.101
###############################################################################
###                           Integration                                   ###
###############################################################################

INTEGRATION_VARS := TEST_INTEGRATION=true

.PHONY: test-e2e-integration
test-e2e-integration:
	# Assumes cluster created and configured:
	# ```
	# KUSTOMIZE_INSTALLS=akash-operator-inventory make kube-cluster-setup-e2e
	# ```
	$(KIND_VARS) $(INTEGRATION_VARS) go test -count=1 -mod=readonly -p 4 -tags "e2e" -v ./integration/... -run TestIntegrationTestSuite -timeout 1500s

.PHONY: test-e2e-integration-k8s
test-e2e-integration-k8s:
	$(INTEGRATION_VARS) KUBE_NODE_IP="$(KUBE_NODE_IP)" KUBE_INGRESS_IP=127.0.0.1 KUBE_INGRESS_PORT=10080 go test -count=1 -mod=readonly -p 4 -tags "e2e $(BUILD_MAINNET)" -v ./integration/... -run TestIntegrationTestSuite

.PHONY: test-query-app
test-query-app:
	 $(INTEGRATION_VARS) $(KIND_VARS) go test -mod=readonly -p 4 -tags "$(BUILD_TAGS_E2E)" -v ./integration/... -run TestQueryApp

.PHONY: test-k8s-integration
test-k8s-integration:
	# Assumes cluster created and configured:
	# ```
	# KUSTOMIZE_INSTALLS=akash-operator-inventory make kube-cluster-setup-e2e
	# ```
	go test -count=1 -v -tags "$(BUILD_TAGS_K8S_INTEGRATION)" ./pkg/apis/akash.network/v2beta2
	go test -count=1 -v -tags "$(BUILD_TAGS_K8S_INTEGRATION)" ./cluster/kube


###############################################################################
###                           Misc tests                                    ###
###############################################################################

.PHONY: shellcheck
shellcheck:
	docker run --rm \
	--volume ${PWD}:/shellcheck \
	--entrypoint sh \
	koalaman/shellcheck-alpine:stable \
	-x /shellcheck/script/shellcheck.sh

.PHONY: test
test:
	$(GO) test -tags=$(BUILD_MAINNET) -timeout 300s ./...

.PHONY: test-nocache
test-nocache:
	$(GO) test -tags=$(BUILD_MAINNET) -count=1 ./...

.PHONY: test-full
test-full:
	$(GO) test -tags=$(BUILD_TAGS) -race ./...

.PHONY: test-coverage
test-coverage: $(AP_DEVCACHE)
	./script/codecov.sh "$(AP_DEVCACHE_TESTS)" $(BUILD_TAGS_ALL)

.PHONY: test-vet
test-vet:
	$(GO) vet ./...
