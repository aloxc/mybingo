# 配置文件目录

CREATE USER mybingo IDENTIFIED BY 'mybingo';  
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'mybingo'@'%';
-- GRANT ALL PRIVILEGES ON *.* TO 'mybingo'@'%' ;
FLUSH PRIVILEGES;
--------------------- 
作者：flyawayjh 
来源：CSDN 
原文：https://blog.csdn.net/flyawayjh/article/details/80990221 
版权声明：本文为博主原创文章，转载请附上博文链接


 sudo mysql -e "use mysql; update user set authentication_string=PASSWORD('mybingo') where User='mybingo'; update user set plugin='mysql_native_password';FLUSH PRIVILEGES;"
  - 


  CREATE DATABASE mybingo;
  GRANT ALL PRIVILEGES ON mybingo.* TO mybingo@localhost IDENTIFIED BY "Atestmybingo17)";
