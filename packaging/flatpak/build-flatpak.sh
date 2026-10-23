#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 2 || $# -gt 3 ]]; then
  echo "Usage: $0 <linux-binary> <output.flatpak> [branch]" >&2
  exit 2
fi

binary="$(realpath "$1")"
output="$(realpath -m "$2")"
branch="${3:-stable}"
app_id="io.github.matheusrmello.attack-shark-x6"
runtime="org.gnome.Platform"
sdk="org.gnome.Sdk"
runtime_version="50"
script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
output_dir="$(dirname -- "$output")"
mkdir -p "$output_dir"
work_dir="$(mktemp -d "$output_dir/.flatpak-build.XXXXXX")"
repo_dir="$work_dir/repo"
build_dir="$work_dir/build"

cleanup() {
  rm -rf "$work_dir"
}
trap cleanup EXIT

if [[ ! -x "$binary" ]]; then
  echo "Executable not found: $binary" >&2
  exit 1
fi

flatpak build-init --arch=x86_64 "$build_dir" "$app_id" "$sdk//$runtime_version" "$runtime//$runtime_version" "$branch"
install -Dm755 "$binary" "$build_dir/files/bin/attack-shark-linux"
install -Dm644 "$script_dir/$app_id.desktop" "$build_dir/files/share/applications/$app_id.desktop"
install -Dm644 "$script_dir/attack-shark-x6.png" "$build_dir/files/share/icons/hicolor/512x512/apps/$app_id.png"
install -Dm644 "$script_dir/$app_id.metainfo.xml" "$build_dir/files/share/metainfo/$app_id.metainfo.xml"

flatpak build-finish \
  --no-exports \
  --command=attack-shark-linux \
  --require-version=1.4.0 \
  --share=ipc \
  --share=network \
  --socket=fallback-x11 \
  --socket=wayland \
  --device=all \
  --talk-name=org.freedesktop.Notifications \
  --talk-name=org.freedesktop.portal.* \
  --talk-name=org.kde.StatusNotifierWatcher \
  "$build_dir"

flatpak build-export --disable-sandbox --no-update-summary "$repo_dir" "$build_dir" "$branch"
rm -f "$output"
flatpak build-bundle \
  --runtime-repo=https://flathub.org/repo/flathub.flatpakrepo \
  "$repo_dir" "$output" "$app_id" "$branch"

echo "Built $output"
