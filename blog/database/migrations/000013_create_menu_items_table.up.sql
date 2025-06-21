CREATE TABLE menu_items (
    id VARCHAR(26) NOT NULL UNIQUE PRIMARY KEY,
    menu_id VARCHAR(26) NOT NULL,
    parent_id VARCHAR(26),
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255),
    page_id VARCHAR(26),
    order_index INT DEFAULT 0,
    FOREIGN KEY (menu_id) REFERENCES menus(id),
    FOREIGN KEY (parent_id) REFERENCES menu_items(id),
    FOREIGN KEY (page_id) REFERENCES pages(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;