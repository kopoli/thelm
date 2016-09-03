#!/bin/sh

usage() {
    cat <<EOF
Usage: $0 [build | release]

Creates a build or a release.
EOF
    exit 1
}

prog=$(basename $(pwd))
archs="amd64"
oss="linux windows"
extra_files="README.md LICENSE"

genversion() {
    git describe --always --tags --dirty
}

compile() {
    go build -v -ldflags "-X main.timestamp=$(date --rfc-3339=seconds | tr ' ' '_') -X main.version=$(genversion)"
}

release() {
    version=$(genversion)
    for arch in $archs; do
        for os in $oss; do
            suffix=""
            echo "Generating a release for $os/$arch"
            (
                export GOARCH=$arch
                export GOOS=$os
                compile
            ) || exit 1
            test "$os" = "windows" && suffix=".exe"
            bsdtar --format zip -s ",^,$prog-$version-$arch-$os/," -cf $prog-$version-$arch-$os.zip \
                   $prog$suffix $extra_files
            done
    done
}

case "$1" in
     build)
         compile
         ;;
     release)
         release
         ;;
     *)
         usage
         ;;
esac
