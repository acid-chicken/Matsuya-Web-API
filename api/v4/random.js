import express from "express";
import dbcontroller from "../../modules/dbcontroller";
const router = express.Router();

router.get('/', function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    dbcontroller.menu.find({ type: type }, function (e, docs) {
      const length = docs.length;
      const random = Math.floor(Math.random() * length);
      res.send(docs[random]);
    });
  }
  else if (query.name) {
    const name = query.name.split(',');
    dbcontroller.menu.find({ name: name }, function (e, docs) {
      const length = docs.length;
      const random = Math.floor(Math.random() * length);
      res.send(docs[random]);
    });
  }
  else {
    dbcontroller.menu.find({  }, function (e, docs) {
      const length = docs.length;
      const random = Math.floor(Math.random() * length);
      res.send(docs[random]);
    });
  }
});

module.exports = router;
