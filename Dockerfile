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
RUN /app/rustdesk-api-server-pro sync
EXPOSE 8080
CMD [ "/app/rustdesk-api-server-pro", "start" ]