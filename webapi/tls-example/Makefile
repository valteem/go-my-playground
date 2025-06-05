GENERATE_CERT=$(shell find / -name "generate_cert.go" -print -quit)

.PHONY: cert
cert:
	go run ${GENERATE_CERT} -host=localhost