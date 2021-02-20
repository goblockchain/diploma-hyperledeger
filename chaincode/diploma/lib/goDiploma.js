 'use strict';

const { Contract } = require('fabric-contract-api');

class GoDiploma extends Contract {

    async initLedger(ctx) {
        console.info('Initialized Ledger');
    }

    async queryDiploma(ctx, key, collection) {
        const apostaAsBytes = await ctx.getPrivateData(collection, key); // get the car from chaincode state
        if (!apostaAsBytes || apostaAsBytes.length === 0) {
            throw new Error(`${key} does not exist`);

        }
        console.info(apostaAsBytes.toString());
        return apostaAsBytes.toString();
    }

    async insertDiploma(ctx, id, jsonPayload, collection) {

            const PublicDetailsDiploma = {
                UniversityId   = jsonPayload.universityId,
                UniversitName  = jsonPayload.universityName,
                DiplomaId      = jsonPayload.diplomaId,
                StudentName    = jsonPayload.studentName,
                CourseName     = jsonPayload.courseName,
                EndDate        = jsonPayload.endDate, 
                StudentCpf     = jsonPayload.studentCPF,
            };

            await ctx.stub.putState('collectionDiploma', id, JSON.stringify(PublicDetailsDiploma));
            
            await ctx.stub.putState(collection, id, JSON.stringify(jsonPayload));
                
        console.info('Diploma created');
    }

    async salvarResultado(ctx, idPartida, resultado) {

        const apostaAsBytes = await ctx.stub.getState(idPartida);
        if (!apostaAsBytes || apostaAsBytes.length === 0) {
            throw new Error(`${idPartida} nao existe`);
        }

        const partidaString = JSON.parse(apostaAsBytes.toString());
        const partida = JSON.parse(partidaString);
        partida.idResultadoPartida = resultado;

        await ctx.stub.putState(idPartida, Buffer.from(JSON.stringify(partida)));
        console.info('Partida encerrada');
    }

    async buscarResultadoPartida(ctx, key) {

        const apostaAsBytes = await ctx.stub.getState(key); // get the car from chaincode state
        if (!apostaAsBytes || apostaAsBytes.length === 0) {
            throw new Error(`${key} does not exist`);

        }
        
        const partidaString = JSON.parse(apostaAsBytes.toString());
        const partida = JSON.parse(partidaString);
        const resultado = partida.idResultadoPartida;

        console.info(apostaAsBytes.toString());
        return resultado.toString();
    }

    async buscarOpcaoAposta(ctx, key) {

        const apostaAsBytes = await ctx.stub.getState(key); // get the car from chaincode state
        if (!apostaAsBytes || apostaAsBytes.length === 0) {
            throw new Error(`${key} does not exist`);

        }
        
        const partidaString = JSON.parse(apostaAsBytes.toString());
        const partida = JSON.parse(partidaString);
        const opcaoAposta = partida.tipoAposta;

        console.info(apostaAsBytes.toString());
        return opcaoAposta.toString();
    }

    
    async queryTimes(ctx, args, thisClass) {
    //   0
    // 'queryString'
    if (args.length < 1) {
        throw new Error('Incorrect number of arguments. Expecting queryString');
    }
    let queryString = args[0];
    if (!queryString) {
        throw new Error('queryString must not be empty');
    }
    let method = thisClass['getQueryResultForQueryString'];
    let queryResults = await method(ctx, queryString, thisClass);
    return queryResults;
    }
    
    async getQueryResultForQueryString(ctx, queryString, thisClass) {

        console.info('- getQueryResultForQueryString queryString:\n' + queryString)
        let resultsIterator = await ctx.getQueryResult(queryString);
        let method = thisClass['getAllResults'];
    
        let results = await method(resultsIterator, false);
    
        return Buffer.from(JSON.stringify(results));
      }
    
    async getHistoryForMarble(ctx, args, thisClass) {

    if (args.length < 1) {
        throw new Error('Incorrect number of arguments. Expecting 1')
    }
    let marbleName = args[0];
    console.info('- start getHistoryForMarble: %s\n', marbleName);

    let resultsIterator = await ctx.getHistoryForKey(marbleName);
    let method = thisClass['getAllResults'];
    let results = await method(resultsIterator, true);

    return Buffer.from(JSON.stringify(results));
    }

    async queryAllCars(ctx) {
        const startKey = 'CAR0';
        const endKey = 'CAR999';

        const iterator = await ctx.stub.getStateByRange(startKey, endKey);

        const allResults = [];
        while (true) {
            const res = await iterator.next();

            if (res.value && res.value.value.toString()) {
                console.log(res.value.value.toString('utf8'));

                const Key = res.value.key;
                let Record;
                try {
                    Record = JSON.parse(res.value.value.toString('utf8'));
                } catch (err) {
                    console.log(err);
                    Record = res.value.value.toString('utf8');
                }
                allResults.push({ Key, Record });
            }
            if (res.done) {
                console.log('end of data');
                await iterator.close();
                console.info(allResults);
                return JSON.stringify(allResults);
            }
        }
    }

    async changeCarOwner(ctx, carNumber, newOwner) {
        console.info('============= START : changeCarOwner ===========');

        const carAsBytes = await ctx.stub.getState(carNumber); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carNumber} does not exist`);
        }
        const car = JSON.parse(carAsBytes.toString());
        car.owner = newOwner;

        await ctx.stub.putState(carNumber, Buffer.from(JSON.stringify(car)));
        console.info('============= END : changeCarOwner ===========');
    }
}

module.exports = GoDiploma;
