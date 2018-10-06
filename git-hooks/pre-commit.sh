#!/bin/sh

gometalinter         \
  --vendor           \
  --disable-all      \
  --enable=gofmt     \
  --enable=golint    \
  --enable=goimports \
  --enable=vetshadow \
  --enable=misspell  \
  --enable=vet       \
  --min-confidence=1 ./...

exit $?
