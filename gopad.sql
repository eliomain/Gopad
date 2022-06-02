/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-06-02 13:20:40
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for `go_options`
-- ----------------------------
DROP TABLE IF EXISTS `go_options`;
CREATE TABLE `go_options` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `val` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_options
-- ----------------------------
INSERT INTO `go_options` VALUES ('1', 'webtitle', 'websiteName');
INSERT INTO `go_options` VALUES ('2', 'companyname', 'company');
INSERT INTO `go_options` VALUES ('3', 'siteurl', 'http://127.0.0.1');
INSERT INTO `go_options` VALUES ('4', 'keywords', 'keywords');
INSERT INTO `go_options` VALUES ('5', 'description', 'description');
INSERT INTO `go_options` VALUES ('6', 'admin_email', '123@qq.com');
INSERT INTO `go_options` VALUES ('7', 'language', '1');

-- ----------------------------
-- Table structure for `go_postcate`
-- ----------------------------
DROP TABLE IF EXISTS `go_postcate`;
CREATE TABLE `go_postcate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `pid` int(10) unsigned NOT NULL DEFAULT '0',
  `sort` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_postcate
-- ----------------------------
INSERT INTO `go_postcate` VALUES ('1', '测试分类', '0', '0');

-- ----------------------------
-- Table structure for `go_posts`
-- ----------------------------
DROP TABLE IF EXISTS `go_posts`;
CREATE TABLE `go_posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL DEFAULT '1',
  `title` text NOT NULL,
  `content` longtext NOT NULL,
  `thumbnail` varchar(128) NOT NULL DEFAULT '',
  `description` varchar(521) NOT NULL DEFAULT '',
  `url2` varchar(256) NOT NULL DEFAULT '#',
  `author` int(11) NOT NULL DEFAULT '1',
  `time` varchar(16) NOT NULL DEFAULT '0',
  `ctime` varchar(16) NOT NULL DEFAULT '0',
  `empty` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_posts
-- ----------------------------
INSERT INTO `go_posts` VALUES ('1', '0', 'hello world', '<p>hello world</p>', '20220602123709.png', '', '', '1', '1654144629', '0', '0');

-- ----------------------------
-- Table structure for `go_users`
-- ----------------------------
DROP TABLE IF EXISTS `go_users`;
CREATE TABLE `go_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(64) NOT NULL DEFAULT '',
  `name` varchar(64) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `login_ip` varchar(32) NOT NULL DEFAULT '',
  `register_time` int(11) NOT NULL DEFAULT '0',
  `login_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_users
-- ----------------------------
INSERT INTO `go_users` VALUES ('1', 'root', '21232f297a57a5a743894a0e4a801fc3', '', 'admin@email.com', '', '0', '0');
INSERT INTO `go_users` VALUES ('2', '10086', '21232f297a57a5a743894a0e4a801fc3', '', 'user@email.com', '', '0', '0');
