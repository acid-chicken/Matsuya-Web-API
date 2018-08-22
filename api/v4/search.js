import express from "express";
import dbSearch from "../../helper/dbSearch";
const router = express.Router();

router.get('/', async function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    const q = { type: type };
    await dbSearch(q)
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
  else if (query.name) {
    const name = query.name.split(',');
    const q = { name: new RegExp(name) };
    await dbSearch(q)
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
  else {
    await dbSearch({})
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
});

module.exports = router;
