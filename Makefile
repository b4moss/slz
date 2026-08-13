# Thin entrypoint for this repository.
# Product repos should include ruleset.mk from their own Makefile instead
# of merging this file. See https://github.com/b4moss/repo-ruleset/issues/2

SHELL := /bin/bash
include ruleset.mk
.DEFAULT_GOAL := ruleset-help
