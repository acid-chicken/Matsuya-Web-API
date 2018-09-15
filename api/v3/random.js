import express from 'express'
var router = express.Router()
import {JsonUtil} from '../../helper/json'

router.get('/', function (req, res) {
  const jsonUtil = new JsonUtil()
  const menu = jsonUtil.getAllMenu()
  const keys = Object.keys(menu)

  const random = Math.floor(Math.random() * (keys.length - 0)) + 0
  const innerKeys = Object.keys(menu[keys[random]])
  const innerRandom = Math.floor(Math.random() * (innerKeys.length - 0)) + 0
  const selected = menu[keys[random]][innerKeys[innerRandom]]
  res.header('Content-Type', 'application/json; charset=utf-8')
  res.send(selected)
})

module.exports = router
