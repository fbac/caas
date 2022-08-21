# chaoscmd

Chaos as a Service for Kubernetes clusters

- [chaoscmd](#chaoscmd)
  - [cmd options](#cmd-options)
    - [Usage](#usage)
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

### Usage

  ```bash
  chaoscmd -kubeconfig=/path/to/kubeconfig -dry-run=false -pods-to-kill=4
  ```

## TODO

- go test, go lint and go vet
- Dockerfile to containerize app
- Makefile target to build container
- github workflow to build container and push to ghcr.io
- Asset yaml yo run container as a cronjob in a Kubernetes cluster
