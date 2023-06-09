FROM registry.ci.openshift.org/openshift/release:golang-1.16 AS builder

ENV PKG=/go/src/github.com/openshift/osde2e-acm-test-harness/
WORKDIR ${PKG}

# compile test binary
COPY . .
RUN make

FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

COPY --from=builder /go/src/github.com/openshift/osde2e-acm-test-harness/osde2e-acm-test-harness.test osde2e-acm-test-harness.test

ENTRYPOINT [ "/osde2e-acm-test-harness.test" ]

