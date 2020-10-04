# 项目介绍

## 主要内容

- 基于gin、vue框架从0-1完成web应用开发（需要了解gin和vue的基本语法），类似于flask+vue的方式
- 实现基于容器的本地开发环境搭建和说明（主要指的是Dockerfile和本地数据库的搭建）
- 完成基于k8s+gitlab-CICD流水线的各类环境发布工作
- 完成基于prometheus、efk、Jaeger的全面监控体系以及日志集中管理体系
- 

## 主要面向群体
- 主要基于运维开发视角
- 主要考虑转入云原生

## 如何使用
- 已经添加`tag`: `lessonxx-xxx`， 通过切换到不同的tag，查看整个开发过程中代码的变化情况


## Tag说明
- lesson01:  使用gin框架启动一个http服务,访问/目录，返回"pong"
- lesson02:  完成Register功能，添加gitignore以及安装Mysql服务器和配置postman
- lesson03:  按照MVC，优化项目目录结构
- lesson04:  完成用户登陆以及注册用户密码加密功能
- lesson05:  添加json web token功能，用户登陆成功后，下发token
- lesson06:  完成中间件鉴权用户token以及封装响应体内容并替换相应旧代码