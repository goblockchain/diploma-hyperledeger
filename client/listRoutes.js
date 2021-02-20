'use strict';
module.exports = function(app) {
    var diplomaPost = require('./controllers/createDiploma');
    var diplomaGet = require('./controllers/queryDiploma');
    var diplomaGetAll = require('./controllers/queryAllDiplomas');
    var diplomaSearch = require('./controllers/searchDiplomas');

    //org2 route
    var diplomaPostOrg2 = require('./controllers/org2-scripts/createDiploma');
    var diplomaGetOrg2 = require('./controllers/org2-scripts/queryDiploma');

     //obs1 route
     var diplomaGetObs1 = require('./controllers/obs1-scripts/queryDiploma');

    // Route the webservices
    app.route('/diploma-create').post(diplomaPost.save_diploma_transaction);
    app.route('/diploma-get').post(diplomaGet.get_diploma_transaction);
    app.route('/diploma-get-all').post(diplomaGetAll.get_all_diplomas);
    app.route('/diploma-search').post(diplomaSearch.search_diploma_transaction);

    //org2 route
    app.route('/diploma-create-org2').post(diplomaPostOrg2.save_diploma_transaction);
    app.route('/diploma-get-org2').post(diplomaGetOrg2.get_diploma_transaction);

    //obs1 route
    app.route('/diploma-get-obs1').post(diplomaGetObs1.get_diploma_transaction);
}   
