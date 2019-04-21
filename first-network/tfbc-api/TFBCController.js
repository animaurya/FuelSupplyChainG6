var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var TFBC = require("./FabricHelper")


// Request LC
router.post('/addFuel', function (req, res) {

TFBC.addFuel(req, res);

});

// Issue LC
router.post('/move', function (req, res) {

    TFBC.move(req, res);
    
});

// Accept LC
router.post('/update', function (req, res) {

    TFBC.updateC(req, res);
    
});

// Get LC
router.post('/transfer', function (req, res) {

    TFBC.transfer(req, res);
    
});

// Get LC history
router.get('/viewStatus', function (req, res) {

    TFBC.viewStatus(req, res);
    
});


module.exports = router;
