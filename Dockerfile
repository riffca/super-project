FROM node:boron


RUN npm install npm -g

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/package.json
RUN cd /usr/src/app && npm install

COPY . /usr/src/app

EXPOSE 8080

CMD ["npm", "run", "dev"]
