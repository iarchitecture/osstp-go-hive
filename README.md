# osstp-go-hive

# 安装工具
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
或 		   GOPROXY='https://proxy.golang.com.cn,direct'
go mod init osstp-go-hive
go mod tidy
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/swaggo/swag/cmd/swag
go get -u -v github.com/swaggo/gin-swagger
go get -u -v github.com/swaggo/files
go get -u -v github.com/alecthomas/template
go install github.com/swaggo/swag/cmd/swag@latest
swag fmt
swag init
go work init
go get -u github.com/gin-gonic/gin
