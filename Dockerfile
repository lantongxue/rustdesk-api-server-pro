FROM golang:alpine AS golang
WORKDIR /backend
COPY ./backend .
RUN go build


FROM node:20-alpine AS node
WORKDIR /frontend
COPY ./soybean-admin .
RUN rm -rf node_modules
RUN npm install -g pnpm
RUN pnpm i && pnpm build

FROM alpine:3.20.3
WORKDIR /app
COPY --from=golang /backend/rustdesk-api-server-pro .
COPY --from=golang /backend/server.yaml .
COPY --from=node /frontend/dist ./dist
RUN apk add tzdata
RUN ln -s /app/rustdesk-api-server-pro /usr/local/bin/rustdesk-api-server-pro
RUN ./rustdesk-api-server-pro sync
EXPOSE 8080
CMD [ "./rustdesk-api-server-pro", "start" ]