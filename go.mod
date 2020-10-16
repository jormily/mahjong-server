module mahjong-server

go 1.15

require (
	github.com/anacrolix/log v0.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/core v0.6.3
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.4.2
	github.com/kudoochui/kudos v1.0.2
	github.com/sirupsen/logrus v1.2.0
	github.com/smallnest/rpcx v0.0.0-20200924044220-f2cdd4dea15a // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/kudoochui/kudos => C:\\GoPath\\src\\github.com\\kudoochui\\kudos
