import dbcontroller from '../modules/dbcontroller'

const findRandom = (query) => {
  return new Promise((resolve, reject) => {
    dbcontroller.menu.find(query, function (e, docs) {
      const length = docs.length
      const random = Math.floor(Math.random() * length)
      resolve(docs[random])
    }, err => {
      reject(err)
    })
  })
}

export default findRandom