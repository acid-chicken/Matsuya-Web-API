import {buildSchema} from 'graphql';
import {JsonUtil} from '../../helper/json';

exports.Schema = buildSchema(`
  type RandomMenu {
    name: String!
    calorie: String
    protein: String
    lipid: String
    carbohydrate: String
    sodium: String
    SaltEquivalent: String
    description: String
    imageURL: String
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
    constructor() {
        const jsonUtil = new JsonUtil();
        const menu = jsonUtil.getAllMenu();
        const keys = Object.keys(menu);

        const random = Math.floor(Math.random() * (keys.length - 0)) + 0;
        const innerKeys = Object.keys(menu[keys[random]]);
        const innerRandom = Math.floor(Math.random() * (innerKeys.length - 0)) + 0;
        this.selected = menu[keys[random]][innerKeys[innerRandom]];
        console.log(this.selected);
    }

    name() {
        return this.selected['name'];
    }
    calorie() {
        return this.selected['calorie'];
    }
    protein() {
        return this.selected['protein'];
    }
    lipid() {
        return this.selected['lipid'];
    }
    carbohydrate() {
        return this.selected['carbohydrate'];
    }
    sodium() {
        return this.selected['sodium'];
    }
    SaltEquivalent() {
        return this.selected['SaltEquivalent'];
    }
    description() {
        return this.selected['description'];
    }
    imageURL() {
        return this.selected['imageURL'];
    }
    
}
