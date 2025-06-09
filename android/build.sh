#!/bin/sh
# Build Govinci Android AAR using gomobile
set -e

if ! command -v gomobile >/dev/null; then
  echo "gomobile is required. Install with 'go install golang.org/x/mobile/cmd/gomobile@latest'" >&2
  exit 1
fi

gomobile bind -target=android -tags govinci -o app/libs/govinci.aar ..
