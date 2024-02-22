#! /usr/bin/env bash

# This script is used to audit the licences of the project. It will check the
# licences of all the dependencies, including indirect ones, and print a report
# of the licences found.

# This script is meant to be run from the root of the project.

# go-licence-detector install instructions: https://github.com/elastic/go-licence-detector

go list -m -json all | go-licence-detector -includeIndirect \
    -depsTemplate=audit/templates/dependencies.asciidoc.tmpl -depsOut=dependencies.asciidoc \
    -noticeTemplate=audit/templates/NOTICE.txt.tmpl -noticeOut=NOTICE.txt \
    -overrides=overrides.txt -rules=rules.json