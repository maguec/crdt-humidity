#!/bin/bash

go build humidity.go
go build humidity-watcher.go

zip crdt-humidity.zip humidity humidity-watcher

cd docs &&
dot -T png diagram.dot > diagram.png

