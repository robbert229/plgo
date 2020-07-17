#!/usr/bin/env bash

cd ../plgo && \
  go install && \
  cd ../test && \
  plgo && \
  cd build && \
  sudo make install with_llvm=no && \
  cd .. && \
  psql -U root -c "drop extension if exists test" postgres && \
  psql -U root -c "create extension test" postgres && \
  psql -U root -c "select plgotest()" postgres
