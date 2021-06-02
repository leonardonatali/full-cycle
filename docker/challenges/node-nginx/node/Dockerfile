FROM node:14-alpine

WORKDIR /app

RUN npm install -g nodemon

COPY src/ /app/

ENTRYPOINT [ "node /app/server.js" ]