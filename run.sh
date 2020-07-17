#!/usr/bin/env bash

go run ./plgo ./example && \
  cd ./build && \
  sudo make install with_llvm=no
