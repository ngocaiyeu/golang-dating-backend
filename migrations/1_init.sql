-- +migrate Up
CREATE TABLE "users"
(
    "id"         varchar(64) PRIMARY KEY,
    "full_name"  varchar(20),
    "phone"      varchar(15) UNIQUE,
    "password"   varchar(255),
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "userinfo"
(
    "id"         varchar(64) PRIMARY KEY,
    "avatar_url" varchar(255) DEFAULT 'https://img2.thuthuatphanmem.vn/uploads/2019/01/04/anh-co-gai-de-thuong_025058983.jpg',
    "full_name"  varchar(20),
    "age"        smallint     DEFAULT 18,
    "sex"        varchar(4)   DEFAULT 'Nam',
    "marriage"   varchar(20)  DEFAULT 'Độc thân',
    "address"    varchar(50)  DEFAULT 'TP HCM'
);

CREATE TABLE "userprofile"
(
    "id"             varchar(64) PRIMARY KEY,
    "avatar_url"     varchar(255) DEFAULT 'https://img2.thuthuatphanmem.vn/uploads/2019/01/04/anh-co-gai-de-thuong_025058983.jpg',
    "full_name"      varchar(20),
    "age"            smallint     DEFAULT 18,
    "sex"            varchar(4)   DEFAULT 'Nam',
    "height"         smallint     DEFAULT 150,
    "job"            varchar(255) DEFAULT 'IT',
    "income"         varchar(50)  DEFAULT 'Từ 8 - 12 triệu',
    "marriage"       varchar(20)  DEFAULT 'Độc thân',
    "children"       varchar(40)  DEFAULT 'Chưa có',
    "home"           varchar(40)  DEFAULT 'Thuê nhà',
    "zodiac"         varchar(50)  DEFAULT 'Song tử',
    "status"         varchar(255) DEFAULT 'Nghiêm túc',
    "formality"      varchar(255) DEFAULT 'Hẹn gặp mặt',
    "link_fb"        varchar(255) DEFAULT '',
    "link_is"        varchar(255) DEFAULT '',
    "zl_phone"       varchar(14)  DEFAULT '',
    "address"        varchar(50)  DEFAULT 'TP HCM',
    "target"         varchar(100) DEFAULT 'Tìm bạn đời',
    "about"          varchar(255) DEFAULT 'Tấm chiếu mới',
    "count_follower"  integer      DEFAULT 0,
    "count_following" integer      DEFAULT 0,
    "count_like"      integer      DEFAULT 0,
    "created_at"     TIMESTAMPTZ NOT NULL,
    "updated_at"     TIMESTAMPTZ NOT NULL
);

CREATE TABLE "userrela"
(
    "id"        varchar(64) PRIMARY KEY,
    "follower"  text,
    "following" text
);

CREATE TABLE "likes"
(
    "id"      varchar(64) PRIMARY KEY,
    "user_id" text,
    "post_id" text
);

CREATE TABLE "posts"
(
    "id"              varchar(64) PRIMARY KEY,
    "user_id"         varchar(64),
    "access_modifier" varchar(1),
    "content"         varchar(3000),
    "image_url"       varchar(255),
    "like_count"      integer DEFAULT 0,
    "comment_count"   integer DEFAULT 0,
    "created_at"      TIMESTAMPTZ NOT NULL
);

CREATE TABLE "comments"
(
    "id"         varchar(64) PRIMARY KEY,
    "post_id"    varchar(64),
    "user_id"    varchar(64),
    "content"    varchar(1000),
    "created_at" TIMESTAMPTZ NOT NULL
);

ALTER TABLE "userinfo"
    ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "userprofile"
    ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "userrela"
    ADD FOREIGN KEY ("id") REFERENCES "userinfo" ("id");
ALTER TABLE "posts"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "likes"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "likes"
    ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
ALTER TABLE "comments"
    ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
ALTER TABLE "comments"
    ADD FOREIGN KEY ("user_id") REFERENCES "userinfo" ("id");


-- +migrate Down
DROP TABLE userprofile;
DROP TABLE userrela;
DROP TABLE likes;
DROP TABLE comments;
DROP TABLE userinfo;
DROP TABLE posts;
DROP TABLE users;

