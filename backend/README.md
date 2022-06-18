# backend

## version
go: 1.18.3

## init
```:.env.dev
APP_ENV=`"local"`
APP_HOST=`"localhost"`
ALLOW_ORIGIN=`"http://localhost:3000"`
CREDENTIALS_PATH=`"credentials/"service account json file`
DB_NAME=`"puzzle"`
DB_HOST=`"mysql"`
DB_USER=`"root"`
DB_PASSWORD=`password`
DB_PORT=`"3306"`
DB_QUERY=`"parseTime=true"`
APP_ORIGIN=`"http://localhost:8080"`
EMAIL_FROM=`Gmail`
EMAIL_USERNAME=`Gmail`
EMAIL_PASSWORD=`Gmail password`
```
```
$ docker exec -it puzzles_backend go run main.go watch
```

## test
```
```

## deploy
```:.env.dev
APP_ENV=`"production"`
APP_HOST=`app host`
ALLOW_ORIGIN=`app origin`
PROJECT=`GCP project`
DB=`"CLOUDSQL"`
DB_NAME=`db name`
DB_HOST=`db host`
DB_USER=`db user`
DB_PASSWORD=`db password`
DB_PORT=`"3306"`
DB_QUERY=`"parseTime=true"`
SESSION=`"DATASTORE"`
APP_ORIGIN=`"api origin"`
```
```
$ go run main.go deploy
```