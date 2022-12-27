### 简介

```
workpool
限流器
```

### 使用

- 开启调度

```
# work_num 任务并发数（工作数）
dispathcher := NewDispathcher(num int)
```

- 加入 job

```
# job 接口
dispathcher.AddJob(job)
```

- 关闭调度

```
dispathcher.Stop()
```
