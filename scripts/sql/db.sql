/*
SQLyog v10.2 
MySQL - 5.6.35-log : Database - deviceshouse
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`deviceshouse` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `deviceshouse`;

/*Table structure for table `dev_boiler_a_info` */

DROP TABLE IF EXISTS `dev_boiler_a_info`;

CREATE TABLE `dev_boiler_a_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `start` varchar(64) DEFAULT NULL,
  `turn_fire` varchar(64) DEFAULT NULL,
  `stop` varchar(64) DEFAULT NULL,
  `gas_open` varchar(64) DEFAULT NULL,
  `gas_feedback` varchar(64) DEFAULT NULL,
  `smoke_loop` varchar(64) DEFAULT NULL,
  `steam_pressure` varchar(64) DEFAULT NULL,
  `fan_freq` varchar(64) DEFAULT NULL,
  `freq_feedback` varchar(64) DEFAULT NULL,
  `throttle_open` varchar(64) DEFAULT NULL,
  `throttle_feedback` varchar(64) DEFAULT NULL,
  `big_fire` varchar(64) DEFAULT NULL,
  `small_fire` varchar(64) DEFAULT NULL,
  `water_pump` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `dev_boiler_a_info` */

insert  into `dev_boiler_a_info`(`id`,`device_sn`,`start`,`turn_fire`,`stop`,`gas_open`,`gas_feedback`,`smoke_loop`,`steam_pressure`,`fan_freq`,`freq_feedback`,`throttle_open`,`throttle_feedback`,`big_fire`,`small_fire`,`water_pump`) values (1,'boiler1','3.5','4.0','5.0','23.0','21.8','0.0','0.0','35','0.8','12.0','11.7','30.0','35.0','1');

/*Table structure for table `dev_boiler_b_info` */

DROP TABLE IF EXISTS `dev_boiler_b_info`;

CREATE TABLE `dev_boiler_b_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `start_temp` varchar(64) DEFAULT NULL,
  `target_temp` varchar(64) DEFAULT NULL,
  `stop_temp` varchar(64) DEFAULT NULL,
  `out_water_temp` varchar(64) DEFAULT NULL,
  `back_water_temp` varchar(64) DEFAULT NULL,
  `smoke_temp` varchar(64) DEFAULT NULL,
  `boiler_load` varchar(64) DEFAULT NULL,
  `gas` varchar(64) DEFAULT NULL,
  `throttle` varchar(64) DEFAULT NULL,
  `smoke` varchar(64) DEFAULT NULL,
  `freq` varchar(64) DEFAULT NULL,
  `run_state` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Data for the table `dev_boiler_b_info` */

insert  into `dev_boiler_b_info`(`id`,`device_sn`,`start_temp`,`target_temp`,`stop_temp`,`out_water_temp`,`back_water_temp`,`smoke_temp`,`boiler_load`,`gas`,`throttle`,`smoke`,`freq`,`run_state`) values (1,'boiler2','50.0','55.0','60.0','40.6','50.0','40.0','0.0','18.0','0.0','0.0','40.0','1');

/*Table structure for table `dev_device` */

DROP TABLE IF EXISTS `dev_device`;

CREATE TABLE `dev_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  `type_name` varchar(64) DEFAULT NULL,
  `dev_name` varchar(64) DEFAULT NULL,
  `is_online` int(11) DEFAULT NULL,
  `activetime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `dev_device` */

insert  into `dev_device`(`id`,`device_sn`,`type_code`,`type_name`,`dev_name`,`is_online`,`activetime`) values (4,'boiler1',31,'boiler_a','boiler1',0,'2017-07-20 11:38:47'),(5,'boiler2',32,'bolier_b','boiler2',0,'2017-07-20 11:38:47');

/*Table structure for table `dev_deviceattr` */

DROP TABLE IF EXISTS `dev_deviceattr`;

CREATE TABLE `dev_deviceattr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `attr_name` varchar(64) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

/*Data for the table `dev_deviceattr` */

insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`) values (2,'开关',1001),(3,'开度',1002),(4,'浓度',1003),(5,'总量',1004),(6,'亮度',1005),(7,'颜色',1006),(8,'色温',1007),(9,'温度',1008),(10,'湿度',1009),(11,'窗帘开关停',1010),(12,'报警',1011);

/*Table structure for table `dev_deviceattrinfo` */

DROP TABLE IF EXISTS `dev_deviceattrinfo`;

CREATE TABLE `dev_deviceattrinfo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  `attr_name` varchar(64) DEFAULT NULL,
  `attr_value_cur` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `dev_deviceattrinfo` */

insert  into `dev_deviceattrinfo`(`id`,`device_sn`,`attr_code`,`attr_name`,`attr_value_cur`) values (1,'robot1',1001,'开关','0'),(2,'robot2',1001,'开关','0'),(3,'robot3',1001,'开关','0');

/*Table structure for table `dev_devicetype` */

DROP TABLE IF EXISTS `dev_devicetype`;

CREATE TABLE `dev_devicetype` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicetype` */

insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (2,'窗帘',2),(3,'PM2.5',3),(4,'甲醛报警',4),(5,'燃气报警',5),(6,'烟雾报警',6),(7,'电表',7),(8,'水表',8),(9,'温湿度',9),(10,'亮度色温灯',15),(11,'开关',16),(12,'插座',17),(13,'门磁',18),(14,'人体红外',19),(15,'燃气表',20),(16,'热表',21),(22,'摄像头',10),(23,'机器人',22),(32,'boiler_a',31),(33,'boiler_b',32);

/*Table structure for table `dev_user_bind_device` */

DROP TABLE IF EXISTS `dev_user_bind_device`;

CREATE TABLE `dev_user_bind_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `device_sn` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `dev_user_bind_device` */

insert  into `dev_user_bind_device`(`id`,`username`,`device_sn`) values (1,'test','boiler1'),(2,'test','boiler2');

/*Table structure for table `usr_user` */

DROP TABLE IF EXISTS `usr_user`;

CREATE TABLE `usr_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `phone` varchar(32) NOT NULL,
  `mail` varchar(64) DEFAULT NULL,
  `registertime` datetime DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8;

/*Data for the table `usr_user` */

insert  into `usr_user`(`id`,`username`,`password`,`phone`,`mail`,`registertime`,`state`) values (123,'test','123456','18610330033','luz@1234.com','2016-12-01 18:43:12',10);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
