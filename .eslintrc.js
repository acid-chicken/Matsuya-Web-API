module.exports = {
  "env": {
    "es6": true,
    "node": true,
    jasmine: true
  },
  "extends": ["eslint:recommended"],
  "plugins": ["jasmine"],
  "parserOptions": {
    "ecmaVersion": 2018,
    "sourceType": "module"
  },
  "rules": {
    "indent": [
      "error",
      2
    ],
    "no-console": [
      "off"
    ],
    "linebreak-style": [
      "error",
      "unix"
    ],
    "quotes": [
      "error",
      "single"
    ],
    "semi": [
      "error",
      "never"
    ]
  }
}
