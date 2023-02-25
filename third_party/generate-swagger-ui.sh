#!/bin/sh

set -e

SWAGGER_UI_VERSION="v4.15.5"
SWAGGER_UI_GIT="https://github.com/swagger-api/swagger-ui.git"
SWAGGER_DIR="./swagger-ui"
CACHE_DIR="./.cache/swagger-ui/$SWAGGER_UI_VERSION"
OPENAPI_URL="/doc/openapi.json"

escape_str() {
  echo "$1" | sed -e 's/[]\/$*.^[]/\\&/g'
}

# do caching if there's no cache yet
if [ ! -d "$CACHE_DIR" ]; then
  mkdir -p "$CACHE_DIR"
  tmp="$(mktemp -d)"
  git clone --depth 1 --branch "$SWAGGER_UI_VERSION" "$SWAGGER_UI_GIT" "$tmp"
  cp -r "$tmp/dist/"* "$CACHE_DIR"
  cp -r "$tmp/LICENSE" "$CACHE_DIR"
  rm -rf "$tmp"
fi

# recreate dist dir
rm -rf "$SWAGGER_DIR"
mkdir -p "$SWAGGER_DIR"
cp -r "$CACHE_DIR/"* "$SWAGGER_DIR"

# replace the default URL
line="$(grep <"$SWAGGER_DIR/swagger-initializer.js" -n "url" | cut -f1 -d:)"
url="    url: \"$OPENAPI_URL\","
escaped_tmp="$(escape_str "$url")"
sed -i'' -e "$line s/^.*$/$escaped_tmp/" "$SWAGGER_DIR/swagger-initializer.js"
rm -f "$SWAGGER_DIR/swagger-initializer.js-e"
