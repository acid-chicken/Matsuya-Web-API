import express from 'express';
var router = express.Router();
import fs from 'fs';
const all = JSON.parse(fs.readFileSync('menu/all.json', 'utf8'));

router.get('/', function (req, res) {
  const random = Math.floor(Math.random() * (44 + 1 - 0)) + 0;
  res.header("Content-Type", "application/json; charset=utf-8");
  res.send("{\"menu\":\"" + all[random] + "\"}");
});

module.exports = router;
