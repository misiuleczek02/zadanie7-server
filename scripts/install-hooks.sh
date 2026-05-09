#!/usr/bin/env bash

set -e
git config core.hooksPath .githooks
chmod +x .githooks/pre-commit .githooks/pre-push 2>/dev/null || true
echo "Git hooks installed (core.hooksPath=.githooks)"
