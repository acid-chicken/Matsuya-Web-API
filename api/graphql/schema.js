import {buildSchema} from 'graphql';
import fs from 'fs';

exports.Schema = buildSchema(`
  type RandomMenu {
    name: String!
  }

  type Query {
    random: RandomMenu
  }
`);

exports.Root = {
    random: function () {
        return new RandomMenu();
    }
};

class RandomMenu {
    constructor() {}

    name() {
        const all = JSON.parse(fs.readFileSync('menu/all.json', 'utf8'));
        const random = Math.floor(Math.random() * (44 + 1 - 0)) + 0;
        return all[random];
    }
}
