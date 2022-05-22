CREATE TABLE `roles` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `short_codes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `short_code` varchar(255) UNIQUE NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `client_id` varchar(255) NOT NULL,
  `client_secret` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `contact_person` varchar(255) NOT NULL,
  `role_id` int NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `services` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `shortcode_id` int NOT NULL,
  `user_id` int NOT NULL,
  `service_name` varchar(255) NOT NULL DEFAULT '',
  `service_id` varchar(255) NOT NULL DEFAULT '',
  `service_interface` ENUM ('SOAP', 'HTTP', 'SMPP'),
  `service` ENUM ('SMS','USSD','VOICE','SUBSCRIPTION'),
  `service_type` ENUM ('DAILY', 'WEEKLY','BI-WEEKLY','MONTHLY','ON-DEMAND'), 
  `product_id` varchar(255) NOT NULL DEFAULT '',
  `node_id` varchar(255) NOT NULL DEFAULT '',
  `subscription_id` varchar(255) NOT NULL DEFAULT '',
  `subscription_description` varchar(255) NOT NULL DEFAULT '',
  `base_url` varchar(255) NOT NULL DEFAULT '',
  `datasync_endpoint` varchar(255) NOT NULL DEFAULT '',
  `notification_endpoint` varchar(255) NOT NULL DEFAULT '',
  `network_type` ENUM ('MTN', 'AIRTEL', 'GLO', '9MOBILE'),
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE `transactions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `shortcode` int NOT NULL,
  `msisdn` varchar(255) NOT NULL,
  `gateway` ENUM ('SMS','USSD','WEB') DEFAULT 'SMS',
  `transaction_id` varchar(255) NOT NULL DEFAULT '',
  `subscription_id` varchar(255) NOT NULL DEFAULT '',
  `subscription_description` varchar(255) NOT NULL DEFAULT '',
  `charge_amount` varchar(255) NOT NULL DEFAULT '',
  `status_code` varchar(255) NOT NULL DEFAULT '',
  `charging_mode` varchar(255) NOT NULL DEFAULT '',
  `validity_type` varchar(255) NOT NULL DEFAULT '',
  `operation_id` varchar(255) NOT NULL DEFAULT '',
  `validity_days` varchar(255) NOT NULL DEFAULT '',
  `status_message` varchar(255) NOT NULL DEFAULT 'ACCEPTED',
  `datasync_status_code` ENUM ('PENDING', 'SENT', 'FAILED'),
  `network_type` ENUM ('MTN', 'AIRTEL', 'GLO', '9MOBILE') DEFAULT 'MTN',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `short_codes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `services` ADD FOREIGN KEY (`shortcode_id`) REFERENCES `short_codes` (`id`);

ALTER TABLE `services` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
