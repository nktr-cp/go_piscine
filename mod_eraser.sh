#!/bin/bash

DIR=${1:-.}

find "$DIR" -type f -name "go.mod" -exec rm -f {} \;

echo "All go.mod files have been deleted from $DIR"

