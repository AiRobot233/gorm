/*
 Navicat Premium Data Transfer

 Source Server         : phpstudy127
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : gorm

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 21/03/2023 09:10:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary`  (
  `id` int(25) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(25) NOT NULL,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `value` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '值',
  `sort` int(20) NOT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO `dictionary` VALUES (1, 0, 'mingzu', '民族', 0);
INSERT INTO `dictionary` VALUES (3, 1, 'hanzu', '汉族', 0);
INSERT INTO `dictionary` VALUES (4, 1, 'huizu', '回族', 0);
INSERT INTO `dictionary` VALUES (5, 1, 'zangzu', '藏族', 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(25) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(25) NOT NULL COMMENT '上级id',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `rule` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '权限',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `is_system` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否root账户 1 否 2是',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 0, 'root', '*', '2022-07-18 14:06:17', '2022-07-18 14:06:19', NULL, 2);
INSERT INTO `role` VALUES (2, 1, '系统管理员1', '1,2,3,5,6,7,14,18', '2022-07-19 14:17:49', '2022-12-28 17:09:53', NULL, 1);
INSERT INTO `role` VALUES (5, 1, '系统管理员2', '2,1', '2022-08-12 10:26:35', '2022-08-12 10:26:35', NULL, 1);

-- ----------------------------
-- Table structure for rule
-- ----------------------------
DROP TABLE IF EXISTS `rule`;
CREATE TABLE `rule`  (
  `id` int(25) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(25) NOT NULL COMMENT '上级id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '规则名称',
  `type` enum('page','api') CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '类型',
  `router` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '地址/路由',
  `sort` int(25) NOT NULL DEFAULT 0 COMMENT '排序',
  `method` enum('GET','POST','PUT','DELETE') CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '请求类型',
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '标识 按钮权限',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rule
-- ----------------------------
INSERT INTO `rule` VALUES (1, 0, 'Root', 'page', 'root', 255, NULL, NULL);
INSERT INTO `rule` VALUES (2, 1, '首页', 'page', 'home', 254, NULL, NULL);
INSERT INTO `rule` VALUES (3, 18, '规则配置', 'page', 'rule', 0, NULL, NULL);
INSERT INTO `rule` VALUES (4, 18, '角色配置', 'page', 'role', 0, NULL, NULL);
INSERT INTO `rule` VALUES (5, 18, '用户管理', 'page', 'user', 0, NULL, NULL);
INSERT INTO `rule` VALUES (6, 3, '列表', 'api', '/admin/rule', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (7, 3, '添加', 'api', '/admin/rule', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (8, 3, '修改', 'api', '/admin/rule', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (9, 3, '删除', 'api', '/admin/rule', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (10, 4, '列表', 'api', '/admin/role', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (11, 4, '添加', 'api', '/admin/role', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (12, 4, '修改', 'api', '/admin/role', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (13, 4, '删除', 'api', '/admin/role', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (14, 5, '列表', 'api', '/admin/user', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (15, 5, '新增', 'api', '/admin/user', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (16, 5, '修改', 'api', '/admin/user', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (17, 5, '删除', 'api', '/admin/user', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (18, 1, '系统设置', 'page', 'system', 0, NULL, NULL);
INSERT INTO `rule` VALUES (19, 18, '字典管理', 'page', 'dictionary', 0, NULL, NULL);
INSERT INTO `rule` VALUES (20, 19, '列表', 'api', '/admin/dictionary', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (21, 19, '新增', 'api', '/admin/dictionary', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (22, 19, '修改', 'api', '/admin/dictionary', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (23, 19, '删除', 'api', '/admin/dictionary', 0, 'DELETE', 'delete');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(30) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `phone` char(11) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '电话',
  `password` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(5) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '密码盐',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1正常2禁用',
  `role_id` int(25) NOT NULL COMMENT '对应角色id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'root', '13979123001', '084f64cf4424d8ce530fee60109476a0', '13f88', '2022-07-15 15:36:19', '2022-09-16 14:58:58', NULL, 1, 1);
INSERT INTO `user` VALUES (4, '王五', '13979123002', '4f753121769612c4110b52aab55bb452', '4b9ce', '2022-07-15 17:03:36', '2022-09-16 15:01:39', NULL, 1, 2);
INSERT INTO `user` VALUES (5, '李四', '13979123005', '7ea8300d09f8a10a463652fe341fc7ea', '918dd', '2022-07-15 17:20:38', '2022-07-15 17:20:38', NULL, 1, 2);
INSERT INTO `user` VALUES (9, '赵六', '13979123006', '4bb0e1f7613f9408f98f44d86f57d63c', '59b73', '2022-07-20 10:36:57', '2022-07-20 10:36:57', NULL, 2, 2);
INSERT INTO `user` VALUES (10, '赵六', '13979123006', '9792581e26b1cbfac5fa1950bf31a48d', '71f5b', '2023-02-24 08:00:05', '2023-02-24 08:49:06', '2023-02-24 08:49:06', 2, 2);

SET FOREIGN_KEY_CHECKS = 1;
