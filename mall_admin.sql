/*
Navicat MariaDB Data Transfer

Source Server         : 个人服务器
Source Server Version : 100317
Source Host           : 42.51.231.174:3306
Source Database       : golang

Target Server Type    : MariaDB
Target Server Version : 100317
File Encoding         : 65001

Date: 2020-09-04 11:44:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for mall_admin
-- ----------------------------
DROP TABLE IF EXISTS `mall_admin`;
CREATE TABLE `mall_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(30) NOT NULL DEFAULT '' COMMENT '密码盐',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(100) DEFAULT '' COMMENT '电子邮箱',
  `loginfailure` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '失败次数',
  `logintime` int(10) DEFAULT NULL COMMENT '登录时间',
  `loginip` varchar(50) DEFAULT NULL COMMENT '登录IP',
  `createtime` int(10) DEFAULT NULL COMMENT '创建时间',
  `updatetime` int(10) DEFAULT NULL COMMENT '更新时间',
  `token` varchar(59) NOT NULL DEFAULT '' COMMENT 'Session标识',
  `status` varchar(30) NOT NULL DEFAULT 'normal' COMMENT '状态',
  `user_id` int(11) unsigned DEFAULT NULL COMMENT '用户id',
  `company_id` int(11) unsigned DEFAULT NULL,
  `mall_id` int(11) unsigned DEFAULT NULL,
  `shop_id` int(10) DEFAULT NULL,
  `phone` char(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Records of mall_admin
-- ----------------------------
INSERT INTO `mall_admin` VALUES ('51', 'wuwu', 'wuwu', 'ff528bfabb4ccac9b94945ee2d892960', 'aadsa', '/assets/img/avatar.png', '', '0', '1599012858', '183.53.190.236', '1597374156', '1599012874', '', 'normal', null, null, null, null, '18818653680');
INSERT INTO `mall_admin` VALUES ('52', 'yubo', 'yubo', '78de6b8dec4d2e71a692f6aabb8bb406', 'sadfs', '/assets/img/avatar.png', '1231231@qq.com', '0', '1599124885', '223.73.147.78', '1598233595', '1599124885', '678f6cea-0c84-45f3-8ec8-92d209ed9be7', 'normal', null, null, null, null, null);
