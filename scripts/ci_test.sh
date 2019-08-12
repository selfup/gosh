docker build -t gosh .

GOSH=$(docker images gosh -q)

docker rmi $GOSH
