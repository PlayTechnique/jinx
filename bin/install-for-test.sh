#!/bin/bash -el

THIS_SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "${THIS_SCRIPT_DIR}/.."

go install .

tmpdir=$(mktemp -d)
cd "${tmpdir}"

set +e
jinx new roflcopter --container-name hippolols
set -e
bash


rm -rf "${tmpdir}"
