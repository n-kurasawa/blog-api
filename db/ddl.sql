CREATE DATABASE blog;
CREATE USER blog_user IDENTIFIED BY '****';
GRANT ALL PRIVILEGES ON `blog`.* TO blog_user;

CREATE TABLE contents (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` text NOT NULL,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE articles (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `slug` varchar(128) NOT NULL,
  `title` varchar(128) NOT NULL,
  `date` datetime NOT NULL,
  `cover_image` varchar(128) NOT NULL,
  `description` varchar(1024) NOT NULL,
  `content_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY idx_name (`slug`),
  CONSTRAINT `fk_content_id` FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
