# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

## 2.0.0 - 2020-03-16
### Changed
- Upgrade to golang 1.14
- Ignore bin/, pkg/, and src/
- Migrate packer from mitchellh/packer to hashicorp/packer
- Migrate golint from github.com/golang/lint/golint to golang.org/x/lint/golint
- Move post-processor and updater from main package to json-updater package
- Migrate deps management from go get to go modules
- Move golint and gox tools installation to go modules

### Removed
- Remove godep usage due to being retired since golang 1.8

## 1.2 - 2017-11-28
### Changed
- A new JSON file containing empty object will be created when file to be updated does not exist

## 1.1 - 2016-03-14
### Added
- Add variable interpolation support to template key [Michael Bicz](https://github.com/bemehow)

## 1.0 - 2015-04-01
### Added
- Initial version
