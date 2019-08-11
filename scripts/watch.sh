#!/usr/bin/env bash

set -e

find . -name \*.go -print | entr -r go test 
