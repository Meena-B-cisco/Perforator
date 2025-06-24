# Install Perforator CLI

Build `perforator/cmd/cli` from the root of the repository.

# Configure Perforator CLI

* Set `PERFORATOR_ENDPOINT` environment variable to specify the Perforator server URL once. Otherwise you need to use `--url` flag for each command.
* Set `PERFORATOR_SECURE` to enable or disable TLS. By default, TLS is enabled.


```console
export PERFORATOR_ENDPOINT="<your perforator endpoint>:<your grpc port>"
export PERFORATOR_SECURE=true
```

