package commands

var df = `
linux中df命令的功能是用来检查linux服务器的文件系统的磁盘空间占用情况。可以利用该命令来获取硬盘被占用了多少空间，目前还剩下多少空间等信息。

1．命令格式：

df [选项] [文件]

2．命令功能：

显示指定磁盘文件的可用空间。如果没有文件名被指定，则所有当前被挂载的文件系统的可用空间将被显示。默认情况下，磁盘空间将以 1KB 为单位进行显示，除非环境变量 POSIXLY_CORRECT 被指定，那样将以512字节为单位进行显示

3．命令参数：

必要参数：

-a 全部文件系统列表

-h 方便阅读方式显示

-H 等于“-h”，但是计算式，1K=1000，而不是1K=1024

-i 显示inode信息

-k 区块为1024字节

-l 只显示本地文件系统

-m 区块为1048576字节

--no-sync 忽略 sync 命令

-P 输出格式为POSIX

--sync 在取得磁盘信息前，先执行sync命令

-T 文件系统类型

选择参数：

--block-size=<区块大小> 指定区块大小

-t<文件系统类型> 只显示选定文件系统的磁盘信息

-x<文件系统类型> 不显示选定文件系统的磁盘信息

--help 显示帮助信息

--version 显示版本信息

4．使用实例：

实例1：显示磁盘使用情况

命令：

df

输出：

[root@CT1190 log]# df

文件系统               1K-块        已用     可用 已用% 挂载点

/dev/sda7             19840892    890896  17925856   5% /

/dev/sda9            203727156 112797500  80413912  59% /opt

/dev/sda8              4956284    570080   4130372  13% /var

/dev/sda6             19840892   1977568  16839184  11% /usr

/dev/sda3               988116     23880    913232   3% /boot

tmpfs                 16473212         0  16473212   0% /dev/shm

说明：

linux中df命令的输出清单的第1列是代表文件系统对应的设备文件的路径名（一般是硬盘上的分区）；第2列给出分区包含的数据块（1024字节）的数目；第3，4列分别表示已用的和可用的数据块数目。用户也许会感到奇怪的是，第3，4列块数之和不等于第2列中的块数。这是因为缺省的每个分区都留了少量空间供系统管理员使用。即使遇到普通用户空间已满的情况，管理员仍能登录和留有解决问题所需的工作空间。清单中Use% 列表示普通用户空间使用的百分比，即使这一数字达到100％，分区仍然留有系统管理员使用的空间。最后，Mounted on列表示文件系统的挂载点。

实例2：以inode模式来显示磁盘使用情况

命令：

df -i

输出：

[root@CT1190 log]# df -i

文件系统               Inode (I)已用 (I)可用 (I)已用% 挂载点

/dev/sda7            5124480    5560 5118920    1% /

/dev/sda9            52592640   50519 52542121    1% /opt

/dev/sda8            1280000    8799 1271201    1% /var

/dev/sda6            5124480   80163 5044317    2% /usr

/dev/sda3             255232      34  255198    1% /boot

tmpfs                4118303       1 4118302    1% /dev/shm

说明：

实例3：显示指定类型磁盘

命令：

df -t ext3

输出：

[root@CT1190 log]# df -t ext3

文件系统               1K-块        已用     可用 已用% 挂载点

/dev/sda7             19840892    890896  17925856   5% /

/dev/sda9            203727156  93089700 100121712  49% /opt

/dev/sda8              4956284    570104   4130348  13% /var

/dev/sda6             19840892   1977568  16839184  11% /usr

/dev/sda3               988116     23880    913232   3% /boot

说明：

实例4：列出各文件系统的i节点使用情况

命令：

df -ia

输出：

[root@CT1190 log]# df -ia

文件系统               Inode (I)已用 (I)可用 (I)已用% 挂载点

/dev/sda7            5124480    5560 5118920    1% 

/proc                       0       0       0    -  /proc

sysfs                      0       0       0    -  /sys

devpts                     0       0       0    -  /dev/pts

/dev/sda9            52592640   50519 52542121    1% /opt

/dev/sda8            1280000    8799 1271201    1% /var

/dev/sda6            5124480   80163 5044317    2% /usr

/dev/sda3             255232      34  255198    1% /boot

tmpfs                4118303       1 4118302    1% /dev/shm

none                       0       0       0    -  /proc/sys/fs/binfmt_misc

说明：

实例5：列出文件系统的类型

命令：

df -T

输出：

root@CT1190 log]# df -T

文件系统      类型     1K-块        已用     可用 已用% 挂载点

/dev/sda7     ext3    19840892    890896  17925856   5% /

/dev/sda9     ext3   203727156  93175692 100035720  49% /opt

/dev/sda8     ext3     4956284    570104   4130348  13% /var

/dev/sda6     ext3    19840892   1977568  16839184  11% /usr

/dev/sda3     ext3      988116     23880    913232   3% /boot

tmpfs        tmpfs    16473212         0  16473212   0% /dev/shm

说明：

实例6：以更易读的方式显示目前磁盘空间和使用情况 

命令：

输出：

[root@CT1190 log]# df -h

文件系统              容量  已用 可用 已用% 挂载点

/dev/sda7              19G  871M   18G   5% /

/dev/sda9             195G   89G   96G  49% /opt

/dev/sda8             4.8G  557M  4.0G  13% /var

/dev/sda6              19G  1.9G   17G  11% /usr

/dev/sda3             965M   24M  892M   3% /boot

tmpfs                  16G     0   16G   0% /dev/shm

[root@CT1190 log]# df -H

文件系统               容量   已用  可用 已用% 挂载点

/dev/sda7               21G   913M    19G   5% /

/dev/sda9              209G    96G   103G  49% /opt

/dev/sda8              5.1G   584M   4.3G  13% /var

/dev/sda6               21G   2.1G    18G  11% /usr

/dev/sda3              1.1G    25M   936M   3% /boot

tmpfs                   17G      0    17G   0% /dev/shm

[root@CT1190 log]# df -lh

文件系统              容量  已用 可用 已用% 挂载点

/dev/sda7              19G  871M   18G   5% /

/dev/sda9             195G   89G   96G  49% /opt

/dev/sda8             4.8G  557M  4.0G  13% /var

/dev/sda6              19G  1.9G   17G  11% /usr

/dev/sda3             965M   24M  892M   3% /boot

tmpfs                  16G     0   16G   0% /dev/shm

[root@CT1190 log]# df -k

文件系统               1K-块        已用     可用 已用% 挂载点

/dev/sda7             19840892    890896  17925856   5% /

/dev/sda9            203727156  93292572  99918840  49% /opt

/dev/sda8              4956284    570188   4130264  13% /var

/dev/sda6             19840892   1977568  16839184  11% /usr

/dev/sda3               988116     23880    913232   3% /boot

tmpfs                 16473212         0  16473212   0% /dev/shm

说明：

-h更具目前磁盘空间和使用情况 以更易读的方式显示

-H根上面的-h参数相同,不过在根式化的时候,采用1000而不是1024进行容量转换

-k以单位显示磁盘的使用情况

-l显示本地的分区的磁盘空间使用率,如果服务器nfs了远程服务器的磁盘,那么在df上加上-l后系统显示的是过滤nsf驱动器后的结果

-i显示inode的使用情况。linux采用了类似指针的方式管理磁盘空间影射.这也是一个比较关键应用

`