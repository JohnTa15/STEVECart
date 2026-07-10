# https://docs.docker.com/guides/vuejs/
FROM node:22-alpine AS builder

WORKDIR /app

COPY package*.json ./

# install dependencies
RUN --mount=type=cache,target=/root/.npm npm ci

# copy the code of the app into the container 
COPY . .

RUN npm run build

FROM nginx:alpine AS runner

COPY nginx.conf /etc/nginx/nginx.conf

COPY --chown=nginx:nginx --from=builder /app/dist /usr/share/nginx/html

USER root 

# nginx listens to 8080 port instead of the default 80
EXPOSE 5173

# starts nginx
ENTRYPOINT ["nginx", "-c", "/etc/nginx/nginx.conf"]

CMD ["-g", "daemon off;"]