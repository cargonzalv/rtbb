# Deployment

Service is hosted in the a Kubernetes cluster. We support following environments:

- dev - Staging environment
- canary - Canary environment. *Coming soon*
- production - Production environment. *Coming soon*

The deployment is an automation process controlled by github actions. 

## Kubernetes

[Kubernetes](https://kubernetes.io/), a.k.a. k8s, is an opensource system for automating deployment, scalling, and management of containerized applications.

We are using the [Rancher](https://rancher-management.k8s.adgear.com/dashboard/auth/login?timed-out) - k8s clusters manager.

Some usefull k8s commands:

```shell
# switch context to use1-rdev cluster
kubectl config use-context use1-rdev

# retrieve current context
kubectl config current-context

# get pods running in rtb-bidder namespace 
kubectl get pods -n rtb-bidder

# get services running in rtb-bidder namespace
kubectl get services -n rtb-bidder
```

## Helm package

We are using `helm` to install the application.

[Helm](https://helm.sh/) is the package manager for Kubernetes. 
The `helm package` located in [ci/helm/service-chart] folder and contains:

- templates folder - The directory contains templates of the k8s configuration files;
- Chart.yaml - The package description file;
- values files - yaml files which have values using in the templates;

values are located in values.yaml file and yaml file specific for the environment. dev.yaml as an example.

values.yaml is the values file supplying values for all environment. The environment specific file is overwriting those values. 

Example of helm commands:

```shell
# dry run command will evaluate helm package and will not change config files in the k8s 
helm deploy --dry-run --name rtb-bidder -n rtb-bidder --kube-context use1-rdev -f values.yaml -f dev.yaml .

# this command will deploy the helm package in the k8s
helm deploy --name rtb-bidder -n rtb-bidder --kube-context use1-rdev -f values.yaml -f dev.yaml .
```

## Docker

We are using docker container to deploy the service.
The [Docker](https://docs.docker.com/get-started/what-is-a-container/) is an isolated environment for our service. The container has everything required to run our code. We are using multistage Docker container.

The Docker file has 2 stages: 

- build container
- run container

Build container has everyting required to build the rtb-bidder service. After build completed, the artifact copied to the run container. The run container has only minimum packages installed which allow to run the service. 

## CI/CD

The information on CI/CD piplines will be available soon.

## Manual build and deployment

In most cases, you will not deploy service manually, but if necessary for testing, you could deploy to a dev environment.
To deploy manually, follow this steps:

Usefull docker commands:

Build the docker container.

 ```shell
 DOCKER_BUILDKIT=1 docker build --no-cache --ssh default -t adgear-docker.jfrog.io/adgear/rtb-bidder:"$TAG" .
 ```

 Now you can run the container.

 ```shell
 docker run -it -p 8085:8085 rtb-bidder:0.0.1
 ```
 or with docker-compose

 ```shell
 docker-compose up
 ```

Now you can run culr command to test it:

```shell
curl -kv http://localhost:8085/health/liveness
```

or naviagate to this url [http://localhost:8085/info] in the browser to see build information page.

 ### Build the `Docker` image and push it to the `Jfrog` repo

Use this command to build the docker image

 ```shell
 DOCKER_BUILDKIT=1 docker build --no-cache --ssh default -t adgear-docker.jfrog.io/adgear/rtb-bidder:"$TAG" .
 ```
After build completed, push the image to `Jfrog` docker repo with the command:

```shell
docker push adgear-docker.jfrog.io/adgear/rtb-bidder:"$TAG"
```

Next we need to modify the environment specific value file.
The tag value which we use to build and publish the docker container.

```yaml
image:
  tag: "0.0.2"
```

Now we can run the helm package to deploy the service to kubernetes.

Change directory to `ci/helm/service-chart`. From this directory run:

```shell
# using dev environment and dev.yaml value file.
helm deploy --name rtb-bidder -n rtb-bidder --kube-context use1-rdev -f values.yaml -f dev.yaml .
```

or use makefile

```shell
make install-chart-dev
```

```shell
# get pods running in rtb-bidder namespace 
kubectl get pods -n rtb-bidder
```

you should see, output similar to:

```shell
❯ kubectl get pods -n rtb-bidder
NAME           READY   STATUS    RESTARTS   AGE
rtb-bidder-0   2/2     Running   0          8d
rtb-bidder-1   2/2     Running   0          8d
```

```shell
# get services running in rtb-bidder namespace
kubectl get services -n rtb-bidder
```

you should see, output similar to:

```shell
❯ kubectl get services -n rtb-bidder
NAME                          TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
rtb-bidder                    NodePort    10.43.155.1   <none>        8080:30218/TCP   10d
rtb-bidder-headless-service   ClusterIP   None          <none>        8080/TCP         10d
```