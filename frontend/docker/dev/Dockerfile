FROM node:16-alpine3.14

WORKDIR /app
COPY . /app
RUN apk update && \
    npm install -g npm 

EXPOSE 3000
ENV HOST 0.0.0.0
CMD ["npm", "run", "dev"]