## Connection check tool
Simple tool for checking that OTel endpoint is reachable.
Published and used inside k8s docker images for use in kubernetes.

Tool exits with non-zero exit codes when connection cannot be established, or with zero exit code when connection is successful.

## Running the Tool
Example command.
```
go run . -clusteruid=1234321 -endpoint=CLUSTER_URL:443 -apitoken=SOME_TOKEN -insecure=false
```
