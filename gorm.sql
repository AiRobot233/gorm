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

 Date: 15/08/2024 17:51:16
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE,
  INDEX `value`(`value`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO `dictionary` VALUES (1, 0, '单位类型', '单位类型', 0);
INSERT INTO `dictionary` VALUES (2, 1, '厅机关', '厅机关', 0);
INSERT INTO `dictionary` VALUES (3, 1, '司法行政单位', '司法行政单位', 0);
INSERT INTO `dictionary` VALUES (4, 1, '监狱系统', '监狱系统', 0);
INSERT INTO `dictionary` VALUES (5, 1, '戒毒系统', '戒毒系统', 0);

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
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 0, 'root', '*', '2022-07-18 14:06:17', '2022-07-18 14:06:19', NULL, 2);
INSERT INTO `role` VALUES (2, 1, '司法局管理员', '1,2,3,5,6,7,14,18', '2022-07-19 14:17:49', '2022-12-28 17:09:53', NULL, 1);

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
  `method` varchar(25) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '请求类型',
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '按钮权限',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `rule` VALUES (8, 3, '修改', 'api', '/admin/rule/:id', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (9, 3, '删除', 'api', '/admin/rule/:id', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (10, 4, '列表', 'api', '/admin/role', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (11, 4, '添加', 'api', '/admin/role', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (12, 4, '修改', 'api', '/admin/role/:id', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (13, 4, '删除', 'api', '/admin/role/:id', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (14, 5, '列表', 'api', '/admin/user', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (15, 5, '新增', 'api', '/admin/user', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (16, 5, '修改', 'api', '/admin/user/:id', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (17, 5, '删除', 'api', '/admin/user/:id', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (18, 1, '系统设置', 'page', 'system', 0, NULL, NULL);
INSERT INTO `rule` VALUES (19, 18, '字典管理', 'page', 'dictionary', 0, NULL, NULL);
INSERT INTO `rule` VALUES (20, 19, '列表', 'api', '/admin/dictionary', 0, 'GET', NULL);
INSERT INTO `rule` VALUES (21, 19, '新增', 'api', '/admin/dictionary', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (22, 19, '修改', 'api', '/admin/dictionary/:id', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (23, 19, '删除', 'api', '/admin/dictionary/:id', 0, 'DELETE', 'delete');
INSERT INTO `rule` VALUES (24, 18, '单位管理', 'page', 'unit', 0, '', '');
INSERT INTO `rule` VALUES (25, 24, '列表', 'api', '/admin/unit', 0, 'GET', '');
INSERT INTO `rule` VALUES (26, 24, '新增', 'api', '/admin/unit', 0, 'POST', 'create');
INSERT INTO `rule` VALUES (27, 24, '修改', 'api', '/admin/unit/:id', 0, 'PUT', 'update');
INSERT INTO `rule` VALUES (28, 24, '删除', 'api', '/admin/unit/:id', 0, 'DELETE', 'delete');

-- ----------------------------
-- Table structure for unit
-- ----------------------------
DROP TABLE IF EXISTS `unit`;
CREATE TABLE `unit`  (
  `id` int(25) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `type` varchar(30) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '单位类型',
  `pid` int(25) NOT NULL DEFAULT 0 COMMENT '上级id',
  `sort` tinyint(255) UNSIGNED NOT NULL DEFAULT 255,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `is_unit` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否单位 1是2否',
  `is_register` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否注册显示 1是 2否',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 181 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of unit
-- ----------------------------
INSERT INTO `unit` VALUES (2, '江西省监狱管理局', '监狱系统', 151, 255, '2024-05-27 17:36:20', '2024-06-05 11:13:27', NULL, 1, 1);
INSERT INTO `unit` VALUES (5, '江西省司法厅', '厅机关', 0, 0, '2024-05-27 17:38:30', '2024-08-15 17:00:16', NULL, 2, 1);
INSERT INTO `unit` VALUES (6, '南昌市', '', 153, 255, '2024-05-27 17:38:46', '2024-06-05 14:57:09', NULL, 2, 1);
INSERT INTO `unit` VALUES (7, '赣州市', '', 153, 255, '2024-05-27 17:39:23', '2024-06-05 14:57:17', NULL, 2, 1);
INSERT INTO `unit` VALUES (8, '上饶市', '', 153, 255, '2024-05-27 17:39:36', '2024-06-05 14:57:26', NULL, 2, 1);
INSERT INTO `unit` VALUES (9, '九江市', '', 153, 255, '2024-05-27 17:39:44', '2024-06-05 14:57:36', NULL, 2, 1);
INSERT INTO `unit` VALUES (10, '吉安市', '', 153, 255, '2024-05-27 17:39:53', '2024-06-05 14:57:41', NULL, 2, 1);
INSERT INTO `unit` VALUES (11, '抚州市', '', 153, 255, '2024-05-27 17:40:02', '2024-06-05 14:57:46', NULL, 2, 1);
INSERT INTO `unit` VALUES (12, '萍乡市', '', 153, 255, '2024-05-27 17:40:11', '2024-06-05 14:57:50', NULL, 2, 1);
INSERT INTO `unit` VALUES (13, '景德镇市', '', 153, 255, '2024-05-27 17:40:30', '2024-06-05 14:57:54', NULL, 2, 1);
INSERT INTO `unit` VALUES (14, '鹰潭市', '', 153, 255, '2024-05-27 17:40:43', '2024-06-05 14:57:58', NULL, 2, 1);
INSERT INTO `unit` VALUES (15, '宜春市', '', 153, 255, '2024-05-27 17:41:11', '2024-06-05 14:58:02', NULL, 2, 1);
INSERT INTO `unit` VALUES (16, '新余市', '', 153, 255, '2024-05-27 17:41:20', '2024-06-05 14:58:06', NULL, 2, 1);
INSERT INTO `unit` VALUES (18, '江西省戒毒管理局', '戒毒系统', 152, 255, '2024-05-27 17:42:38', '2024-06-05 11:13:52', NULL, 1, 1);
INSERT INTO `unit` VALUES (20, '江西省赣江公证处', '司法行政单位', 153, 0, '2024-05-27 17:43:24', '2024-06-12 09:32:08', NULL, 1, 1);
INSERT INTO `unit` VALUES (22, '南昌市司法局', '司法行政单位', 6, 255, '2024-05-27 17:44:14', '2024-05-27 17:44:14', NULL, 1, 1);
INSERT INTO `unit` VALUES (23, '赣州市司法局', '司法行政单位', 7, 255, '2024-05-27 17:44:19', '2024-05-27 17:44:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (24, '上饶市司法局', '司法行政单位', 8, 255, '2024-05-27 17:44:29', '2024-05-27 17:44:29', NULL, 1, 1);
INSERT INTO `unit` VALUES (25, '九江市司法局', '司法行政单位', 9, 255, '2024-05-27 17:44:37', '2024-05-27 17:44:37', NULL, 1, 1);
INSERT INTO `unit` VALUES (26, '吉安市司法局', '司法行政单位', 10, 255, '2024-05-27 17:44:46', '2024-05-27 17:44:46', NULL, 1, 1);
INSERT INTO `unit` VALUES (27, '抚州市司法局', '司法行政单位', 11, 255, '2024-05-27 17:44:51', '2024-05-27 17:44:51', NULL, 1, 1);
INSERT INTO `unit` VALUES (28, '萍乡市司法局', '司法行政单位', 12, 255, '2024-05-27 17:44:57', '2024-05-27 17:44:57', NULL, 1, 1);
INSERT INTO `unit` VALUES (29, '景德镇市司法局', '司法行政单位', 13, 255, '2024-05-27 17:45:03', '2024-05-27 17:45:03', NULL, 1, 1);
INSERT INTO `unit` VALUES (30, '鹰潭市司法局', '司法行政单位', 14, 255, '2024-05-27 17:45:09', '2024-05-27 17:45:09', NULL, 1, 1);
INSERT INTO `unit` VALUES (31, '宜春市司法局', '司法行政单位', 15, 255, '2024-05-27 17:45:14', '2024-05-27 17:45:14', NULL, 1, 1);
INSERT INTO `unit` VALUES (32, '新余市司法局', '司法行政单位', 16, 255, '2024-05-27 17:45:19', '2024-05-27 17:45:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (33, '南昌市东湖区司法局', '司法行政单位', 6, 255, '2024-05-27 17:45:30', '2024-05-27 17:45:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (34, '南昌市西湖区司法局', '司法行政单位', 6, 255, '2024-05-27 17:45:38', '2024-05-27 17:45:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (35, '南昌市青山湖区司法局', '司法行政单位', 6, 255, '2024-05-27 17:45:46', '2024-05-27 17:45:46', NULL, 1, 1);
INSERT INTO `unit` VALUES (36, '南昌市青云谱区司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:00', '2024-05-27 17:46:00', NULL, 1, 1);
INSERT INTO `unit` VALUES (37, '南昌市新建区司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:09', '2024-05-27 17:46:09', NULL, 1, 1);
INSERT INTO `unit` VALUES (38, '南昌市红谷滩区司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:19', '2024-05-27 17:46:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (39, '南昌市南昌县司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:33', '2024-05-27 17:46:33', NULL, 1, 1);
INSERT INTO `unit` VALUES (40, '南昌市进贤县司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:45', '2024-05-27 17:46:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (41, '南昌市安义县司法局', '司法行政单位', 6, 255, '2024-05-27 17:46:55', '2024-05-27 17:46:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (42, '赣州市章贡区司法局', '司法行政单位', 7, 255, '2024-05-27 17:48:10', '2024-05-27 17:48:10', NULL, 1, 1);
INSERT INTO `unit` VALUES (43, '赣州市南康区司法局', '司法行政单位', 7, 255, '2024-05-27 17:48:47', '2024-05-27 17:48:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (44, '赣州赣县区司法局', '司法行政单位', 7, 255, '2024-05-27 17:49:38', '2024-05-27 17:49:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (45, '赣州市信丰县司法局', '司法行政单位', 7, 255, '2024-05-27 17:50:07', '2024-05-27 17:50:07', NULL, 1, 1);
INSERT INTO `unit` VALUES (46, '赣州市大余县司法局', '司法行政单位', 7, 255, '2024-05-27 17:50:18', '2024-05-27 17:50:18', NULL, 1, 1);
INSERT INTO `unit` VALUES (47, '赣州市上余县司法局', '司法行政单位', 7, 255, '2024-05-27 17:50:36', '2024-05-27 17:50:36', NULL, 1, 1);
INSERT INTO `unit` VALUES (48, '赣州市上犹县司法局', '司法行政单位', 7, 255, '2024-05-27 17:50:43', '2024-05-27 17:50:43', NULL, 1, 1);
INSERT INTO `unit` VALUES (49, '赣州市安远县司法局', '司法行政单位', 7, 255, '2024-05-27 17:50:55', '2024-05-27 18:03:48', NULL, 1, 1);
INSERT INTO `unit` VALUES (50, '上饶市广丰区司法局', '司法行政单位', 8, 255, '2024-05-27 17:51:14', '2024-05-27 17:54:14', NULL, 1, 1);
INSERT INTO `unit` VALUES (51, '赣州市龙南市司法局', '司法行政单位', 7, 255, '2024-05-28 09:17:43', '2024-05-28 09:17:43', NULL, 1, 1);
INSERT INTO `unit` VALUES (52, '赣州市全南县司法局', '司法行政单位', 7, 255, '2024-05-28 09:17:55', '2024-05-28 09:17:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (53, '赣州市定南县司法局', '司法行政单位', 7, 255, '2024-05-28 09:18:06', '2024-05-28 09:18:06', NULL, 1, 1);
INSERT INTO `unit` VALUES (54, '赣州市兴国县司法局', '司法行政单位', 7, 255, '2024-05-28 09:18:15', '2024-05-28 09:18:15', NULL, 1, 1);
INSERT INTO `unit` VALUES (55, '赣州市宁都县司法局', '司法行政单位', 7, 255, '2024-05-28 09:18:30', '2024-05-28 09:18:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (56, '赣州市于都县司法局', '司法行政单位', 7, 255, '2024-05-28 09:18:39', '2024-05-28 09:18:39', NULL, 1, 1);
INSERT INTO `unit` VALUES (57, '赣州市会昌县司法局', '司法行政单位', 7, 255, '2024-05-28 09:18:47', '2024-05-28 09:18:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (58, '赣州市寻乌县司法局', '司法行政单位', 7, 255, '2024-05-28 09:19:05', '2024-05-28 09:19:05', NULL, 1, 1);
INSERT INTO `unit` VALUES (59, '赣州市石城县司法局', '司法行政单位', 7, 255, '2024-05-28 09:19:18', '2024-05-28 09:19:18', NULL, 1, 1);
INSERT INTO `unit` VALUES (60, '赣州市赣州溶江新区司法分局', '司法行政单位', 7, 255, '2024-05-28 09:19:42', '2024-05-28 09:19:42', NULL, 1, 1);
INSERT INTO `unit` VALUES (61, '上饶市广信区司法局', '司法行政单位', 8, 255, '2024-05-28 09:21:24', '2024-05-28 09:21:24', NULL, 1, 1);
INSERT INTO `unit` VALUES (62, '上饶市玉山县司法局', '司法行政单位', 8, 255, '2024-05-28 09:21:35', '2024-05-28 09:21:35', NULL, 1, 1);
INSERT INTO `unit` VALUES (63, '上饶市横峰县司法局', '司法行政单位', 8, 255, '2024-05-28 09:21:43', '2024-05-28 09:21:43', NULL, 1, 1);
INSERT INTO `unit` VALUES (64, '上饶市鄱阳县司法局', '司法行政单位', 8, 255, '2024-05-28 09:21:55', '2024-05-28 09:21:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (65, '上饶市万年县司法局', '司法行政单位', 8, 255, '2024-05-28 09:22:03', '2024-05-28 09:22:03', NULL, 1, 1);
INSERT INTO `unit` VALUES (66, '上饶市婺源县司法局', '司法行政单位', 8, 255, '2024-05-28 09:22:12', '2024-05-28 09:22:12', NULL, 1, 1);
INSERT INTO `unit` VALUES (67, '上饶市德兴市司法局', '司法行政单位', 8, 255, '2024-05-28 09:22:24', '2024-05-28 09:22:24', NULL, 1, 1);
INSERT INTO `unit` VALUES (68, '九江市濂溪区司法局', '司法行政单位', 9, 255, '2024-05-28 09:22:40', '2024-05-28 09:22:40', NULL, 1, 1);
INSERT INTO `unit` VALUES (69, '九江市浔阳区司法局', '司法行政单位', 9, 255, '2024-05-28 09:22:49', '2024-05-28 09:22:49', NULL, 1, 1);
INSERT INTO `unit` VALUES (70, '九江市采桑区司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:01', '2024-05-28 09:23:01', NULL, 1, 1);
INSERT INTO `unit` VALUES (71, '九江市武宁县司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:14', '2024-05-28 09:23:14', NULL, 1, 1);
INSERT INTO `unit` VALUES (72, '九江市修水县司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:22', '2024-05-28 09:23:22', NULL, 1, 1);
INSERT INTO `unit` VALUES (73, '九江市永修县司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:32', '2024-05-28 09:23:32', NULL, 1, 1);
INSERT INTO `unit` VALUES (74, '九江市德安县司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:42', '2024-05-28 09:23:42', NULL, 1, 1);
INSERT INTO `unit` VALUES (75, '九江市都昌县司法局', '司法行政单位', 9, 255, '2024-05-28 09:23:54', '2024-05-28 09:23:54', NULL, 1, 1);
INSERT INTO `unit` VALUES (76, '九江市湖口县司法局', '司法行政单位', 9, 255, '2024-05-28 09:24:01', '2024-05-28 09:24:01', NULL, 1, 1);
INSERT INTO `unit` VALUES (77, '九江市彭泽县司法局', '司法行政单位', 9, 255, '2024-05-28 09:24:11', '2024-05-28 09:24:33', NULL, 1, 1);
INSERT INTO `unit` VALUES (78, '九江市瑞昌市司法局', '司法行政单位', 9, 255, '2024-05-28 09:24:42', '2024-05-28 09:24:42', NULL, 1, 1);
INSERT INTO `unit` VALUES (79, '九江市共青城市司法局', '司法行政单位', 9, 255, '2024-05-28 09:24:57', '2024-05-28 09:24:57', NULL, 1, 1);
INSERT INTO `unit` VALUES (80, '九江市庐山市司法局', '司法行政单位', 9, 255, '2024-05-28 09:25:10', '2024-05-28 09:25:10', NULL, 1, 1);
INSERT INTO `unit` VALUES (81, '吉安市吉州区司法局', '司法行政单位', 10, 255, '2024-05-28 09:25:34', '2024-05-28 09:25:34', NULL, 1, 1);
INSERT INTO `unit` VALUES (82, '吉安市青原区司法局', '司法行政单位', 10, 255, '2024-05-28 09:25:42', '2024-05-28 09:25:42', NULL, 1, 1);
INSERT INTO `unit` VALUES (83, '吉安市吉安县司法局', '司法行政单位', 10, 255, '2024-05-28 09:25:51', '2024-05-28 09:25:51', NULL, 1, 1);
INSERT INTO `unit` VALUES (84, '吉安市吉水县司法局', '司法行政单位', 10, 255, '2024-05-28 09:26:01', '2024-05-28 09:26:01', NULL, 1, 1);
INSERT INTO `unit` VALUES (85, '吉安市峡江县司法局', '司法行政单位', 10, 255, '2024-05-28 09:26:17', '2024-05-28 09:26:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (86, '吉安市遂川县司法局', '司法行政单位', 10, 255, '2024-05-28 09:26:38', '2024-05-28 09:26:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (87, '吉安市新干县司法局', '司法行政单位', 10, 255, '2024-05-28 09:26:44', '2024-05-28 09:26:44', NULL, 1, 1);
INSERT INTO `unit` VALUES (88, '吉安市永丰县司法局', '司法行政单位', 10, 255, '2024-05-28 09:26:52', '2024-05-28 09:26:52', NULL, 1, 1);
INSERT INTO `unit` VALUES (89, '吉安市泰和县司法局', '司法行政单位', 10, 255, '2024-05-28 09:27:04', '2024-05-28 09:27:04', NULL, 1, 1);
INSERT INTO `unit` VALUES (90, '吉安市万安县司法局', '司法行政单位', 10, 255, '2024-05-28 09:27:20', '2024-05-28 09:27:20', NULL, 1, 1);
INSERT INTO `unit` VALUES (91, '吉安市安福县司法局', '司法行政单位', 10, 255, '2024-05-28 09:27:28', '2024-05-28 09:27:28', NULL, 1, 1);
INSERT INTO `unit` VALUES (92, '吉安市永新县司法局', '司法行政单位', 10, 255, '2024-05-28 09:27:35', '2024-05-28 09:27:35', NULL, 1, 1);
INSERT INTO `unit` VALUES (93, '吉安市井冈山市司法局', '司法行政单位', 10, 255, '2024-05-28 09:27:47', '2024-05-28 09:27:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (94, '抚州市临川区司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:05', '2024-05-28 09:28:05', NULL, 1, 1);
INSERT INTO `unit` VALUES (95, '抚州市东乡区司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:15', '2024-05-28 09:28:15', NULL, 1, 1);
INSERT INTO `unit` VALUES (96, '抚州市南城县司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:26', '2024-05-28 09:28:26', NULL, 1, 1);
INSERT INTO `unit` VALUES (97, '抚州市黎川县司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:36', '2024-05-28 09:28:36', NULL, 1, 1);
INSERT INTO `unit` VALUES (98, '抚州市南丰县司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:45', '2024-05-28 09:28:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (99, '抚州市崇仁县司法局', '司法行政单位', 11, 255, '2024-05-28 09:28:53', '2024-05-28 09:28:53', NULL, 1, 1);
INSERT INTO `unit` VALUES (100, '抚州市乐安县司法局', '司法行政单位', 11, 255, '2024-05-28 09:29:01', '2024-05-28 09:29:01', NULL, 1, 1);
INSERT INTO `unit` VALUES (101, '抚州市宜黄县司法局', '司法行政单位', 11, 255, '2024-05-28 09:29:11', '2024-05-28 09:29:11', NULL, 1, 1);
INSERT INTO `unit` VALUES (102, '抚州市金溪县司法局', '司法行政单位', 11, 255, '2024-05-28 09:29:22', '2024-05-28 09:29:22', NULL, 1, 1);
INSERT INTO `unit` VALUES (103, '抚州市资溪县司法局', '司法行政单位', 11, 255, '2024-05-28 09:29:30', '2024-05-28 09:29:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (104, '抚州市广昌县司法局', '司法行政单位', 11, 255, '2024-05-28 09:29:51', '2024-05-28 09:29:51', NULL, 1, 1);
INSERT INTO `unit` VALUES (105, '萍乡市安源区司法局', '司法行政单位', 12, 255, '2024-05-28 09:30:10', '2024-05-28 09:30:10', NULL, 1, 1);
INSERT INTO `unit` VALUES (106, '萍乡市湘东区司法局', '司法行政单位', 12, 255, '2024-05-28 09:30:22', '2024-05-28 09:30:22', NULL, 1, 1);
INSERT INTO `unit` VALUES (107, '萍乡市莲花县司法局', '司法行政单位', 12, 255, '2024-05-28 09:30:30', '2024-05-28 09:30:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (108, '萍乡市上栗县司法局', '司法行政单位', 12, 255, '2024-05-28 09:30:38', '2024-05-28 09:30:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (109, '萍乡市芦溪县司法局', '司法行政单位', 12, 255, '2024-05-28 09:30:48', '2024-05-28 09:30:48', NULL, 1, 1);
INSERT INTO `unit` VALUES (110, '景德镇市昌江区司法局', '司法行政单位', 13, 255, '2024-05-28 09:31:19', '2024-05-28 09:31:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (111, '景德镇市珠山区司法局', '司法行政单位', 13, 255, '2024-05-28 09:31:30', '2024-05-28 09:31:30', NULL, 1, 1);
INSERT INTO `unit` VALUES (112, '景德镇市浮梁县司法局', '司法行政单位', 13, 255, '2024-05-28 09:31:39', '2024-05-28 09:31:39', NULL, 1, 1);
INSERT INTO `unit` VALUES (113, '景德镇市乐平市司法局', '司法行政单位', 13, 255, '2024-05-28 09:31:48', '2024-05-28 09:31:48', NULL, 1, 1);
INSERT INTO `unit` VALUES (114, '鹰潭市月湖区司法局', '司法行政单位', 14, 255, '2024-05-28 09:32:11', '2024-05-28 09:32:11', NULL, 1, 1);
INSERT INTO `unit` VALUES (115, '鹰潭市贵溪市司法局', '司法行政单位', 14, 255, '2024-05-28 09:32:20', '2024-05-28 09:32:20', NULL, 1, 1);
INSERT INTO `unit` VALUES (116, '鹰潭市余江区司法局', '司法行政单位', 14, 255, '2024-05-28 09:32:29', '2024-05-28 09:32:29', NULL, 1, 1);
INSERT INTO `unit` VALUES (117, '宜春市袁州区司法局', '司法行政单位', 15, 255, '2024-05-28 09:32:47', '2024-05-28 09:32:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (118, '宜春市奉新县司法局', '司法行政单位', 15, 255, '2024-05-28 09:32:55', '2024-05-28 09:32:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (119, '宜春市万载县司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:04', '2024-05-28 09:33:04', NULL, 1, 1);
INSERT INTO `unit` VALUES (120, '宜春市上高县司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:11', '2024-05-28 09:33:11', NULL, 1, 1);
INSERT INTO `unit` VALUES (121, '宜春市宜丰县司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:19', '2024-05-28 09:33:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (122, '宜春市靖安县司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:29', '2024-05-28 09:33:29', NULL, 1, 1);
INSERT INTO `unit` VALUES (123, '宜春市铜鼓县司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:37', '2024-05-28 09:33:37', NULL, 1, 1);
INSERT INTO `unit` VALUES (124, '宜春市丰城市司法局', '司法行政单位', 15, 255, '2024-05-28 09:33:45', '2024-05-28 09:33:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (125, '宜春市樟树市司法局', '司法行政单位', 15, 255, '2024-05-28 09:34:31', '2024-05-28 09:34:31', NULL, 1, 1);
INSERT INTO `unit` VALUES (126, '宜春市高安市司法局', '司法行政单位', 15, 255, '2024-05-28 09:34:42', '2024-05-28 09:34:42', NULL, 1, 1);
INSERT INTO `unit` VALUES (127, '新余市渝水区司法局', '司法行政单位', 16, 255, '2024-05-28 09:35:02', '2024-05-28 09:35:02', NULL, 1, 1);
INSERT INTO `unit` VALUES (128, '新余市分宜县司法局', '司法行政单位', 16, 255, '2024-05-28 09:35:09', '2024-05-28 09:35:09', NULL, 1, 1);
INSERT INTO `unit` VALUES (129, '江西省南昌监狱', '监狱系统', 2, 255, '2024-05-28 09:36:55', '2024-05-28 09:36:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (130, '江西省赣州监狱', '监狱系统', 2, 255, '2024-05-28 09:37:07', '2024-05-28 09:37:07', NULL, 1, 1);
INSERT INTO `unit` VALUES (131, '江西省景德镇监狱', '监狱系统', 2, 255, '2024-05-28 09:37:28', '2024-05-28 09:37:28', NULL, 1, 1);
INSERT INTO `unit` VALUES (132, '江西省洪城监狱', '监狱系统', 2, 255, '2024-05-28 09:37:38', '2024-05-28 09:37:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (133, '江西省饶州监狱', '监狱系统', 2, 255, '2024-05-28 09:37:47', '2024-05-28 09:37:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (134, '江西省洪都监狱', '监狱系统', 2, 255, '2024-05-28 09:37:55', '2024-05-28 09:37:55', NULL, 1, 1);
INSERT INTO `unit` VALUES (135, '江西省赣江监狱', '监狱系统', 2, 255, '2024-05-28 09:38:03', '2024-05-28 09:38:03', NULL, 1, 1);
INSERT INTO `unit` VALUES (136, '江西省赣西监狱', '监狱系统', 2, 255, '2024-05-28 09:38:12', '2024-05-28 09:38:12', NULL, 1, 1);
INSERT INTO `unit` VALUES (137, '江西省豫章监狱', '监狱系统', 2, 255, '2024-05-28 09:38:20', '2024-05-28 09:38:20', NULL, 1, 1);
INSERT INTO `unit` VALUES (138, '江西省温圳监狱', '监狱系统', 2, 255, '2024-05-28 09:38:29', '2024-05-28 09:38:29', NULL, 1, 1);
INSERT INTO `unit` VALUES (139, '江西省女子监狱', '监狱系统', 2, 255, '2024-05-28 09:38:38', '2024-05-28 09:38:38', NULL, 1, 1);
INSERT INTO `unit` VALUES (140, '江西省吉安监狱', '监狱系统', 2, 255, '2024-05-28 09:38:45', '2024-05-28 09:38:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (141, '江西省未管所', '司法行政单位', 2, 255, '2024-05-28 09:38:58', '2024-05-28 09:38:58', NULL, 1, 1);
INSERT INTO `unit` VALUES (142, '江西省新康监狱', '监狱系统', 2, 255, '2024-05-28 09:39:05', '2024-05-28 09:39:05', NULL, 1, 1);
INSERT INTO `unit` VALUES (143, '江西省永桥强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:41:07', '2024-05-28 09:41:07', NULL, 1, 1);
INSERT INTO `unit` VALUES (144, '江西省赣西强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:41:16', '2024-05-28 09:41:16', NULL, 1, 1);
INSERT INTO `unit` VALUES (145, '江西省强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:41:23', '2024-05-28 09:41:23', NULL, 1, 1);
INSERT INTO `unit` VALUES (146, '江西省女子强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:41:45', '2024-05-28 09:41:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (147, '江西省上饶强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:42:01', '2024-05-28 09:42:01', NULL, 1, 1);
INSERT INTO `unit` VALUES (148, '江西省赣南强制戒毒所', '戒毒系统', 18, 255, '2024-05-28 09:42:12', '2024-05-28 09:42:12', NULL, 1, 1);
INSERT INTO `unit` VALUES (149, '江西省病残戒毒人员收治所', '戒毒系统', 18, 255, '2024-05-28 09:42:39', '2024-05-28 09:42:39', NULL, 1, 1);
INSERT INTO `unit` VALUES (151, '监狱系统', '监狱系统', 0, 3, '2024-06-05 11:13:15', '2024-06-11 09:33:23', NULL, 2, 1);
INSERT INTO `unit` VALUES (152, '戒毒系统', '戒毒系统', 0, 4, '2024-06-05 11:13:35', '2024-06-11 09:33:42', NULL, 2, 1);
INSERT INTO `unit` VALUES (153, '司法行政单位', '司法行政单位', 0, 2, '2024-06-05 14:56:37', '2024-06-11 09:33:33', NULL, 2, 1);
INSERT INTO `unit` VALUES (155, '省委全面依法治省委员会办公室秘书处', '厅机关', 5, 255, '2024-06-11 09:35:46', '2024-06-11 09:36:19', NULL, 1, 1);
INSERT INTO `unit` VALUES (156, '办公室', '厅机关', 5, 255, '2024-06-11 09:36:32', '2024-06-11 09:36:37', NULL, 1, 1);
INSERT INTO `unit` VALUES (157, '指挥中心', '厅机关', 5, 255, '2024-06-11 09:36:52', '2024-06-11 09:36:52', NULL, 1, 1);
INSERT INTO `unit` VALUES (158, '人事警务处', '厅机关', 5, 255, '2024-06-11 09:37:08', '2024-06-11 09:37:13', NULL, 1, 1);
INSERT INTO `unit` VALUES (159, '宣传培训处', '厅机关', 5, 255, '2024-06-11 09:37:26', '2024-06-11 09:37:26', NULL, 1, 1);
INSERT INTO `unit` VALUES (160, '法治调研处', '厅机关', 5, 255, '2024-06-11 09:37:36', '2024-06-11 09:37:45', NULL, 1, 1);
INSERT INTO `unit` VALUES (161, '法治督察处', '厅机关', 5, 255, '2024-06-11 09:37:58', '2024-06-11 09:37:58', NULL, 1, 1);
INSERT INTO `unit` VALUES (162, '立法一处', '厅机关', 5, 255, '2024-06-11 09:38:10', '2024-06-11 09:38:10', NULL, 1, 1);
INSERT INTO `unit` VALUES (163, '立法二处', '厅机关', 5, 255, '2024-06-11 09:38:18', '2024-06-11 09:38:18', NULL, 1, 1);
INSERT INTO `unit` VALUES (164, '合法性审查处', '厅机关', 5, 255, '2024-06-11 09:38:32', '2024-06-11 09:38:32', NULL, 1, 1);
INSERT INTO `unit` VALUES (165, '社区矫正工作处', '厅机关', 5, 255, '2024-06-11 09:38:43', '2024-06-11 09:38:56', NULL, 1, 1);
INSERT INTO `unit` VALUES (166, '监所工作处', '厅机关', 5, 255, '2024-06-11 09:39:07', '2024-06-11 09:39:07', NULL, 1, 1);
INSERT INTO `unit` VALUES (167, '行政复议审理处', '厅机关', 5, 255, '2024-06-11 09:39:21', '2024-06-11 09:39:21', NULL, 1, 1);
INSERT INTO `unit` VALUES (168, '行政复议受理与应诉处', '厅机关', 5, 255, '2024-06-11 09:39:40', '2024-06-11 09:39:40', NULL, 1, 1);
INSERT INTO `unit` VALUES (169, '行政执法协调监督处', '厅机关', 5, 255, '2024-06-11 09:40:02', '2024-06-11 09:40:02', NULL, 1, 1);
INSERT INTO `unit` VALUES (170, '普法与依法治理处', '厅机关', 5, 255, '2024-06-11 09:40:25', '2024-06-11 09:40:25', NULL, 1, 1);
INSERT INTO `unit` VALUES (171, '人民参与和促进法治处', '厅机关', 5, 255, '2024-06-11 09:40:47', '2024-06-11 09:40:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (172, '公共法律服务管理处', '厅机关', 5, 255, '2024-06-11 09:41:04', '2024-06-11 09:41:04', NULL, 1, 1);
INSERT INTO `unit` VALUES (173, '律师工作处', '厅机关', 5, 255, '2024-06-11 09:41:16', '2024-06-11 09:41:16', NULL, 1, 1);
INSERT INTO `unit` VALUES (174, '公证仲裁工作处', '厅机关', 5, 255, '2024-06-11 09:41:34', '2024-06-11 09:41:34', NULL, 1, 1);
INSERT INTO `unit` VALUES (175, '司法鉴定工作处', '厅机关', 5, 255, '2024-06-11 09:41:47', '2024-06-11 09:41:47', NULL, 1, 1);
INSERT INTO `unit` VALUES (176, '法律职业资格管理处', '厅机关', 5, 255, '2024-06-11 09:42:08', '2024-06-11 09:42:08', NULL, 1, 1);
INSERT INTO `unit` VALUES (177, '装备财务保障处', '厅机关', 5, 255, '2024-06-11 09:42:24', '2024-06-11 09:42:24', NULL, 1, 1);
INSERT INTO `unit` VALUES (178, '直属机关党委', '厅机关', 5, 255, '2024-06-11 09:42:41', '2024-06-11 09:42:41', NULL, 1, 1);
INSERT INTO `unit` VALUES (179, '江西司法警官学院', '厅机关', 153, 1, '2024-06-12 09:59:58', '2024-06-13 17:25:46', NULL, 1, 2);
INSERT INTO `unit` VALUES (180, '省监所安全服务保障中心', '厅机关', 153, 2, '2024-06-12 10:03:47', '2024-06-13 17:25:49', NULL, 1, 2);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(30) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '名称',
  `phone` char(11) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '电话',
  `password` varchar(150) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(5) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '密码盐',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1正常2禁用',
  `role_id` int(25) NOT NULL COMMENT '对应角色id',
  `unit_id` int(25) NOT NULL COMMENT '单位id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE,
  INDEX `unit_id`(`unit_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'root', '13979123001', '084f64cf4424d8ce530fee60109476a0', '13f88', '2022-07-15 15:36:19', '2022-09-16 14:58:58', NULL, 1, 1, 0);
INSERT INTO `user` VALUES (2, '江西省司法厅', '13333333333', 'e0dc1740242eb862a620701f8e53a198', '031b9', '2024-08-15 17:17:16', '2024-08-15 17:39:41', NULL, 1, 2, 5);

SET FOREIGN_KEY_CHECKS = 1;
