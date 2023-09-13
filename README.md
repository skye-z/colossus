# Colossus - SSH Remote Server Management Tool

[中文](README_zh.md)

> Currently only supports Simplified Chinese, other languages are waiting for development.

[![](https://img.shields.io/badge/Go-1.21+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/Version-0.0.1%20beta1-green)](control)
[![CodeQL](https://github.com/skye-z/colossus/workflows/CodeQL/badge.svg)](https://github.com/skye-z/colossus/security/code-scanning)


[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=skye-z_colossus&metric=bugs)](https://sonarcloud.io/summary/new_code?id=skye-z_colossus)

## Features

* Compact: application size is less than 30MB, core process memory overhead is less than 50MB
* Extension: built-in extension tools dynamic loader, can quickly develop extension tools
* Efficient: built-in shortcut commands, file management and other efficient tools
* Security: support for SSH certificate authentication, and built-in certificate manager
* Aesthetics: optimized for MacOS, providing system-specific styles

## Extension Development

If you want to develop extension tools for Colossus, check out this [guide](https://github.com/skye-z/colossus-frontend/blob/main/src/tools/README.md).