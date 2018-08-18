FROM node:10.9.0-alpine

WORKDIR /usr/local/src
COPY . .

RUN yarn
CMD yarn start

EXPOSE 3000
