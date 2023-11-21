# 百灵快传 桌面版

## About

基于 [Wails](https://wails.io/docs/gettingstarted/installation) 开发的 [百灵快传](https://github.com/bitepeng/b0pass) 桌面版程序


## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.

Or build with [UPX](https://upx.github.io/) for better distribution size: ``wails build -upx -upxflags="--best --lzma"``

To use ``UPX``, you need to download and at least put the path in the ``System Enviroment Variables``