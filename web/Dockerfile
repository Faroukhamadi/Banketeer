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