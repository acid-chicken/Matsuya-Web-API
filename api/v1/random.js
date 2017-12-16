import express from 'express';
var router = express.Router();
import fs from 'fs';
var all = JSON.parse(fs.readFileSync('menu/all.json', 'utf8'));

router.get('/', function (req, res) {
  var random = Math.floor(Math.random() * (44 + 1 - 0)) + 0;
  res.header("Content-Type", "application/json; charset=utf-8");
  res.send("[\"" + all[random] + "\"]");
});

module.exports = router;
