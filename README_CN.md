Rustdesk Api Server Pro
============
[English](https://github.com/rustdesk/rustdesk) | [简体中文](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/README_CN.md)

这是一个基于开源RustDesk客户端的开源Api服务器，实现了客户端所有Api接口，并提供一个Web-UI用于管理数据。

![Preview](./img/login.png "Preview")

> 我们致力于用最简单的代码和结构实现功能！

## 特性
- 同步RuskDesk版本（当前适配客户端：1.2.7）
- 纯Go实现所有接口
- 可视化管理界面
    - 国际化支持
    - 统计面板
    - 用户管理
    - 会话管理
    - 日志审计
    - hbbr&hbbs管理
- 轻量化&跨平台
    - 最小sqlite即可
    - 支持主流操作系统和架构

![Dashboard](./img/dashboard.png "Dashboard")

![Users](./img/users.png "Users")

![Sessions](./img/sessions.png "Sessions")

![Audit](./img/audit.png "Audit")

## 使用Docker部署
1. 拉取镜像
```shell
docker pull ghcr.io/lantongxue/rustdesk-api-server-pro:latest
```
2. 创建配置
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
3. 运行镜像
```shell
docker run \
--name rustdesk-api-server-pro \
-d \
-e ADMIN_USER=admin \ #管理员账号（可选）
-e ADMIN_PASS=yourpassword \ #管理员密码（可选）
-p 8080:8080 \
-v /your/path/server.yaml:/app/server.yaml \
ghcr.io/lantongxue/rustdesk-api-server-pro:latest
```
4. 添加管理员账号（如果设置了用于初始化管理员账号密码的环境变量，此步骤可以忽略，但我仍推荐你使用此方式创建管理员账号，而不是通过环境变量初始化）
```shell
docker exec rustdesk-api-server-pro rustdesk-api-server-pro user add admin yourpassword --admin
```
> 容器镜像默认监听`8080`端口

> 默认配置文件路径`/app/server.yaml`，可以通过`-v`指定您自己的配置文件

### Docker compose
```yaml
services:
  rustdesk-api-server-pro:
    container_name: rustdesk-api-server-pro
    image: ghcr.io/lantongxue/rustdesk-api-server-pro:latest
    environment:
      - "ADMIN_USER=youruser"
      - "ADMIN_PASS=yourpassword"
    volumes:
      - ./server.yaml:/app/server.yaml
    network_mode: host
    restart: unless-stopped
```

### 环境变量

| 变量  | 默认值 | 说明 |
| :------------: | :------------: | :------------: |
|ADMIN_USER|-|默认管理员账号|
|ADMIN_PASS|-|默认管理员密码|

## 源代码编译
### 必要环境
- Golang >= 1.21.4
- NodeJs ~= 推荐最新LTS版本
- pnpm ~= 最新版

### 编译
1. 获取源码

```shell
git clone https://github.com/lantongxue/rustdesk-api-server-pro.git
```

2. 编译api-server

```shell
cd backend && go build
```

3. 编译前端
```shell
cd soybean-admin && pnpm i && pnpm build --outDir dist/admin
```

### 运行

#### api-server
1. 同步数据表结构
```shell
rustdesk-api-server-pro.exe sync
```

2. 创建第一个账号
```shell
rustdesk-api-server-pro.exe user add admin yourpassword --admin
```
> --admin 是一个可选项，启用后添加的用户为管理员用户，否则是普通用户

3. 启动
```shell
rustdesk-api-server-pro.exe start
```
> 默认监听`8080`端口

#### Web管理界面
此步骤你需要一个WEB服务器软件（例如：nginx、apache等），通过将打包后的产物复制到WEB根目录即可。

一般情况下，打包后的产物在`soybean-admin/dist`目录中。

反向代理配置，你需要将在`nginx`或其他WEB服务器中配置反向代理，通过反向代理服务端才能正确访问到接口地址。

下面是`nginx`反向代理的参考配置：
```nginx
#PROXY-START
location ~ ^/(api|admin)
{
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_http_version 1.1;
    # proxy_hide_header Upgrade;

    add_header X-Cache $upstream_cache_status;
}
#PROXY-END/
```

## CLI命令行
```shell
Usage:
  rustdesk-api-server-pro [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  rustdesk    About rustdesk-server command
  start       Start the api-server
  sync        The api-server database synchronization
  user        User management

Flags:
  -h, --help   help for rustdesk-api-server-pro

Use "rustdesk-api-server-pro [command] --help" for more information about a command.
```

## 后续计划
后续会持续跟进RustDesk客户端，并实现对应接口，这将作为一个长期计划。

## 赞助

如果您觉得此项目对您有所帮助，不妨请开发者喝杯咖啡 :)

![Sponsorship](./soybean-admin/src/assets/imgs/sponsorships.png "Sponsorship")

**感谢您的赞助**

## 开源许可
>您可以查看完整的许可证 [这里](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/LICENSE)

本项目采用**MIT**许可条款。
