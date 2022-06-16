# fix this if new one doesn't work
# FROM node:17 AS builder
# WORKDIR /app
# COPY . .
# RUN npm install && npm run-script build

# FROM nginx:alpine
# WORKDIR /etc/nginx/conf.d
# RUN rm -f default.conf
# COPY ./nginx /etc/nginx/conf.d
# WORKDIR /usr/share/nginx/html
# RUN rm -rf ./*
# RUN ls -l /app
# COPY --from=builder /app/package.json /app/build ./
# ENTRYPOINT ["nginx", "-g", "daemon off;"]

FROM node:17
WORKDIR /app
COPY . .
RUN npm ci
RUN npm audit fix
RUN npm run build
FROM node:17
WORKDIR /app
COPY --from=0 /app/package*.json ./
RUN npm ci --production --ignore-scripts
RUN npm audit fix
COPY --from=0 /app/build ./
EXPOSE 3000
CMD ["node", "./index.js"]