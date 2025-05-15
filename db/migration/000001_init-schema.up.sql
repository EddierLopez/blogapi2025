CREATE TABLE users(
	id 				int(255) AUTO_INCREMENT NOT NULL,
    name			varchar(100) NOT NULL,
    last_name		varchar(150) NOT NULL,
    role			varchar(30) NOT NULL,
    email			varchar(150) UNIQUE NOT NULL,
    password 		varchar(255) NOT NULL,
    description		text,
    image			varchar(255),
    created_at		datetime DEFAULT NULL,
    updated_at		datetime DEFAULT NULL,
    remember_token	varchar(255),
    CONSTRAINT 	pk_users PRIMARY KEY(id)
)ENGINE=INNODB;

CREATE TABLE categories(
	id 				int(255) AUTO_INCREMENT NOT NULL,
    name			varchar(100) NOT NULL,
    created_at		datetime DEFAULT NULL,
    updated_at		datetime DEFAULT NULL,
    CONSTRAINT pk_categories PRIMARY KEY(id)
    
)ENGINE=INNODB;

CREATE TABLE posts(
	id				int(255) AUTO_INCREMENT NOT NULL,
    user_id			int(255) NOT NULL,
    category_id		int(255) NOT NULL,
    title			varchar(150) NOT NULL,
    content			text NOT NULL,
    image			varchar(255) NOT NULL,
    created_at		datetime DEFAULT NULL,
    updated_at		datetime DEFAULT NULL,
    CONSTRAINT pk_posts PRIMARY KEY(id),
    CONSTRAINT fk_post_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_post_category FOREIGN KEY(category_id) REFERENCES categories(id)
    
)ENGINE=INNODB;