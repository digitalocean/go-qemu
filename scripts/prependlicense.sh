#!/bin/bash

# License block to be prepended to file
BASEDIR=$(dirname "$0")
LICENSE=$(cat $BASEDIR/license.txt)

if [ -z "$1" ]; then
	echo "missing filename argument"
	echo "usage: prependlicense.sh [filename]"
	exit 1
fi

# Prepend license block to file
echo -e "$LICENSE\n" | cat - $1 > .prependlicense && mv .prependlicense $1
