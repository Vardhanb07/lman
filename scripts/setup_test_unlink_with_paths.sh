#!/bin/bash

mkdir ./test_unlink_with_paths
ln -s ./testdata/files/test{1..4} ./test_unlink_with_paths
rm ./test_unlink_with_paths/test3
