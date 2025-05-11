FROM node:lts as build-stage

WORKDIR /app
COPY package*.json ./
#RUN npm install --location=global @angular-cli
RUN npm install
COPY . .
#RUN npm --version

RUN npm run build  --configuration=prod
#RUN npm run build --openssl-legacy-provider
#RUN ng build --prod
#RUN npm ionic build --prod

# production stage
#for caddy2
FROM caddy:alpine as production-stage
RUN apk update
WORKDIR /app
COPY Caddyfile .
COPY --from=build-stage /app/dist/ngapp .
EXPOSE 80
CMD ["caddy", "run"]

