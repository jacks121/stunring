-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.32 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  12.2.0.6576
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 swetelove 的数据库结构
DROP DATABASE IF EXISTS `swetelove`;
CREATE DATABASE IF NOT EXISTS `swetelove` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `swetelove`;

-- 导出  表 swetelove.advertisements 结构
DROP TABLE IF EXISTS `advertisements`;
CREATE TABLE IF NOT EXISTS `advertisements` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  swetelove.advertisements 的数据：~0 rows (大约)
DELETE FROM `advertisements`;
INSERT INTO `advertisements` (`id`, `code`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'banner', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(2, 'category_banner', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL);

-- 导出  表 swetelove.attributes 结构
DROP TABLE IF EXISTS `attributes`;
CREATE TABLE IF NOT EXISTS `attributes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `attribute_name` varchar(255) DEFAULT NULL COMMENT '属性名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性表';

-- 正在导出表  swetelove.attributes 的数据：~3 rows (大约)
DELETE FROM `attributes`;
INSERT INTO `attributes` (`id`, `attribute_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(4, 'Color', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL),
	(5, 'Style', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL),
	(6, 'Size', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL);

-- 导出  表 swetelove.attribute_values 结构
DROP TABLE IF EXISTS `attribute_values`;
CREATE TABLE IF NOT EXISTS `attribute_values` (
  `id` int NOT NULL AUTO_INCREMENT,
  `attribute_id` int NOT NULL,
  `value` varchar(255) DEFAULT NULL COMMENT '属性值',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性值表';

-- 正在导出表  swetelove.attribute_values 的数据：~6 rows (大约)
DELETE FROM `attribute_values`;
INSERT INTO `attribute_values` (`id`, `attribute_id`, `value`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(9, 4, 'Red', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
	(10, 4, 'Blue', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
	(11, 5, 'haha', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
	(12, 5, 'heihei', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
	(13, 6, 'Small', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
	(14, 6, 'Large', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL);

-- 导出  表 swetelove.categories 结构
DROP TABLE IF EXISTS `categories`;
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT '0',
  `category_name` varchar(255) DEFAULT NULL COMMENT '分类名称',
  `url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类表';

-- 正在导出表  swetelove.categories 的数据：~31 rows (大约)
DELETE FROM `categories`;
INSERT INTO `categories` (`id`, `parent_id`, `category_name`, `url`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 0, 'Wedding', '/wedding.html', NULL, NULL, NULL),
	(2, 1, 'Engagement Rings', '/wedding/engagement-rings.html', NULL, NULL, NULL),
	(3, 1, 'Wedding Band', '/wedding/wedding-band.html', NULL, NULL, NULL),
	(4, 1, 'Wedding Set', '/wedding/wedding-set.html', NULL, NULL, NULL),
	(5, 1, 'Wedding Rings', '/wedding/wedding-rings.html', NULL, NULL, NULL),
	(6, 1, 'Jewelry Set', '/wedding/jewelry-set.html', NULL, NULL, NULL),
	(7, 1, 'Gift Packaging', '/wedding/gift-packaging.html', NULL, NULL, NULL),
	(8, 0, 'Earrings', '/earrings.html', NULL, NULL, NULL),
	(9, 8, 'Studs', '/earrings/studs.html', NULL, NULL, NULL),
	(10, 8, 'Drops', '/earrings/drops.html', NULL, NULL, NULL),
	(11, 8, 'Hoops', '/earrings/hoops.html', NULL, NULL, NULL),
	(12, 0, 'Necklaces', '/necklaces.html', NULL, NULL, NULL),
	(13, 12, 'Pendants', '/necklaces/pendants.html', NULL, NULL, NULL),
	(14, 12, 'Chokers & Tennis Necklace', '/necklaces/chokers-tennis-necklace.html', NULL, NULL, NULL),
	(15, 12, 'Lariat & Y Necklace', '/necklaces/lariat-y-necklace.html', NULL, NULL, NULL),
	(16, 12, 'Layered Necklace', '/necklaces/layered-necklace.html', NULL, NULL, NULL),
	(17, 12, 'Pear Necklace', '/necklaces/statement-necklace.html', NULL, NULL, NULL),
	(18, 0, 'Bracelets', '/bracelets.html', NULL, NULL, NULL),
	(19, 0, 'Men\'s', '/men.html', NULL, NULL, NULL),
	(20, 19, 'Men\'s Rings', '/men/men-s-rings.html', NULL, NULL, NULL),
	(21, 19, 'Men\'s Necklaces', '/men/men-s-necklaces.html', NULL, NULL, NULL),
	(22, 19, 'Cufflinks', '/men/cufflinks.html', NULL, NULL, NULL),
	(23, 0, 'COLLECTION', '/collection.html', NULL, NULL, NULL),
	(24, 23, 'Back to School', '/collection/back-to-school.html', NULL, NULL, NULL),
	(25, 23, 'The Spirit Snake', '/collection/the-spirit-snake.html', NULL, NULL, NULL),
	(26, 23, 'The Vintage Art Deco', '/collection/the-vintage-art-deco-ring.html', NULL, NULL, NULL),
	(27, 23, 'The Snow Band Ring', '/collection/the-snow-band-ring.html', NULL, NULL, NULL),
	(28, 23, 'The Starry Night', '/collection/the-starry-night.html', NULL, NULL, NULL),
	(29, 23, 'Mother\'s Day', '/collection/mother-s-day.html', NULL, NULL, NULL),
	(30, 23, 'Merry Christmas', '/collection/merry-christmas.html', NULL, NULL, NULL),
	(31, 23, 'Father\'s Day', '/collection/father-s-day.html', NULL, NULL, NULL);

-- 导出  表 swetelove.collections 结构
DROP TABLE IF EXISTS `collections`;
CREATE TABLE IF NOT EXISTS `collections` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` int DEFAULT NULL,
  `rule` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  swetelove.collections 的数据：~0 rows (大约)
DELETE FROM `collections`;

-- 导出  表 swetelove.currencies 结构
DROP TABLE IF EXISTS `currencies`;
CREATE TABLE IF NOT EXISTS `currencies` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `exchange` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  swetelove.currencies 的数据：~9 rows (大约)
DELETE FROM `currencies`;
INSERT INTO `currencies` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `exchange`) VALUES
	(1, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'AUD', 5.00),
	(2, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'GBP', 9.00),
	(3, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'CAD', 5.50),
	(4, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'EUR', 8.50),
	(5, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'JPY', 0.07),
	(6, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'MXN', 0.40),
	(7, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'NZD', 4.50),
	(8, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'PHP', 0.16),
	(9, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'SGD', 5.00);

-- 导出  表 swetelove.images 结构
DROP TABLE IF EXISTS `images`;
CREATE TABLE IF NOT EXISTS `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `image_url` varchar(255) DEFAULT NULL COMMENT '图片链接',
  `link` varchar(255) DEFAULT NULL COMMENT '图片链接',
  `imageable_id` int DEFAULT NULL COMMENT '图片关联对象的ID',
  `imageable_type` varchar(255) DEFAULT NULL COMMENT '图片关联对象的类型',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图片表';

-- 正在导出表  swetelove.images 的数据：~18 rows (大约)
DELETE FROM `images`;
INSERT INTO `images` (`id`, `image_url`, `link`, `imageable_id`, `imageable_type`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(7, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_BAND.jpg', NULL, 3, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
	(8, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/ENGAGEMENT_RINGS.jpg', NULL, 2, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
	(9, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_SET.jpg', NULL, 4, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
	(10, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/EARRINGS.jpg', NULL, 8, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
	(11, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/NECKLACES.jpg', NULL, 12, 'categories', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(12, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/BRACELETS.jpg', NULL, 18, 'categories', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(13, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/MEN_S.jpg', NULL, 19, 'categories', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(14, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_RINGS.jpg', NULL, 5, 'categories', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(15, 'http://product1.image', NULL, 8, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(16, 'http://product2.image', NULL, 4, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(17, 'http://product1.image', NULL, 5, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(18, 'http://product2.image', NULL, 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(19, 'http://product1.image', NULL, 7, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(20, 'http://product2.image', NULL, 8, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(21, 'http://product1.image', NULL, 9, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(22, 'http://product2.image', NULL, 9, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(23, 'http://product1.image', NULL, 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(24, 'http://product2.image', NULL, 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
	(25, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/Wedding_banner-pc.jpg', 'https://www.stunring.com/early-black-friday-sale.html', 1, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(26, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/ENGAGEMENT_RINGS.jpg', 'https://www.stunring.com/wedding/engagement-rings.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(27, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_BAND.jpg', 'https://www.stunring.com/wedding/wedding-band.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(28, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_SET.jpg', 'https://www.stunring.com/wedding/wedding-set.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(29, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/EARRINGS.jpg', 'https://www.stunring.com/earrings.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(30, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/NECKLACES.jpg', 'https://www.stunring.com/necklaces.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(31, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/BRACELETS.jpg', 'https://www.stunring.com/bracelets.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(32, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/MEN_S.jpg', 'https://www.stunring.com/men.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
	(33, 'https://cdn.stunring.com/media/wysiwyg/2023_wedding_sale/WEDDING_RINGS.jpg', 'https://www.stunring.com/wedding/wedding-rings.html', 2, 'advertisements', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL);

-- 导出  表 swetelove.products 结构
DROP TABLE IF EXISTS `products`;
CREATE TABLE IF NOT EXISTS `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `original_price` decimal(10,2) DEFAULT NULL COMMENT '原价',
  `current_price` decimal(10,2) DEFAULT NULL COMMENT '现价',
  `on_sale` tinyint(1) DEFAULT NULL COMMENT '是否上架',
  `description` text COMMENT '商品描述',
  `video_url` varchar(255) DEFAULT NULL COMMENT '商品视频链接',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';

-- 正在导出表  swetelove.products 的数据：~21 rows (大约)
DELETE FROM `products`;
INSERT INTO `products` (`id`, `product_name`, `original_price`, `current_price`, `on_sale`, `description`, `video_url`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(4, 'Product 1', 100.00, 90.00, 1, 'This is product 1', 'http://product1.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(5, 'Product 24343', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(6, 'Product 3', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(7, 'Product 4', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(8, 'Product 5', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(9, 'Product 6', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(10, 'Product 7', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(11, 'Product 8', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(12, 'Product 9', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(13, 'Product 10', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(14, 'Product d', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(15, 'Product 32', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(16, 'Product fd4', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(17, 'Product 35', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(18, 'Product jk89', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(19, 'Product d6', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(20, 'Product l;9', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(21, 'Product d5', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(22, 'Product j90', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(23, 'Product ds7', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
	(24, 'Product p9', 110.00, 99.00, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL);

-- 导出  表 swetelove.product_attributes 结构
DROP TABLE IF EXISTS `product_attributes`;
CREATE TABLE IF NOT EXISTS `product_attributes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL COMMENT '商品ID',
  `attribute_id` int NOT NULL COMMENT '属性ID',
  `value_id` int NOT NULL COMMENT '属性值ID',
  `price_adjustment` decimal(10,2) DEFAULT NULL COMMENT '价格调整',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';

-- 正在导出表  swetelove.product_attributes 的数据：~2 rows (大约)
DELETE FROM `product_attributes`;
INSERT INTO `product_attributes` (`id`, `product_id`, `attribute_id`, `value_id`, `price_adjustment`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(6, 4, 1, 1, 0.00, '2023-05-18 14:11:54', '2023-05-18 14:11:54', NULL),
	(7, 5, 2, 3, 10.00, '2023-05-18 14:11:54', '2023-05-18 14:11:54', NULL);

-- 导出  表 swetelove.product_categories 结构
DROP TABLE IF EXISTS `product_categories`;
CREATE TABLE IF NOT EXISTS `product_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL COMMENT '商品ID',
  `category_id` int NOT NULL COMMENT '分类ID',
  `parent_id` int DEFAULT NULL COMMENT '父级分类ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品分类表';

-- 正在导出表  swetelove.product_categories 的数据：~2 rows (大约)
DELETE FROM `product_categories`;
INSERT INTO `product_categories` (`id`, `product_id`, `category_id`, `parent_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(7, 4, 1, NULL, '2023-05-18 14:11:11', '2023-05-18 14:11:11', NULL),
	(8, 5, 2, NULL, '2023-05-18 14:11:11', '2023-05-18 14:11:11', NULL);

-- 导出  表 swetelove.reviews 结构
DROP TABLE IF EXISTS `reviews`;
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL COMMENT '商品ID',
  `user_id` int DEFAULT NULL COMMENT '用户ID',
  `rating` int DEFAULT NULL COMMENT '评分',
  `review_text` text COMMENT '评价文本',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评价表';

-- 正在导出表  swetelove.reviews 的数据：~2 rows (大约)
DELETE FROM `reviews`;
INSERT INTO `reviews` (`id`, `product_id`, `user_id`, `rating`, `review_text`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(7, 4, 1, 5, 'Great product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL),
	(8, 5, 2, 4, 'Good product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
