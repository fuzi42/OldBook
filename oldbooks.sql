/*
 Navicat Premium Data Transfer

 Source Server         : article
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost:3306
 Source Schema         : oldbooks

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 13/12/2019 06:27:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book_lists
-- ----------------------------
DROP TABLE IF EXISTS `book_lists`;
CREATE TABLE `book_lists`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `messages` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `owner_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `kind` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `images` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of book_lists
-- ----------------------------
INSERT INTO `book_lists` VALUES ('1573889752', 'C语言', 'C语言是世界上最好的语言', 100.00, '12.46.07pm', '计算机IT', 'O1CN01cWbyBJ2GS0HjnutV6_!!4263749013.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1573899257', '历代文论', '历代文论', 50.00, '1573899090', '文学', 'O1CN01fqkR1D1VdC0wTjXoh_!!0-item_pic.jpg_430x430q90.jpg');
INSERT INTO `book_lists` VALUES ('1573899330', '环境微生物学', '五成新微生学', 30.00, '1573899090', '生物', 'O1CN01CgjSyc1q8bS2UqxIV_!!0-item_pic.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1573899368', '大物理实验书', '九成新', 20.00, '1573899090', '物理', 'O1CN01f1GjNE1Y5FWhCABro_!!0-item_pic.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1573901741', '计算机操作系统', '计算机学操作找欧德', 50.00, '1573899090', '计算机IT', 'O1CN01bsTPiI1q8bVT6URdM_!!2-item_pic.png');
INSERT INTO `book_lists` VALUES ('1573901796', '普通生物化学', '统统58，只要58，全部带回家', 58.00, '1573899090', '生物', 'O1CN01McZpbJ1KRJzUseu3M_!!0-item_pic.jpg');
INSERT INTO `book_lists` VALUES ('1573901843', '大学应用物理', '学不懂的大学应用物理九成新', 38.00, '1573899090', '物理', 'O1CN018uQyrW1q8bUxSBqTK_!!2-item_pic.png_400x400.jpg.png');
INSERT INTO `book_lists` VALUES ('1573901890', '可用MySQL', '从建库到删库跑路', 35.00, '1573899090', '计算机IT', 'O1CN01WpQwLs1NjacsEqGN5_!!2-item_pic.png_400x400.jpg.png');
INSERT INTO `book_lists` VALUES ('1574059816', '玩转Linux', '玩电脑就玩Linux', 30.00, '12.46.07pm', '计算机IT', 'O1CN01fX9Ucj27pW66ZxfI7_!!0-item_pic.jpg');
INSERT INTO `book_lists` VALUES ('1574126674', 'HelloWorld！！！', '354146', 90.00, '12.46.07pm', '计算机IT', 'O1CN01Tc002w1VfwOw8jxgA_!!2-item_pic.png');
INSERT INTO `book_lists` VALUES ('1574938730', '好吃生物学', '好吃的生物学！只要88，统统带回家', 88.00, '1573899124', '生物', 'TB2hBL2AylnpuFjSZFgXXbi7FXa_!!3319364627.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1574938764', '高代和答案', '高代和答案', 99.00, '1573899124', '数学', 'O1CN01GhczWF1HrwAkTfFIY_!!2-item_pic.png_400x400.jpg.png');
INSERT INTO `book_lists` VALUES ('1575578015', '超级物理学词典', '遇事不决，量子力学！', 22.00, '1573899090', '文学', 'O1CN014eAZRn2K28OguxtVD_!!0-item_pic.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1575578094', '好英语书', '只要998，统统带回家', 98.00, '1573899090', '文学', 'O1CN01AoE8091GKDLjD2uVB_!!1993220603.jpg_400x400.jpg');
INSERT INTO `book_lists` VALUES ('1575942209', 'hnjk', 'jkl', 88.00, '12.46.07pm', '文学', 'O1CN01cb9zrJ228OiUWpKZG_!!4015457075.jpg_400x400.jpg');

-- ----------------------------
-- Table structure for order_lists
-- ----------------------------
DROP TABLE IF EXISTS `order_lists`;
CREATE TABLE `order_lists`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `buyer_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `seller_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `boss` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `collects` int(255) NULL DEFAULT NULL,
  `follows` int(255) NULL DEFAULT NULL,
  `userimage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES ('12.46.07pm', 'test', '123456', NULL, NULL, NULL, 'images/demo.jpg');
INSERT INTO `userinfo` VALUES ('12.48.50pm', '是的发送到', 'sss', NULL, NULL, NULL, NULL);
INSERT INTO `userinfo` VALUES ('12.49.44pm', '是', '是', NULL, NULL, NULL, NULL);
INSERT INTO `userinfo` VALUES ('1573899090', 'demo1', '123456', NULL, NULL, NULL, 'images/demo.jpg');
INSERT INTO `userinfo` VALUES ('1573899124', 'demo', '123456', NULL, NULL, NULL, 'images/demo.jpg');

SET FOREIGN_KEY_CHECKS = 1;
