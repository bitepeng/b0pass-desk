#! /bin/bash

# windows
 wails build -upx -webview2 embed -nsis
# mac
 wails build -webview2 embed -platform darwin/universal