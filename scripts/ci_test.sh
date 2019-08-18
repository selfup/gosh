docker build -t gosh .

GOSH=$(docker images gosh -q)

echo $GOSH

docker rmi $GOSH
