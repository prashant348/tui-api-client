#!/bin/sh

echo "Running pre-commit checks..."

echo "1. Formatting code..."
gofmt -w .

echo "2. Tidying go modules..."
go mod tidy

echo "3. Verifying code..."
if ! go vet ./...; then
    echo "❌ Error: Code verification failed (go vet). Fix errors before committing."
    exit 1
fi

echo "4. Staging all changes..."
git add .

echo "✅ Ready to commit!"