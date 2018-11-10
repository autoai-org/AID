FROM node:latest

ARG AWS_ACCESS_KEY AWS_SECRET_ACCESS_KEY LOGSENE_TOKEN

RUN export AWS_ACCESS_KEY=$AWS_ACCESS_KEY AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY LOGSENE_TOKEN=$LOGSENE_TOKEN

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/
RUN npm install

COPY . /usr/src/app
RUN npm run build

EXPOSE 3000
CMD ["npm", "start"]