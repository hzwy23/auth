-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: sso
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.14-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_domain_define`
--

DROP TABLE IF EXISTS `sys_domain_define`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_define` (
  `domain_id` varchar(30) NOT NULL,
  `domain_name` varchar(300) NOT NULL,
  `domain_status_id` char(1) NOT NULL,
  `create_time` datetime NOT NULL,
  `modify_time` datetime DEFAULT NULL,
  `modify_user` varchar(30) DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`domain_id`),
  KEY `fk_sys_idx_05` (`domain_status_id`),
  CONSTRAINT `fk_sys_idx_05` FOREIGN KEY (`domain_status_id`) REFERENCES `sys_domain_status_attr` (`domain_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='域管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_define`
--

LOCK TABLES `sys_domain_define` WRITE;
/*!40000 ALTER TABLE `sys_domain_define` DISABLE KEYS */;
INSERT INTO `sys_domain_define` VALUES ('abc','中国农业银行','0','2017-08-19 23:45:30','2017-08-19 23:45:30','ccbc_admin','ccbc_admin'),('ccb','测试银行','0','2017-08-20 00:11:47','2017-08-20 00:11:47','ccbc_admin','ccbc_admin'),('ccbc','中国工商银行','0','2017-08-19 14:11:43','2017-08-19 14:11:43','admin','admin'),('icbc','爱存不存','0','2017-08-19 23:45:46','2017-08-19 23:45:46','ccbc_admin','ccbc_admin'),('vertex_root','超级管理域','0','2016-12-26 16:43:19','2017-08-19 14:13:25','admin','admin');
/*!40000 ALTER TABLE `sys_domain_define` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_domain_status_attr`
--

DROP TABLE IF EXISTS `sys_domain_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_status_attr` (
  `domain_status_id` char(1) NOT NULL,
  `domain_status_name` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`domain_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_status_attr`
--

LOCK TABLES `sys_domain_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_domain_status_attr` DISABLE KEYS */;
INSERT INTO `sys_domain_status_attr` VALUES ('0','正常'),('1','锁定');
/*!40000 ALTER TABLE `sys_domain_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_handle_logs`
--

DROP TABLE IF EXISTS `sys_handle_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_handle_logs` (
  `uuid` varchar(60) NOT NULL,
  `user_id` varchar(30) DEFAULT NULL,
  `handle_time` datetime DEFAULT NULL,
  `client_ip` varchar(30) DEFAULT NULL,
  `status_code` varchar(10) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  `data` varchar(3000) DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_handle_logs`
--

LOCK TABLES `sys_handle_logs` WRITE;
/*!40000 ALTER TABLE `sys_handle_logs` DISABLE KEYS */;
INSERT INTO `sys_handle_logs` VALUES ('0c82398e-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/','{}'),('0c8286c1-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','POST','/login','{\"duration\":\"43200\",\"password\":\"hzwy23\",\"username\":\"admin\"}'),('0c82c51a-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/HomePage','{}'),('0c82fac5-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/main/menu','{\"Id\":\"-1\",\"TypeId\":\"0\"}'),('0c832f96-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/menu/all/except/button','{\"resId\":\"0100000000\"}'),('0c8345e7-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/batch/page','{}'),('0c835bdf-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/role/page','{}'),('0c836d1f-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/role/get','{\"domain_id\":\"\",\"limit\":\"10\",\"offset\":\"0\",\"order\":\"asc\"}'),('0c837ceb-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/user/page','{}'),('0c8391a7-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/user/get','{\"domain_id\":\"\",\"limit\":\"10\",\"offset\":\"0\",\"order\":\"asc\"}'),('0c83aa9f-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/resource/org/page','{}'),('0c83eafd-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/domain/page','{}'),('0c84021f-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/domain/get','{\"limit\":\"40\",\"offset\":\"0\",\"order\":\"asc\"}'),('0c8416f4-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/resource/org/page','{}'),('0c842e23-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/user/page','{}'),('0c84435f-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/user/get','{\"domain_id\":\"\",\"limit\":\"10\",\"offset\":\"0\",\"order\":\"asc\"}'),('0c845bac-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/resource/service','{}'),('0c84733d-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/menu/all/except/button','{\"resId\":\"-1\"}'),('0c8487ce-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/resource/page','{}'),('0c849863-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/menu/all/except/button','{\"resId\":\"-1\"}'),('0c84a72c-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/v1/auth/resource/func','{\"resId\":\"0101010000\",\"themeId\":\"1001\"}'),('0c84b68e-87b5-11e7-9d7c-a0c58951c8d5','admin','2017-08-23 11:42:04','127.0.0.1','200','GET','/logout','{}');
/*!40000 ALTER TABLE `sys_handle_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_index_page`
--

DROP TABLE IF EXISTS `sys_index_page`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_index_page` (
  `theme_id` varchar(30) NOT NULL,
  `res_url` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`theme_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_index_page`
--

LOCK TABLES `sys_index_page` WRITE;
/*!40000 ALTER TABLE `sys_index_page` DISABLE KEYS */;
INSERT INTO `sys_index_page` VALUES ('1001','./views/hauth/theme/default/index.tpl'),('1002','./views/hauth/theme/blue/index.tpl'),('1003','./views/hauth/theme/apple/index.tpl'),('1004','./views/hauth/theme/cyan/index.tpl'),('1005','./views/hauth/theme/tradition/index.tpl');
/*!40000 ALTER TABLE `sys_index_page` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_org_info`
--

DROP TABLE IF EXISTS `sys_org_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_org_info` (
  `org_unit_id` varchar(66) NOT NULL,
  `org_unit_desc` varchar(300) NOT NULL,
  `up_org_id` varchar(66) NOT NULL,
  `create_date` date NOT NULL,
  `maintance_date` date NOT NULL,
  `create_user` varchar(30) NOT NULL,
  `maintance_user` varchar(30) NOT NULL,
  PRIMARY KEY (`org_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_org_info`
--

LOCK TABLES `sys_org_info` WRITE;
/*!40000 ALTER TABLE `sys_org_info` DISABLE KEYS */;
INSERT INTO `sys_org_info` VALUES ('200','中国农业银行总行','vertex_root_join_vertex_root','2017-08-19','2017-08-19','ccbc_admin','ccbc_admin'),('20010','中国农业银行湖北省分行','200','2017-08-19','2017-08-19','ccbc_admin','ccbc_admin'),('20020','中国农业银行四川省分行','200','2017-08-19','2017-08-19','ccbc_admin','ccbc_admin'),('20030','中国农业银行上海省分行','200','2017-08-19','2017-08-19','ccbc_admin','ccbc_admin'),('20040','中国农业银行湖南省分行','200','2017-08-19','2017-08-19','ccbc_admin','ccbc_admin'),('20050','中国农业银行重庆市分行','200','2017-08-19','2017-08-22','ccbc_admin','admin'),('ccbc_join_100','中国工商银行总行','vertex_root_join_vertex_root','2017-08-19','2017-08-19','admin','ccbc_admin'),('ccbc_join_10010','中国工商银行湖北省分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10020','中国工商银行四川省分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10030','中国工商银行重庆市分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10040','中国工商银行上海分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10050','中国工商银行北京分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10060','中国工商银行江苏省分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('ccbc_join_10070','中国工商银行浙江省分行','ccbc_join_100','2017-08-19','2017-08-19','admin','admin'),('vertex_root_join_vertex_root','某某联合社','root_vertex_system','2016-01-01','2017-08-19','sys','ccbc_admin');
/*!40000 ALTER TABLE `sys_org_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_privilege`
--

DROP TABLE IF EXISTS `sys_privilege`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_privilege` (
  `privilege_id` varchar(30) NOT NULL,
  `privilege_desc` varchar(200) DEFAULT NULL,
  `privilege_type` varchar(10) DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`privilege_id`),
  KEY `pk_sys_privilege01` (`privilege_type`),
  CONSTRAINT `pk_sys_privilege01` FOREIGN KEY (`privilege_type`) REFERENCES `sys_privilege_type_attr` (`privilege_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_privilege`
--

LOCK TABLES `sys_privilege` WRITE;
/*!40000 ALTER TABLE `sys_privilege` DISABLE KEYS */;
INSERT INTO `sys_privilege` VALUES ('domain_auth_001','域访问权限','101','admin','2017-01-01 12:12:12');
/*!40000 ALTER TABLE `sys_privilege` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_privilege_domain`
--

DROP TABLE IF EXISTS `sys_privilege_domain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_privilege_domain` (
  `uuid` varchar(66) NOT NULL,
  `privilege_id` varchar(66) DEFAULT NULL,
  `domain_id` varchar(30) DEFAULT NULL,
  `permission` varchar(5) DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `pk_sys_privilege_domain01` (`privilege_id`),
  KEY `pk_sys_privilege_domain02` (`domain_id`),
  CONSTRAINT `pk_sys_privilege_domain01` FOREIGN KEY (`privilege_id`) REFERENCES `sys_privilege` (`privilege_id`),
  CONSTRAINT `pk_sys_privilege_domain02` FOREIGN KEY (`domain_id`) REFERENCES `sys_domain_define` (`domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_privilege_domain`
--

LOCK TABLES `sys_privilege_domain` WRITE;
/*!40000 ALTER TABLE `sys_privilege_domain` DISABLE KEYS */;
INSERT INTO `sys_privilege_domain` VALUES ('312321','domain_auth_001','icbc','2','admin','2017-01-01 12:12:12');
/*!40000 ALTER TABLE `sys_privilege_domain` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_privilege_role`
--

DROP TABLE IF EXISTS `sys_privilege_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_privilege_role` (
  `uuid` varchar(66) NOT NULL,
  `privilege_id` varchar(66) DEFAULT NULL,
  `role_id` varchar(66) DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `pk_sys_privilege_role01` (`privilege_id`),
  KEY `pk_sys_privilege_role02` (`role_id`),
  CONSTRAINT `pk_sys_privilege_role01` FOREIGN KEY (`privilege_id`) REFERENCES `sys_privilege` (`privilege_id`),
  CONSTRAINT `pk_sys_privilege_role02` FOREIGN KEY (`role_id`) REFERENCES `sys_role_define` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_privilege_role`
--

LOCK TABLES `sys_privilege_role` WRITE;
/*!40000 ALTER TABLE `sys_privilege_role` DISABLE KEYS */;
INSERT INTO `sys_privilege_role` VALUES ('3232','domain_auth_001','20010','admin','2017-01-01 12:12:12');
/*!40000 ALTER TABLE `sys_privilege_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_privilege_type_attr`
--

DROP TABLE IF EXISTS `sys_privilege_type_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_privilege_type_attr` (
  `privilege_type` varchar(10) NOT NULL,
  `privilege_type_desc` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`privilege_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_privilege_type_attr`
--

LOCK TABLES `sys_privilege_type_attr` WRITE;
/*!40000 ALTER TABLE `sys_privilege_type_attr` DISABLE KEYS */;
INSERT INTO `sys_privilege_type_attr` VALUES ('101','域权限类型');
/*!40000 ALTER TABLE `sys_privilege_type_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_info`
--

DROP TABLE IF EXISTS `sys_resource_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_info` (
  `res_id` varchar(30) NOT NULL COMMENT '菜单资源编码',
  `res_name` varchar(300) DEFAULT NULL COMMENT '资源名称',
  `res_attr` char(1) DEFAULT NULL COMMENT '资源属性 0：结点 1：叶子',
  `res_up_id` varchar(30) DEFAULT NULL COMMENT '上级资源编码，-1 表示顶层编码',
  `res_type` char(1) DEFAULT NULL COMMENT '资源类型：\n0：系统首页\n1：菜单页面\n2：功能服务\n3：功能服务结点\n4：菜单结点',
  `sys_flag` char(1) DEFAULT NULL COMMENT '0：系统内置菜单',
  `inner_flag` varchar(5) DEFAULT NULL COMMENT 'true：内部路由\nfalse：外部路由',
  `service_cd` varchar(30) DEFAULT NULL COMMENT '所属系统标识，为空表示为内部资源',
  PRIMARY KEY (`res_id`),
  KEY `fk_sys_idx_13` (`res_type`),
  KEY `fk_sys_idx_14` (`res_attr`),
  CONSTRAINT `fk_sys_idx_13` FOREIGN KEY (`res_type`) REFERENCES `sys_resource_type_attr` (`res_type`),
  CONSTRAINT `fk_sys_idx_14` FOREIGN KEY (`res_attr`) REFERENCES `sys_resource_info_attr` (`res_attr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_info`
--

LOCK TABLES `sys_resource_info` WRITE;
/*!40000 ALTER TABLE `sys_resource_info` DISABLE KEYS */;
INSERT INTO `sys_resource_info` VALUES ('0100000000','系统管理','0','-1','0','0','true',NULL),('0101000000','系统审计','0','0100000000','4','0','true',NULL),('0101010000','操作查询','1','0101000000','1','0','true',NULL),('0101010100','查看操作日志权限','1','0101010000','2','0','true',NULL),('0101010200','下载操作日志按钮','1','0101010000','2','0','true',NULL),('0101010300','搜索日志信息按钮','1','0101010000','2','0','true',NULL),('0102000000','反向代理','0','0100000000','4',NULL,'true',''),('0103000000','权限管理','0','0100000000','4',NULL,'true',''),('0103010000','按钮资源','1','0104000000','1','0','true',''),('0103010100','查询资源信息','1','0103010000','2','0','true',NULL),('0103010200','新增资源信息按钮','1','0103010000','2','0','true',NULL),('0103010300','编辑资源信息按钮','1','0103010000','2','0','true',NULL),('0103010400','删除资源信息按钮','1','0103010000','2','0','true',NULL),('01030104001','删除资源信息按钮','1','0101010000','2','0','true',NULL),('0103010500','配置主题信息按钮','1','0103010000','2','0','true',NULL),('0103020000','组织管理','1','0103000000','1','0','true',''),('0103020100','查询组织架构功能服务','1','0103020000','2','0','true',''),('0103020200','新增组织架构信息按钮','1','0103020000','2','0','true',NULL),('0103020300','更新组织架构信息按钮','1','0103020000','2','0','true',NULL),('0103020400','删除组织架构信息按钮','1','0103020000','2','0','true',NULL),('0103020500','导出组织架构信息按钮','1','0103020000','2','0','true',NULL),('0103030000','本地静态路由','1','0102000000','1',NULL,'true',''),('0103030100','查询共享域信息服务','1','0104010200','2','0','true',''),('0103030200','新增共享域信息按钮','1','0104010200','2','0','true',NULL),('0103030300','删除共享域信息按钮','1','0104010200','2','0','true',NULL),('0103030400','更新共享域信息按钮','1','0104010200','2','0','true',NULL),('0103040000','子系统管理','1','0102000000','1',NULL,'true',''),('0103050000','代理子系统API','1','0102000000','1',NULL,'true',''),('0103060000','代理静态路由','1','0102000000','1',NULL,'true',''),('0103070000','菜单管理','1','0104000000','1',NULL,'true',''),('0103070100','查询菜单资源','1','0103070000','2',NULL,'true',''),('0104000000','开发者选项','0','0100000000','4',NULL,'true',''),('0104010000','域定义','1','0103000000','1','0','true',''),('0104010100','查询域信息功能服务','1','0104010000','2','0','true',''),('0104010200','共享域管理按钮','1','0104010000','2','0','true',''),('0104010300','编辑域信息按钮','1','0104010000','2','0','true',NULL),('0104010400','删除域信息按钮','1','0104010000','2','0','true',NULL),('0104010500','新增域信息按钮','1','0104010000','2','0','true',NULL),('0105010000','用户管理','1','0103000000','1','0','true',''),('0105010100','查询用户信息','1','0105010000','2','0','true',NULL),('0105010200','新增用户信息按钮','1','0105010000','2','0','true',NULL),('0105010300','编辑用户信息按钮','1','0105010000','2','0','true',NULL),('0105010400','删除用户信息按钮','1','0105010000','2','0','true',NULL),('0105010500','修改用户密码按钮','1','0105010000','2','0','true',NULL),('0105010600','修改用户状态按钮','1','0105010000','2','0','true',NULL),('0105020000','角色管理','1','0103000000','1','0','true',''),('0105020100','查询角色信息','1','0105020000','2','0','true',NULL),('0105020200','新增角色信息按钮','1','0105020000','2','0','true',NULL),('0105020300','更新角色信息按钮','1','0105020000','2','0','true',NULL),('0105020400','删除角色信息按钮','1','0105020000','2','0','true',NULL),('0105020500','角色资源管理','1','0105020000','2','0','true',NULL),('0105020510','查询角色资源信息','1','0105020500','2','0','true',NULL),('0105020520','修改角色资源信息','1','0105020500','2','0','true',NULL),('0105040000','授权管理','1','0103000000','1','0','true',''),('0105040100','授予权限按钮','1','0105040000','2','0','true',NULL),('0105040200','移除权限','1','0105040000','2','0','true',NULL),('0200000000','成本分摊','0','-1','0',NULL,'false','ca'),('0201000000','维度信息管理','0','0200000000','4',NULL,'false','ca'),('0201010000','责任中心','1','0201000000','1',NULL,'false','ca'),('0201030000','成本类别','1','0201000000','1',NULL,'false','ca'),('0201040000','动因定义','1','0201000000','1',NULL,'false','ca'),('0201060000','成本池管理','1','0201000000','1',NULL,'false','ca'),('0202000000','分摊规则管理','0','0200000000','4',NULL,'false','ca'),('0202010000','静态分摊定义','1','0202000000','1',NULL,'false','ca'),('0202020000','分摊规则定义','1','0202000000','1',NULL,'false','ca'),('0202030000','静态分摊比例管理','1','0202010000','2',NULL,'false','ca'),('0202040000','分摊规则组定义','1','0202000000','1',NULL,'false','ca'),('0203000000','批次综合管理','0','0200000000','4',NULL,'false','ca'),('0203010000','批次定义与管理','1','0203000000','1',NULL,'false','ca'),('0203020000','批次运行历史查询','1','0203000000','1',NULL,'false','ca'),('0203040000','直接费用查询与管理','1','0203000000','1',NULL,'false','ca'),('0203050000','动因值查询与管理','1','0203000000','1',NULL,'false','ca'),('0400000000','公共信息','0','-1','0',NULL,'true',''),('0401000000','条线信息','1','0410000000','1',NULL,'true',''),('0402000000','产品信息','1','0410000000','1',NULL,'true',''),('0403000000','科目信息','1','0410000000','1',NULL,'true',''),('0404000000','币种信息','1','0410000000','1',NULL,'true',''),('0410000000','基础维度','0','0400000000','4',NULL,'true',''),('0500000000','批次调度系统','0','-1','0',NULL,'false','dispatch'),('0501000000','调度参数配置','0','0500000000','4',NULL,'false','dispatch'),('0501010000','任务参数定义','1','0501000000','1',NULL,'false','dispatch'),('0501020000','调度核心参数管理','1','0501000000','1',NULL,'false','dispatch'),('0502000000','任务与任务组配置','0','0500000000','4',NULL,'false','dispatch'),('0502010000','任务定义','1','0502000000','1',NULL,'false','dispatch'),('0502020000','任务组定义','1','0502000000','1',NULL,'false','dispatch'),('0503000000','批次配置管理','0','0500000000','4',NULL,'false','dispatch'),('0503010000','批次定义','1','0503000000','1',NULL,'false','dispatch'),('0503020000','批次监控','1','0503000000','1',NULL,'false','dispatch'),('0503030000','批次历史信息','1','0503000000','1',NULL,'false','dispatch'),('0600000000','数据迁移','0','-1','0',NULL,'true',''),('0601000000','数据源管理','0','0600000000','4',NULL,'true',''),('0601010000','数据源配置','1','0601000000','1',NULL,'true',''),('0601020000','数据表定义','1','0601000000','1',NULL,'true',''),('0601030000','任务定义','1','0601000000','1',NULL,'true',''),('1100000000','系统帮助','0','-1','0',NULL,'true',NULL),('1101000000','系统管理帮助','0','1100000000','4',NULL,'true',NULL),('1101010000','系统维护帮助信息','1','1101000000','1',NULL,'true',NULL),('1101020000','API文档','1','1101000000','1',NULL,'true',NULL),('1102000000','管理会计帮助文档','0','1100000000','4',NULL,'true',NULL),('1103000000','公共信息帮助','0','1100000000','4',NULL,'true',NULL),('1200000000','组织首页','0','-1','0',NULL,'false','baidu');
/*!40000 ALTER TABLE `sys_resource_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_info_attr`
--

DROP TABLE IF EXISTS `sys_resource_info_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_info_attr` (
  `res_attr` char(1) NOT NULL,
  `res_attr_desc` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`res_attr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_info_attr`
--

LOCK TABLES `sys_resource_info_attr` WRITE;
/*!40000 ALTER TABLE `sys_resource_info_attr` DISABLE KEYS */;
INSERT INTO `sys_resource_info_attr` VALUES ('0','结点'),('1','叶子');
/*!40000 ALTER TABLE `sys_resource_info_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_type_attr`
--

DROP TABLE IF EXISTS `sys_resource_type_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_type_attr` (
  `res_type` char(1) NOT NULL,
  `res_type_desc` varchar(90) DEFAULT NULL,
  PRIMARY KEY (`res_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_type_attr`
--

LOCK TABLES `sys_resource_type_attr` WRITE;
/*!40000 ALTER TABLE `sys_resource_type_attr` DISABLE KEYS */;
INSERT INTO `sys_resource_type_attr` VALUES ('0','首页菜单'),('1','子系统菜单'),('2','功能服务'),('3','功能服务结点'),('4','虚拟节点');
/*!40000 ALTER TABLE `sys_resource_type_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_define`
--

DROP TABLE IF EXISTS `sys_role_define`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_define` (
  `role_id` varchar(66) NOT NULL,
  `role_name` varchar(300) NOT NULL,
  `create_user` varchar(30) NOT NULL,
  `create_time` datetime NOT NULL,
  `role_status_id` char(1) NOT NULL,
  `modify_time` datetime NOT NULL,
  `modify_user` varchar(30) NOT NULL,
  PRIMARY KEY (`role_id`),
  KEY `fk_sys_idx_11` (`role_status_id`),
  CONSTRAINT `fk_sys_idx_11` FOREIGN KEY (`role_status_id`) REFERENCES `sys_role_status_attr` (`role_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_define`
--

LOCK TABLES `sys_role_define` WRITE;
/*!40000 ALTER TABLE `sys_role_define` DISABLE KEYS */;
INSERT INTO `sys_role_define` VALUES ('20010','中国农业银行管理员橘色','admin','2017-08-20 10:59:09','0','2017-08-20 10:59:22','admin'),('20020','中国农业银行上海分行管理员角色','admin','2017-08-20 10:59:48','0','2017-08-20 10:59:48','admin'),('20030','中国农业银行湖北省分行管理员角色','admin','2017-08-20 11:43:34','0','2017-08-20 11:43:34','admin'),('20040','中国农业银行四川省分行管理员角色','admin','2017-08-20 11:43:59','0','2017-08-20 11:43:59','admin'),('20050','中国农业银行湖南省分行管理员角色','admin','2017-08-20 11:44:18','0','2017-08-20 11:44:18','admin');
/*!40000 ALTER TABLE `sys_role_define` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_resource`
--

DROP TABLE IF EXISTS `sys_role_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_resource` (
  `uuid` varchar(60) NOT NULL DEFAULT 'uuid()',
  `role_id` varchar(66) DEFAULT NULL,
  `res_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_idx_06` (`res_id`),
  KEY `fk_sys_role_res_01_idx` (`role_id`),
  CONSTRAINT `fk_sys_idx_06` FOREIGN KEY (`res_id`) REFERENCES `sys_resource_info` (`res_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_role_res_01` FOREIGN KEY (`role_id`) REFERENCES `sys_role_define` (`role_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_resource`
--

LOCK TABLES `sys_role_resource` WRITE;
/*!40000 ALTER TABLE `sys_role_resource` DISABLE KEYS */;
INSERT INTO `sys_role_resource` VALUES ('20010_join_0100000000','20010','0100000000'),('20010_join_0101000000','20010','0101000000'),('20010_join_0101010000','20010','0101010000'),('20010_join_0101010100','20010','0101010100'),('20010_join_0101010200','20010','0101010200'),('20010_join_0101010300','20010','0101010300'),('20010_join_0102000000','20010','0102000000'),('20010_join_0103000000','20010','0103000000'),('20010_join_0103010000','20010','0103010000'),('20010_join_0103010100','20010','0103010100'),('20010_join_0103010200','20010','0103010200'),('20010_join_0103010300','20010','0103010300'),('20010_join_0103010400','20010','0103010400'),('20010_join_01030104001','20010','01030104001'),('20010_join_0103010500','20010','0103010500'),('20010_join_0103020000','20010','0103020000'),('20010_join_0103020100','20010','0103020100'),('20010_join_0103020200','20010','0103020200'),('20010_join_0103020300','20010','0103020300'),('20010_join_0103020400','20010','0103020400'),('20010_join_0103020500','20010','0103020500'),('20010_join_0103030000','20010','0103030000'),('20010_join_0103030100','20010','0103030100'),('20010_join_0103030200','20010','0103030200'),('20010_join_0103030300','20010','0103030300'),('20010_join_0103030400','20010','0103030400'),('20010_join_0103040000','20010','0103040000'),('20010_join_0103050000','20010','0103050000'),('20010_join_0103060000','20010','0103060000'),('20010_join_0103070000','20010','0103070000'),('20010_join_0103070100','20010','0103070100'),('20010_join_0104000000','20010','0104000000'),('20010_join_0104010000','20010','0104010000'),('20010_join_0104010100','20010','0104010100'),('20010_join_0104010200','20010','0104010200'),('20010_join_0104010300','20010','0104010300'),('20010_join_0104010400','20010','0104010400'),('20010_join_0104010500','20010','0104010500'),('20010_join_0105010000','20010','0105010000'),('20010_join_0105010100','20010','0105010100'),('20010_join_0105010200','20010','0105010200'),('20010_join_0105010300','20010','0105010300'),('20010_join_0105010400','20010','0105010400'),('20010_join_0105010500','20010','0105010500'),('20010_join_0105010600','20010','0105010600'),('20010_join_0105020000','20010','0105020000'),('20010_join_0105020100','20010','0105020100'),('20010_join_0105020200','20010','0105020200'),('20010_join_0105020300','20010','0105020300'),('20010_join_0105020400','20010','0105020400'),('20010_join_0105020500','20010','0105020500'),('20010_join_0105020510','20010','0105020510'),('20010_join_0105020520','20010','0105020520'),('20010_join_0105040000','20010','0105040000'),('20010_join_0105040100','20010','0105040100'),('20010_join_0105040200','20010','0105040200'),('20010_join_0400000000','20010','0400000000'),('20010_join_0401000000','20010','0401000000'),('20010_join_0402000000','20010','0402000000'),('20010_join_0403000000','20010','0403000000'),('20010_join_0404000000','20010','0404000000'),('20010_join_0410000000','20010','0410000000'),('20010_join_0600000000','20010','0600000000'),('20010_join_0601000000','20010','0601000000'),('20010_join_0601010000','20010','0601010000'),('20010_join_0601020000','20010','0601020000'),('20010_join_0601030000','20010','0601030000'),('20010_join_1100000000','20010','1100000000'),('20010_join_1101000000','20010','1101000000'),('20010_join_1101010000','20010','1101010000'),('20010_join_1101020000','20010','1101020000'),('20010_join_1102000000','20010','1102000000'),('20010_join_1103000000','20010','1103000000'),('20010_join_1200000000','20010','1200000000'),('20020_join_0100000000','20020','0100000000'),('20020_join_0101000000','20020','0101000000'),('20020_join_0101010000','20020','0101010000'),('20020_join_0101010100','20020','0101010100'),('20020_join_0101010200','20020','0101010200'),('20020_join_0101010300','20020','0101010300'),('20020_join_0102000000','20020','0102000000'),('20020_join_0103000000','20020','0103000000'),('20020_join_0103010000','20020','0103010000'),('20020_join_0103010100','20020','0103010100'),('20020_join_0103010200','20020','0103010200'),('20020_join_0103010300','20020','0103010300'),('20020_join_0103010400','20020','0103010400'),('20020_join_01030104001','20020','01030104001'),('20020_join_0103010500','20020','0103010500'),('20020_join_0103020000','20020','0103020000'),('20020_join_0103020100','20020','0103020100'),('20020_join_0103020200','20020','0103020200'),('20020_join_0103020300','20020','0103020300'),('20020_join_0103020400','20020','0103020400'),('20020_join_0103020500','20020','0103020500'),('20020_join_0103030000','20020','0103030000'),('20020_join_0103030100','20020','0103030100'),('20020_join_0103030200','20020','0103030200'),('20020_join_0103030300','20020','0103030300'),('20020_join_0103030400','20020','0103030400'),('20020_join_0103040000','20020','0103040000'),('20020_join_0103050000','20020','0103050000'),('20020_join_0103060000','20020','0103060000'),('20020_join_0103070000','20020','0103070000'),('20020_join_0103070100','20020','0103070100'),('20020_join_0104000000','20020','0104000000'),('20020_join_0104010000','20020','0104010000'),('20020_join_0104010100','20020','0104010100'),('20020_join_0104010200','20020','0104010200'),('20020_join_0104010300','20020','0104010300'),('20020_join_0104010400','20020','0104010400'),('20020_join_0104010500','20020','0104010500'),('20020_join_0105010000','20020','0105010000'),('20020_join_0105010100','20020','0105010100'),('20020_join_0105010200','20020','0105010200'),('20020_join_0105010300','20020','0105010300'),('20020_join_0105010400','20020','0105010400'),('20020_join_0105010500','20020','0105010500'),('20020_join_0105010600','20020','0105010600'),('20020_join_0105020000','20020','0105020000'),('20020_join_0105020100','20020','0105020100'),('20020_join_0105020200','20020','0105020200'),('20020_join_0105020300','20020','0105020300'),('20020_join_0105020400','20020','0105020400'),('20020_join_0105020500','20020','0105020500'),('20020_join_0105020510','20020','0105020510'),('20020_join_0105020520','20020','0105020520'),('20020_join_0105040000','20020','0105040000'),('20020_join_0105040100','20020','0105040100'),('20020_join_0105040200','20020','0105040200'),('20020_join_0200000000','20020','0200000000'),('20020_join_0202000000','20020','0202000000'),('20020_join_0202010000','20020','0202010000'),('20020_join_0202020000','20020','0202020000'),('20020_join_0202030000','20020','0202030000'),('20020_join_0202040000','20020','0202040000'),('20020_join_0203000000','20020','0203000000'),('20020_join_0203010000','20020','0203010000'),('20020_join_0203020000','20020','0203020000'),('20020_join_0203040000','20020','0203040000'),('20020_join_0203050000','20020','0203050000'),('20020_join_0400000000','20020','0400000000'),('20020_join_0401000000','20020','0401000000'),('20020_join_0402000000','20020','0402000000'),('20020_join_0403000000','20020','0403000000'),('20020_join_0404000000','20020','0404000000'),('20020_join_0410000000','20020','0410000000'),('20020_join_0500000000','20020','0500000000'),('20020_join_0501000000','20020','0501000000'),('20020_join_0501010000','20020','0501010000'),('20020_join_0501020000','20020','0501020000'),('20020_join_0502000000','20020','0502000000'),('20020_join_0502010000','20020','0502010000'),('20020_join_0502020000','20020','0502020000'),('20020_join_0503000000','20020','0503000000'),('20020_join_0503010000','20020','0503010000'),('20020_join_0503020000','20020','0503020000'),('20020_join_0503030000','20020','0503030000'),('20020_join_0600000000','20020','0600000000'),('20020_join_0601000000','20020','0601000000'),('20020_join_0601010000','20020','0601010000'),('20020_join_0601020000','20020','0601020000'),('20020_join_0601030000','20020','0601030000'),('20020_join_1100000000','20020','1100000000'),('20020_join_1101000000','20020','1101000000'),('20020_join_1101010000','20020','1101010000'),('20020_join_1101020000','20020','1101020000'),('20020_join_1102000000','20020','1102000000'),('20020_join_1103000000','20020','1103000000'),('20020_join_1200000000','20020','1200000000');
/*!40000 ALTER TABLE `sys_role_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_status_attr`
--

DROP TABLE IF EXISTS `sys_role_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_status_attr` (
  `role_status_id` char(1) NOT NULL,
  `role_status_desc` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`role_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_status_attr`
--

LOCK TABLES `sys_role_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_role_status_attr` DISABLE KEYS */;
INSERT INTO `sys_role_status_attr` VALUES ('0','正常'),('1','锁定');
/*!40000 ALTER TABLE `sys_role_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_user`
--

DROP TABLE IF EXISTS `sys_role_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_user` (
  `uuid` varchar(60) NOT NULL,
  `role_id` varchar(66) DEFAULT NULL,
  `user_id` varchar(30) DEFAULT NULL,
  `create_date` date DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_idx_03` (`user_id`),
  KEY `fk_sys_role_user_01_idx` (`role_id`),
  CONSTRAINT `fk_sys_idx_03` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_role_user_01` FOREIGN KEY (`role_id`) REFERENCES `sys_role_define` (`role_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_user`
--

LOCK TABLES `sys_role_user` WRITE;
/*!40000 ALTER TABLE `sys_role_user` DISABLE KEYS */;
INSERT INTO `sys_role_user` VALUES ('admin_join_20010','20010','admin','2017-08-22','admin'),('admin_join_20020','20020','admin','2017-08-22','admin'),('ccbc_admin_join_20010','20010','ccbc_admin','2017-08-22','admin'),('ccbc_admin_join_20020','20020','ccbc_admin','2017-08-22','admin'),('ccbc_admin_join_20030','20030','ccbc_admin','2017-08-21','admin'),('ccbc_admin_join_20040','20040','ccbc_admin','2017-08-21','admin'),('ccbc_admin_join_20050','20050','ccbc_admin','2017-08-21','admin');
/*!40000 ALTER TABLE `sys_role_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_sec_user`
--

DROP TABLE IF EXISTS `sys_sec_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_sec_user` (
  `user_id` varchar(30) NOT NULL,
  `user_passwd` varchar(30) DEFAULT NULL,
  `status_id` char(1) DEFAULT NULL,
  `continue_error_cnt` int(11) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_idx_02` (`status_id`),
  CONSTRAINT `fk_sys_idx_01` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_idx_02` FOREIGN KEY (`status_id`) REFERENCES `sys_user_status_attr` (`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_sec_user`
--

LOCK TABLES `sys_sec_user` WRITE;
/*!40000 ALTER TABLE `sys_sec_user` DISABLE KEYS */;
INSERT INTO `sys_sec_user` VALUES ('admin','rVbaiQ3XuCj8aCnhIL1KAA==','0',0),('ccbc_admin','CguSVgQY2Df4LxG0UT/xwA==','0',0);
/*!40000 ALTER TABLE `sys_sec_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_theme_define`
--

DROP TABLE IF EXISTS `sys_theme_define`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_theme_define` (
  `theme_id` varchar(30) NOT NULL,
  `theme_desc` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`theme_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_theme_define`
--

LOCK TABLES `sys_theme_define` WRITE;
/*!40000 ALTER TABLE `sys_theme_define` DISABLE KEYS */;
INSERT INTO `sys_theme_define` VALUES ('1001','绿色主题'),('1002','深蓝主题'),('1003','粉色主题'),('1004','青色主题'),('1005','传统布局');
/*!40000 ALTER TABLE `sys_theme_define` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_theme_resource`
--

DROP TABLE IF EXISTS `sys_theme_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_theme_resource` (
  `uuid` varchar(60) NOT NULL COMMENT '系统内唯一编码',
  `theme_id` varchar(30) DEFAULT NULL COMMENT '主题编码',
  `res_id` varchar(30) DEFAULT NULL COMMENT '资源编码',
  `res_url` varchar(120) DEFAULT NULL COMMENT '资源地址',
  `res_open_type` varchar(5) DEFAULT NULL COMMENT '打开方式：0 内部区域打开，1  新建选项卡打开',
  `res_bg_color` varchar(30) DEFAULT NULL COMMENT '菜单背景色（Metro风格主题使用）',
  `res_class` varchar(90) DEFAULT NULL COMMENT '菜单样式类（Metro风格主题使用）',
  `group_id` char(1) DEFAULT NULL COMMENT '组编号',
  `res_img` varchar(200) DEFAULT NULL COMMENT '图标',
  `sort_id` decimal(10,0) DEFAULT NULL COMMENT '组内排序号',
  `new_iframe` varchar(5) DEFAULT NULL COMMENT '使用使用iframe打开，true 使用 false 不使用',
  PRIMARY KEY (`uuid`),
  KEY `pk_sys_theme_value_01` (`uuid`),
  KEY `fk_sys_theme_resource01_idx` (`res_id`),
  CONSTRAINT `fk_sys_theme_resource01` FOREIGN KEY (`res_id`) REFERENCES `sys_resource_info` (`res_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_theme_resource`
--

LOCK TABLES `sys_theme_resource` WRITE;
/*!40000 ALTER TABLE `sys_theme_resource` DISABLE KEYS */;
INSERT INTO `sys_theme_resource` VALUES ('00714873-07ed-11e7-952f-a0c58951c8d5','1001','0105010600','/v1/auth/user/modify/status','0','','','','',0,'false'),('0130f7d4-7d9a-11e7-97cd-a0c58951c8d5','1005','0203010000','/v1/ca/dispatch/page','0','#666699','tile tile-wide','3','/static/images/mdui/mdl-005.png',3,'false'),('015376ca-2b2b-11e7-9c7e-a0c58951c8d5','1002','0400000000','./views/mas/common/blue/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('0287ee48-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020400','/v1/auth/role/delete','0','','tile','','',0,'false'),('04625d55-83c4-11e7-9b5b-a0c58951c8d5','1005','0601010000','/v1/asofdate/data/migration/datasource/page','0','#339999','tile tile-wide','1','/static/images/mdui/mdl-004.png',1,'false'),('052dc4ac-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020500','/v1/auth/role/resource/details','0','','tile','','',0,'false'),('0574add7-07e7-11e7-952f-a0c58951c8d5','1001','0103020300','/v1/auth/resource/org/update','0','','','','',0,'false'),('06cd1570-720f-11e7-963e-a0c58951c8d5','1002','0103050000','/v1/sso/subsystem/api/page','0','#ed9f86','tile','1','/static/images/ca_icon/index.png',4,'false'),('07ab049a-2b2b-11e7-9c7e-a0c58951c8d5','1002','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1,'false'),('0875a5f3-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020510','/v1/auth/role/resource/get','0','','tile','','',0,'false'),('08de2dad-720f-11e7-963e-a0c58951c8d5','1003','0103050000','/v1/sso/subsystem/api/page','0','#ed9f86','tile','1','/static/images/ca_icon/index.png',4,'false'),('098dd130-2b2b-11e7-9c7e-a0c58951c8d5','1002','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2,'false'),('0a6ed393-720f-11e7-963e-a0c58951c8d5','1004','0103050000','/v1/sso/subsystem/api/page','0','#ed9f86','tile','1','/static/images/ca_icon/index.png',4,'false'),('0a964ef9-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020520','/v1/auth/role/resource/rights','0','','tile','','',0,'false'),('0bbabbfb-2b2b-11e7-9c7e-a0c58951c8d5','1002','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1,'false'),('0db2afab-2b2b-11e7-9c7e-a0c58951c8d5','1002','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1,'false'),('0e4ca28b-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4,'false'),('0f9303e2-07f2-11e7-952f-a0c58951c8d5','1001','0105040100','/v1/auth/user/roles/auth','0','','','','',0,'false'),('1001-0100000000','1001','0100000000','./views/hauth/theme/default/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1,'false'),('1001-0101010000','1001','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1,'false'),('1001-0103010000','1001','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',5,'false'),('1001-0103020000','1001','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','2','/static/images/hauth/org.png',2,'false'),('1001-0104010000','1001','0104010000','/v1/auth/domain/page','0','#0099CC','tile','2','/static/images/hauth/domain.png',1,'false'),('1001-0105010000','1001','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',3,'false'),('1001-0105020000','1001','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',4,'false'),('1001-0105040000','1001','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4,'false'),('107e273d-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040100','/v1/auth/user/roles/auth','0','','tile','','',0,'false'),('12cd5409-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040200','/v1/auth/user/roles/revoke','0','','tile','','',0,'false'),('16094161-83c4-11e7-9b5b-a0c58951c8d5','1005','0601020000','/v1/asofdate/data/migration/table/page','0','#339999','tile tile-wide','1','/static/images/mdui/mdl-004.png',2,'false'),('167bd2eb-7d99-11e7-97cd-a0c58951c8d5','1005','0201060000','/v1/ca/cost/page','0','#e4d690','tile tile-wide','1','/static/images/mdui/mdl-003.png',4,'false'),('1bde8991-07e9-11e7-952f-a0c58951c8d5','1001','0103010100','/v1/auth/resource/get','0','','','','',0,'false'),('1bf270aa-07e7-11e7-952f-a0c58951c8d5','1001','0103020400','/v1/auth/resource/org/delete','0','','','','',0,'false'),('1c30f988-07e2-11e7-952f-a0c58951c8d5','1001','0103030400','/v1/auth/domain/share/put','0','','','','',0,'false'),('1fc99c96-7d9a-11e7-97cd-a0c58951c8d5','1005','0203020000','/v1/ca/dispatch/history/page','0','#339999','tile tile-wide','3','/static/images/mdui/mdl-006.png',4,'false'),('25165700-07f2-11e7-952f-a0c58951c8d5','1001','0105040200','/v1/auth/user/roles/revoke','0','','','','',0,'false'),('264e0d55-71fb-11e7-963e-a0c58951c8d5','1001','0103030000','/v1/sso/local/static/page','0','#009966','tile','1','/static/images/ca_icon/index.png',4,'false'),('287973cc-83c4-11e7-9b5b-a0c58951c8d5','1005','0601030000','/v1/asofdate/data/migration/job/page','0','#339999','tile tile-wide','1','/static/images/mdui/mdl-004.png',3,'false'),('316dda9d-71fb-11e7-963e-a0c58951c8d5','1002','0103030000','/v1/sso/local/static/page','0','#009966','tile','1','/static/images/ca_icon/index.png',4,'false'),('32cb5534-2b45-11e7-9c7e-a0c58951c8d5','1003','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1,'false'),('33388d48-71fb-11e7-963e-a0c58951c8d5','1003','0103030000','/v1/sso/local/static/page','0','#009966','tile','1','/static/images/ca_icon/index.png',4,'false'),('33b9cb0c-07e9-11e7-952f-a0c58951c8d5','1001','0103010200','/v1/auth/resource/post','0','','','','',0,'false'),('3403b3b7-2b44-11e7-9c7e-a0c58951c8d5','1003','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1,'false'),('350dd891-2b45-11e7-9c7e-a0c58951c8d5','1003','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2,'false'),('353f9592-71fb-11e7-963e-a0c58951c8d5','1004','0103030000','/v1/sso/local/static/page','0','#009966','tile','1','/static/images/ca_icon/index.png',4,'false'),('354f834f-7d9a-11e7-97cd-a0c58951c8d5','1005','0203040000','/v1/ca/cost/manage/page','0','#FFCC33','tile','3','/static/images/mdui/mdl-007.png',1,'false'),('377ea0e5-7e58-11e7-97cd-a0c58951c8d5','1005','0103070000','/v1/auth/resource/service','0','#009966','tile tile-wide','1','/static/images/mdui/mdl-001.png',5,'false'),('37ceac85-2b44-11e7-9c7e-a0c58951c8d5','1003','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1,'false'),('3a0e741e-2b44-11e7-9c7e-a0c58951c8d5','1003','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1,'false'),('3d237ba7-07e7-11e7-952f-a0c58951c8d5','1001','0103020500','/v1/auth/resource/org/download','0','','','','',0,'false'),('3e3b64e8-2b45-11e7-9c7e-a0c58951c8d5','1003','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1,'false'),('3ee58eee-7d92-11e7-97cd-a0c58951c8d5','1005','0201010000','/v1/ca/responsibility/page','0','#6fc07c','tile tile-wide','1','/static/images/mdui/mdl-004.png',1,'false'),('40813c9f-2b45-11e7-9c7e-a0c58951c8d5','1003','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1,'false'),('40940862-720f-11e7-963e-a0c58951c8d5','1001','0103060000','/v1/sso/proxy/static/page','0','#FFCC33','tile','1','/static/images/dispatch_icon/etl.png',3,'false'),('43ad2a9a-07f1-11e7-952f-a0c58951c8d5','1001','0105020510','/v1/auth/role/resource/get','0','','','','',0,'false'),('44ee39cc-720f-11e7-963e-a0c58951c8d5','1002','0103060000','/v1/sso/proxy/static/page','0','#FFCC33','tile','1','/static/images/dispatch_icon/etl.png',3,'false'),('46c22c38-720f-11e7-963e-a0c58951c8d5','1003','0103060000','/v1/sso/proxy/static/page','0','#FFCC33','tile','1','/static/images/dispatch_icon/etl.png',3,'false'),('48236448-83c4-11e7-9b5b-a0c58951c8d5','1005','0600000000','www.asofdate.com','0','#339999','tile tile-wide','2','/static/images/hauth/system.png',3,'false'),('48460086-07e9-11e7-952f-a0c58951c8d5','1001','0103010300','/v1/auth/resource/update','0','','','','',0,'false'),('4892c00e-720f-11e7-963e-a0c58951c8d5','1004','0103060000','/v1/sso/proxy/static/page','0','#FFCC33','tile','1','/static/images/dispatch_icon/etl.png',3,'false'),('4b32d9c0-7d90-11e7-97cd-a0c58951c8d5','1005','0501010000','/v1/dispatch/argument/page','0','#009966','tile','1','/static/images/mdui/mdl-001.png',1,'false'),('4f87bd68-7d9a-11e7-97cd-a0c58951c8d5','1005','0203050000','/v1/ca/driver/manage/page','0','#CC6600','tile','3','/static/images/mdui/mdl-008.png',2,'false'),('4fd8fdcf-2b42-11e7-9c7e-a0c58951c8d5','1002','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1,'false'),('51a5bff2-2b28-11e7-9c7e-a0c58951c8d5','1002','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',3,'false'),('527f4d10-8250-11e7-9b5b-a0c58951c8d5','1005','0103070100','/v1/auth/menu/all/except/button','0','','','','',0,'false'),('5a7d8dbf-07f1-11e7-952f-a0c58951c8d5','1001','0105020520','/v1/auth/role/resource/rights','0','','','','',0,'false'),('5dcfdfc0-2b42-11e7-9c7e-a0c58951c8d5','1002','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1,'false'),('60c6e788-2b42-11e7-9c7e-a0c58951c8d5','1002','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1,'false'),('624b90c0-0278-11e7-9b60-a0c58951c8d5','1002','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1,'false'),('64644fc7-71d9-11e7-963e-a0c58951c8d5','1003','1200000000','www.asofdate.com','0','#666699','tile tile-wide','1','/static/images/hauth/api.png',2,'true'),('686b1d59-824d-11e7-9b5b-a0c58951c8d5','1001','0103070100','/v1/auth/menu/all/except/button','0','','','','',0,'false'),('6ad656f4-7d93-11e7-97cd-a0c58951c8d5','1005','0200000000','./views/mas/ca/apple/ca.tpl','0','#666699','tile tile-wide','2','/static/images/hauth/system.png',2,'false'),('6bb7b2c8-07e9-11e7-952f-a0c58951c8d5','1001','0103010400','/v1/auth/resource/delete','0','','','','',0,'false'),('6c7f5772-250a-11e7-9c7e-a0c58951c8d5','1001','01030104001','/v1/auth/resource/org/page','0','','','','',0,'false'),('7d73058c-07ec-11e7-952f-a0c58951c8d5','1001','0105010100','/v1/auth/user/get','0','','','','',0,'false'),('8024ac09-07d8-11e7-952f-a0c58951c8d5','1001','0104010300','/v1/auth/domain/update','0','','','','',0,'false'),('824c0d97-04a3-11e7-9b60-a0c58951c8d5','1001','0400000000','./views/mas/common/green/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('8ca386d8-07e5-11e7-952f-a0c58951c8d5','1001','0101010200','/v1/auth/handle/logs/download','0','','','','',0,'false'),('8e2d2ae7-1c0a-11e7-9d82-a0c58951c8d5','1004','0101010100','/v1/auth/handle/logs','0','','tile tile-large','','',0,'false'),('946658e9-07d5-11e7-952f-a0c58951c8d5','1001','0104010200','/v1/auth/domain/share/page','0','','','','',0,'false'),('948f67dc-024a-11e7-9b60-a0c58951c8d5','1001','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1,'false'),('952d95e3-7d99-11e7-97cd-a0c58951c8d5','1005','0202010000','/v1/ca/static/radio/page','0','#92cdd2','tile tile-wide','2','/static/images/mdui/mdl-003.png',1,'false'),('9705437b-07d8-11e7-952f-a0c58951c8d5','1001','0104010400','/v1/auth/domain/delete','0','','','','',0,'false'),('974ce1fd-07ec-11e7-952f-a0c58951c8d5','1001','0105010200','/v1/auth/user/post','0','','','','',0,'false'),('991641c3-0d55-11e7-964b-a0c58951c8d5','1004','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1,'false'),('99164f5c-0d55-11e7-964b-a0c58951c8d5','1004','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',5,'false'),('9916502d-0d55-11e7-964b-a0c58951c8d5','1004','0104010000','/v1/auth/domain/page','0','#0099CC','tile','2','/static/images/hauth/domain.png',1,'false'),('991650a9-0d55-11e7-964b-a0c58951c8d5','1004','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',3,'false'),('9916512d-0d55-11e7-964b-a0c58951c8d5','1004','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',4,'false'),('9916519c-0d55-11e7-964b-a0c58951c8d5','1004','0100000000','./views/hauth/theme/cyan/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1,'false'),('99165203-0d55-11e7-964b-a0c58951c8d5','1004','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4,'false'),('9916525c-0d55-11e7-964b-a0c58951c8d5','1004','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','2','/static/images/hauth/org.png',2,'false'),('9917f6e5-0d55-11e7-964b-a0c58951c8d5','1004','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1,'false'),('99180bfa-0d55-11e7-964b-a0c58951c8d5','1004','0400000000','./views/mas/common/cyan/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('99180c36-0d55-11e7-964b-a0c58951c8d5','1004','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1,'false'),('99180c72-0d55-11e7-964b-a0c58951c8d5','1004','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2,'false'),('99180ca9-0d55-11e7-964b-a0c58951c8d5','1004','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1,'false'),('99180ced-0d55-11e7-964b-a0c58951c8d5','1004','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1,'false'),('99180d65-0d55-11e7-964b-a0c58951c8d5','1004','0104010100','/v1/auth/domain/get','0','','','','',0,'false'),('99180da1-0d55-11e7-964b-a0c58951c8d5','1004','0104010200','/v1/auth/domain/share/page','0','','','','',0,'false'),('99180ddc-0d55-11e7-964b-a0c58951c8d5','1004','0104010300','/v1/auth/domain/update','0','','','','',0,'false'),('99180e14-0d55-11e7-964b-a0c58951c8d5','1004','0104010400','/v1/auth/domain/delete','0','','','','',0,'false'),('99180e4f-0d55-11e7-964b-a0c58951c8d5','1004','0104010500','/v1/auth/domain/post','0','','','','',0,'false'),('99180e87-0d55-11e7-964b-a0c58951c8d5','1004','0103030100','/v1/auth/domain/share/get','0','','','','',0,'false'),('99180ec3-0d55-11e7-964b-a0c58951c8d5','1004','0103030200','/v1/auth/domain/share/post','0','','','','',0,'false'),('99180efa-0d55-11e7-964b-a0c58951c8d5','1004','0103030300','/v1/auth/domain/share/delete','0','','','','',0,'false'),('99180f32-0d55-11e7-964b-a0c58951c8d5','1004','0103030400','/v1/auth/domain/share/put','0','','','','',0,'false'),('99180fa1-0d55-11e7-964b-a0c58951c8d5','1004','0101010200','/v1/auth/handle/logs/download','0','','','','',0,'false'),('99180fdc-0d55-11e7-964b-a0c58951c8d5','1004','0101010300','/v1/auth/handle/logs/search','0','','','','',0,'false'),('99181014-0d55-11e7-964b-a0c58951c8d5','1004','0103020100','/v1/auth/resource/org/get','0','','','','',0,'false'),('9918104b-0d55-11e7-964b-a0c58951c8d5','1004','0103020200','/v1/auth/resource/org/insert','0','','','','',0,'false'),('99181087-0d55-11e7-964b-a0c58951c8d5','1004','0103020300','/v1/auth/resource/org/update','0','','','','',0,'false'),('991810be-0d55-11e7-964b-a0c58951c8d5','1004','0103020400','/v1/auth/resource/org/delete','0','','','','',0,'false'),('991810fe-0d55-11e7-964b-a0c58951c8d5','1004','0103020500','/v1/auth/resource/org/download','0','','','','',0,'false'),('9918113a-0d55-11e7-964b-a0c58951c8d5','1004','0103010100','/v1/auth/resource/get','0','','','','',0,'false'),('99181176-0d55-11e7-964b-a0c58951c8d5','1004','0103010200','/v1/auth/resource/post','0','','','','',0,'false'),('991811ad-0d55-11e7-964b-a0c58951c8d5','1004','0103010300','/v1/auth/resource/update','0','','','','',0,'false'),('991811e1-0d55-11e7-964b-a0c58951c8d5','1004','0103010400','/v1/auth/resource/delete','0','','','','',0,'false'),('99181218-0d55-11e7-964b-a0c58951c8d5','1004','0103010500','/v1/auth/resource/config/theme','0','','','','',0,'false'),('9918124f-0d55-11e7-964b-a0c58951c8d5','1004','0105010100','/v1/auth/user/get','0','','','','',0,'false'),('9918128b-0d55-11e7-964b-a0c58951c8d5','1004','0105010200','/v1/auth/user/post','0','','','','',0,'false'),('991812c3-0d55-11e7-964b-a0c58951c8d5','1004','0105010300','/v1/auth/user/put','0','','','','',0,'false'),('991812fa-0d55-11e7-964b-a0c58951c8d5','1004','0105010400','/v1/auth/user/delete','0','','','','',0,'false'),('99181332-0d55-11e7-964b-a0c58951c8d5','1004','0105010500','/v1/auth/user/modify/passwd','0','','','','',0,'false'),('99181365-0d55-11e7-964b-a0c58951c8d5','1004','0105010600','/v1/auth/user/modify/status','0','','','','',0,'false'),('9918139c-0d55-11e7-964b-a0c58951c8d5','1004','0105020100','/v1/auth/role/get','0','','','','',0,'false'),('991813d4-0d55-11e7-964b-a0c58951c8d5','1004','0105020200','/v1/auth/role/post','0','','','','',0,'false'),('9918140b-0d55-11e7-964b-a0c58951c8d5','1004','0105020300','/v1/auth/role/update','0','','','','',0,'false'),('99181443-0d55-11e7-964b-a0c58951c8d5','1004','0105020400','/v1/auth/role/delete','0','','','','',0,'false'),('99181476-0d55-11e7-964b-a0c58951c8d5','1004','0105020500','/v1/auth/role/resource/details','0','','','','',0,'false'),('991814ad-0d55-11e7-964b-a0c58951c8d5','1004','0105020510','/v1/auth/role/resource/get','0','','','','',0,'false'),('991814f2-0d55-11e7-964b-a0c58951c8d5','1004','0105020520','/v1/auth/role/resource/rights','0','','','','',0,'false'),('9918152d-0d55-11e7-964b-a0c58951c8d5','1004','0105040100','/v1/auth/user/roles/auth','0','','','','',0,'false'),('99181569-0d55-11e7-964b-a0c58951c8d5','1004','0105040200','/v1/auth/user/roles/revoke','0','','','','',0,'false'),('991815e1-0d55-11e7-964b-a0c58951c8d5','1004','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1,'false'),('9b081aec-2b27-11e7-9c7e-a0c58951c8d5','1002','0100000000','./views/hauth/theme/blue/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1,'false'),('a0e208f2-20f8-11e7-966c-a0c58951c8d5','1001','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1,'false'),('a265597d-07ed-11e7-952f-a0c58951c8d5','1001','0105020100','/v1/auth/role/get','0','','','','',0,'false'),('a29fba3f-07e5-11e7-952f-a0c58951c8d5','1001','0101010300','/v1/auth/handle/logs/search','0','','','','',0,'false'),('a343cbfc-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010100','/v1/auth/handle/logs','0','','tile','','',0,'false'),('a65d91b0-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010200','/v1/auth/handle/logs/download','0','','tile','','',0,'false'),('a8854ec0-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010300','/v1/auth/handle/logs/search','0','','tile','','',0,'false'),('a96e70c9-7da3-11e7-97cd-a0c58951c8d5','1005','0202030000','/v1/ca/static/config/page','0','','tile','','',0,'true'),('aabbbd36-2b27-11e7-9c7e-a0c58951c8d5','1002','01030104001','/v1/auth/resource/org/page','0','','tile','','',0,'false'),('ad3e295c-07d8-11e7-952f-a0c58951c8d5','1001','0104010500','/v1/auth/domain/post','0','','','','',0,'false'),('adb32811-7b84-11e7-97cd-a0c58951c8d5','1005','0105010600','/v1/auth/user/modify/status','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb358a6-7b84-11e7-97cd-a0c58951c8d5','1005','0103020300','/v1/auth/resource/org/update','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3b14d-7b84-11e7-97cd-a0c58951c8d5','1005','0105040100','/v1/auth/user/roles/auth','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3b473-7b84-11e7-97cd-a0c58951c8d5','1005','0100000000','./views/hauth/theme/default/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1,'false'),('adb3e7e2-7b84-11e7-97cd-a0c58951c8d5','1005','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/mdui/mdl-001.png',1,'false'),('adb3efea-7b84-11e7-97cd-a0c58951c8d5','1005','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/mdui/mdl-002.png',6,'false'),('adb3f326-7b84-11e7-97cd-a0c58951c8d5','1005','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','2','/static/images/mdui/mdl-007.png',2,'false'),('adb3f44d-7b84-11e7-97cd-a0c58951c8d5','1005','0104010000','/v1/auth/domain/page','0','#0099CC','tile','2','/static/images/mdui/mdl-008.png',1,'false'),('adb3f56f-7b84-11e7-97cd-a0c58951c8d5','1005','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/mdui/mdl-009.png',3,'false'),('adb3f666-7b84-11e7-97cd-a0c58951c8d5','1005','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/mdui/mdl-010.png',4,'false'),('adb3f784-7b84-11e7-97cd-a0c58951c8d5','1005','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/mdui/mdl-011.png',5,'false'),('adb3f8d9-7b84-11e7-97cd-a0c58951c8d5','1005','0103010100','/v1/auth/resource/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3fa08-7b84-11e7-97cd-a0c58951c8d5','1005','0103020400','/v1/auth/resource/org/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3fb26-7b84-11e7-97cd-a0c58951c8d5','1005','0103030400','/v1/auth/domain/share/put','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3fc62-7b84-11e7-97cd-a0c58951c8d5','1005','0105040200','/v1/auth/user/roles/revoke','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb3fd77-7b84-11e7-97cd-a0c58951c8d5','1005','0103030000','/v1/sso/local/static/page','0','#009966','tile','1','/static/images/mdui/mdl-003.png',4,'false'),('adb3ff11-7b84-11e7-97cd-a0c58951c8d5','1005','0103010200','/v1/auth/resource/post','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40162-7b84-11e7-97cd-a0c58951c8d5','1005','0103020500','/v1/auth/resource/org/download','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40288-7b84-11e7-97cd-a0c58951c8d5','1005','0103060000','/v1/sso/proxy/static/page','0','#FFCC33','tile','1','/static/images/mdui/mdl-006.png',3,'false'),('adb403a2-7b84-11e7-97cd-a0c58951c8d5','1005','0105020510','/v1/auth/role/resource/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb404bc-7b84-11e7-97cd-a0c58951c8d5','1005','0103010300','/v1/auth/resource/update','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40673-7b84-11e7-97cd-a0c58951c8d5','1005','0105020520','/v1/auth/role/resource/rights','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40795-7b84-11e7-97cd-a0c58951c8d5','1005','0103010400','/v1/auth/resource/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40880-7b84-11e7-97cd-a0c58951c8d5','1005','01030104001','/v1/auth/resource/org/page','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40984-7b84-11e7-97cd-a0c58951c8d5','1005','0105010100','/v1/auth/user/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40a88-7b84-11e7-97cd-a0c58951c8d5','1005','0104010300','/v1/auth/domain/update','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40b37-7b84-11e7-97cd-a0c58951c8d5','1005','0400000000','./views/mas/common/green/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('adb40bef-7b84-11e7-97cd-a0c58951c8d5','1005','0101010200','/v1/auth/handle/logs/download','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40c91-7b84-11e7-97cd-a0c58951c8d5','1005','0104010200','/v1/auth/domain/share/page','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40d37-7b84-11e7-97cd-a0c58951c8d5','1005','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('adb40dde-7b84-11e7-97cd-a0c58951c8d5','1005','0104010400','/v1/auth/domain/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40e80-7b84-11e7-97cd-a0c58951c8d5','1005','0105010200','/v1/auth/user/post','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb40f22-7b84-11e7-97cd-a0c58951c8d5','1005','1101020000','/v1/auth/swagger/page','0','#339999','tile tile-wide','2','/static/images/mdui/mdl-001.png',1,'true'),('adb40fc8-7b84-11e7-97cd-a0c58951c8d5','1005','0105020100','/v1/auth/role/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41077-7b84-11e7-97cd-a0c58951c8d5','1005','0101010300','/v1/auth/handle/logs/search','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41119-7b84-11e7-97cd-a0c58951c8d5','1005','0104010500','/v1/auth/domain/post','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb411cd-7b84-11e7-97cd-a0c58951c8d5','1005','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/mdui/mdl-001.png',1,'false'),('adb41273-7b84-11e7-97cd-a0c58951c8d5','1005','0105010300','/v1/auth/user/put','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41326-7b84-11e7-97cd-a0c58951c8d5','1005','0103010500','/v1/auth/resource/config/theme','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb413d9-7b84-11e7-97cd-a0c58951c8d5','1005','0105020200','/v1/auth/role/post','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41488-7b84-11e7-97cd-a0c58951c8d5','1005','0101010100','/v1/auth/handle/logs','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41533-7b84-11e7-97cd-a0c58951c8d5','1005','0103030100','/v1/auth/domain/share/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb415de-7b84-11e7-97cd-a0c58951c8d5','1005','0105020500','/v1/auth/role/resource/details','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41684-7b84-11e7-97cd-a0c58951c8d5','1005','0105010400','/v1/auth/user/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41722-7b84-11e7-97cd-a0c58951c8d5','1005','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/mdui/mdl-001.png',2,'false'),('adb417cd-7b84-11e7-97cd-a0c58951c8d5','1005','0105020300','/v1/auth/role/update','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41873-7b84-11e7-97cd-a0c58951c8d5','1005','0103030200','/v1/auth/domain/share/post','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41911-7b84-11e7-97cd-a0c58951c8d5','1005','0103020100','/v1/auth/resource/org/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb419af-7b84-11e7-97cd-a0c58951c8d5','1005','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/mdui/mdl-001.png',1,'false'),('adb41a59-7b84-11e7-97cd-a0c58951c8d5','1005','0103040000','/v1/sso/subsystem/page','0','#339999','tile tile-wide','1','/static/images/mdui/mdl-004.png',1,'false'),('adb41b00-7b84-11e7-97cd-a0c58951c8d5','1005','0105020400','/v1/auth/role/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41baf-7b84-11e7-97cd-a0c58951c8d5','1005','0105010500','/v1/auth/user/modify/passwd','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41c66-7b84-11e7-97cd-a0c58951c8d5','1005','0103020200','/v1/auth/resource/org/insert','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41d08-7b84-11e7-97cd-a0c58951c8d5','1005','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-wide','2','/static/images/hauth/system.png',1,'false'),('adb41db7-7b84-11e7-97cd-a0c58951c8d5','1005','0104010100','/v1/auth/domain/get','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb41e59-7b84-11e7-97cd-a0c58951c8d5','1005','1200000000','www.asofdate.com','1','#666699','tile tile-wide','1','/static/images/hauth/api.png',2,'true'),('adb41f04-7b84-11e7-97cd-a0c58951c8d5','1005','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/mdui/mdl-001.png',1,'false'),('adb41faf-7b84-11e7-97cd-a0c58951c8d5','1005','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/mdui/mdl-001.png',1,'false'),('adb42051-7b84-11e7-97cd-a0c58951c8d5','1005','0103030300','/v1/auth/domain/share/delete','0','','','','/static/images/mdui/mdl-001.png',0,'false'),('adb420f7-7b84-11e7-97cd-a0c58951c8d5','1005','0103050000','/v1/sso/subsystem/api/page','0','#ed9f86','tile','1','/static/images/mdui/mdl-005.png',2,'false'),('af0e054c-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',5,'false'),('b1314131-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010100','/v1/auth/resource/get','0','','tile','','',0,'false'),('b2561d1e-04a3-11e7-9b60-a0c58951c8d5','1001','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1,'false'),('b3c7c6a6-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010200','/v1/auth/resource/post','0','','tile','','',0,'false'),('b3f18e0b-20f8-11e7-966c-a0c58951c8d5','1004','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1,'false'),('b58002f6-07ec-11e7-952f-a0c58951c8d5','1001','0105010300','/v1/auth/user/put','0','','','','',0,'false'),('b6372ff3-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010300','/v1/auth/resource/update','0','','tile','','',0,'false'),('b8d3d1c1-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010400','/v1/auth/resource/delete','0','','tile','','',0,'false'),('b8df0cd7-07e9-11e7-952f-a0c58951c8d5','1001','0103010500','/v1/auth/resource/config/theme','0','','','','',0,'false'),('ba15af88-2b44-11e7-9c7e-a0c58951c8d5','1003','0400000000','./views/mas/common/apple/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1,'false'),('ba50273b-7d99-11e7-97cd-a0c58951c8d5','1005','0202020000','/v1/ca/amart/rules/page','0','#3399CC','tile tile-wide','2','/static/images/mdui/mdl-008.png',2,'false'),('bb9fc76f-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010500','/v1/auth/resource/config/theme','0','','tile','','',0,'false'),('bd264fd7-07ed-11e7-952f-a0c58951c8d5','1001','0105020200','/v1/auth/role/post','0','','','','',0,'false'),('bea9df22-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020100','/v1/auth/resource/org/get','0','','tile','','',0,'false'),('becde5db-0eb9-11e7-9612-a0c58951c8d5','1001','0101010100','/v1/auth/handle/logs','0','','','','',0,'false'),('c1174621-07e1-11e7-952f-a0c58951c8d5','1001','0103030100','/v1/auth/domain/share/get','0','','','','',0,'false'),('c15e0f8b-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020200','/v1/auth/resource/org/insert','0','','tile','','',0,'false'),('c37806f8-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020300','/v1/auth/resource/org/update','0','','tile','','',0,'false'),('c3bad47b-07ee-11e7-952f-a0c58951c8d5','1001','0105020500','/v1/auth/role/resource/details','0','','','','',0,'false'),('c59c3303-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020400','/v1/auth/resource/org/delete','0','','tile','','',0,'false'),('c77c6ed0-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020500','/v1/auth/resource/org/download','0','','tile','','',0,'false'),('c988bb89-07ec-11e7-952f-a0c58951c8d5','1001','0105010400','/v1/auth/user/delete','0','','','','',0,'false'),('cb4afcc4-04a3-11e7-9b60-a0c58951c8d5','1001','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2,'false'),('cc8891d2-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','2','/static/images/hauth/org.png',2,'false'),('cd5ca47f-7d98-11e7-97cd-a0c58951c8d5','1005','0201030000','/v1/ca/cost/direction/page','0','#FF9999','tile','1','/static/images/mdui/mdl-002.png',2,'false'),('d1f01d28-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010000','/v1/auth/domain/page','0','#0099CC','tile','2','/static/images/hauth/domain.png',1,'false'),('d4605f79-2b43-11e7-9c7e-a0c58951c8d5','1003','0100000000','./views/hauth/theme/apple/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1,'false'),('d4bfa83c-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010100','/v1/auth/domain/get','0','','tile','','',0,'false'),('d517aab8-07ed-11e7-952f-a0c58951c8d5','1001','0105020300','/v1/auth/role/update','0','','','','',0,'false'),('d767f63e-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010200','/v1/auth/domain/share/page','0','','tile','','',0,'false'),('d7867a7a-2b43-11e7-9c7e-a0c58951c8d5','1003','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1,'false'),('d8fccbcb-07e1-11e7-952f-a0c58951c8d5','1001','0103030200','/v1/auth/domain/share/post','0','','','','',0,'false'),('da84a5e1-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030100','/v1/auth/domain/share/get','0','','tile','','',0,'false'),('daadf91b-07e6-11e7-952f-a0c58951c8d5','1001','0103020100','/v1/auth/resource/org/get','0','','','','',0,'false'),('dc65642a-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030200','/v1/auth/domain/share/post','0','','tile','','',0,'false'),('dd972a84-2b43-11e7-9c7e-a0c58951c8d5','1003','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',5,'false'),('de8f9fcb-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030300','/v1/auth/domain/share/delete','0','','tile','','',0,'false'),('dee49b86-7d99-11e7-97cd-a0c58951c8d5','1005','0202040000','/v1/ca/amart/group/page','0','#99CC33','tile tile-wide','2','/static/images/mdui/mdl-007.png',3,'false'),('e007d284-2b43-11e7-9c7e-a0c58951c8d5','1003','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','2','/static/images/hauth/org.png',2,'false'),('e0a10dc4-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030400','/v1/auth/domain/share/put','0','','tile','','',0,'false'),('e224205c-2b43-11e7-9c7e-a0c58951c8d5','1003','0104010000','/v1/auth/domain/page','0','#0099CC','tile','2','/static/images/hauth/domain.png',1,'false'),('e2e782c4-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010300','/v1/auth/domain/update','0','','tile','','',0,'false'),('e4ac3710-2b43-11e7-9c7e-a0c58951c8d5','1003','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',3,'false'),('e4e17463-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010400','/v1/auth/domain/delete','0','','tile','','',0,'false'),('e6191fef-04a3-11e7-9b60-a0c58951c8d5','1001','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1,'false'),('e716b0a1-2b43-11e7-9c7e-a0c58951c8d5','1003','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',4,'false'),('e777d2c2-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010500','/v1/auth/domain/post','0','','tile','','',0,'false'),('e7e063d6-7207-11e7-963e-a0c58951c8d5','1001','0103040000','/v1/sso/subsystem/page','0','#339999','tile tile-wide','1','/static/images/ca_icon/dispatch_manage.png',2,'false'),('ea237b6a-07ed-11e7-952f-a0c58951c8d5','1001','0105020400','/v1/auth/role/delete','0','','','','',0,'false'),('ea4b0eda-2b43-11e7-9c7e-a0c58951c8d5','1003','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4,'false'),('eb13f0e9-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010100','/v1/auth/user/get','0','','tile','','',0,'false'),('ec5cb33a-07ec-11e7-952f-a0c58951c8d5','1001','0105010500','/v1/auth/user/modify/passwd','0','','','','',0,'false'),('ed148f2a-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010200','/v1/auth/user/post','0','','tile','','',0,'false'),('ee765e9a-07e6-11e7-952f-a0c58951c8d5','1001','0103020200','/v1/auth/resource/org/insert','0','','','','',0,'false'),('ef613f0c-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010300','/v1/auth/user/put','0','','tile','','',0,'false'),('f0ccdcc7-4666-11e7-9beb-a0c5895118d5','1001','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-large','2','/static/images/dispatch_icon/etl.png',1,'false'),('f15862dd-7207-11e7-963e-a0c58951c8d5','1002','0103040000','/v1/sso/subsystem/page','0','#339999','tile tile-wide','1','/static/images/ca_icon/dispatch_manage.png',2,'false'),('f19af335-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010400','/v1/auth/user/delete','0','','tile','','',0,'false'),('f2e81083-07d2-11e7-95d9-a0c58951c8d5','1001','0104010100','/v1/auth/domain/get','0','','','','',0,'false'),('f2f35399-7207-11e7-963e-a0c58951c8d5','1003','0103040000','/v1/sso/subsystem/page','0','#339999','tile tile-wide','1','/static/images/ca_icon/dispatch_manage.png',2,'false'),('f3959708-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010500','/v1/auth/user/modify/passwd','0','','tile','','',0,'false'),('f39a8178-71da-11e7-963e-a0c58951c8d5','1001','1200000000','www.asofdate.com','0','#666699','tile tile-wide','1','/static/images/hauth/api.png',2,'true'),('f4e79c0b-7207-11e7-963e-a0c58951c8d5','1004','0103040000','/v1/sso/subsystem/page','0','#339999','tile tile-wide','1','/static/images/ca_icon/dispatch_manage.png',2,'false'),('f5a0999f-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010600','/v1/auth/user/modify/status','0','','tile','','',0,'false'),('f5fe13cd-71da-11e7-963e-a0c58951c8d5','1002','1200000000','www.asofdate.com','0','#666699','tile tile-wide','1','/static/images/hauth/api.png',2,'true'),('f6a6448b-04a3-11e7-9b60-a0c58951c8d5','1001','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1,'false'),('f6f73744-81ca-11e7-9b5b-a0c58951c8d5','1001','0103070000','/v1/auth/resource/service','0','#009966','tile tile-wide','1','/static/images/mdui/mdl-001.png',6,'false'),('f7c96a3c-71da-11e7-963e-a0c58951c8d5','1004','1200000000','www.asofdate.com','0','#666699','tile tile-wide','1','/static/images/hauth/api.png',2,'true'),('f7cdda80-7d98-11e7-97cd-a0c58951c8d5','1005','0201040000','/v1/ca/driver/page','0','#b4d39e','tile','1','/static/images/mdui/mdl-001.png',3,'false'),('f87a9123-0991-11e7-952f-a0c58951c8d5','1001','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1,'false'),('f94b4a93-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020100','/v1/auth/role/get','0','','tile','','',0,'false'),('f9b4686e-81ca-11e7-9b5b-a0c58951c8d5','1002','0103070000','/v1/sso/resource/service','0','#009966','tile tile-wide','1','/static/images/mdui/mdl-001.png',6,'false'),('fa2728fc-4666-11e7-9beb-a0c5895128d5','1002','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-large','2','/static/images/dispatch_icon/etl.png',1,'false'),('fa50fe43-70f1-11e7-963e-a0c58951c8d5','1003','0101010100','/v1/auth/handle/logs','0','','tile','','',0,'false'),('fb975107-07e1-11e7-952f-a0c58951c8d5','1001','0103030300','/v1/auth/domain/share/delete','0','','','','',0,'false'),('fbcd8b0b-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',4,'false'),('fc0531b3-81ca-11e7-9b5b-a0c58951c8d5','1003','0103070000','/v1/sso/resource/service','0','#009966','tile tile-wide','1','/static/images/mdui/mdl-001.png',6,'false'),('fc276e32-4666-11e7-9beb-a0c5895138d5','1003','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-large','2','/static/images/dispatch_icon/etl.png',1,'false'),('fdb44348-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020200','/v1/auth/role/post','0','','tile','','',0,'false'),('fe095c6e-81ca-11e7-9b5b-a0c58951c8d5','1004','0103070000','/v1/sso/resource/service','0','#009966','tile tile-wide','1','/static/images/mdui/mdl-001.png',6,'false'),('fe6968d2-4666-11e7-9beb-a0c5895148d5','1004','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-large','2','/static/images/dispatch_icon/etl.png',1,'false'),('ff313d37-720e-11e7-963e-a0c58951c8d5','1001','0103050000','/v1/sso/subsystem/api/page','0','#ed9f86','tile','1','/static/images/ca_icon/index.png',4,'false'),('ff9f6773-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020300','/v1/auth/role/update','0','','tile','','',0,'false');
/*!40000 ALTER TABLE `sys_theme_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_info`
--

DROP TABLE IF EXISTS `sys_user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_info` (
  `user_id` varchar(30) NOT NULL COMMENT '账号',
  `user_name` varchar(300) DEFAULT NULL COMMENT '账号描述',
  `user_email` varchar(30) DEFAULT NULL COMMENT '用户邮箱',
  `user_phone` decimal(15,0) DEFAULT NULL COMMENT '用户电话',
  `org_unit_id` varchar(66) DEFAULT NULL COMMENT '所属机构/部门',
  `create_user` varchar(30) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `modify_user` varchar(30) DEFAULT NULL COMMENT '修改人',
  `modify_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_user_org_idx` (`org_unit_id`),
  CONSTRAINT `fk_sys_user_org` FOREIGN KEY (`org_unit_id`) REFERENCES `sys_org_info` (`org_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_info`
--

LOCK TABLES `sys_user_info` WRITE;
/*!40000 ALTER TABLE `sys_user_info` DISABLE KEYS */;
INSERT INTO `sys_user_info` VALUES ('admin','超级管理员','hzwy23@sina.com',18107217021,'vertex_root_join_vertex_root','sys','2016-01-01 00:00:00','ccbc_admin','2017-08-19 23:09:42'),('ccbc_admin','中国工商银行超级管理员','hzwy23@163.com',18107217021,'ccbc_join_100','admin','2017-08-19 15:34:45','admin','2017-08-19 15:34:45');
/*!40000 ALTER TABLE `sys_user_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_status_attr`
--

DROP TABLE IF EXISTS `sys_user_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_status_attr` (
  `status_id` char(1) NOT NULL,
  `status_desc` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_status_attr`
--

LOCK TABLES `sys_user_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_user_status_attr` DISABLE KEYS */;
INSERT INTO `sys_user_status_attr` VALUES ('0','正常'),('1','失效');
/*!40000 ALTER TABLE `sys_user_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_theme`
--

DROP TABLE IF EXISTS `sys_user_theme`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_theme` (
  `user_id` varchar(30) NOT NULL,
  `theme_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `pk_sys_user_theme_01` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_theme`
--

LOCK TABLES `sys_user_theme` WRITE;
/*!40000 ALTER TABLE `sys_user_theme` DISABLE KEYS */;
INSERT INTO `sys_user_theme` VALUES ('admin','1005'),('ccbc_admin','1005');
/*!40000 ALTER TABLE `sys_user_theme` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `v_sys_privilege_user_domain`
--

DROP TABLE IF EXISTS `v_sys_privilege_user_domain`;
/*!50001 DROP VIEW IF EXISTS `v_sys_privilege_user_domain`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `v_sys_privilege_user_domain` AS SELECT 
 1 AS `user_id`,
 1 AS `privilege_id`,
 1 AS `role_id`,
 1 AS `domain_id`,
 1 AS `permission`,
 1 AS `role_status_id`*/;
SET character_set_client = @saved_cs_client;

--
-- Dumping routines for database 'sso'
--

--
-- Final view structure for view `v_sys_privilege_user_domain`
--

/*!50001 DROP VIEW IF EXISTS `v_sys_privilege_user_domain`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_sys_privilege_user_domain` AS select `u`.`user_id` AS `user_id`,`r`.`privilege_id` AS `privilege_id`,`r`.`role_id` AS `role_id`,`d`.`domain_id` AS `domain_id`,`d`.`permission` AS `permission`,`e`.`role_status_id` AS `role_status_id` from (((`sys_privilege_role` `r` join `sys_privilege_domain` `d` on((`r`.`privilege_id` = `d`.`privilege_id`))) join `sys_role_user` `u` on((`r`.`role_id` = `u`.`role_id`))) join `sys_role_define` `e` on((`r`.`role_id` = `e`.`role_id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-08-23 13:32:20
