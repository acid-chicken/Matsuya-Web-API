import express from "express";
import findRandom from "../../helper/findRandom";
const router = express.Router();

router.get('/', async function (req, res) {
  const query = req.query;
  if (query.type) {
    const type = query.type.split(',');
    const q = { type: new RegExp(type) };
    await findRandom(q)
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
  else if (query.name) {
    const name = query.name.split(',');
    const q = { type: new RegExp(name) };
    await findRandom(q)
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
  else {
    await findRandom({})
      .then(menu =>res.send(menu))
      .catch(ex => res.send(ex));
  }
});

module.exports = router;
