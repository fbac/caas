FROM frolvlad/alpine-glibc
COPY ./kubeconfig /
COPY ./bin/chaoscmd /
ENTRYPOINT ["/chaoscmd"]