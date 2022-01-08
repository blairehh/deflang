#!/bin/bash

DEFLANG_SESSION_ID=$(./deflang-cmd)
# PATH="$PATH:/tmp/deflang/${DEFLANG_SESSION_ID}/vars"
echo "${DEFLANG_SESSION_ID}"

function def() {
  ./def-cmd "--session ${DEFLANG_SESSION_ID} $*"
}