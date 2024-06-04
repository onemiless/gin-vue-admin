打包前后端
前端到web下执行 npm run build 打包前端程序
后端到go下执行go build 打包后端程序
mac执行
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
windows执行
set GOARCH=amd64
set GOOS=linux
go build