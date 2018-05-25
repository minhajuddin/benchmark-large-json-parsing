#!/bin/sh

cd "$(dirname "$0")"

cmake -H. -B_builds -DCMAKE_BUILD_TYPE=Release
cmake --build _builds --config Release

