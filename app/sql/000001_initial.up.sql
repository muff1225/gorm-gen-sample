CREATE TABLE IF NOT EXISTS `users` (
	`id` int UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`created_at` datetime (3) NOT NULL DEFAULT current_timestamp(3),
`updated_at` datetime (3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE CURRENT_TIMESTAMP(3),
`deleted_at` datetime (3),
PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `purchase_orders` (
	`id` int UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL,
	`created_at` datetime (3) NOT NULL DEFAULT current_timestamp(3),
`updated_at` datetime (3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE CURRENT_TIMESTAMP(3),
`deleted_at` datetime (3),
PRIMARY KEY (`id`),
FOREIGN KEY `fk_purchase_orders_user_id` (`user_id`) REFERENCES `users` (`id`)
);