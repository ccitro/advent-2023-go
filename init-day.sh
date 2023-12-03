#!/bin/bash

# Check if a day number is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <day-number>"
  exit 1
fi

DAY_NUM=$(printf "%02d" $1)
DAY_DIR="internal/day$DAY_NUM"

# Copy template directory
cp -r internal/dayXX $DAY_DIR

# Add import statement to main.go
IMPORT_STATEMENT="\t\"github.com/ccitro/advent-2023-go/$DAY_DIR\""
sed -i "/import (/a $IMPORT_STATEMENT" main.go

# Add entry to the dayMethods map
MAP_ENTRY="\t\"day$DAY_NUM\": {LoadPuzzle: day$DAY_NUM.LoadPuzzle, Part1: day$DAY_NUM.Part1, Part2: day$DAY_NUM.Part2},"
sed -i "/var dayMethods = map\[string\]DayMethods{/a $MAP_ENTRY" main.go
