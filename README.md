### pstatus

pstatus是一个简单进程状态查看工具, 可以查看进程的CPU内存使用率、打开的连接数、打开的文件数等.

效果如下:

```shell script
[xxxx@xxxx_ansible_001 xxxx]$ sudo ./pstatus -t 20
┌─────┬──────────────────┬────────┬─────────────────────┬────────────────┬───────────────────┬───────────────────────┐
│ CPU │ CPU Used Percent │ Memory │ Memory Used Percent │ Process Number │ System Open Files │ Max System Open Files │
├─────┼──────────────────┼────────┼─────────────────────┼────────────────┼───────────────────┼───────────────────────┤
│ 2   │ 7.46%            │ 16     │ 68.20%              │ 130            │ 5312              │ 3259355               │
└─────┴──────────────────┴────────┴─────────────────────┴────────────────┴───────────────────┴───────────────────────┘
┌───────┬───────────────────┬──────────┬─────────────────────────┬───────┬────────┬─────────────┬────────────┬────────────────┐
│ PID   │ Name              │ Username │ Exe                     │ CPU   │ Mem    │ Connections │ Open Files │ Max Open Files │
├───────┼───────────────────┼──────────┼─────────────────────────┼───────┼────────┼─────────────┼────────────┼────────────────┤
│ 9665  │ java              │ jenkins  │ /usr/lib/jvm/java-1.... │ 0.63% │ 10.43% │ 3           │ 344        │ 65535          │
│ 1633  │ master            │ root     │ /usr/libexec/postfix... │ 0.00% │ 0.01%  │ 24          │ 88         │ 4096           │
│ 13075 │ mongod            │ mongod   │ /usr/bin/mongod         │ 0.27% │ 1.52%  │ 4           │ 73         │ 65535          │
│ 17392 │ AliYunDun         │ root     │ /usr/local/aegis/aeg... │ 0.66% │ 0.54%  │ 10          │ 51         │ 1024           │
│ 15484 │ java              │ root     │ /usr/local/cloudmoni... │ 0.46% │ 1.00%  │ 1           │ 43         │ 4096           │
│ 17373 │ AliYunDunUpdate   │ root     │ /usr/local/aegis/aeg... │ 0.03% │ 0.05%  │ 7           │ 28         │ 1024           │
│ 21559 │ supervisord       │ root     │ /usr/bin/python-2.6.... │ 0.04% │ 0.20%  │ 1           │ 23         │ 65535          │
│ 10974 │ devops-dns-server │ root     │ /data/app/devops-dns... │ 0.02% │ 0.16%  │ 1           │ 17         │ 65535          │
│ 9575  │ sshd              │ wenba    │ /usr/sbin/sshd          │ 0.00% │ 0.02%  │ 3           │ 14         │ 65535          │
│ 1106  │ caddy             │ root     │ /data/soft/caddy/cad... │ 0.06% │ 0.25%  │ 5           │ 14         │ 65535          │
│ 18893 │ rinetd            │ root     │ /usr/sbin/rinetd        │ 0.00% │ 0.01%  │ 8           │ 13         │ 65535          │
│ 6376  │ sshd              │ wenba    │ /usr/sbin/sshd          │ 0.00% │ 0.02%  │ 4           │ 12         │ 65535          │
│ 8876  │ sshd              │ wenba    │ /usr/sbin/sshd          │ 0.00% │ 0.02%  │ 3           │ 11         │ 65535          │
│ 23670 │ sshd              │ wenba    │ /usr/sbin/sshd          │ 0.00% │ 0.02%  │ 3           │ 11         │ 65535          │
│ 11884 │ pritunl-http-api  │ wenba    │ /data/app/pritunl-ht... │ 0.01% │ 0.14%  │ 1           │ 10         │ 65535          │
│ 527   │ udevd             │ root     │ /sbin/udevd             │ 0.00% │ 0.01%  │ 2           │ 10         │ 1024           │
│ 20394 │ pickup            │ postfix  │ /usr/libexec/postfix... │ 0.00% │ 0.04%  │ 2           │ 9          │ 4096           │
│ 6307  │ sshd              │ wenba    │ /usr/sbin/sshd          │ 0.00% │ 0.03%  │ 2           │ 9          │ 65535          │
└───────┴───────────────────┴──────────┴─────────────────────────┴───────┴────────┴─────────────┴────────────┴────────────────┘
```

### 系统支持

CentOS

### 用法

```shell script
[xxxx@xxx_ansible_001 xxxx]$ ./pstatus --help
Usage of ./pstatus:
  -p string
    	process pid, like: 1234 or 123,456
  -s string
    	sort field, value: openfile|cpu|mem|conn (default "openfile")
  -t int
    	display top x

-p 指定进程的PID,多个用逗号分开，如果不指定默认显示全部的PID
-s 指定排序的字段, openfile|cpu|mem|conn
-t 可以只显示top x 的进程
```

持续监控可以使用

```shell script
watch -n 1 "sudo ./pstatus -t 20 -s cpu"
```