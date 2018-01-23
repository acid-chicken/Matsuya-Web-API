import mongoose from "mongoose";
require("dotenv").config();

// DB関連設定
const dburl = process.env.NODE_MONGODB_CONNECTION_URL;
const dbport = process.env.NODE_MONGODB_CONNECTION_PORT;
const dbuser = process.env.NODE_MONGODB_CONNECTION_USER;
const dbpass = process.env.NODE_MONGODB_CONNECTION_PASS;
const dbconnectionstring = "mongodb://" + dbuser + ":" + dbpass + "@" + dburl + ":" + dbport + "/matsuya";

const Schema = mongoose.Schema;

const menu = new Schema ({
  _id: Number,
  name: String,
  type: String,
  price: Number,
  calorie: Number,
  protein: Number,
  lipid: Number,
  carbohydrate: Number,
  sodium: Number,
  saltEquivalent: Number,
  description: String,
  imageURL: String
});

let menudata, data;

menudata = mongoose.model("menus", menu);

data = { menu: menudata };

mongoose.connect(dbconnectionstring, function (error) {
  if (error) {
    console.log(error);
  } else {
    console.log("DB connection success!");
  }
});

module.exports = data;
