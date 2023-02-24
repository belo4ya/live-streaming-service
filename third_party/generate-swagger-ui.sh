#!/bin/sh

set -e

SWAGGER_UI_VERSION="v4.15.5"
SWAGGER_UI_GIT="https://github.com/swagger-api/swagger-ui.git"
SWAGGER_DIR="./swagger-ui"
CACHE_DIR="./.cache/swagger-ui/$SWAGGER_UI_VERSION"
OPENAPI_DIR="../api"

escape_str() {
  echo "$1" | sed -e 's/[]\/$*.^[]/\\&/g'
}

# do caching if there's no cache yet
if [[ ! -d "$CACHE_DIR" ]]; then
  mkdir -p "$CACHE_DIR"
  tmp="$(mktemp -d)"
  git clone --depth 1 --branch "$SWAGGER_UI_VERSION" "$SWAGGER_UI_GIT" "$tmp"
  cp -r "$tmp/dist/"* "$CACHE_DIR"
  cp -r "$tmp/LICENSE" "$CACHE_DIR"
  rm -rf "$tmp"
fi

# populate swagger.json
tmp="    urls: ["
for i in $(find "$OPENAPI_DIR" -name "*.swagger.json"); do
  escaped_proto_dir="$(escape_str "$OPENAPI_DIR/")"
  path="$(echo $i | sed -e "s/escaped_proto_dir//g")"
  tmp="$tmp{\"url\":\"$path\",\"name\":\"$path\"},"
done
# delete last characters from $tmp
tmp=$(echo "$tmp" | sed 's/.$//')
tmp="$tmp],"

# recreate swagger-ui, delete all except swagger.json
find "$SWAGGER_DIR" -type f -not -name "*.swagger.json" -delete
mkdir -p "$SWAGGER_DIR"
cp -r "$CACHE_DIR/"* "$SWAGGER_DIR"

# replace the default URL
line="$(cat "$SWAGGER_DIR/swagger-initializer.js" | grep -n "url" | cut -f1 -d:)"
escaped_tmp="$(escape_str "$tmp")"
sed -i'' -e "$line s/^.*$/$escaped_tmp/" "$SWAGGER_DIR/swagger-initializer.js"
rm -f "$SWAGGER_DIR/swagger-initializer.js-e"
