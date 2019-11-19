### Pre-requisites

1. `minikube` (`brew install minikube`)
1. `helm` (`brew install kubernetes-helm`)
1. Add the `minikube`'s IP to the list of insecure docker registries of your
   local docker daemon (Either using the Docker Desktop or editing
   `~/.docker/daemon.json`):
```
{
  "insecure-registries" : ["minikubeIP:port"]
}
```


### Installing dependencies and CAPI

CAPI requires a database and blobstore.  We chose to use Postgres and Minio for
those dependencies, respectively.  Both have stable `helm` charts, so that is
the approach we use to install them.


1. `minikube start` to make sure `minikube` is up and running
1. `./dev/deploy.sh` to deploy the dependencies and CAPI


### Rolling out changes to CAPI

1. `./dev/build-and-rollout-capi.sh` will take the `cloud_controller_ng` code in
   the `src/cloud_controller_ng` submodule, build a docker image with it, and
   roll the new image out to the `minikube` cluster.
