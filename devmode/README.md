
1. Install a newer version of curl from 

https://curl.haxx.se/windows/dl-7.62.0/curl-7.62.0-win64-mingw.zip

(Install in C/Windows/System32)

2. Execute the following:

> git config --global core.autocrlf false
> git config --global core.longpaths true

3. Now execute the following command in the directory you wish to install the fabric samples:

> curl -sSL http://bit.ly/2ysbOFE | bash -s 1.3.0

4. Now this:

> export MSYS_NO_PATHCONV=1
> export COMPOSE_CONVERT_WINDOWS_PATHS=1

5. Now to start the network...

> winpty docker-compose -f docker-compose-devmode.yaml up

from the devmode directory

6. Open a new terminal and execute the following:

> winpty docker exec -it chaincode bash

7. Now, compile your chaincode:

> cd sacc
> go build

8. Now run the chaincode:

> CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./sacc

9. Open another terminal...

> winpty docker exec -it cli bash

Within the container:

> peer chaincode install -p chaincodedev/chaincode/sacc -n mycc -v 0
> peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc

10. Now issue an invoke to change the value of “a” to “20”.

> peer chaincode invoke -n mycc -c '{"Args":["set", "a", "20"]}' -C myc

11. Finally, query a. We should see a value of 20

> peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc
