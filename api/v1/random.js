var express = require('express');
var router = express.Router();
var fs = require('fs');
var all = JSON.parse(fs.readFileSync('menu/all.json', 'utf8'));

router.get('/', function (req, res) {
  var random = Math.floor(Math.random() * (44 + 1 - 0)) + 0;
  // console.log(random + '/' + 44);
  res.header("Content-Type", "application/json; charset=utf-8");
  res.send("[\"" + all[random] + "\"]");
});

module.exports = router;
