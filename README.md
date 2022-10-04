# Docker
A wheel of Docker that implement on linux enviroment

## Change logs

---
### V 0.10
1. 将 containerName 生成放在run命令部分，可以自定义容器ID
2. 调整目标container结构，在containers 下用containerId文件夹包含log文件与rootfs
3. 实现了 detach -d 参数，使容器后台运行挂载在 init 进程下
4. 从 使用 home-made image 转为 从docker hub 导出的 busybox 镜像
5. 将错误代码与提示信息集成到 alert 包下