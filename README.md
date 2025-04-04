# Inventory `0.0.1`

> Your universal model catalog, everything, everywhere, all at once.

> [Live API Docs](https://inventory.moeru.ai/apis/docs/v1)

Or you can try it with our instance:

```shell
curl -L -X GET -v 'https://inventory.moeru.ai/api/v1/common-tasks/models' | jq
```

## Getting Started

### Build

> require `go` 1.23+

```bash
git clone https://github.com/moeru-ai/inventory.git
cd inventory
go build -o ./result/inventory ./cmd/api-server
```

### Run

```bash
# http server started on [::]:7921
# grpc server started on [::]:7922
./result/api-server
```
