var express = require('express');
var router = express.Router();
var fs = require('fs');
var all = JSON.parse(fs.readFileSync('../menu/all.json', 'utf8'));

router.get('/', function (req, res, next) {
  var length = Object.keys(all).length;
  var random = Math.floor(Math.random() * (length + 1 - 0)) + 0;
  console.log(random + '/' + length);
  res.send(all[random]);
});

module.exports = router;
