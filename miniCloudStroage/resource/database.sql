CREATE TABLE users (
                       user_id BIGINT PRIMARY KEY AUTO_INCREMENT,  -- 用户ID，主键，自增
                       username VARCHAR(50) NOT NULL,              -- 用户名
                       email VARCHAR(100) NOT NULL UNIQUE,         -- 用户邮箱，唯一
                       password_hash VARCHAR(255) NOT NULL,        -- 密码哈希值
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- 用户注册时间
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 用户信息更新时间
);

-- 文件表
CREATE TABLE files (
                       file_id BIGINT PRIMARY KEY AUTO_INCREMENT,  -- 文件ID，主键，自增
                       file_name VARCHAR(255) NOT NULL,            -- 文件名
                       file_size BIGINT NOT NULL,                  -- 文件大小（字节）
                       file_type VARCHAR(50) NOT NULL,             -- 文件类型（如jpg、pdf）
                       storage_path VARCHAR(255) NOT NULL,         -- 文件存储路径（如S3路径）
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- 文件创建时间
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 文件更新时间
                       owner_id BIGINT NOT NULL                   -- 文件所有者ID（外键）
);

-- 目录表
CREATE TABLE directories (
                             directory_id BIGINT PRIMARY KEY AUTO_INCREMENT, -- 目录ID，主键，自增
                             directory_name VARCHAR(255) NOT NULL,       -- 目录名
                             parent_id BIGINT,                           -- 父目录ID（外键）
                             owner_id BIGINT NOT NULL,                   -- 目录所有者ID（外键）
                             created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- 目录创建时间
                             updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 目录更新时间
);

-- 文件目录关系表
CREATE TABLE file_directory_mapping (
                                        mapping_id BIGINT PRIMARY KEY AUTO_INCREMENT, -- 映射ID，主键，自增
                                        file_id BIGINT NOT NULL,                    -- 文件ID（外键）
                                        directory_id BIGINT NOT NULL,               -- 目录ID（外键）
                                        UNIQUE (file_id, directory_id)              -- 确保文件与目录的唯一映射
);

-- 权限表
CREATE TABLE permissions (
                             permission_id BIGINT PRIMARY KEY AUTO_INCREMENT, -- 权限ID，主键，自增
                             user_id BIGINT NOT NULL,                     -- 用户ID（外键）
                             resource_id BIGINT NOT NULL,                 -- 资源ID（文件或目录ID）
                             resource_type ENUM('file', 'directory') NOT NULL, -- 资源类型（文件或目录）
                             permission ENUM('read', 'write', 'delete') NOT NULL, -- 权限类型
                             created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- 权限创建时间
                             updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 权限更新时间
                             UNIQUE (user_id, resource_id, resource_type) -- 确保用户对同一资源的唯一权限
);

-- 分享表
CREATE TABLE shares (
                        share_id BIGINT PRIMARY KEY AUTO_INCREMENT,  -- 分享ID，主键，自增
                        file_id BIGINT NOT NULL,                     -- 文件ID（外键）
                        owner_id BIGINT NOT NULL,                    -- 分享者ID（外键）
                        shared_with BIGINT NOT NULL,                 -- 被分享者ID（外键）
                        share_token VARCHAR(255) NOT NULL UNIQUE,    -- 分享令牌（用于公开链接）
                        expires_at DATETIME,                         -- 分享过期时间
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP -- 分享创建时间
);

-- 索引设计
CREATE INDEX idx_files_owner_id ON files(owner_id);
CREATE INDEX idx_directories_parent_id ON directories(parent_id);
CREATE INDEX idx_permissions_user_id ON permissions(user_id);
CREATE INDEX idx_permissions_resource_id ON permissions(resource_id);
CREATE INDEX idx_shares_file_id ON shares(file_id);
CREATE INDEX idx_shares_share_token ON shares(share_token);