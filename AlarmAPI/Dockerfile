FROM node:latest AS builder

RUN mkdir /app
WORKDIR /app
COPY . ./
RUN npm install


FROM node:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/. ./
CMD ["npm", "start"]
