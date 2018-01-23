import express from "express";
import dbcontroller from "../../modules/dbcontroller";
const router = express.Router();

router.get('/', function (req, res) {
  dbcontroller.menu.find({  }, function (e, docs) {
    const length = docs.length;
    const random = Math.floor(Math.random() * (length + 1 - 0)) + 0;
    res.send(docs[random]);
  });
});

module.exports = router;
