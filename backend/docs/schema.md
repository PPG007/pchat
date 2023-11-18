# DB schema

## setting

```json5
{
  "_id": "ObjectId",
  "updatedAt": "DateTime",
  "emailSetting": {
    // 邮件服务地址
    "server": "String",
    // 邮件服务端口号
    "port": "number",
    // 发件人用户名
    "username": "String",
    // 发件人密码
    "password": "String",
    // 离线后通过邮箱发送聊天内容
    "sendEmailIfNotOnline": "Boolean"
  },
  "openAISetting": {
    // openAI API key
    "key": "String",
    // openAI 代理设置
    "proxy": "String",
    // 是否启用 OpenAI
    "isEnabled": "Boolean"
  },
  "accessTokenSetting": {
    // JWT token 密钥
    "key": "String",
    // JWT token 有效期
    "expiredSecond": "number"
  },
  "ossSetting": {
    // OSS 服务提供方，支持 minio
    "provider": "String",
    "bucket": "String",
    // 分享链接过期时间
    "expiredSecond": "number",
    // oss url
    "url": "String"
  },
  "chatSetting": {
    // 是否显示消息已读状态
    "showMessageReadStatus": "Boolean",
    // 是否允许撤回消息
    "allowRollback": "Boolean",
    // 注册完成前是否需要审核
    "mustBeApprovedBeforeRegister": "Boolean"
  }
}
```

## permission

```json5
{
  _id: "ObjectId",
  // 资源名称
  name: "String"
}
```

## role

```json5
{
  _id: "ObjectId",
  // 权限名称集合
  permissions: "[permission.name]"
}
```

## user

```json5
{
  _id: "ObjectId",
  // 昵称
  name: "String",
  // 密码
  password: "String",
  // 邮箱
  email: "String",
  // 角色列表
  roles: "[role._id]",
  createdAt: "Date",
  updatedAt: "Date",
  // 状态，blocked、activated、auditing
  status: "String",
  // 头像
  avatar: "String",
  // 在线状态，online、offline、leaving、busy
  chatStatus: "String"
}
```
