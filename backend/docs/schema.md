# DB schema

## setting

```json5
{
  "_id": "ObjectId",
  "updatedAt": "DateTime",
  "smtp": {
    // 邮件服务地址
    "host": "String",
    // 邮件服务端口号
    "port": "number",
    // 发件人用户名
    "username": "String",
    // 发件人密码
    "password": "String",
    // 发件人名称
    "senderName": "String",
    // 协议，smtp、smtps
    "protocol": "String",
  },
  "ai": {
    // 服务提供者，OpenAI
    "provider": "String",
    "apiKey": "String",
    "proxy": "String",
    "isEnabled": "Boolean",
    "model": "String"
  },
  "account": {
    // token 有效时间
    "TokenValidSecond": "Long",
    // 密码设置
    "password": {
      "isEnabled": "Boolean",
      // 最小长度
      "minLength": "Long",
      // 最大长度
      "maxLength": "Long",
      // 必须包含小写字母
      "mustHasLowerCase": "Boolean",
      // 必须包含大写字母
      "mustHasUpperCase": "Boolean",
      // 必须包含数字
      "mustHasNumber": "Boolean",
      // 必须包含特殊字符
      "mustHasSpecialCode": "Boolean",
    },
    // 注册设置
    "register": {
      // 注册完成前是否需要审核
      "mustBeApprovedBeforeRegister": "Boolean"
    }
  },
  "oss": {
    // OSS 服务提供方，支持 minio
    "provider": "String",
    "publicBucket": "String",
    "privateBucket": "String",
    // 分享链接过期时间
    "validSecond": "number",
    "endpoint": "String",
    "accessKey": "String",
    "secretAccessKey": "String",
    
  },
  "chat": {
    // 是否显示消息已读状态
    "showMessageReadStatus": "Boolean",
    // 是否允许撤回消息
    "allowRollback": "Boolean",
    // 离线后通过邮箱发送聊天内容
    "sendEmailIfNotOnline": "Boolean",
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
  chatStatus: "String",
  // 是否开启双因素认证
  is2FAEnabled: "Boolean",
  // 双因素认证密钥
  otpSecret: "String",
  // 双因素恢复密码
  recoveryCodes: "[String]",
}
```

## chinaHoliday

```json5
{
  _id: "ObjectId",
  date: "Date",
  dateStr: "String",
  isWorkingDay: "Boolean"
}
```

## chat

```json5
{
  _id: "ObjectId",
  isDeleted: "Boolean",
  createdAt: "Date",
  // 群聊：group，私聊：direct
  type: "String",
  // 聊天成员
  members: ["ObjectId"],
  // 群头像
  avatar: "String",
  // 是否非公开
  isPrivate: "Boolean",
  lastMessage: {
    id: "ObjectId",
    createdAt: "Date",
    sender: "ObjectId",
    content: "String"
  }
}
```

## message

```json5
{
  _id: "ObjectId",
  createdAt: "Date",
  isDeleted: "Boolean",
  chatId: "ObjectId",
  // 消息发送者
  sender: "ObjectId",
  // 如果此消息是对某个消息的回复，那么此字段值为被回复消息的 id
  replyTo: "ObjectId",
  // 是否被编辑过
  hasBeenEdited: "Boolean",
  // 如果是讨论串消息，那么此字段值为讨论串 id
  threadId: "ObjectId",
  // 当前消息是否在讨论串内
  isInThread: "Boolean",
  // isInThread 为 true 时，此字段表示是否将消息显示在外部消息中
  showThreadInChat: "Boolean",
  // 对消息的 emoji 回应
  responseEmojis: ["String"],
  // 文本：text，文件：file
  type: "String",
  // 正文，如果是文件那么就是文件名
  content: "String",
  // 文件的 url
  fileUrl: "String",
  // 被提及的人
  mentionedUsers: ["ObjectId"]
}
```

## readStats

```json5
// 消息阅读状态
{
  _id: "ObjectId",
  messageId: "ObjectId",
  userId: "ObjectId",
  createdAt: "Date"
}
```

## userFavorChat

```json5
{
  _id: "ObjectId",
  chatId: "ObjectId",
  userId: "ObjectId"
}
```

## userUnreadMessage

```json5
{
  _id: "ObjectId",
  messageId: "ObjectId",
  chatId: "ObjectId",
  userId: "ObjectId",
}
```

## userSubscribedThread

```json5
{
  _id: "ObjectId",
  chatId: "ObjectId",
  threadId: "ObjectId",
}
```

## registerApplication

```json5
{
  _id: "ObjectId",
  createdAt: "Date",
  reason: "String",
  // 状态：rejected（拒绝），approved（同意），pending（等待）
  status: "String",
  rejectReason: "String",
  updatedAt: "Date",
  email: "String",
  userId: "ObjectId"
}
```


## Todo

### todo

```json5
{
  id: 'ObjectId',
  createdAt: 'DateTime',
  updatedAt: 'DateTime',
  isDeleted: 'Boolean',
  needRemind: 'Boolean',
  content: 'String',
  userId: 'ObjectId',
  images: '[String]',
  remindSetting: {
    remindAt: 'DateTime',
    isRepeatable: 'Boolean',
    lastRemindAt: 'DateTime',
    repeatType: 'String',
    repeatDateOffset: 'Long',
  }
}
```

### todoRecord

```json5
{
  id: 'ObjectId',
  isDeleted: 'Boolean',
  createdAt: 'DateTime',
  updatedAt: 'DateTime',
  remindAt: 'DateTime',
  hasBeenDone: 'Boolean',
  content: 'String',
  todoId: 'ObjectId',
  userId: 'ObjectId',
  hasBeenReminded: 'Boolean',
  images: '[String]',
  doneAt: 'DateTime',
}
```
