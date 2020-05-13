FROM node:latest as builder
WORKDIR /app
COPY ./package.json ./
RUN npm install
COPY . .
RUN npm run build

# Second Stage : Setup command to run your app
FROM node:latest
WORKDIR /app
COPY --from=builder /app ./
EXPOSE 3000
CMD ["npm", "run", "serve:prod"]