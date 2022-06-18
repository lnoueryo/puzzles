# puzzles
backend及びfrontendで分割

## version
go: 1.18.3
node: 16.15.1
npm: 8.12.1
nuxt: v2.15.8
vue: 2.6.14

## docker
docker-compose up -d

## docker frontend
docker exec -it puzzles_frontend npm run dev

## docker backend
docker exec -it puzzles_backend go run main.go watch