# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.2] 2022-03-18

### Added
* The CI process now runs the tool in the context of a GitHub Actions runner, which is an explicitly desired support case.

### Fixed
* The tool no longer checks if a version of Terraform is installed before executing the `tfenv install` command.
  When run in the context of a machine where there are no Terraform versions already installed by `tfenv`, the tool 
  would fail because `tfenv` returns a non zero exit code in this case. Rather than introduce logic to try and capture
  this edge case, the "check-before-install" logic has just been removed, since the `install` command is idempotent.
  
## [0.1.1] 2022-03-17

### Fixed
* StdError for invoked processes is now always logged to the console. Previously it was only logged if a method input was set to `true`.

## [0.1.0] 2022-03-16

This is the initial release for the tool.
