#!/bin/bash

# Function to get the base version prefix
get_base_version() {
  local version=$(git describe --tags --match 'v*' --abbrev=0 2>/dev/null)
  if [ -z "$version" ]; then
    echo "v0.0.0"
  else
    echo "$version"
  fi
}

# Function to get the timestamp (yymmddhhmmss)
get_timestamp() {
  # git log -1 --format=%ct | date -u "+%Y%m%d%H%M%S"
  git log -1 --format=%ad --date=format:"%Y%m%d%H%M%S"
}

# Function to get the revision identifier (abcdefabcdef)
get_revision_identifier() {
  git rev-parse --short=12 HEAD 
}

# Assemble the version string
assemble_version() {
  local base_version timestamp revision_identifier
  base_version=$(get_base_version)
  timestamp=$(get_timestamp)
  revision_identifier=$(get_revision_identifier)
  echo "${base_version}-${timestamp}-${revision_identifier}"
}

# Print the assembled version
assemble_version

