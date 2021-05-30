FROM node:14-alpine

WORKDIR /app

COPY . .

RUN npm install

EXPOSE 80

ENTRYPOINT [ "node", "index.js" ]