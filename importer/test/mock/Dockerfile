FROM mhart/alpine-node:8
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN npm install
CMD ["node", "index.js"]