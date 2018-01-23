import express from "express";
import dbcontroller from "../../modules/dbcontroller";
const router = express.Router();

router.get('/type/:type', function (req, res) {
  const type = req.params.type;
  dbcontroller.menu.find({ type: type }, function (e, docs) {
    if (e) {
      console.log(e);
      return;
    }
    if (docs.length === 0) {
      console.log('docs not found');
      res.sendStatus(404);
      return;
    }
    res.send(docs);
  });
});

router.get('/name/:name', function (req, res) {
  const name = req.params.name;
  dbcontroller.menu.find({ name: name }, function (e, docs) {
    if (e) {
      console.log(e);
      return;
    }
    if (docs.length === 0) {
      console.log('docs not found');
      res.sendStatus(404);
      return;
    }
    res.send(docs);
  });
});

module.exports = router;
