FROM node:9.4.0-alpine

WORKDIR /usr/local/src
COPY . .

RUN npm install
CMD npm start

EXPOSE 3000
