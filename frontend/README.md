# frontend

## version
node: 16.15.1
npm: 8.12.1
nuxt: v2.15.8
vue: 2.6.14

## init
```
$ npm install
$ docker exec -it puzzles_frontend npm run dev
```

## test
```
$ npm run cypress:open
```

## deploy
```:.env
API_URL=`URL for API`
BASE_URL=`base URL`
MEDIA_URL=`URL for media storage`
```
```
$ npm run generate
$ firebase deploy --only hosting
```

## Account
organizations: quaereoli5ath2iPiequaidoo
email: jounetsism@gmail.com

