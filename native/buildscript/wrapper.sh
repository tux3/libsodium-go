#!/bin/bash/
echo "[BUILDSCRIPT] Build wrapper called, we have successfully hijacked go build!"

# Run the original compiler command
$@

# If this is the first run, run the build script
STAMP_FILE="build_stamp"
if [ -f $STAMP_FILE ]; then
    echo "[BUILDSCRIPT] Nothing to be done. Remove $PWD/$STAMP_FILE to trigger a rebuild"
    exit 0
else
    echo "[BUILDSCRIPT] Triggering a new build!"
    bash build.sh
    echo "Remove this file to trigger a new build" > $PWD/$STAMP_FILE
fi

