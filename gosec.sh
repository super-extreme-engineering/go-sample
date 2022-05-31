#!/usr/bin/env bash

_file=(
pkg/downloader/manager.go
pkg/postrender/exec.go
pkg/chart/loader/archive.go
pkg/repo/chartrepo.go
pkg/repo/repotest/server.go
cmd/helm/load_plugins.go
cmd/helm/helm_test.go
internal/tlsutil/cfg.go
internal/tlsutil/tls.go
pkg/getter/httpgetter.go
pkg/getter/plugingetter.go
pkg/registry/constants.go
pkg/registry/client_test.go
pkg/kube/ready.go
pkg/kube/ready_test.go
pkg/release/mock.go
pkg/plugin/installer/http_installer.go
pkg/plugin/installer/installer.go
)

rm -rf ~/tempDir/go
mkdir -p ~/tempDir/go
for file in ${_file[*]}; do
  echo "File name is $file"
  rsync -Rv $file ~/tempDir/go/
  rm $file
done
