
地址：
  https://trpc.group/zh/docs/languages/go/
  https://github.com/trpc-group/trpc-cmdline/blob/main/README.zh_CN.md

1. 安装 trpc-cmdline
PS E:\> go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
PS E:\> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

2. 安装依赖
PS E:\> trpc setup
[setup][Info][setup.go:41] Initializing setup && Installing dependency tools
[setup][Debug][dependencies.go:85] mockgen.exe check passed
[setup][Debug][dependencies.go:85] protoc.exe check passed
[setup][Debug][dependencies.go:85] flatc.exe check passed
[setup][Debug][dependencies.go:85] protoc-gen-go.exe check passed
[setup][Debug][dependencies.go:85] protoc-gen-validate.exe check passed
[setup][Debug][dependencies.go:85] protoc-gen-validate-go.exe check passed
[setup][Debug][dependencies.go:85] goimports.exe check passed
[setup][Info][setup.go:57] Setup completed

3. 创建proto文件
E:\trpc_hello.proto

4. 生成生成完整项目
PS E:\> trpc create -p trpc_hello.proto -o TrpcToday
[create] Create tRPC project `trpc_hello`: succeed! ヾ(@^▽^@)ノ
[create] Create tRPC project `trpc_hello` post process: succeed! (〃'▽'〃)

5. 修改trpc_go.yml
  server -> service -> protocol: http
6. 启动服务
PS E:\TrpcToday> go run .
2024-04-13 01:14:09.723 DEBUG   maxprocs/maxprocs.go:47 maxprocs: Leaving GOMAXPROCS=20: CPU quota undefined
2024-04-13 01:14:09.724 INFO    server/service.go:167   process:9372, trpc service:trpchello.HelloService launch success, tcp:127.0.0.1:8000, serving ...

7. http访问

    curl -X POST -H 'Content-Type: application/json' -i 'http://127.0.0.1:8000/trpchello.HelloService/HandleFirst' --data '{"message":"13345566"}'

    response: 
        {
        "code": 100,
        "message": "return client message:13345566",
        "timestamp": "1712942933"
        }

8. 修改proto文件，更新桩代码：
    注意，修改的是E:\trpc_hello.proto。proto有go_package 的git仓库更新，则没有这个问题，可直接在项目目录更新
PS E:\> trpc create -p trpc_hello.proto -o .\TrpcToday\stub\github.com\trpcproto\trpchello\ --rpconly -f
    执行桩代码更新后，proto会复制到stub目录。proto有go_package 的git仓库更新，则没有这个问题，可直接在项目目录更新
9. 访问新接口：

    curl -X POST -H 'Content-Type: application/json' -i 'http://127.0.0.1:8000/trpchello.HelloService/HandleSeconds' --data '{"message":"13345566"}'
    response：
        {
        "code": 200,
        "message": "second success",
        "timestamp": "1712945732"
        }
