#!/bin/bash

# License block to be prepended to file
read -r -d '' LICENSE <<EndOfLicense
// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
EndOfLicense

if [ -z "$1" ]; then
	echo "missing filename argument"
	echo "usage: prependlicense.sh [filename]"
	exit 1
fi

# Prepend license block to file
echo -e "$LICENSE\n" | cat - $1 > .prependlicense && mv .prependlicense $1
