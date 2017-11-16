var express = require('express');
var router = express.Router();
var fs = require('fs');
var all = JSON.parse(fs.readFileSync('./menu/all.json', 'utf8'));

router.get('/', function (req, res, next) {
  var random = Math.floor(Math.random() * (45 + 1 - 0)) + 0;
  console.log(random + '/' + 45);
  res.send(all[random]);
});

module.exports = router;
