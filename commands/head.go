package commands

var head = `
head 与 tail 就像它的名字一样的浅显易懂，它是用来显示开头或结尾某个数量的文字区块，head 用来显示档案的开头至标准输出中，而 tail 想当然尔就是看档案的结尾。 

1．命令格式：

head [参数]... [文件]...  

2．命令功能：

head 用来显示档案的开头至标准输出中，默认head命令打印其相应文件的开头10行。 

3．命令参数：

-q 隐藏文件名

-v 显示文件名

-c<字节> 显示字节数

-n<行数> 显示的行数

4．使用实例：

实例1：显示文件的前n行

命令：

head -n 5 log2014.log

输出：

[root@localhost test]# cat log2014.log 

2014-01

2014-02

2014-03

2014-04

2014-05

2014-06

2014-07

2014-08

2014-09

2014-10

2014-11

2014-12

==============================

[root@localhost test]# head -n 5 log2014.log 

2014-01

2014-02

2014-03

2014-04

2014-05[root@localhost test]#

实例2：显示文件前n个字节

命令：

head -c 20 log2014.log

输出：

[root@localhost test]# head -c 20 log2014.log

2014-01

2014-02

2014

[root@localhost test]#

实例3：文件的除了最后n个字节以外的内容 

命令：

head -c -32 log2014.log

输出：

[root@localhost test]# head -c -32 log2014.log

2014-01

2014-02

2014-03

2014-04

2014-05

2014-06

2014-07

2014-08

2014-09

2014-10

2014-11

2014-12[root@localhost test]#

实例4：输出文件除了最后n行的全部内容

命令：

head -n -6 log2014.log

输出：

[root@localhost test]# head -n -6 log2014.log

2014-01

2014-02

2014-03

2014-04

2014-05

2014-06

2014-07[root@localhost test]#
`
