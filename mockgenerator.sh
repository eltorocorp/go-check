#!/bin/bash

################################################################################
# This script searches for all files named *iface.go and uses mockgen to 
# generate mocks for the interfaces in each matching file.
# The resulting mock files are stored in mocks/original/path/to/fileiface.go
# The mocks directory is initially purged before any new mocks are generated.
################################################################################

# using this shell function rather than 'install' for better portability.
# https://stackoverflow.com/a/24666836/3434541
mkfileP() { mkdir -p "$(dirname "$1")" || return; touch "$1"; }

rm -drf mocks
files=$(find . -type f -name "*iface.go")
for file in $files; do
    mkfileP "mocks/$file"
    mockgen -source "$file" > "mocks/$file"
done