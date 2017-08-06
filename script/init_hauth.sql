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
-- Table structure for table `sso_local_static_route`
--

DROP TABLE IF EXISTS `sso_local_static_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sso_local_static_route` (
  `uuid` varchar(60) NOT NULL,
  `url` varchar(200) DEFAULT NULL COMMENT '本地服务路由前缀',
  `path` varchar(200) DEFAULT NULL COMMENT '本地静态文件路径',
  `create_user` varchar(30) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `modify_user` varchar(30) DEFAULT NULL,
  `modify_time` datetime DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='静态资源管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sso_proxy_static_route`
--

DROP TABLE IF EXISTS `sso_proxy_static_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sso_proxy_static_route` (
  `uuid` varchar(60) NOT NULL,
  `register_url` varchar(200) DEFAULT NULL COMMENT 'api接口',
  `route_desc` varchar(100) DEFAULT NULL COMMENT '服务名称',
  `service_cd` varchar(45) DEFAULT NULL,
  `remote_url` varchar(200) DEFAULT NULL COMMENT '目标系统url',
  `create_user` varchar(45) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `modify_user` varchar(45) DEFAULT NULL,
  `modify_time` datetime DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sso_proxy_static001_idx` (`service_cd`),
  CONSTRAINT `fk_sso_proxy_static001` FOREIGN KEY (`service_cd`) REFERENCES `sso_remote_service` (`service_cd`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sso_remote_service`
--

DROP TABLE IF EXISTS `sso_remote_service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sso_remote_service` (
  `service_cd` varchar(60) NOT NULL COMMENT '子系统编码',
  `remote_scheme` varchar(5) DEFAULT NULL COMMENT '协议名称，如http,https',
  `remote_host` varchar(100) DEFAULT NULL COMMENT '子系统主机名',
  `remote_port` varchar(5) DEFAULT NULL COMMENT '端口号',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `modify_time` datetime DEFAULT NULL COMMENT '修改日期',
  `create_user` varchar(45) DEFAULT NULL COMMENT '创建人',
  `modify_user` varchar(45) DEFAULT NULL COMMENT '最近修改人',
  `service_name` varchar(200) DEFAULT NULL COMMENT '服务名称',
  `prefix_url` varchar(200) DEFAULT NULL COMMENT '表示这个子系统中所有服务的前缀',
  PRIMARY KEY (`service_cd`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统注册';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sso_route_register`
--

DROP TABLE IF EXISTS `sso_route_register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sso_route_register` (
  `uuid` varchar(60) NOT NULL COMMENT '系统唯一标识',
  `service_cd` varchar(60) DEFAULT NULL COMMENT '远程服务器编码',
  `register_url` varchar(200) DEFAULT NULL COMMENT 'api接口',
  `route_desc` varchar(100) DEFAULT NULL COMMENT '服务名称',
  `remote_url` varchar(200) DEFAULT NULL COMMENT '目标系统url',
  `create_user` varchar(45) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `modify_user` varchar(30) DEFAULT NULL COMMENT '最近修改人',
  `modify_time` datetime DEFAULT NULL COMMENT '最近修改时间',
  PRIMARY KEY (`uuid`),
  KEY `fk_sso_route_reg001_idx` (`service_cd`),
  CONSTRAINT `fk_sso_route_reg001` FOREIGN KEY (`service_cd`) REFERENCES `sso_remote_service` (`service_cd`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='api注册';
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Table structure for table `sys_domain_share_info`
--

DROP TABLE IF EXISTS `sys_domain_share_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_share_info` (
  `uuid` varchar(66) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `target_domain_id` varchar(30) NOT NULL,
  `Authorization_level` char(1) NOT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_date` date DEFAULT NULL,
  `modify_date` date DEFAULT NULL,
  `modify_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_domain_share_info_01_idx` (`domain_id`),
  CONSTRAINT `fk_sys_domain_share_info_01` FOREIGN KEY (`domain_id`) REFERENCES `sys_domain_define` (`domain_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `domain_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Table structure for table `sys_org_info`
--

DROP TABLE IF EXISTS `sys_org_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_org_info` (
  `org_unit_id` varchar(66) NOT NULL,
  `org_unit_desc` varchar(300) NOT NULL,
  `up_org_id` varchar(66) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `create_date` date NOT NULL,
  `maintance_date` date NOT NULL,
  `create_user` varchar(30) NOT NULL,
  `maintance_user` varchar(30) NOT NULL,
  `code_number` varchar(66) NOT NULL,
  PRIMARY KEY (`org_unit_id`),
  KEY `pk_sys_org_info_03_idx` (`domain_id`),
  CONSTRAINT `fk_sys_org_info_01` FOREIGN KEY (`domain_id`) REFERENCES `sys_domain_define` (`domain_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sys_resource_info`
--

DROP TABLE IF EXISTS `sys_resource_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_info` (
  `res_id` varchar(30) NOT NULL COMMENT '菜单资源编码',
  `res_name` varchar(300) DEFAULT NULL,
  `res_attr` char(1) DEFAULT NULL,
  `res_up_id` varchar(30) DEFAULT NULL,
  `res_type` char(1) DEFAULT NULL,
  `sys_flag` char(1) DEFAULT NULL COMMENT '0：系统内置菜单',
  `inner_flag` varchar(5) DEFAULT NULL COMMENT 'true：内部路由\nfalse：外部路由',
  `service_cd` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`res_id`),
  KEY `fk_sys_idx_13` (`res_type`),
  KEY `fk_sys_idx_14` (`res_attr`),
  CONSTRAINT `fk_sys_idx_13` FOREIGN KEY (`res_type`) REFERENCES `sys_resource_type_attr` (`res_type`),
  CONSTRAINT `fk_sys_idx_14` FOREIGN KEY (`res_attr`) REFERENCES `sys_resource_info_attr` (`res_attr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Table structure for table `sys_role_define`
--

DROP TABLE IF EXISTS `sys_role_define`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_define` (
  `role_id` varchar(66) NOT NULL,
  `role_name` varchar(300) NOT NULL,
  `role_owner` varchar(30) NOT NULL,
  `role_create_date` datetime NOT NULL,
  `role_status_id` char(1) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `role_maintance_date` datetime NOT NULL,
  `role_maintance_user` varchar(30) NOT NULL,
  `code_number` varchar(66) NOT NULL,
  PRIMARY KEY (`role_id`),
  KEY `fk_sys_idx_11` (`role_status_id`),
  CONSTRAINT `fk_sys_idx_11` FOREIGN KEY (`role_status_id`) REFERENCES `sys_role_status_attr` (`role_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Table structure for table `sys_theme_resource`
--

DROP TABLE IF EXISTS `sys_theme_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_theme_resource` (
  `uuid` varchar(60) NOT NULL,
  `theme_id` varchar(30) DEFAULT NULL,
  `res_id` varchar(30) DEFAULT NULL,
  `res_url` varchar(120) DEFAULT NULL,
  `res_type` varchar(5) DEFAULT NULL,
  `res_bg_color` varchar(30) DEFAULT NULL,
  `res_class` varchar(90) DEFAULT NULL,
  `group_id` char(1) DEFAULT NULL,
  `res_img` varchar(200) DEFAULT NULL,
  `sort_id` decimal(10,0) DEFAULT NULL,
  `new_iframe` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `pk_sys_theme_value_01` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sys_user_info`
--

DROP TABLE IF EXISTS `sys_user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_info` (
  `user_id` varchar(30) NOT NULL,
  `user_name` varchar(300) DEFAULT NULL,
  `user_email` varchar(30) DEFAULT NULL,
  `user_phone` decimal(15,0) DEFAULT NULL,
  `org_unit_id` varchar(66) DEFAULT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `modify_user` varchar(30) DEFAULT NULL,
  `modify_time` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_user_org_idx` (`org_unit_id`),
  CONSTRAINT `fk_sys_user_org` FOREIGN KEY (`org_unit_id`) REFERENCES `sys_org_info` (`org_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Dumping routines for database 'sso'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-08-06 18:57:28
