#!/bin/bash
mongoimport -d matsuya -c menus --jsonArray --file /docker-entrypoint-initdb.d/menu.json
