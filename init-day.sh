#!/usr/bin/env bash
set -euo pipefail

# Check if a day number is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <day-number>"
  exit 1
fi

DAY_NUM=$(printf "%02d" $1)
DAY_DIR="internal/day$DAY_NUM"
TEMPLATE_DIR="internal/dayXX"

# Copy template directory
cp -r $TEMPLATE_DIR $DAY_DIR

# Rename the Go file in the new day directory
GO_FILE="$DAY_DIR/day$DAY_NUM.go"
mv $DAY_DIR/dayXX.go $GO_FILE

# Update package name in the Go file
if [[ "$OSTYPE" == "darwin"* ]]; then
  # macOS requires an explicit argument for in-place editing
  sed -i '' "s/package dayXX/package day$DAY_NUM/" $GO_FILE
else
  # Linux
  sed -i "s/package dayXX/package day$DAY_NUM/" $GO_FILE
fi

# Add import statement to main.go
IMPORT_STATEMENT="\"github.com/ccitro/advent-2023-go/$DAY_DIR\""
awk -v s="$IMPORT_STATEMENT" '/import \(/ { print; print "\t"s; next }1' main.go > tmp.go && mv tmp.go main.go

# Add entry to the dayMethods map
MAP_ENTRY="\"day$DAY_NUM\": {LoadPuzzle: day$DAY_NUM.LoadPuzzle, Part1: day$DAY_NUM.Part1, Part2: day$DAY_NUM.Part2},"
awk -v m="$MAP_ENTRY" '/var dayMethods = map\[string\]DayMethods{/ { print; print "\t"m; next }1' main.go > tmp.go && mv tmp.go main.go

# Format main.go to order the imports
go fmt main.go
