#!/bin/bash

# 定义输出目录
OUTPUT_DIR="build"
mkdir -p "$OUTPUT_DIR"

# 定义目标平台
declare -a TARGETS=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
  "windows/386"
)

# 编译每个目标平台
all_success=true
for target in "${TARGETS[@]}"; do
  # 使用parameter expansion分割target
  os="${target%%/*}"
  arch="${target#*/}"

  echo "Building for $os/$arch..."

  # 执行go build命令并检查是否成功
  if GOOS="$os" GOARCH="$arch" go build -o "$OUTPUT_DIR/mcp-ip_$os""_$arch"; then
    echo "Successfully built for $os/$arch."
  else
    echo "Failed to build for $os/$arch." >&2
    all_success=false
  fi
done

if $all_success; then
  echo "Build completed successfully. Output files are in $OUTPUT_DIR/"
else
  echo "Build completed with some failures. Check the output for details." >&2
fi