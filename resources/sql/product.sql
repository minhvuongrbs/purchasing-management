CREATE TABLE product
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(1000)                      NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
)
    ENGINE = innodb
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_520_ci;

-- name: CreateProduct :execrows
INSERT INTO product (name)
VALUES (?);

-- name: GetProducts :many
SELECT *
FROM product
ORDER BY updated_at DESC
LIMIT 100;