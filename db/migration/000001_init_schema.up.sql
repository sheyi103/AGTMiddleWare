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
  `password` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `contact_person` varchar(255) NOT NULL,
  `role_id` int NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `services` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `client_id` varchar(255) NOT NULL,
  `client_secret` varchar(255) NOT NULL,
  `shortcode_id` int NOT NULL,
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  `service_name` varchar(255),
  `service_id` varchar(255),
  `service_interface` ENUM ('SOAP', 'HTTP', 'SMPP'),
  `service` ENUM ('SMS','USSD','VOICE','SUBSCRIPTION'),
  `service_type` ENUM ('DAILY', 'WEEKLY','MONTHLY','ON-DEMAND'),
  `product_id` varchar(255),
  `node_id` varchar(255),
  `subscription_id` varchar(255),
  `subscription_description` varchar(255),
  `base_url` varchar(255),
  `datasync_endpoint` varchar(255),
  `notification_endpoint` varchar(255),
  `network_type` ENUM ('MTN', 'AIRTEL', 'GLO', '9MOBILE'),
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `services` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `services` ADD FOREIGN KEY (`shortcode_id`) REFERENCES `short_codes` (`id`);

ALTER TABLE `services` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `services` ADD CONSTRAINT `short_code_network_type_key` UNIQUE (`shortcode_id`,`network_type`);
