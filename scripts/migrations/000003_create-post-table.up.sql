CREATE TABLE IF NOT EXISTS posts(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    post_title varchar(250) NOT NULL,
    post_content LONGTEXT NOT NULL,
    post_hastags LONGTEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL    
)