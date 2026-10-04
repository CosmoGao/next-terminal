# Next Terminal

基于 [dushixiang/next-terminal](https://github.com/dushixiang/next-terminal) v1.3.9（提交 `27ca72d`）的独立分支。v1.3.9 是上游最后一个 AGPL-3.0 版本，v2 及之后的后端不再开源，本仓库不包含那些代码。

镜像：`ghcr.io/cosmogao/next-terminal:main`

推送到 `main` 后，GitHub Actions 会构建前端并发布 linux/amd64 镜像。

## 现在能做什么

这是交互审计系统，支持 RDP、SSH、VNC、Telnet、Kubernetes。v1.3.9 里已经有：

- 资产、授权凭证和用户分组
- 在线会话监控、强制断开，以及离线录屏
- 批量命令、计划任务和登录策略
- 内置 SSH 入口（密码登录）

还没做的，也是这个分支接下来要补的：SSH 公钥登录、按用户授权资产、以及一个对外的 SSH 入口。guacd 和会话层先不动。

## 测试部署

在威联通的 Portainer 里新建一个 stack，内容见 [deploy/portainer-stack.yml](./deploy/portainer-stack.yml)。

- 网页：`http://<主机>:18088`，默认账号 `admin` / `admin`，第一次登录后改掉。
- SSH 入口：`<主机>:18089`，同样是这个账号的密码。
- 数据在 `/share/Container/next-terminal-dev/data`，和正在使用的 Next Terminal 分开。
- `guacd` 请换成你现在能连上 Windows 的那份镜像。`dushixiang/guacd:latest` 对应 1.6.0，连家里的 Windows 会在 RDP 安全协商阶段失败。

网页里的 SSH、RDP、VNC 走 guacd。两个容器必须把同一份数据挂到同一个绝对路径 `/usr/local/next-terminal/data`，录屏和网盘文件才写得进去。

## 本地编译

前端用 Node 22。`react-scripts` 5 在 Node 17 及以后需要旧的 OpenSSL 摘要。

```shell
cd web
yarn install
NODE_OPTIONS=--openssl-legacy-provider yarn build
cd ..
cp -r web/build server/resource/
CGO_ENABLED=0 go build -ldflags '-s -w' -o next-terminal main.go
```

Go 版本按 `Dockerfile.ci`，用 1.20。仓库里没有 `yarn.lock`，依赖会按 `package.json` 的范围重新解析。

## 协议

[AGPL-3.0](./LICENSE)。不能改成别的协议。使用、修改或分发前需要遵守该协议，本项目不提供担保。
