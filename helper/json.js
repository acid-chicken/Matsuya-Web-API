import fs from 'fs'

export class JsonUtil {
  getAllMenu() {
    const filenames = fs.readdirSync('menu')
    const files = filenames.map(filename => {
      const ext = filename.slice((filename.lastIndexOf('.') - 1 >>> 0) + 2)
      if (ext === 'json') {
        return JSON.parse(fs.readFileSync('menu/' + filename, 'utf8'))
      }
    })

    return files.reduce((prev, current) => {
      return Object.assign(prev, current)
    })
  }
}
