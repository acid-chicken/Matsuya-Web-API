FROM node:9.4.0-alpine

WORKDIR /usr/local/src
COPY . .

RUN yarn install
CMD yarn start

EXPOSE 3000
