CREATE DATABASE IF NOT EXISTS inventory_db;

CREATE TABLE IF NOT EXISTS `stock` (
    `id`            VARCHAR(36)   NOT NULL,
    `name`          VARCHAR(100)  NOT NULL,
    `price`         DECIMAL(18,4) NOT NULL,
    `availability`  INT(11)       NOT NULL,
    `is_active`     BOOLEAN       NOT NULL DEFAULT FALSE,
    `created_at`    TIMESTAMP(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`    TIMESTAMP(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY(`id`),
    INDEX(`name`)
) ENGINE = InnoDB;