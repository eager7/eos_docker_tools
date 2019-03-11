#!/bin/sh
curl -X POST http://localhost:13141/file_update \
-H "Content-Type: multipart/form-data" \
-F "account=pct" \
-F "file=@./CMakeLists.txt" \
-F "file=@./hello.cpp" \
-F "file=@./hello.hpp"