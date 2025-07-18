#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# Short-circuit if protoc.sh has already been sourced
[[ $(type -t onex::protoc::loaded) == function ]] && return 0

# The root of the build/dist directory
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/../..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"

#PROTOC_VERSION=23.4
PROTOC_VERSION=31.1

# Generates $1/api.pb.go from the protobuf file $1/api.proto
# and formats it correctly
# $1: Full path to the directory where the api.proto file is
function onex::protoc::generate_proto() {
  onex::golang::setup_env
  GO111MODULE=on GOPROXY=off go install k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo

  onex::protoc::check_protoc

  local package=${1}
  onex::protoc::protoc "${package}"
  onex::protoc::format "${package}"
}

# Checks that the current protoc version matches the required version and
# exit 1 if it's not the case
function onex::protoc::check_protoc() {
  if [[ -z "$(/usr/bin/which protoc)" || "$(protoc --version)" != "libprotoc ${PROTOC_VERSION}"* ]]; then
    echo "Generating protobuf requires protoc ${PROTOC_VERSION}."
    echo "Run scripts/install-protoc.sh or download and install the"
    echo "platform-appropriate Protobuf package for your OS from"
    echo "https://github.com/protocolbuffers/protobuf/releases"
    return 1
  fi
}

# Generates $1/api.pb.go from the protobuf file $1/api.proto
# $1: Full path to the directory where the api.proto file is
function onex::protoc::protoc() {
  local package=${1}
  gogopath=$(dirname "$(onex::util::find-binary "protoc-gen-gogo")")

  (
    cd "${package}"

    # This invocation of --gogo_out produces its output in the current
    # directory (despite gogo docs saying it would be source-relative, it
    # isn't).  The inputs to this function do not all have a common root, so
    # this works best for all inputs.
    PATH="${gogopath}:${PATH}" protoc \
      --proto_path="$(pwd -P)" \
      --proto_path="${PROJ_ROOT_DIR}/vendor" \
      --proto_path="${PROJ_ROOT_DIR}/third_party/protobuf" \
      --gogo_out=paths=source_relative,plugins=grpc:. \
      api.proto
  )
}

# Formats $1/api.pb.go, adds the boilerplate comments and run gofmt on it
# $1: Full path to the directory where the api.proto file is
function onex::protoc::format() {
  local package=${1}

  # Update boilerplate for the generated file.
  cat hack/boilerplate/boilerplate.generatego.txt "${package}/api.pb.go" > tmpfile && mv tmpfile "${package}/api.pb.go"

  # Run gofmt to clean up the generated code.
  onex::golang::verify_go_version
  gofmt -s -w "${package}/api.pb.go"
}

# Compares the contents of $1 and $2
# Echo's $3 in case of error and exits 1
function onex::protoc::diff() {
  local ret=0
  diff -I "gzipped FileDescriptorProto" -I "0x" -Naupr "${1}" "${2}" || ret=$?
  if [[ ${ret} -ne 0 ]]; then
    echo "${3}"
    exit 1
  fi
}

function onex::protoc::install() {
  # run in a subshell to isolate caller from directory changes
  (
    local os
    local arch
    local download_folder
    local download_file

    os=$(onex::util::host_os)
    arch=$(onex::util::host_arch)
    download_folder="protoc-v${PROTOC_VERSION}-${os}-${arch}"
    download_file="${download_folder}.zip"

    cd "${HOME}/bin" || return 1
    if [[ ! -f protoc  ]]; then
      local url
      if [[ ${os} == "darwin" ]]; then
        # TODO: switch to universal binary when updating to 3.20+
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip"
      elif [[ ${os} == "linux" && ${arch} == "amd64" ]]; then
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
      elif [[ ${os} == "linux" && ${arch} == "arm64" ]]; then
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip"
      else
        onex::log::info "This install script does not support ${os}/${arch}"
        return 1
      fi
      onex::util::download_file "${url}" "${download_file}"
      unzip -o "${download_file}" -d "${download_folder}"
      mv "${download_folder}/bin/protoc" $HOME/bin/protoc
      chmod +rX $HOME/bin/protoc
      rm -rvf "${download_file}" "${download_folder}"
    fi
    onex::log::info "$HOME/bin/protoc v${PROTOC_VERSION} installed."
  )
}

# Marker function to indicate protoc.sh has been fully sourced
onex::protoc::loaded() {
  return 0
}
