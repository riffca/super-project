docker build -t foo/nginx .
docker run -p 8000:80 --link node-app:app --name nginx-proxy foo/nginx
