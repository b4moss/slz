# GitHub repository ruleset helpers.
# Real work lives in scripts/; this file is a thin wrapper.
# Target/variable names are prefixed with ruleset- / RULESET_ so this
# file can be vendored via git subtree without colliding with host Makefiles.
#
# Product repos: append one line at the end of the host Makefile:
#   include ruleset.mk
#   # or, if vendored under a subdirectory:
#   # include path/to/ruleset.mk
#
# Do not set .DEFAULT_GOAL here — the host Makefile owns the default target.
# Recipes expect bash (host should set SHELL := /bin/bash if needed).

RULESET_ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
RULESET_SCRIPTS := $(RULESET_ROOT_DIR)/scripts

RULESET_REPO ?=
RULESET_BRANCH ?= main
RULESET_VISIBILITY ?= public
RULESET_CREATE_FLAGS ?=

.PHONY: ruleset-help ruleset-create ruleset-apply ruleset-check

ruleset-help: ## Show ruleset helper usage
	@printf '%s\n' \
		'Targets:' \
		'' \
		'  make ruleset-create RULESET_REPO=OWNER/NAME [RULESET_VISIBILITY=public] [RULESET_CREATE_FLAGS="--clone"]' \
		'      Create a GitHub repo and apply rulesets.' \
		'' \
		'  make ruleset-apply RULESET_REPO=OWNER/NAME' \
		'      Apply/update rulesets on an existing repo.' \
		'' \
		'  make ruleset-check RULESET_REPO=OWNER/NAME [RULESET_BRANCH=main]' \
		'      List rulesets and check which rules apply to RULESET_BRANCH.' \
		'' \
		'Notes:' \
		'  - GitHub Free (org): rulesets work on public repos only.' \
		'  - Requires gh (repo admin) and jq.' \
		'  - Namespaced as ruleset-* / RULESET_* for subtree-safe includes.'

ruleset-create: ## Create a GitHub repo and apply rulesets
	@if [[ -z "$(RULESET_REPO)" ]]; then \
		echo "error: RULESET_REPO=OWNER/NAME is required" >&2; \
		echo "example: make ruleset-create RULESET_REPO=my-org/new-app" >&2; \
		exit 1; \
	fi
	@$(RULESET_SCRIPTS)/create-repo-with-rulesets.sh "$(RULESET_REPO)" "--$(RULESET_VISIBILITY)" $(RULESET_CREATE_FLAGS)

ruleset-apply: ## Apply/update rulesets on an existing repo
	@if [[ -z "$(RULESET_REPO)" ]]; then \
		echo "error: RULESET_REPO=OWNER/NAME is required" >&2; \
		echo "example: make ruleset-apply RULESET_REPO=my-org/existing-app" >&2; \
		exit 1; \
	fi
	@$(RULESET_SCRIPTS)/apply-rulesets.sh --repo "$(RULESET_REPO)"

ruleset-check: ## List rulesets and check branch applicability
	@if [[ -z "$(RULESET_REPO)" ]]; then \
		echo "error: RULESET_REPO=OWNER/NAME is required" >&2; \
		echo "example: make ruleset-check RULESET_REPO=my-org/existing-app RULESET_BRANCH=main" >&2; \
		exit 1; \
	fi
	@echo "== ruleset list =="
	@gh ruleset list -R "$(RULESET_REPO)"
	@echo
	@echo "== ruleset check $(RULESET_BRANCH) =="
	@gh ruleset check "$(RULESET_BRANCH)" -R "$(RULESET_REPO)"
