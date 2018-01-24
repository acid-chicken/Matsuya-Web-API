import express from "express";
import dbcontroller from "../../modules/dbcontroller";
const router = express.Router();

router.get('/', function (req, res) {
  dbcontroller.menu.find({  }, function (e, docs) {
    const length = docs.length;
    const random = Math.floor(Math.random() * length);
    res.send(docs[random]);
  });
});

router.get('/type', function (req, res) {
  const type = req.query.type.split(',');
  dbcontroller.menu.find({ type: type }, function (e, docs) {
    const length = docs.length;
    const random = Math.floor(Math.random() * length);
    res.send(docs[random]);
  });
});

module.exports = router;
