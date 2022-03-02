#!/usr/bin/env bash -el

JINX="jinx"
THIS_SCRIPT_DIR=$(cd $(dirname $0) && pwd)

EXIT_COMMANDS=()
function at_exit() {
  for command in "${EXIT_COMMANDS}"; do
    eval $command
  done
}

trap at_exit EXIT

cd ${THIS_SCRIPT_DIR}/../..

# Some containers are running. Bail and ask the user to fix this.
if [[ $(docker ps --format "{{.Names}}") = "jinkies" ]]; then
  echo "jinkies containers are running. Please calm this situation and try again." >&2
  exit 1
fi

TEST_CONFIG_FILE="testHostConfig.yml"
EXIT_COMMANDS+=("rm ${TEST_CONFIG_FILE}")

echo "Writing host config file"
echo "---" >> ${TEST_CONFIG_FILE}
echo "AutoRemove: false" >> ${TEST_CONFIG_FILE}


echo "Verifying jinx can start a container with ${TEST_CONFIG_FILE}"

# Verify jinx can start a container
OUTPUT=$(./jinx serve start -o ${TEST_CONFIG_FILE})
# Some containers are running. Bail and ask the user to fix this.
if [[ $? != 0 ]] || [[ $(docker ps --format "{{.Names}}") != "jinkies" ]]; then
  echo "'jinx serve start' failed" >&2
  echo "Containers not started." >&2
  echo ${OUTPUT} >&2
  exit 2
fi

# Verify jinx can start a container
OUTPUT=$(./jinx serve stop)
if [[ $? != 0 ]] || [[ $(docker ps --format "{{.Names}}") == "jinkies" ]]; then
  echo "'jinx serve stop' failed" >&2
  echo "Containers not stopped." >&2
  echo ${OUTPUT} >&2
  exit 2
fi

# The test host config sets autoremove to false, so there should be a stopped container named jinkies!
OUTPUT=$(docker container ls -a --format "{{.Names}}")
if [[ $? != 0 ]] || [[ $OUTPUT != "jinkies" ]]; then
  echo "'No container 'jinkies' found in the stopped containers" >&2
  echo "Rats!" >&2
  echo ${OUTPUT} >&2
  exit 3
else
  # test succeeded! Remove the stopped container.
  docker container rm jinkies
fi