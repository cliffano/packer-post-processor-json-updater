source "amazon-ebs" "example" {
  ami_name = "{{user `ami_name`}}-hcl"
  region = "ap-southeast-2"
  source_ami = "ami-11032472"
  ssh_username = "ec2-user"
  tags {
    project = "packer-post-processor-json-updater"
  }
}

build {
  sources = ["source.amazon-ebs.example"]

  variables {
    ami_name = "packer-post-process-json-updater {{env `BUILD_ID`}}"
    interpolation_postfix = "with_interpolation"
  }

  post-processor "json-updater" {
    ami_id = {
      testdata/test-integration/success_result.json = [
        "variables.built_ami",
        "variables.built_ami_{{user `interpolation_postfix` }}",
      ]
    }
  }
}
