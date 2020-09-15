#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


run_pod_tests() {
  set -o nounset
  set -o errexit

  hypercloud::log::status "Testing kubectl(v1:pods)"

  #hypercloud::test::get_object_assert pods "{{range.items}}{{$id_field}}:{{end}}" ''

  kubectl create -f - << __EOF__
{
  "kind": "Pod",
  "apiVersion": "v1",
  "metadata": {
    "name": "test",
    "namespace": "kind"
  },
  "spec": {
    "containers": [
      {
        "name": "nginx",
        "image": "nginx"
      }
    ]
  }
}
__EOF__

  hypercloud::test::get_object_assert pod/test "{{$id_field}}" 'test' "-n kind"

  kubectl delete pod test -n kind
  hypercloud::test::get_object_assert pod/test "{{$id_field}}" '' "-n kind"



}

