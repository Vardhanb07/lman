#!/usr/bin/env bash

mkdir test_unlink
mkdir test_unlink/links
touch test_unlink/test
ln -s test_unlink/test test_unlink/links/test
touch test_unlink/lman.config.toml
printf "[[links]]\nfilepath = \"./test\"\nlinkpath = \"./links\"" > test_unlink/lman.config.toml
