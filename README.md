<img align="right" src="https://raw.github.com/cliffano/packer-post-processor-json-updater/master/avatar.jpg" alt="Avatar"/>

[![Build Status](https://secure.travis-ci.org/cliffano/packer-post-processor-json-updater.png?branch=master)](http://travis-ci.org/cliffano/packer-post-processor-json-updater)

Packer Post-Processor JSON Updater
----------------------------------

A Packer Post-Processor plugin for updating JSON files.

This is handy when you want to update one or more JSON files after Packer already finish executing the builders.

An example scenario is when you have several Packer templates where the built image of a template becomes the source of a number of other templates.

Installation
------------

Download the plugin binaries from [Releases page](https://github.com/cliffano/packer-post-processor-json-updater/releases), or build from source using `make build` or other Go build tools of your choice. Then follow [Packer plugins installation guide](https://www.packer.io/docs/extend/plugins.html).

If you downloaded one of the pre-built binaries, make sure that you rename the binary file to `packer-post-processor-json-updater` (i.e. remove the _os_arch postfix).

Usage
-----

Add `json-updater` type to Packer template's post-processor section:

    "post-processors": [
        {
            "type": "json-updater",
            "ami_id": {
                "templates/child_template_1.json": [
                    "variables.aws_source_ami"
                ],
                "templates/child_template_2.json": [
                    "variables.aws_source_ami"
                ]
            }
        }
    ]

The above `ami_id` configuration indicates that the ID of the newly created AWS AMI will be set as the value of `variables.aws_source_ami` key in `templates/child_template_1.json` and `templates/child_template_2.json` files.

    "variables": {
        "aws_source_ami": "<ami_id>"
    }

Variable interpolation is also supported:

    "variables": {
        "build_id": "build_123"
    },
    "post-processors": [
        {
            "type": "json-updater",
            "ami_id": {
                "templates/child_template_3.json": [
                    "variables.aws_source_ami_{{user `build_id` }}"
                ],
            }
        }
    ]

The above `ami_id` configuration indicates that the ID of the newly created AWS AMI will be set as the value of `variables.aws_source_ami_build_123` key in `templates/child_template_3.json` file.

    "variables": {
        "aws_source_ami_build_123": "<ami_id>"
    }

Colophon
--------

Articles:

* [Custom AWS AMIs Dependency Tree](http://blog.cliffano.com/2016/03/19/custom-aws-amis-dependency-tree/)
