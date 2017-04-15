docker build -t foo/node .
docker run -d -p 3000:3000 --name node-app foo/node
docker run --rm -p 8000:80 nginx
