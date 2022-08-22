# chaoscmd

Chaos as a Service for Kubernetes clusters

- [chaoscmd](#chaoscmd)
  - [cmd options](#cmd-options)
    - [Usage as binary](#usage-as-binary)
    - [Usage as container](#usage-as-container)
    - [Usage in kubernetes](#usage-in-kubernetes)
  - [TODO](#todo)

## cmd options

```bash
Usage:
  -dry-run
        run chaosnetes as dry-run. (default true)
  -kubeconfig string
        absolute path to the kubeconfig file (default "kubeconfig")
  -namespace string
        namespace to use (default "workloads")
  -pods-to-kill int
        set maximum amount of pods to kill (default 1)
```

### Usage as binary

  ```bash
  chaoscmd -kubeconfig=/path/to/kubeconfig -dry-run=false -pods-to-kill=4
  ```

### Usage as container

- Place your kubeconfig in the root path of the repository
- Run `make image`
- Run `docker run fbac/chaoscmd:latest <OPTIONS>`

### Usage in kubernetes

- There's a cronjob in assets to run this in a cluster

## TODO

- Test suite: fmt, vet, lint and test using v1 api mocks
- Critical change: Do not upload kubeconfig in the image, instead pass it inside a volume to the container
- Change build logic: binary should be built inside the container
- Authenticate to cluster using serviceAccount if running in kubernetes
- Github Actions workflow to build container and push to ghcr.io as OCI object
- Set flags as env vars and adapt the cronjob to use them
