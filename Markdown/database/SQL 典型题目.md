# SQL 典型题目

参考：[sql语句练习50题(Mysql版)][1]



### 建表

登录 mysql 并创建一个数据库：`create database course_info`；

保存下面 sql 语句到一个文本文件中，暂时命名为 `course_info.sql`

```mysql
CREATE TABLE `Student`(
	`s_id` VARCHAR(20),
	`s_name` VARCHAR(20) NOT NULL DEFAULT '',
	`s_birth` VARCHAR(20) NOT NULL DEFAULT '',
	`s_sex` VARCHAR(10) NOT NULL DEFAULT '',
	PRIMARY KEY(`s_id`)
);
CREATE TABLE `Course`(
	`c_id`  VARCHAR(20),
	`c_name` VARCHAR(20) NOT NULL DEFAULT '',
	`t_id` VARCHAR(20) NOT NULL,
	PRIMARY KEY(`c_id`)
);
CREATE TABLE `Teacher`(
	`t_id` VARCHAR(20),
	`t_name` VARCHAR(20) NOT NULL DEFAULT '',
	PRIMARY KEY(`t_id`)
);
CREATE TABLE `Score`(
	`s_id` VARCHAR(20),
	`c_id`  VARCHAR(20),
	`s_score` INT(3),
	PRIMARY KEY(`s_id`,`c_id`)
);
insert into Student values('01' , '赵雷' , '1990-01-01' , '男');
insert into Student values('02' , '钱电' , '1990-12-21' , '男');
insert into Student values('03' , '孙风' , '1990-05-20' , '男');
insert into Student values('04' , '李云' , '1990-08-06' , '男');
insert into Student values('05' , '周梅' , '1991-12-01' , '女');
insert into Student values('06' , '吴兰' , '1992-03-01' , '女');
insert into Student values('07' , '郑竹' , '1989-07-01' , '女');
insert into Student values('08' , '王菊' , '1990-01-20' , '女');
insert into Course values('01' , '语文' , '02');
insert into Course values('02' , '数学' , '01');
insert into Course values('03' , '英语' , '03');

insert into Teacher values('01' , '张三');
insert into Teacher values('02' , '李四');
insert into Teacher values('03' , '王五');

insert into Score values('01' , '01' , 80);
insert into Score values('01' , '02' , 90);
insert into Score values('01' , '03' , 99);
insert into Score values('02' , '01' , 70);
insert into Score values('02' , '02' , 60);
insert into Score values('02' , '03' , 80);
insert into Score values('03' , '01' , 80);
insert into Score values('03' , '02' , 80);
insert into Score values('03' , '03' , 80);
insert into Score values('04' , '01' , 50);
insert into Score values('04' , '02' , 30);
insert into Score values('04' , '03' , 20);
insert into Score values('05' , '01' , 76);
insert into Score values('05' , '02' , 87);
insert into Score values('06' , '01' , 31);
insert into Score values('06' , '03' , 34);
insert into Score values('07' , '02' , 89);
insert into Score values('07' , '03' , 98);
```

使用下面命令执行 sql 文件，向数据库中插入测试数据：

`mysql -h host -P 10008 -uroot -p -Dcourse_info</path/to/course_info.sql`

### 题目

* *查询"01"课程比"02"课程成绩高的学生的信息及课程分数*

```sql
select a.* ,b.s_score as 01_score,c.s_score as 02_score from 
Student a join Score b on a.s_id=b.s_id and b.c_id='01'
left join Score c on a.s_id=c.s_id and c.c_id='02' or c.c_id = NULL where b.s_score>c.s_score
```









































[1]:<https://blog.csdn.net/fashion2014/article/details/78826299>