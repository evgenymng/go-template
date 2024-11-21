# Linters & Formatters

format:
    @goimports -l -w cmd internal
    @gofumpt -l -w cmd internal
    @golines -w -m 80 cmd internal

lint:
    @golangci-lint run ./cmd/... ./internal/...

gen-docs:
    @swag init \
        -d ./cmd/app \
        --collectionFormat multi \
        --parseInternal \
        -o ./docs \
        --outputTypes go \
        --packageName docs

format-docs:
    @swag fmt \
        -d ./cmd/app,./internal/app/routes

pre-commit: format format-docs gen-docs

# Tests

test:
    @go test ./cmd/... ./internal/... -v

# Hooks

update-tag:
    #!/bin/bash
    git_hash=$(git rev-parse --short HEAD)
    release_date=$(date +"%Y-%m-%d")
    sed -i.back "/GIT_HASH/d" .env
    sed -i.back "/RELEASE/d" .env
    echo "GIT_HASH=${git_hash}" | tee -a .env
    echo "RELEASE=${release_date}:${git_hash}" | tee -a .env

setup-hooks:
    #!/bin/sh
    if [ -f .git/hooks/post-merge ]; then
        sed -i '' "/just update-tag/d" .git/hooks/post-merge
    fi
    echo "just update-tag" | tee -a .git/hooks/post-merge > /dev/null

    if [ -f .git/hooks/post-commit ]; then
        sed -i '' "/just update-tag/d" .git/hooks/post-merge
    fi
    echo "just update-tag" | tee -a .git/hooks/post-commit > /dev/null

    chmod u+x .git/hooks/post-merge
    chmod u+x .git/hooks/post-commit
    chmod u+x .git/hooks/pre-commit
