#!/bin/bash
mongoimport -d matsuya -c menus --file /docker-entrypoint-initdb.d/menu.json
