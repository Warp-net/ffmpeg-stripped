#!/bin/sh
set -eu

cd "$(dirname "$0")/.."
mkdir -p bin

# -n omits the source filename and timestamp for reproducible archives.
for os in linux darwin windows; do
    source_os=$os
    if [ "$os" = darwin ]; then source_os=macos; fi
    for arch in amd64 arm64; do
        source_arch=$arch
        if [ "$arch" = amd64 ]; then source_arch=x86_64; fi
        suffix=
        if [ "$os" = windows ]; then suffix=.exe; fi
        gzip -n -9 -c "$source_os-$source_arch/ffmpeg$suffix" > "bin/ffmpeg-$os-$arch$suffix.gz"
    done
done
