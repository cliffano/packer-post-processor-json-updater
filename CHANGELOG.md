### 2.0.0-pre.0
* Upgrade to golang 1.14
* Ignore bin/, pkg/, and src/
* Migrate packer from mitchellh/packer to hashicorp/packer
* Migrate golint from github.com/golang/lint/golint to golang.org/x/lint/golint
* Move golint and gox usage to use the binaries on bin/
* Remove godep usage due to being retired since golang 1.8
* Move post-processor and updater from main package to json-updater package
* Migrate deps management from go get to go modules

### 1.2
* A new JSON file containing empty object will be created when file to be updated does not exist

### 1.1
* Add variable interpolation support to template key [Michael Bicz](https://github.com/bemehow)

### 1.0
* Initial version
