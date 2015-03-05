FROM nginx
MAINTAINER Alex Welch

ADD . /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
