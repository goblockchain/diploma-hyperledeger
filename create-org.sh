!/bin/bash

set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

export COMPANY_DOMAIN=goblockchain.com
export PEER_NUMBER=0
export ORGANIZATION_NAME=GoBlockchain
export ORGANIZATION_NAME2=goblockchain.com

# Create config and crypto-config if not exists
# mkdir -p config/
# mkdir -p crypto-config/

# remove previous crypto material and config transactions
# rm -fr config/*
# rm -fr crypto-config/*

rm -rf ./client/controllers/users/*

docker-compose -f docker-compose.yml down

docker-compose -p diploma -f docker-compose.yml up -d 

export FABRIC_START_TIMEOUT=2

sleep ${FABRIC_START_TIMEOUT}

# Create the channel
docker exec cli.${ORGANIZATION_NAME2} peer channel create -o orderer.${ORGANIZATION_NAME2}:7050 --tls --cafile /opt/peer/crypto/ordererOrganizations/goblockchain.com/tlsca/tlsca.goblockchain.com-cert.pem -c diplomachannel -f /etc/hyperledger/configtx/channel.tx

# Join peer0.goblockchain.com to the channel.
docker exec cli.${ORGANIZATION_NAME2} peer channel join -b diplomachannel.block

sleep ${FABRIC_START_TIMEOUT}

docker exec cli.${ORGANIZATION_NAME2} peer channel update -o orderer.${ORGANIZATION_NAME2}:7050 -c diplomachannel --tls --cafile /opt/peer/crypto/ordererOrganizations/goblockchain.com/tlsca/tlsca.goblockchain.com-cert.pem -f /etc/hyperledger/configtx/${ORGANIZATION_NAME}MSPanchors.tx

docker exec cli.${ORGANIZATION_NAME2} peer chaincode install -n diploma -v 1.0 -p github.com/chaincode -l golang

docker exec cli.${ORGANIZATION_NAME2} peer chaincode instantiate -o orderer.${ORGANIZATION_NAME2}:7050 -C diplomachannel --tls --cafile /opt/peer/crypto/ordererOrganizations/goblockchain.com/tlsca/tlsca.goblockchain.com-cert.pem -n diploma -l golang -v 1.0 -c '{"Args":[]}' -P "OR('${ORGANIZATION_NAME}MSP.member')" --collections-config /opt/gopath/src/github.com/chaincode/collections_config.json

#Enroll admins and register users to interact with the network
node ./client/controllers/enrollAdmin.js
sleep 1
node ./client/controllers/registerUser.js
sleep 1
