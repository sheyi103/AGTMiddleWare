CREATE TABLE `roles` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `short_codes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `short_code` varchar(255) UNIQUE NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `contact_person` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `user_credentials` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `client_id` varchar(255) NOT NULL,
  `client_secret` varchar(255) NOT NULL,
  `shortcode_id` int NOT NULL,
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  `service_name` varchar(255),
  `service_id` varchar(255),
  `service` ENUM ('SMS', 'USSD', 'VOICE'),
  `service_type` ENUM ('DAILY', 'WEEKLY', 'MONTHLY', 'ON_DEMAND'),
  `product_id` varchar(255),
  `node_id` varchar(255),
  `subscription_id` varchar(255),
  `subscription_description` varchar(255),
  `base_url` varchar(255),
  `datasyn_endpoint` varchar(255),
  `notification_endpoint` varchar(255),
  `network_type` ENUM ('MTN', 'AIRTEL', 'GLO', '9MOBILE'),
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `user_credentials` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `user_credentials` ADD FOREIGN KEY (`shortcode_id`) REFERENCES `short_codes` (`id`);

ALTER TABLE `user_credentials` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
