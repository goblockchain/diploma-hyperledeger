set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
export FABRIC_START_TIMEOUT=2

COMPANY_DOMAIN=goblockchain.com

docker rm -f $(docker ps -aq)

sudo rm -rf crypto-config/ $COMPANY_DOMAIN/ config/*

docker-compose -f docker-compose.yml down

docker-compose -p diploma -f docker-compose.yml up -d ca

docker exec ca.goblockchain.com sh -c '/etc/hyperledger/scripts/idsGeneration.sh'

sudo cp -r ./goblockchain/client/crypto-config ./

sudo ./generate.sh

docker-compose -p diploma -f docker-compose.yml up -d orderer peer cli couchdb

sleep ${FABRIC_START_TIMEOUT}

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=GoBlockchainMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/peer/crypto/peerOrganizations/users/Admin@peer.goblockchain.com/msp" cli.goblockchain.com peer channel create -o orderer.goblockchain.com:7050 --tls --cafile /opt/peer/crypto/ordererOrganizations/users/Admin@orderer.goblockchain.com/tls/tlscacerts/tls-ca-goblockchain-com-7150.pem -c diplomachannel -f /etc/hyperledger/configtx/channel.tx


# Join peer0.goblockchain.com to the channel.
docker exec cli.${COMPANY_DOMAIN} peer channel join -b diplomachannel.block

sleep ${FABRIC_START_TIMEOUT}

docker exec cli.${COMPANY_DOMAIN} peer channel update -o orderer.${COMPANY_DOMAIN}:7050 -c diplomachannel --tls --cafile /opt/peer/crypto/ordererOrganizations/goblockchain.com/tlsca/tlsca.goblockchain.com-cert.pem -f /etc/hyperledger/configtx/${ORGANIZATION_NAME}MSPanchors.tx

docker exec cli.${COMPANY_DOMAIN} peer chaincode install -n diploma -v 1.0 -p github.com/chaincode -l golang

docker exec cli.${COMPANY_DOMAIN} peer chaincode instantiate -o orderer.${COMPANY_DOMAIN}:7050 -C diplomachannel --tls --cafile /opt/peer/crypto/ordererOrganizations/goblockchain.com/tlsca/tlsca.goblockchain.com-cert.pem -n diploma -l golang -v 1.0 -c '{"Args":[]}' -P "OR('${ORGANIZATION_NAME}MSP.member')" --collections-config /opt/gopath/src/github.com/chaincode/collections_config.json


docker exec cli.${COMPANY_DOMAIN} peer channel create -o orderer.${COMPANY_DOMAIN}:7050 --tls --cafile /opt/peer/crypto/ordererOrganizations/users/Admin@orderer.goblockchain.com/tls/tlscacerts/tls-ca-goblockchain-com-7150.pem -c diplomachannel -f /etc/hyperledger/configtx/channel.tx