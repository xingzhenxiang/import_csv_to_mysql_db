# import_csv_to_mysql_db
import_csv_to_mysql_db


 go run main.go --table=t1  --host=192.168.0.170 --user=xzx --password= --dbname=mysqlslap --csvPath=/root/t12.csv
 
 mysql> show create table t1;
+-------+-------------------------------------------------------------------------------------------------------------------------------------+
| Table | Create Table                                                                                                                        |
+-------+-------------------------------------------------------------------------------------------------------------------------------------+
| t1    | CREATE TABLE `t1` (
  `intcol1` int(32) DEFAULT NULL,
  `charcol1` varchar(128) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 |
+-------+-------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)


 more /root/t12.csv 
"1804289383","mxvtvmC9127qJNm06sGB8R92q2j7vTiiITRDGXM9ZLzkdekbWtmXKwZ2qG1llkRw5m9DHOFilEREk3q7oce8O3BEJC0woJsm6uzFAEynLH2xCsw1KQ1lT4zg9rdxBL","1"
"822890675","97RGHZ65mNzkSrYT3zWoSbg9cNePQr1bzSk81qDgE4Oanw3rnPfGsBHSbnu1evTdFDe83ro9w4jjteQg4yoo9xHck3WNqzs54W5zEm92ikdRF48B2oz3m8gMBAl11W","2"
"1308044878","50w46i58Giekxik0cYzfA8BZBLADEg3JhzGfZDoqvQQk0Akcic7lcJInYSsf9wqin6LDC1vzJLkJXKn5onqOy04MTw1WksCYqPl2Jg2eteqOqTLfGCvE4zTZwWvgMz","3"
"964445884","DPh7kD1E6f4MMQk1ioopsoIIcoD83DD8Wu7689K6oHTAjD3Hts6lYGv8x9G0EL0k87q8G2ExJjz2o3KhnIJBbEJYFROTpO5pNvxgyBT9nSCbNO9AiKL9QYhi0x3hL9","4"
"1586903190","lwRHuWm4HE8leYmg66uGYIp6AnAr0BDd7YmuvYqCfqp9EbhKZRSymA4wx6gpHlJHI53DetH9j7Ixar90Jey5outd1ZIAJdJTjMaD7rMiqYXHFhHaB7Xr1HKuqe51GG","5"

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
 
 
 cat /root/12.txt 
20190413,1116732901359226880,N/A,N/A,32,N/A,829789a0-41a2-440c-b778-c84b49ef4ff3,NULL,d.aiyamayaaiyamaya.com.cn,1830151665811,110.110.110.110,zh-CN,Eaiyamayaaiyamaya,c22bd2b6-0a70-4796-a83f-f45a0a8d7cc8,N/A,1280,720,"a,button,input,span,img,div,i,ul,li,p,label,select,em",哈哈哈,https://d.aiyamayaaiyamaya.com.cn/aiyamayaaiyamaya_activity3.0/6161?utm_source=tencent&utm_medium=Network&utm_term=84001201904AD11004&utm_campaign=Octavia,Dalvik/2.1.0 (Linux; U; Android 5.1; Coolpad 7230S Build/JDQ39),N/A,3.0,N/A
20190413,1116732972968579072,N/A,N/A,32,N/A,987f65e7-766e-45fa-be23-e7a75cd75761,NULL,d.aiyamayaaiyamaya.com.cn,1830151665811,110.110.110.110,zh-CN,Eaiyamayaaiyamaya,b83d5f33-b8d8-458d-ba1e-bbf2063e0497,N/A,1812,1080,"a,button,input,span,img,div,i,ul,li,p,label,select,em",哈哈哈,https://d.aiyamayaaiyamaya.com.cn/aiyamayaaiyamaya_activity3.0/6163?utm_source=tencent&utm_medium=Network&utm_term=84001201904AD11005&utm_campaign=Octavia,"Mozilla/5.0 (Linux; U; Android 5.0; T980 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.36 UCBrowser/10.10.8.822",N/A,3.0,N/A
20190413,1116733022683664384,N/A,N/A,24,N/A,8abdc000-b756-4919-91cd-25a65cb9fb00,NULL,aiyamayaaiyamaya.com.cn,1830151665811,110.110.110.110,zh-CN,Eaiyamayaaiyamaya,c744e65c-c23f-41de-bd7f-c8703ba3875f,N/A,640,360,"a,button,input,span,img,div,i,ul,li,p,label,select,em",笑笑笑,https://bjax.aiyamayaaiyamaya.com.cn/m,"Mozilla/5.0 (Linux; Android 9; MHA-AL00 Build/HUAWEIMHA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044606 Mobile Safari/537.36 MMWEBID/7068 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/4G Language/zh_CN",N/A,3.0,N/A
20190413,1116733069701812224,N/A,N/A,32,N/A,6c8e037c-6096-416f-b8a5-39736800a894,NULL,d.aiyamayaaiyamaya.com.cn,1830151665811,110.110.110.110,zh-CN,Eaiyamayaaiyamaya,d538b28e-858f-42d6-b15c-c1fdfad46c05,N/A,1920,1080,"a,button,input,span,img,div,i,ul,li,p,label,select,em",哈哈哈,https://d.aiyamayaaiyamaya.com.cn/aiyamayaaiyamaya_activity3.0/6163?utm_source=tencent&utm_medium=Network&utm_term=84001201904AD11005&utm_campaign=Octavia,Dalvik/2.1.0 (Linux; U; Android 5.1.1; HM 1SLTETD Build/KOT49H),N/A,3.0,N/A
 
















csv2MySQL

./csv2MySQL --table=t1  --host=192.168.0.170 --user=xzx --password= --dbname=mysqlslap --csvPath=/root/t12.csv
