# Aegis Terminal

开源、轻量、AI Native 的堡垒机，用于管理多台 Windows / Linux 服务器。

本仓库是 [dushixiang/next-terminal](https://github.com/dushixiang/next-terminal) 最后开源版 **v1.3.9**（提交 `27ca72d`）的分叉。上游 v1.3.9 之后的后端不再开源，本仓库不包含那些代码。感谢原作者 [dushixiang](https://github.com/dushixiang)。

AI Native 是目标方向，当前版本尚未内置 AI API。

## 镜像

- 应用：`ghcr.io/cosmogao/aegis-terminal:main`
- guacd：`ghcr.io/cosmogao/guacd:1.4.0`（不用 `:latest`）

推送到 `main` 后，GitHub Actions 会构建并发布 linux/amd64 镜像。guacd 从 Apache guacamole-server 1.4.0 源码编译并装上仓库字体；不用上游 `dushixiang/guacd:latest`（与 1.6.0 同类，连 Windows 时 RDP 安全协商会失败）。

## 已有能力（继承自 Next Terminal v1.3.9）

支持 RDP、SSH、VNC、Telnet、Kubernetes：

- 资产、授权凭证与用户分组
- 在线会话监控、强制断开，以及离线录屏
- 批量命令、计划任务和登录策略
- 内置 SSH 入口（密码登录）

后续计划（尚未实现）：SSH 公钥登录、按用户授权资产、对外 SSH 入口等。会话协议先不动，guacd 固定 1.4.0。

## 测试部署

在威联通 Portainer 中新建 stack，内容见 [deploy/portainer-stack.yml](./deploy/portainer-stack.yml)。

- 网页：`http://<主机>:18088`，默认账号 `admin` / `admin`，首次登录后请改密
- SSH 入口：`<主机>:18089`，同一账号密码
- 数据目录：`/share/Container/next-terminal-dev/data`（与正式使用的 Next Terminal 分开）

网页里的 SSH / RDP / VNC 走 guacd。两个容器须把同一份数据挂到同一绝对路径 `/usr/local/next-terminal/data`，录屏和网盘才能写入。

## 本地编译

前端用 Node 22。`react-scripts` 5 在 Node 17+ 需要旧的 OpenSSL 摘要：

```shell
cd web
yarn install
NODE_OPTIONS=--openssl-legacy-provider yarn build
cd ..
cp -r web/build server/resource/
CGO_ENABLED=0 go build -ldflags '-s -w' -o next-terminal main.go
```

Go 版本见 `Dockerfile.ci`（1.20）。仓库无 `yarn.lock`，依赖按 `package.json` 范围解析。

## 协议

[AGPL-3.0](./LICENSE)。不能改成别的协议。使用、修改或分发前须遵守该协议；本项目不提供担保。
