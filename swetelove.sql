-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by joe at 04-06-2023 21:54.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


-- BEGIN TABLE advertisements
DROP TABLE IF EXISTS advertisements;
CREATE TABLE `advertisements` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Inserting 2 rows into advertisements
-- Insert batch #1
INSERT INTO advertisements (id, code, created_at, updated_at, deleted_at) VALUES
(1, 'banner', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL),
(2, 'category_banner', '2023-06-01 01:28:17', '2023-06-01 01:28:17', NULL);

-- END TABLE advertisements

-- BEGIN TABLE attribute_values
DROP TABLE IF EXISTS attribute_values;
CREATE TABLE `attribute_values` (
  `id` int NOT NULL AUTO_INCREMENT,
  `attribute_id` int NOT NULL,
  `value` varchar(255) DEFAULT NULL COMMENT '属性值',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性值表';

-- Inserting 47 rows into attribute_values
-- Insert batch #1
INSERT INTO attribute_values (id, attribute_id, `value`, created_at, updated_at, deleted_at) VALUES
(1, 1, '5', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(2, 1, '6', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(3, 1, '7', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(4, 1, '8', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(5, 1, '9', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(6, 2, 'Red', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(7, 2, 'Blue', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(8, 2, 'Green', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(9, 2, 'Yellow', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(10, 2, 'Purple', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(11, 3, 'Round', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(12, 3, 'Square', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(13, 3, 'Oval', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(14, 3, 'Heart', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(15, 3, 'Cushion', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(16, 4, 'Platinum', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(17, 4, 'Gold', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(18, 4, 'Silver', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(19, 4, 'Rose Gold', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(20, 4, 'White Gold', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(21, 5, 'Classic', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(22, 5, 'Modern', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(23, 5, 'Vintage', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(24, 5, 'Minimalist', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(25, 5, 'Bohemian', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(26, 6, 'Wedding', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(27, 6, 'Anniversary', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(28, 6, 'Birthday', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(29, 6, 'Valentine''s Day', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(30, 6, 'Christmas', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(31, 7, 'Women''s', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(32, 7, 'Men''s', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(33, 7, 'Unisex', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(34, 8, 'Tiffany & Co.', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(35, 8, 'Cartier', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(36, 8, 'Swarovski', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(37, 8, 'Pandora', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(38, 8, 'Zales', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(39, 9, '1', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(40, 9, '2', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(41, 9, '3', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(42, 9, '4', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(43, 9, '5', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(44, 10, '<$100', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(45, 10, '$100-$200', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(46, 10, '$200-$500', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL),
(47, 10, '$500-$1000', '2023-06-04 13:05:35', '2023-06-04 13:05:35', NULL);

-- END TABLE attribute_values

-- BEGIN TABLE attributes
DROP TABLE IF EXISTS attributes;
CREATE TABLE `attributes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `attribute_name` varchar(255) DEFAULT NULL COMMENT '属性名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性表';

-- Inserting 10 rows into attributes
-- Insert batch #1
INSERT INTO attributes (id, attribute_name, created_at, updated_at, deleted_at) VALUES
(1, 'Size', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(2, 'Color', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(3, 'Shape', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(4, 'Material', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(5, 'Style', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(6, 'Occasion', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(7, 'Gender', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(8, 'Brand', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(9, 'Rating', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL),
(10, 'Price Range', '2023-06-04 13:03:43', '2023-06-04 13:03:43', NULL);

-- END TABLE attributes

-- BEGIN TABLE categories
DROP TABLE IF EXISTS categories;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT '0',
  `category_name` varchar(255) DEFAULT NULL COMMENT '分类名称',
  `url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类表';

-- Inserting 31 rows into categories
-- Insert batch #1
INSERT INTO categories (id, parent_id, category_name, url, created_at, updated_at, deleted_at) VALUES
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
(19, 0, 'Men''s', '/men.html', NULL, NULL, NULL),
(20, 19, 'Men''s Rings', '/men/men-s-rings.html', NULL, NULL, NULL),
(21, 19, 'Men''s Necklaces', '/men/men-s-necklaces.html', NULL, NULL, NULL),
(22, 19, 'Cufflinks', '/men/cufflinks.html', NULL, NULL, NULL),
(23, 0, 'COLLECTION', '/collection.html', NULL, NULL, NULL),
(24, 23, 'Back to School', '/collection/back-to-school.html', NULL, NULL, NULL),
(25, 23, 'The Spirit Snake', '/collection/the-spirit-snake.html', NULL, NULL, NULL),
(26, 23, 'The Vintage Art Deco', '/collection/the-vintage-art-deco-ring.html', NULL, NULL, NULL),
(27, 23, 'The Snow Band Ring', '/collection/the-snow-band-ring.html', NULL, NULL, NULL),
(28, 23, 'The Starry Night', '/collection/the-starry-night.html', NULL, NULL, NULL),
(29, 23, 'Mother''s Day', '/collection/mother-s-day.html', NULL, NULL, NULL),
(30, 23, 'Merry Christmas', '/collection/merry-christmas.html', NULL, NULL, NULL),
(31, 23, 'Father''s Day', '/collection/father-s-day.html', NULL, NULL, NULL);

-- END TABLE categories

-- BEGIN TABLE collections
DROP TABLE IF EXISTS collections;
CREATE TABLE `collections` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `rule` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Inserting 1 row into collections
-- Insert batch #1
INSERT INTO collections (id, created_at, updated_at, deleted_at, name, type, rule) VALUES
(1, '2023-06-04 13:19:20', '2023-06-04 13:19:20', NULL, 'Newest Releases', 'new', '{"limit": 100}');

-- END TABLE collections

-- BEGIN TABLE currencies
DROP TABLE IF EXISTS currencies;
CREATE TABLE `currencies` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `exchange` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Inserting 9 rows into currencies
-- Insert batch #1
INSERT INTO currencies (id, created_at, updated_at, deleted_at, code, exchange) VALUES
(1, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'AUD', 5),
(2, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'GBP', 9),
(3, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'CAD', 5.5),
(4, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'EUR', 8.5),
(5, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'JPY', 0.07),
(6, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'MXN', 0.4),
(7, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'NZD', 4.5),
(8, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'PHP', 0.16),
(9, '0000-00-00 00:00:00.000000', '0000-00-00 00:00:00.000000', NULL, 'SGD', 5);

-- END TABLE currencies

-- BEGIN TABLE images
DROP TABLE IF EXISTS images;
CREATE TABLE `images` (
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

-- Inserting 27 rows into images
-- Insert batch #1
INSERT INTO images (id, image_url, link, imageable_id, imageable_type, created_at, updated_at, deleted_at) VALUES
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

-- END TABLE images

-- BEGIN TABLE product_attributes
DROP TABLE IF EXISTS product_attributes;
CREATE TABLE `product_attributes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL COMMENT '商品ID',
  `attribute_id` int NOT NULL COMMENT '属性ID',
  `value_id` int NOT NULL COMMENT '属性值ID',
  `price_adjustment` decimal(10,2) DEFAULT NULL COMMENT '价格调整',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';

-- Inserting 29 rows into product_attributes
-- Insert batch #1
INSERT INTO product_attributes (id, product_id, attribute_id, value_id, price_adjustment, created_at, updated_at, deleted_at) VALUES
(1, 1, 1, 1, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(2, 1, 1, 2, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(3, 1, 1, 3, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(4, 1, 1, 4, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(5, 1, 1, 5, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(6, 1, 2, 6, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(7, 1, 2, 7, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(8, 1, 2, 8, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(9, 1, 2, 9, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(10, 1, 2, 10, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(11, 2, 3, 11, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(12, 2, 3, 12, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(13, 2, 3, 13, 0, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(14, 2, 3, 14, -20, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(15, 2, 3, 15, -20, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(16, 2, 4, 16, 100, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(17, 2, 4, 17, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(18, 2, 4, 18, -100, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(19, 2, 4, 19, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(20, 2, 4, 20, -50, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(21, 3, 5, 21, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(22, 3, 5, 22, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(23, 3, 5, 23, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(24, 3, 5, 24, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(25, 3, 5, 25, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(26, 3, 6, 26, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(27, 3, 6, 27, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(28, 3, 6, 28, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL),
(29, 3, 6, 29, -10, '2023-06-04 13:11:49', '2023-06-04 13:11:49', NULL);

-- END TABLE product_attributes

-- BEGIN TABLE product_categories
DROP TABLE IF EXISTS product_categories;
CREATE TABLE `product_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL COMMENT '商品ID',
  `category_id` int NOT NULL COMMENT '分类ID',
  `parent_id` int DEFAULT NULL COMMENT '父级分类ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品分类表';

-- Inserting 2 rows into product_categories
-- Insert batch #1
INSERT INTO product_categories (id, product_id, category_id, parent_id, created_at, updated_at, deleted_at) VALUES
(7, 4, 1, NULL, '2023-05-18 14:11:11', '2023-05-18 14:11:11', NULL),
(8, 5, 2, NULL, '2023-05-18 14:11:11', '2023-05-18 14:11:11', NULL);

-- END TABLE product_categories

-- BEGIN TABLE products
DROP TABLE IF EXISTS products;
CREATE TABLE `products` (
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

-- Inserting 21 rows into products
-- Insert batch #1
INSERT INTO products (id, product_name, original_price, current_price, on_sale, description, video_url, created_at, updated_at, deleted_at) VALUES
(4, 'Product 1', 100, 90, 1, 'This is product 1', 'http://product1.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(5, 'Product 24343', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(6, 'Product 3', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(7, 'Product 4', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(8, 'Product 5', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(9, 'Product 6', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(10, 'Product 7', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(11, 'Product 8', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(12, 'Product 9', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(13, 'Product 10', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(14, 'Product d', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(15, 'Product 32', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(16, 'Product fd4', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(17, 'Product 35', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(18, 'Product jk89', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(19, 'Product d6', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(20, 'Product l;9', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(21, 'Product d5', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(22, 'Product j90', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(23, 'Product ds7', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL),
(24, 'Product p9', 110, 99, 1, 'This is product 2', 'http://product2.video', '2023-05-18 14:05:26', '2023-05-18 14:05:26', NULL);

-- END TABLE products

-- BEGIN TABLE reviews
DROP TABLE IF EXISTS reviews;
CREATE TABLE `reviews` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL COMMENT '商品ID',
  `user_id` int DEFAULT NULL COMMENT '用户ID',
  `rating` int DEFAULT NULL COMMENT '评分',
  `review_text` text COMMENT '评价文本',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评价表';

-- Inserting 12 rows into reviews
-- Insert batch #1
INSERT INTO reviews (id, product_id, user_id, rating, review_text, created_at, updated_at, deleted_at) VALUES
(7, 4, 1, 5, 'Great product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL),
(8, 5, 2, 4, 'Good product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL),
(9, 1, 1, 5, 'I love this ring! It is so beautiful and sparkly.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(10, 1, 2, 4, 'This ring is very nice but a bit too expensive for me.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(11, 1, 3, 3, 'The ring is okay but the color is not what I expected.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(12, 1, 4, 2, 'I don’t like this ring at all. It looks cheap and fake.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(13, 1, 5, 1, 'This ring is terrible! It broke after one week of wearing it.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(14, 2, 6, 5, 'This necklace is amazing! It is so elegant and classy.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(15, 2, 7, 4, 'This necklace is pretty but a bit too long for me.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(16, 2, 8, 3, 'The necklace is fine but the pearls are not very shiny.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(17, 2, 9, 2, 'I hate this necklace! It is so heavy and uncomfortable.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL),
(18, 2, 10, 1, 'This necklace is awful! It gave me a rash on my neck.', '2023-06-04 13:14:09', '2023-06-04 13:14:09', NULL);

-- END TABLE reviews

