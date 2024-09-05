Rustdesk Api Server Pro
============
[English](https://github.com/rustdesk/rustdesk) | [简体中文](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/README_CN.md)

这是一个基于开源RustDesk客户端的开源Api服务器，实现了客户端所有Api接口，并提供一个Web-UI用于管理相关数据。

![Preview](./img/login.png)

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

![Dashboard](./img/dashboard.png)

![Users](./img/users.png)

![Sessions](./img/sessions.png)

![Audit](./img/audit.png)

## 编译和运行
### 基础环境
- Golang >= 1.21.4
- NodeJs ~= 最新长期支持版本

### 编译
- 克隆仓库代码
    - `git clone https://github.com/lantongxue/rustdesk-api-server-pro.git`
- 编译服务端程序
    - `cd backend && go build`
- 打包Web管理界面
    - `cd soybean-admin`
    - `npm -g install pnpm`
    - `pnpm build`

### 运行

#### 服务端
假设编译之后的二进制文件叫`rustdesk-api-server-pro.exe`
- 同步数据库表结构
    - `rustdesk-api-server-pro.exe sync`
- 运行服务端
    - `rustdesk-api-server-pro.exe start`
    - 默认监听`8080`端口

#### Web管理界面
此步骤你需要一个WEB服务器软件（例如：nginx、apache等），通过将打包后的产物复制到WEB根目录即可。

一般情况下，打包后的产物在`soybean-admin/dist`目录中。

反向代理配置，你需要将在`nginx`或其他WEB服务器中配置反向代理，通过反向代理服务端才能正确访问到接口地址。

## 后续计划
后续会持续跟进RustDesk客户端，并实现对应接口，这将作为一个长期计划。
### 短期计划
- [ ] 继续实现1.2.7版本中的新接口
- [ ] 增加Device管理
- [ ] 增加用户登录日志
- [ ] 增加hbbr&hbbs管理
- [ ] 增加普通用户登录权限
- [ ] Dockerfile

## 赞助

如果您觉得此项目对您有所帮助，不妨请开发者喝杯咖啡 :)

<img src="./soybean-admin/src/assets/imgs/sponsorships.png" style="height: auto !important;width: 200px;" />

**感谢您的赞助**

## 开源许可
>您可以查看完整的许可证 [这里](https://github.com/lantongxue/rustdesk-api-server-pro/blob/master/LICENSE)

本项目采用**MIT**许可条款。
