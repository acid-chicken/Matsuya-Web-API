import express from "express";
import dbSearch from "../../helper/dbSearch";  
const router = express.Router();

router.get('/', async function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    const q = { type: type };
    try {
      res.send(await dbSearch(q));
    } catch(ex) {
      res.send(ex);
    }
  }
  else if (query.name) {
    const name = query.name.split(',');
    const q = { name: name };
    try {
      res.send(await dbSearch(q));
    } catch(ex) {
      res.send(ex);
    }
  }
  else {
    try {
      res.send(await dbSearch({}));
    } catch(ex) {
      res.send(ex);
    }
  }
});

module.exports = router;
