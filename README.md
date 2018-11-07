## novels

novels 是一个小说/影视更新通知程序, 通知方式: 邮件、微信, 推荐使用邮件方式，配置简单，微信通知要复杂一些.

支持Linux标准的crontab表达式周期性的检查

每天7-22点，每5分钟检查一次，默认是每隔5分钟检查一次  */5 * * * *
```sh
*/5 7-22 * * *
```

每周五的7-22点，每10分钟检查一次
```sh
*/10 7-22 * * 5
```


## 通知效果

**邮件标题**

![凡人修仙传](demo/fanren_subject.png)


**小说通知效果 提供最新的章节链接**

![凡人修仙传](demo/fanren.png)

**动漫通知效果 提供的下载地址**

![星辰变](demo/xcb.png)


## 下载安装

直接 [下载](https://github.com/chanyipiaomiao/novels/releases) 二进制包运行

## 使用步骤

1. 如果使用邮件报警,推荐使用 [SendCloud](https://sendcloud.sohu.com/) 进行邮件通知, 每天可以免费50封邮件，足够使用

2. 注册之后拿到 api_user 和 api_key, 到 conf/app.conf 中配置

```go
[email]
# 发送邮件方式 sendcloud | selfhost
type = sendcloud

# 收件人
to = xxxxx@163.com

[email_sendcloud]
# sendcloud服务

# from 邮件地址
from = xxxxx@xxxxxx.com

# fromName
fromName = 更新通知xxxxx@xxxxxx.com
api_user = xxxxxxxxxxxx
api_key = xxxxxxxxxxxx
api_url = http://api.sendcloud.net/apiv2/mail/send
```

3.使用帮助

```go
D:\workspace\Go\src\novels>novels --help
usage: novels [<flags>] <command> [<args> ...]

novels is a novel update notice program

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  novel --action=ACTION [<flags>]
    novel增删改查

  parser [<flags>]
    解析器

  cron --expr=EXPR
    cron表达式

```

