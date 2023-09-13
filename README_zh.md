# Colossus - SSH远程服务器管理工具

[English](README.md)

> 目前仅支持简体中文, 其他语言等待后续开发

[![](https://img.shields.io/badge/Go-1.21+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/Version-0.0.1%20beta1-green)](control)
[![CodeQL](https://github.com/skye-z/colossus/workflows/CodeQL/badge.svg)](https://github.com/skye-z/colossus/security/code-scanning)


[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=bugs)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)

## 特性

* 小巧: 应用体积小于30MB, 核心进程内存开销小于50MB
* 扩展: 内置扩展工具动态加载器, 可快速开发扩展工具
* 高效: 内置快捷命令、文件管理等高效工具
* 安全: 支持SSH证书认证, 同时内置证书管理器
* 美观: 针对MacOS优化, 提供系统专属样式

## 扩展工具开发

如果你想为Colossus开发扩展工具, 可以查看这一份[指南](https://github.com/skye-z/colossus-frontend/blob/main/src/tools/README_zh.md).