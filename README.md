# 项目名称

这个项目是一个用于自动备份Grafana数据库到华为云对象存储服务(OBS)并将备份状态推送到Prometheus的Go应用程序。

## 功能特性

- 使用华为云OBS SDK进行文件上传。
- 利用Prometheus客户端库推送备份状态指标。
- 支持单次运行备份任务，并通过同步原语`sync.Once`确保任务只执行一次。

## 使用方法

1. **安装依赖**：请确保你已经安装了Go环境，并使用`go get`或`go mod`来安装项目依赖。

```bash
go get -u github.com/huaweicloud/huaweicloud-sdk-go-obs/obs
go get -u github.com/prometheus/client_golang/prometheus
go get -u github.com/prometheus/client_golang/prometheus/push
```

或者使用Go模块（在项目根目录下）：

```bash
go mod init
go mod tidy
```

2. **配置认证信息**：在`main.go`中设置你的华为云OBS的`ak`（Access Key ID）和`sk`（Secret Access Key），以及OBS服务的`endPoint`。

3. **配置Prometheus Pushgateway**：确保Prometheus Pushgateway服务是可用的，并在`pushToPrometheus`函数中配置正确的Pushgateway地址。

4. **运行程序**：直接运行`main.go`，程序将会备份指定的Grafana数据库文件到OBS，并将备份状态推送到Prometheus Pushgateway。

```bash
go run main.go
```

5. **检查结果**：检查OBS存储桶以确认Grafana数据库文件已成功上传，并查看Prometheus以验证备份状态指标是否已被正确推送。

## 注意事项

- 确保你的机器有权限访问Grafana数据库文件、华为云OBS服务以及Prometheus Pushgateway。
- 在生产环境中，不要在代码中硬编码认证信息，应该使用环境变量或配置文件来管理这些敏感信息。
- 此程序为示例代码，可能需要根据你的具体环境进行调整。

## 贡献

如果你有任何改进意见或发现了bug，请随时通过GitHub的Issue跟踪器来报告，或者直接提交Pull Request。

## 许可证

该项目遵循MIT许可证。有关详细信息，请参阅`LICENSE`文件。