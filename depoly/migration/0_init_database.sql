-- +goose Up
CREATE TABLE users
(
    id         int AUTO_INCREMENT PRIMARY KEY,
    name       varchar(80) NOT NULL,
    line_id    varchar(40) NOT NULL,
    updated_at datetime    null,
    created_at datetime    null
);

CREATE TABLE `groups`
(
    id       int AUTO_INCREMENT primary key,
    name     varchar(40) NOT NULL,
    group_id int         NOT NULL
);

CREATE TABLE user_group_relation
(
    user_id  int NOT NULL,
    group_id int NOT NULL
);

CREATE TABLE stores
(
    id         int AUTO_INCREMENT PRIMARY KEY,
    group_name varchar(80) NOT NULL
);

CREATE TABLE branch_stores
(
    id             int AUTO_INCREMENT PRIMARY KEY,
    name           varchar(80)                        NOT NULL,
    address        varchar(100)                       NOT NULL,
    phone          varchar(20)                        NOT NULL,
    store_group_id int                                NOT NULL,
    updated_at     datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at     datetime DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE orders
(
    id              int AUTO_INCREMENT PRIMARY KEY,
    creator         varchar(60) NOT NULL,
    `order`         json        null,
    branch_store_id int         NOT NULL,
    price           int         NOT NULL,
    group_id        varchar(60) NOT NULL,
    status          tinyint     NOT NULL,
    updated_at      datetime    NOT NULL,
    created_at      datetime    NOT NULL
);

CREATE TABLE order_choices
(
    id         int AUTO_INCREMENT PRIMARY KEY,
    order_id   int      NOT NULL,
    user_id    int      NOT NULL,
    drinks     json     NOT NULL,
    updated_at datetime NOT NULL,
    created_at datetime NOT NULL
);

CREATE TABLE menus
(
    id       int AUTO_INCREMENT PRIMARY KEY,
    name     varchar(100) NOT NULL,
    store_id int          NOT NULL,
    price    int          NOT NULL,
    img_uri  varchar(50)  NOT NULL,
    size     varchar(20)  NOT NULL
);

CREATE TABLE extras
(
    id       int AUTO_INCREMENT PRIMARY KEY,
    name     varchar(100) NOT NULL,
    store_id int          NOT NULL,
    price    int          NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS `groups`;
DROP TABLE IF EXISTS user_group_relation;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS branch_stores;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_choices;
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS extras;
