FROM node:16-buster

WORKDIR /usr/src/app

COPY ["personal_page/package.json", "personal_page/yarn.lock", "./"]

RUN yarn install

COPY personal_page/ .

CMD [ "yarn", "start" ]