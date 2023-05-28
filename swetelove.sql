-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by joe at 28-05-2023 20:21.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性值表';

-- Inserting 6 rows into attribute_values
-- Insert batch #1
INSERT INTO attribute_values (id, attribute_id, `value`, created_at, updated_at, deleted_at) VALUES
(9, 4, 'Red', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
(10, 4, 'Blue', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
(11, 5, 'haha', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
(12, 5, 'heihei', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
(13, 6, 'Small', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL),
(14, 6, 'Large', '2023-05-18 14:07:26', '2023-05-18 14:07:26', NULL);

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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性表';

-- Inserting 3 rows into attributes
-- Insert batch #1
INSERT INTO attributes (id, attribute_name, created_at, updated_at, deleted_at) VALUES
(4, 'Color', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL),
(5, 'Style', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL),
(6, 'Size', '2023-05-18 14:06:05', '2023-05-18 14:06:05', NULL);

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

-- BEGIN TABLE images
DROP TABLE IF EXISTS images;
CREATE TABLE `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `image_url` varchar(255) DEFAULT NULL COMMENT '图片链接',
  `imageable_id` int DEFAULT NULL COMMENT '图片关联对象的ID',
  `imageable_type` varchar(255) DEFAULT NULL COMMENT '图片关联对象的类型',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图片表';

-- Inserting 18 rows into images
-- Insert batch #1
INSERT INTO images (id, image_url, imageable_id, imageable_type, created_at, updated_at, deleted_at) VALUES
(7, 'http://category1.image', 1, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
(8, 'http://category2.image', 2, 'categories', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
(9, 'http://review1.image', 7, 'reviews', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
(10, 'http://review2.image', 8, 'reviews', '2023-05-18 14:08:06', '2023-05-18 14:08:06', NULL),
(11, 'http://product1.image', 4, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(12, 'http://product2.image', 5, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(13, 'http://product1.image', 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(14, 'http://product2.image', 7, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(15, 'http://product1.image', 8, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(16, 'http://product2.image', 4, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(17, 'http://product1.image', 5, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(18, 'http://product2.image', 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(19, 'http://product1.image', 7, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(20, 'http://product2.image', 8, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(21, 'http://product1.image', 9, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(22, 'http://product2.image', 9, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(23, 'http://product1.image', 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL),
(24, 'http://product2.image', 6, 'products', '2023-05-18 14:09:39', '2023-05-18 14:09:39', NULL);

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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';

-- Inserting 2 rows into product_attributes
-- Insert batch #1
INSERT INTO product_attributes (id, product_id, attribute_id, value_id, price_adjustment, created_at, updated_at, deleted_at) VALUES
(6, 4, 1, 1, 0, '2023-05-18 14:11:54', '2023-05-18 14:11:54', NULL),
(7, 5, 2, 3, 10, '2023-05-18 14:11:54', '2023-05-18 14:11:54', NULL);

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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评价表';

-- Inserting 2 rows into reviews
-- Insert batch #1
INSERT INTO reviews (id, product_id, user_id, rating, review_text, created_at, updated_at, deleted_at) VALUES
(7, 4, 1, 5, 'Great product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL),
(8, 5, 2, 4, 'Good product!', '2023-05-18 14:10:44', '2023-05-18 14:10:44', NULL);

-- END TABLE reviews

