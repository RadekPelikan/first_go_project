VERSION_FILE="current.version"

if [[ $# -eq 0 ]]; then
  echo "Usage: $0 [-m] [-n] [-r]"
  echo "  -m    Increment major version"
  echo "  -n    Increment minor version"
  echo "  -r    Increment revision version"
  exit 1
fi

tag="1.0.0"
if [ -e $VERSION_FILE ]; then
  tag=$(cat $VERSION_FILE)
else
  touch $VERSION_FILE
fi

increment_version() {
  local version=$1
  local part=$2

  # Regex to capture major, minor, and revision numbers
  if [[ "$version" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    major="${BASH_REMATCH[1]}"
    minor="${BASH_REMATCH[2]}"
    revision="${BASH_REMATCH[3]}"
  else
    echo "Invalid version format: $version"
    exit 1
  fi

  case "$part" in
    major)
      major=$((major + 1))
      minor=0
      revision=0
      ;;
    minor)
      minor=$((minor + 1))
      revision=0
      ;;
    revision)
      revision=$((revision + 1))
      ;;
    *)
      echo "Invalid version part!"
      exit 1
      ;;
  esac

  # Return the updated version
  echo "$major.$minor.$revision"
}

# Parse flags and update version
while getopts "mnr" opt; do
  case $opt in
    m)
      tag=$(increment_version "$tag" "major")
      ;;
    n)
      tag=$(increment_version "$tag" "minor")
      ;;
    r)
      tag=$(increment_version "$tag" "revision")
      ;;
    *)
      echo "Usage: $0 [-m] [-n] [-r]"
      exit 1
      ;;
  esac
done

echo $tag > "$VERSION_FILE"

docker build . -t first-go-app:$tag
docker tag first-go-app:$tag first-go-app:latest