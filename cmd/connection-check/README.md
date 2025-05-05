## Connection check tool
Simple tool for checking that OTel endpoint is reachable.
Published and used inside k8s docker image distribution for use in kubernetes.

Tool either exits with non-zero exit codes when connection cannot be established.

## Running the Tool
Run the tool with the following command, adjusting paths and versions as needed:
```
go run . -clusteruid=1 -endpoint=CLUSTER_URL:443 -apitoken=SOME_TOKEN -insecure=false
```
