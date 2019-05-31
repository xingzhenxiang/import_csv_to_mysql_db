# import_csv_to_mysql_db
import_csv_to_mysql_db


 go run main.go --table=t1  --host=192.168.0.170 --user=xzx --password= --dbname=mysqlslap --csvPath=/root/t12.csv
 
 mysql> show create table t1;
CREATE TABLE `t1` (
  `intcol1` int(32) DEFAULT NULL,
  `charcol1` varchar(128) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 


mysql> show create table tmp_log;
CREATE TABLE `tmp_log` (
  `day` date DEFAULT NULL,
  `logid` bigint(20) DEFAULT NULL,
  `accout` varchar(225) DEFAULT NULL,
  `action` varchar(225) DEFAULT NULL,
  `cd` varchar(50) DEFAULT NULL,
  `collectmark` varchar(50) DEFAULT NULL,
  `cookie` varchar(50) DEFAULT NULL,
  `disp` varchar(50) DEFAULT NULL,
  `domain` varchar(500) DEFAULT NULL,
  `executiontime` bigint(20) DEFAULT NULL,
  `ip` varchar(50) DEFAULT NULL,
  `lang` varchar(50) DEFAULT NULL,
  `product` varchar(50) DEFAULT NULL,
  `query` varchar(50) DEFAULT NULL,
  `referrer` text,
  `sh` int(11) DEFAULT NULL,
  `sw` int(11) DEFAULT NULL,
  `target` varchar(500) DEFAULT NULL,
  `title` varchar(2000) DEFAULT NULL,
  `url` varchar(5000) DEFAULT NULL,
  `useragent` varchar(2000) DEFAULT NULL,
  `userinfo` varchar(2000) DEFAULT NULL,
  `version` varchar(20) DEFAULT NULL,
  `xpath` varchar(1000) DEFAULT NULL,
  KEY `inx_day` (`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

 go run main.go --table=tmp_log  --host=192.168.0.170 --user=root --password=123456 --dbname=mysqlslap --csvPath=/root/12.txt
 
 
 














csv2MySQL

./csv2MySQL --table=t1  --host=192.168.0.170 --user=xzx --password= --dbname=mysqlslap --csvPath=/root/t12.csv
