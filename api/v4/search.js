import express from "express";
import dbSearch from "../../helper/dbSearch";  
const router = express.Router();

router.get('/', async function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    const q = { type: type };
    res.send(await dbSearch(q));
  }
  else if (query.name) {
    const name = query.name.split(',');
    const q = { name: name };
    res.send(await dbSearch(q));
  }
  else {
    res.send(await dbSearch({}));
  }
});

module.exports = router;
