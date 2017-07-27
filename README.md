# Fabric CLI Application

Fabric CLI is a command-line interface built using Fabric SDK GO. It provides the following functionality:

Channel:

- Create - Create a channel
- Join - Join a peer to a channel

Query:

- Info - Displays information about the block, including the block height
- Block - Displays the contents of a given block
- Tx - Displays the contents of a transaction within a block
- Channels - Displays all channels for a peer
- Installed - Displays the chaincodes installed on a peer

Chaincode:

- Install - Installs chaincode
- Instantiate - Instantiates chaincode
- Invoke - Invokes chaincode with a transaction proposal and a commit
- Query - Invokes chaincode without a commit
- Info - Retrieves details about the chaincode

Events:

- Listen Block - Listens for block events and displays the block when a new block is created
- Listen TX - Listens for transaction events
- Listen CC - Listens for chaincode events for a specified chaincode

## Running

Navigate to folder fabcli:

$ fabcli <command> <sub-command> [options]

To display the available commands/options:

$ fabcli

## Compatability

This example is compatible with the following Hyperledger Fabric commit levels:

- fabric: v1.0.0
- fabric-ca: v1.0.0

## Sample Usage

### Channel

###### Create a channel

$ fabcli channel create --cid mychannel --txfile fixtures/channel/mychannel.tx --config fixtures/config/config_test.yaml

###### Join a peer to a channel

$ fabcli channel join --cid mychannel --peer localhost:7051

###### Join all peers in org1 to a channel

$ fabcli channel join --cid mychannel --orgid org1

###### Join all peers to a channel

$ fabcli channel join --cid mychannel

### Query

###### Query info:

$ fabcli query info --cid mychannel

###### Query block by block number:

$ fabcli query block --cid mychannel --num 0

###### Query block by hash:

$ fabcli query block --cid mychannel --hash MKUvwa85E7OvITqBZYmf8yn9QIS5eZkal2xLTleK2AA

###### Query block output in JSON format:

$ fabcli query block --cid mychannel --num 0 --format json

###### Query transaction:

$ fabcli query tx --cid mychannel --txid 29bd4fd03e657da488acfa8ae1740eebf4a6ee81399bfc501f192cb407d2328c

###### Query channels joined by a peer:
$ fabcli query channels --peer localhost:7051

###### Query installed chaincodes on a peer:

$ fabcli query installed --peer localhost:7051

## Chaincode

###### Install chaincode on a peer:

$ fabcli chaincode install --cid=mychannel --peer localhost:7051 --ccp=github.com/user/somecc --ccid=somecc --v v0

###### Install chaincode on all peers of org1:

$ fabcli chaincode install --cid=mychannel --orgid org1 --ccp=github.com/user/somecc --ccid=somecc --v v0

###### Install chaincode on all peers:

$ fabcli chaincode install --cid=mychannel --ccp=github.com/user/somecc --ccid=somecc --v v0

###### Instantiate chaincode with default endorsement policy:

$ fabcli chaincode instantiate --cid mychannel --ccp github.com/user/somecc --ccid somecc --v v0 --args='{"Args":["arg1","arg2"]}'

###### Instantiate chaincode with specified endorsement policy:

$ fabcli chaincode instantiate --cid mychannel --ccp github.com/user/somecc --ccid somecc --v v0 --args='{"Args":["arg1","arg2"]}' --policy "AND('Org1MSP.member','Org2MSP.member')"

###### Retrieve chaincode deployment info:

$ fabcli chaincode info --cid mychannel --ccid somecc

###### Query chaincode on a set of peers:

$ fabcli chaincode query --ccid=somecc --args='{"Func":"query","Args":["a"]}' --peer localhost:7051,localhost:8051

###### Invoke chaincode on all peers:

$ fabcli chaincode invoke --ccid=somecc --args='{"Func":"add","Args":["a","1","11"]}'

###### Invoke chaincode on all peers in org1:

$ fabcli chaincode invoke --ccid=somecc --args='{"Func":"add","Args":["a","1","11"]}' --orgid org1

###### Invoke chaincode 5 times:

$ fabcli chaincode invoke --ccid=somecc --args='{"Func":"add","Args":["a","1","11"]}' --iterations 5

## Event

###### Listen for block events (output in JSON):

$ fabcli event listenblock --format json

###### Listen for chaincode events:

$ fabcli event listencc --ccid=somecc --event=someevent