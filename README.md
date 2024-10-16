Rustdesk Api Server Pro
============

[English](https://github.com/rustdesk/rustdesk) | [简体中文](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/README_CN.md)

This is an open source Api server based on the open source [RustDesk](https://github.com/rustdesk/rustdesk) client, the implementation of the client all Api interfaces, and provides a Web-UI for the management of data.

![Preview](./img/login.png)

> We strive to achieve functionality with the simplest possible code and structure!

## Features
- Synchronized RuskDesk version (Currently adapted client: 1.2.7)
- Pure Go implementation of all interfaces
- Visual management interface
    - Internationalization support
    - Statistics panel
    - User Management
    - Session Management
    - Log Audit
    - hbbr&hbbs management
- Lightweight & Cross Platform
    - Minimal sqlite
    - Support for major operating systems and architectures

![Dashboard](./img/dashboard.png "Dashboard")

![Users](./img/users.png "Users")

![Sessions](./img/sessions.png "Sessions")

![Audit](./img/audit.png "Audit")

## Deploying with Docker
1. pull image
```shell
docker pull ghcr.io/lantongxue/rustdesk-api-server-pro:latest
```
2. create config
```shell
cat > /your/path/server.yaml <<EOF
signKey: "sercrethatmaycontainch@r$32chars" # this is the token signing key. change this before start server
db:
  driver: "sqlite"
  dsn: "./server.db"
  timeZone: "Asia/Shanghai" # setting the time zone fixes the database creation time problem
  showSql: true

  # driver: "mysql"
  # dsn: "root:123@tcp(localhost:3306)/test?charset=utf8mb4"
httpConfig:
  printRequestLog: true
  port: ":8080" # api server port
jobsConfig:
  deviceCheckJob:
    duration: 30
EOF
```
3. run image
```shell
docker run \
--name rustdesk-api-server-pro \
-d \
-e ADMIN_USER=admin \ #Administrator account (optional)
-e ADMIN_PASS=yourpassword \ #Administrator password (optional)
-p 8080:8080 \
-v /your/path/server.yaml:/app/server.yaml \
ghcr.io/lantongxue/rustdesk-api-server-pro:latest
```
4. add your admin account(This step can be ignored if an environment variable is set to initialize the administrator account password, but I still recommend that you create the administrator account this way instead of initializing it with an environment variable)
```shell
docker exec rustdesk-api-server-pro rustdesk-api-server-pro user add admin yourpassword --admin
```

> The container image listens on port `8080` by default.

> Default configuration file path `/app/server.yaml`, you can specify your own configuration file with `-v`.

### Environment variables

| Variables | Default Values | Description |
| :------------: | :------------: | :------------: |
|ADMIN_USER|-|Default administrator account|
|ADMIN_PASS|-|Default administrator password|

## Build from source
### Required
- Golang >= 1.21.4
- NodeJs ~= latest(recommend LTS)version
- pnpm ~= latest

### Build
1. Get source code

```shell
git clone https://github.com/lantongxue/rustdesk-api-server-pro.git
```

2. Build the api-server

```shell
cd backend && go build
```

3. Build the frontend
```shell
cd soybean-admin && pnpm i && pnpm build
```

### Run

#### api-server
Assuming the compiled binary file is called `rustdesk-api-server-pro.exe`.

1. Synchronize the database table structure
```shell
rustdesk-api-server-pro.exe sync
```

2. Add your first user
```shell
rustdesk-api-server-pro.exe user add admin yourpassword --admin
```
> --admin is optional, when enabled the added user is an administrator user, otherwise it is a regular user

3. Start the server
```shell
rustdesk-api-server-pro.exe start
```
> Listening on port `8080` by default

#### Web Management Interface
For this step you need a web server software (e.g. nginx, apache, etc.), by copying the packaged product to the web root directory.

Typically, the packaged product is in the `soybean-admin/dist` directory.

Reverse Proxy Configuration, you need to configure reverse proxy in `nginx` or other WEB servers, through the reverse proxy server can access the interface address correctly.

Here's my backend reverse proxy configuration for you to refer to:
```nginx
#PROXY-START /api for rustdesk client
location ^~ /api
{
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header Host 127.0.0.1;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_http_version 1.1;
    # proxy_hide_header Upgrade;

    add_header X-Cache $upstream_cache_status;
}
#PROXY-END/

#PROXY-START /admin for web-ui
location ^~ /admin
{
    proxy_pass http://127.0.0.1:8080/admin;
    proxy_set_header Host 127.0.0.1;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_http_version 1.1;
    # proxy_hide_header Upgrade;

    add_header X-Cache $upstream_cache_status;
}
#PROXY-END/
```

## CLI help
```shell
Usage:
  rustdesk-api-server [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  rustdesk    About rustdesk-server command
  start       Start the api-server
  sync        The api-server database synchronization
  user        User management

Flags:
  -h, --help   help for rustdesk-api-server

Use "rustdesk-api-server [command] --help" for more information about a command.
```

## Follow-up plan
We will continue to follow up the RustDesk client and implement the corresponding interfaces, which will be a long-term plan.

## Sponsorship

If you found this project helpful, why not buy the developers a cup of coffee :)

![Sponsorship](./soybean-admin/src/assets/imgs/sponsorships.png "Sponsorship")

**Thank you for your sponsorship**

## License
>You can view the full license [here](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/LICENSE)

This project is under the terms of the **MIT** license.