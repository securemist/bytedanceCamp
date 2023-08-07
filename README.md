# 字节第六届青训营

任务三个模块：

- [x] 基础模块
- [ ] 互动模块
- [ ] 社交模块

初步考虑使用`consul`+`grpc`+`protobuf`作为微服务技术栈，使用`gin`作为web框架，`gorm`作为orm框架，其它的后续慢慢完善。。。



> 本项目启动方式

```bash
// 1. 首先克隆项目 
git clone https://github.com/YakultGo/bytedanceCamp.git
// 2. 进入到文件夹中
cd bytedanceCamp
// 3. 运行cmd/run.bash
bash ./cmd/run.bash
// 4. 如果想终止该项目, 则使用kill命令
ps | grep "main" | grep -v "grep" | awk '{print $1}' | xargs kill
```

