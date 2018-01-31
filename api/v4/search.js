import express from "express";
import dbcontroller from "../../modules/dbcontroller";
const router = express.Router();

router.get('/', function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    dbcontroller.menu.find({ type: type }, function (e, docs) {
      res.send(docs);
    });
  }
  else if (query.name) {
    const name = query.name.split(',');
    dbcontroller.menu.find({ name: name }, function (e, docs) {
      res.send(docs);
    });
  }
  else {
    dbcontroller.menu.find({  }, function (e, docs) {
      res.send(docs);
    });
  }
});

module.exports = router;
