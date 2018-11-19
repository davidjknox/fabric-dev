## How to run

In a new terminal execute:
```bash
winpty docker exec -it chaincode bash
```

Within the container:
``` bash
cd fmscc && \
go build && \
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./fmscc
```

In a new terminal:
```bash
winpty docker exec -it cli bash
```

Within the container:
``` bash
peer chaincode install -p chaincodedev/chaincode/fmscc -n mycc -v 0 &&\
peer chaincode instantiate -n mycc -v 0 -c '{"Function": "Init","Args":["{}"]}' -C myc
```

The following adds a new CDR to the block, and as the value (£50.10) exceeds the threshold, an alert (message to the console) is produced:
``` bash
peer chaincode invoke -n mycc -C myc -c '{"Function":"AddCdr", "Args":["2018-11-19 00:00:00","441536200251","441536200256","30","50.1"]}'
```
