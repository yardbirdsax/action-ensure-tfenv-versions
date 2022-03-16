# ensure-tfenv-versions

A tool written in Golang that will recursively look for .terraform-version files and install the requested Terraform version using `tfenv`.

## Usage

* Search for and install all versions of Terraform specified in files found under the current shell's directory: `ensure-tfenv-versions`
* Search for and install all versions of Terraform specified in files found under another directory: `ensure-tfenv-versions -d "./some-directory"`

## Contributions

> **Before submitting a pull request for anything other than a bug fix, please open a GitHub issue so we can discuss if the feature or change is an appropriate one.**

Contributions are welcome, however I can make no guarantees around when they will be reviewed or released. 
