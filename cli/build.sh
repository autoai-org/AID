#!/usr/bin/env bash
package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_name="cvpm"
platforms=("windows/amd64" "darwin/amd64" "linux/amd64" "linux/arm64" "linux/mips" "freebsd/amd64")
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    echo 'Building CVPM for '$GOOS'/'$GOARCH'...'
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o 'dist/'$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done