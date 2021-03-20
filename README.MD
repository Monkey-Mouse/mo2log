# MO2 log service

一个提供log service的grpc微服务  
[mo2](https://github.com/Monkey-Mouse/mo2)项目的依赖项之一

## 用途
各种log，可以完成各种需求。例如用户点赞、打分  

## 使用流程举例
假设文章可以被点赞  
那么点赞时使用这个为服务的流程如下：  
- 调用log方法，level为info，operator设为点赞的人的id，operation设置为1（约定），
  operationTarget设置为被点赞文章的id，operationTargetID设置为被点赞的文章的作者的ID，extra message
  设置为文章标题之类的信息
- 之后用户进来的时候，通过调用Exist获取用户是否已经点赞过文章
- 作者在进主页的时候，调用GetUserNewMsgNum获取未读信息的数量，显示提示
- 用户点进信息列表的时候，调用GetUserMsgs从新到旧获取信息，这个api获取信息之后会自动将之前所有信息设置为已读

## 镜像地址  
```
registry.cn-hangzhou.aliyuncs.com/mo2/mo2log:latest
```

## TODO
- [x] 添加index
- [ ] 完善文档
