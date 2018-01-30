FROM node:9.4.0-alpine

WORKDIR /usr/local/src
COPY . .

RUN yarn
CMD yarn start

EXPOSE 3000
