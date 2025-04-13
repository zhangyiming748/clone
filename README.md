# 0x01 

自行登录GitHub的gh ~~我这边使用最稳妥的token方式~~

# 0x02

获取你自己的全部仓库信息

```bash
gh repo list --limit 150 | tee repo.list
```

# 0x03

运行程序批量克隆仓库

```bash
go run main.go
```

或者如果没有或不想安装go sdk 可以使用对应你系统架构的[release](https://github.com/zhangyiming748/clone/releases)
程序和仓库列表文件`repo.list`放在同一个目录下,运行程序即可

