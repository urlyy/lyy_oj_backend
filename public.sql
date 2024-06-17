/*
 Navicat Premium Data Transfer

 Source Server         : 毕设pg
 Source Server Type    : PostgreSQL
 Source Server Version : 160001 (160001)
 Source Host           : 192.168.88.132:5432
 Source Catalog        : lyy_oj
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160001 (160001)
 File Encoding         : 65001

 Date: 18/06/2024 00:07:28
*/


-- ----------------------------
-- Sequence structure for contest_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."contest_id_seq";
CREATE SEQUENCE "public"."contest_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for discussion_comment_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."discussion_comment_id_seq";
CREATE SEQUENCE "public"."discussion_comment_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for discussion_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."discussion_id_seq";
CREATE SEQUENCE "public"."discussion_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for domain_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."domain_id_seq";
CREATE SEQUENCE "public"."domain_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for domain_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."domain_user_id_seq";
CREATE SEQUENCE "public"."domain_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for homework_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."homework_id_seq";
CREATE SEQUENCE "public"."homework_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for notification_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."notification_id_seq";
CREATE SEQUENCE "public"."notification_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for question_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."question_id_seq";
CREATE SEQUENCE "public"."question_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_id_seq";
CREATE SEQUENCE "public"."role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for submission_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."submission_id_seq";
CREATE SEQUENCE "public"."submission_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_id_seq";
CREATE SEQUENCE "public"."user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS "public"."config";
CREATE TABLE "public"."config" (
  "address_list" varchar[] COLLATE "pg_catalog"."default" NOT NULL,
  "recommend" text COLLATE "pg_catalog"."default" NOT NULL,
  "announce" text COLLATE "pg_catalog"."default" NOT NULL,
  "compilers" varchar(255) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."config"."recommend" IS '默认首页推荐';
COMMENT ON COLUMN "public"."config"."announce" IS '默认首页公告';

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO "public"."config" VALUES ('{192.168.88.132:8800}', '<h3 style="text-align: start;">OJ</h3><p style="text-align: start;"><a href="https://www.luogu.com.cn/" target="_blank">洛谷</a> <a href="https://codeforces.com/" target="_blank">Codeforces </a> <a href="https://atcoder.jp/" target="_blank">AtCoder</a> <a href="https://leetcode.cn/" target="_blank">力扣</a> <a href="https://www.acwing.com/" target="_blank">AcWing</a></p><h3 style="text-align: start;">网站</h3><p style="text-align: start;"><a href="https://oi-wiki.org/" target="_blank">OI Wiki &nbsp;</a><a href="https://stackoverflow.com/" target="_blank">Stack Overflow </a></p><p style="text-align: start;"><a href="https://juejin.cn/" target="_blank">稀土掘金</a> <a href="https://github.com/" target="_blank">Github</a></p>', '<p style="text-align: start;">欢迎来到吉首大学程序设计在线平台（Online Judge，简称OJ）！</p><p style="text-align: start;">这是一个面向高校师生和编程爱好者的专属竞赛和学习平台。</p><p style="text-align: start;">在这里，你可以：</p><ul><li style="text-align: start;">提升学习体验：在线编程评测系统即时反馈代码质量和正确性，让你在实践中快速成长。同时，详尽的题解和学习资料将帮助你理解难题。</li><li style="text-align: start;">与师生互动：在这里，你可以与同学和老师交流经验，分享心得。我们还提供讨论区，让你可以提问、答疑和互相学习。</li></ul><p style="text-align: start;">加入平台，共同成长，体验编程的乐趣和挑战！</p><p style="text-align: start;"><br></p><p style="text-align: start;"><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255);">拙劣的程序员担心代码。好的程序员担心数据结构及它们的关系。——林纳斯•托瓦兹</span></p><p style="text-align: start;"><br></p><p style="text-align: start;">Welcome to the Jishou University Programming Online Platform (Online Judge, abbreviated as OJ)!</p><p style="text-align: start;">This is a dedicated competition and learning platform for university faculty, students, and programming enthusiasts.</p><p style="text-align: start;">Here, you can:</p><p style="text-align: start;">• Enhance your learning experience: The online programming evaluation system provides instant feedback on code quality and correctness, allowing you to grow quickly through practice. Additionally, detailed problem solutions and learning materials will help you understand challenging problems.</p><p style="text-align: start;">• Interact with faculty and students: Here, you can exchange experiences and share insights with peers and teachers. We also offer discussion forums where you can ask questions, get answers, and learn from each other.</p><p style="text-align: start;">Join the platform to grow together and experience the fun and challenges of programming!</p><p style="text-align: start;">Whether you''re a beginner or an experienced developer, this is the ideal place for you to learn programming!</p><p style="text-align: start;"><br></p><p style="text-align: start;"><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255); font-size: 15px;">Bad programmers worry about the code. Good programmers worry about data structures and their relationships. &nbsp;</span><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255);">——</span><span style="color: rgb(13, 13, 13); background-color: rgb(255, 255, 255); font-size: 16px;">Linus Torvalds</span></p>', '[["c","gcc"],["c++","cpp"],["python3.8","python"],["java11","java"]]');

-- ----------------------------
-- Table structure for contest
-- ----------------------------
DROP TABLE IF EXISTS "public"."contest";
CREATE TABLE "public"."contest" (
  "id" int4 NOT NULL DEFAULT nextval('contest_id_seq'::regclass),
  "start_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "domain_id" int4 NOT NULL,
  "creator_id" int4 NOT NULL,
  "public" bool NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "update_time" timestamptz(6) NOT NULL,
  "is_deleted" bool NOT NULL,
  "description" text COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "problem_ids" int4[]
)
;
COMMENT ON COLUMN "public"."contest"."type" IS 'OI ACM';

-- ----------------------------
-- Records of contest
-- ----------------------------
INSERT INTO "public"."contest" VALUES (9, '2024-04-13 00:29:00+00', '2024-04-13 22:29:00+00', '第九次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-14 02:44:22.676103+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6,19}');
INSERT INTO "public"."contest" VALUES (1, '2024-04-13 20:29:00+00', '2024-04-14 22:29:00+00', '比赛的标题噢', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-19 23:20:53.159593+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (8, '2024-04-20 00:29:00+00', '2024-04-20 15:29:00+00', '第八次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-20 00:13:21.186691+00', 'f', '<p>123412341234</p>', 'ACM', '{4,6,5}');
INSERT INTO "public"."contest" VALUES (2, '2024-04-13 20:29:00+00', '2024-04-13 22:29:00+00', '第二次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-11 02:54:32.312652+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (3, '2024-04-13 20:29:00+00', '2024-04-13 22:29:00+00', '第三次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-11 02:54:32.312652+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (4, '2024-04-13 20:29:00+00', '2024-04-13 22:29:00+00', '第四次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-11 02:54:32.312652+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (7, '2024-04-20 14:29:00+00', '2024-04-20 15:29:00+00', '第七次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-19 22:56:14.139652+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (5, '2024-04-13 20:29:00+00', '2024-04-13 22:29:00+00', '第五次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-11 02:54:32.312652+00', 'f', '<p>123412341234</p>', 'IOI', '{4,6}');
INSERT INTO "public"."contest" VALUES (6, '2024-04-12 20:29:00+00', '2024-04-17 00:29:00+00', '第六次比赛', 1, 1, 't', '2024-03-20 12:33:05.37315+00', '2024-04-19 16:49:21.034342+00', 'f', '<p>123412341234</p>', 'ACM', '{4,6}');

-- ----------------------------
-- Table structure for discussion
-- ----------------------------
DROP TABLE IF EXISTS "public"."discussion";
CREATE TABLE "public"."discussion" (
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "domain_id" int4 NOT NULL,
  "creator_id" int4 NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "update_time" timestamptz(6) NOT NULL,
  "is_deleted" bool NOT NULL,
  "id" int4 NOT NULL DEFAULT nextval('discussion_id_seq'::regclass),
  "comment_num" int4 NOT NULL
)
;

-- ----------------------------
-- Records of discussion
-- ----------------------------
INSERT INTO "public"."discussion" VALUES ('1234', '<p>1234</p>', 1, 1, '2024-04-12 17:01:37.873232+00', '2024-04-12 17:01:37.873232+00', 'f', 3, 0);
INSERT INTO "public"."discussion" VALUES ('11234', '<p>去玩儿</p>', 1, 1, '2024-04-12 17:01:44.355364+00', '2024-04-12 17:01:44.355364+00', 'f', 4, 0);
INSERT INTO "public"."discussion" VALUES ('werqwer', '<p>qwerqwer</p>', 1, 1, '2024-04-12 17:01:49.385955+00', '2024-04-12 17:01:49.385955+00', 'f', 5, 0);
INSERT INTO "public"."discussion" VALUES ('asdfasdf', '<p>qwerqwerqwer</p>', 1, 1, '2024-04-12 17:01:52.97975+00', '2024-04-12 17:01:52.97975+00', 'f', 6, 0);
INSERT INTO "public"."discussion" VALUES ('asdfasdf', '<p>qwerqwerqwer</p>', 1, 1, '2024-04-12 17:01:56.040098+00', '2024-04-12 17:01:56.040098+00', 'f', 7, 0);
INSERT INTO "public"."discussion" VALUES ('asdfasdf', '<p>sdfasdfasdf</p>', 1, 1, '2024-04-12 17:01:59.419934+00', '2024-04-12 17:01:59.419934+00', 'f', 8, 0);
INSERT INTO "public"."discussion" VALUES ('qwerqwer', '<p>werqwerqwer</p>', 1, 1, '2024-04-12 17:02:02.687702+00', '2024-04-12 17:02:02.687702+00', 'f', 9, 0);
INSERT INTO "public"."discussion" VALUES ('qwerqwer', '<p>qwerqwerqwerqwerwer</p>', 1, 1, '2024-04-12 17:02:08.288486+00', '2024-04-12 17:02:08.288486+00', 'f', 10, 0);
INSERT INTO "public"."discussion" VALUES ('12341234', '<p>qwerqwerqwerqwer</p>', 1, 1, '2024-04-12 17:02:15.215436+00', '2024-04-12 17:02:15.215436+00', 'f', 11, 0);
INSERT INTO "public"."discussion" VALUES ('qwerqwer', '<p>1234234</p>', 1, 1, '2024-04-12 17:03:24.47863+00', '2024-04-12 17:03:24.47863+00', 'f', 12, 0);
INSERT INTO "public"."discussion" VALUES ('12341234123412', '<p>4123412341<strong>2341</strong>234</p>', 1, 1, '2024-03-20 13:51:26.595112+00', '2024-03-20 13:51:26.595112+00', 'f', 1, 6);
INSERT INTO "public"."discussion" VALUES ('1234', '<p>1234</p>', 1, 1, '2024-04-12 18:30:50.64375+00', '2024-04-12 18:30:50.64375+00', 'f', 13, 13);

-- ----------------------------
-- Table structure for discussion_comment
-- ----------------------------
DROP TABLE IF EXISTS "public"."discussion_comment";
CREATE TABLE "public"."discussion_comment" (
  "creator_id" int4 NOT NULL,
  "discussion_id" int4 NOT NULL,
  "reply_id" int4 NOT NULL,
  "create_time" timestamptz(0) NOT NULL,
  "is_deleted" bool NOT NULL,
  "id" int4 NOT NULL DEFAULT nextval('discussion_comment_id_seq'::regclass),
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "floor_id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of discussion_comment
-- ----------------------------
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 1, '2024-03-21 00:00:00+00', 'f', 1, '<p>123412341234</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 1, '2024-03-21 00:00:00+00', 'f', 3, '<p>卡卡那卡</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 4, '2024-04-13 16:24:10+00', 'f', 6, '<p>三级评论</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 1, '2024-03-21 00:00:00+00', 't', 2, '<p>k可以的</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 2, '2024-03-20 17:55:10+00', 't', 5, '<p>再测试一个</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (1, 1, 2, '2024-03-20 17:48:23+00', 't', 4, '<p>二级评论</p>', 1);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 24, '2024-04-13 19:36:51+00', 'f', 24, '<p>12341234</p>', 24);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 24, '2024-04-13 19:36:56+00', 'f', 25, '<p>12341234123412341</p>', 24);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 26, '2024-04-13 19:37:47+00', 'f', 26, '<p>真的</p>', 26);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 26, '2024-04-13 19:41:13+00', 'f', 27, '<p>123412341234</p>', 26);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 28, '2024-04-13 19:41:22+00', 'f', 28, '<p>qwerqwerqwer</p>', 28);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 29, '2024-04-13 19:45:07+00', 'f', 29, '<p>12341234123412</p>', 29);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 30, '2024-04-13 19:46:04+00', 'f', 30, '<p>1234123412341234</p>', 30);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 31, '2024-04-13 19:47:15+00', 'f', 31, '<p>123412341234234</p>', 31);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 32, '2024-04-13 19:51:28+00', 'f', 32, '<p>123412432341234</p>', 32);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 32, '2024-04-13 19:51:33+00', 'f', 33, '<p>qwerqwerqwer</p>', 32);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 34, '2024-04-13 19:51:36+00', 'f', 34, '<p>ghjmghjkgyhjk</p>', 34);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 32, '2024-04-13 19:52:57+00', 'f', 35, '<p><br></p>', 32);
INSERT INTO "public"."discussion_comment" VALUES (10, 13, 35, '2024-04-13 19:54:17+00', 'f', 36, '<p>12341234</p>', 32);

-- ----------------------------
-- Table structure for domain
-- ----------------------------
DROP TABLE IF EXISTS "public"."domain";
CREATE TABLE "public"."domain" (
  "id" int4 NOT NULL DEFAULT nextval('domain_id_seq'::regclass),
  "owner_id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "announce" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text,
  "is_deleted" bool NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "update_time" timestamptz(6) NOT NULL,
  "recommend" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."domain"."name" IS 'root域为超管域';
COMMENT ON COLUMN "public"."domain"."announce" IS '公告';

-- ----------------------------
-- Records of domain
-- ----------------------------
INSERT INTO "public"."domain" VALUES (2, 1, 'ROOT', '<p>4123412<strong>341</strong>234123</p><p>😀</p><p><span data-w-e-type="formula" data-w-e-is-void data-w-e-is-inline data-value="a_b^c\sum{\frac{m}{n}}"></span></p>', 'f', '2024-04-03 19:58:18+00', '2024-04-03 19:58:23+00', '<h3 style="text-align: start;">OJ推荐</h3><p style="text-align: start;"><a href="https://www.luogu.com.cn/" target="_blank">洛谷</a> <a href="https://codeforces.com/" target="_blank">Codeforces </a> <a href="https://atcoder.jp/" target="_blank">AtCoder</a> <a href="https://leetcode.cn/" target="_blank">力扣</a> <a href="https://www.acwing.com/" target="_blank">AcWing</a></p><h3 style="text-align: start;">网站推荐</h3><p style="text-align: start;"><a href="https://oi-wiki.org/" target="_blank">OI Wiki</a><a href="https://stackoverflow.com/" target="_blank">Stack Overflow </a><a href="https://juejin.cn/" target="_blank">稀土掘金</a> <a href="https://github.com/" target="_blank">Github</a></p>');
INSERT INTO "public"."domain" VALUES (1, 1, 'urlyy的域', '<p style="text-align: start;">欢迎来到吉首大学程序设计在线平台（Online Judge，简称OJ）！</p><p style="text-align: start;">这是一个面向高校师生和编程爱好者的专属竞赛和学习平台。</p><p style="text-align: start;">在这里，你可以：</p><ul><li style="text-align: start;">提升学习体验：在线编程评测系统即时反馈代码质量和正确性，让你在实践中快速成长。同时，详尽的题解和学习资料将帮助你理解难题。</li><li style="text-align: start;">与师生互动：在这里，你可以与同学和老师交流经验，分享心得。我们还提供讨论区，让你可以提问、答疑和互相学习。</li></ul><p style="text-align: start;">加入平台，共同成长，体验编程的乐趣和挑战！</p><p style="text-align: start;"><br></p><p style="text-align: start;"><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255);">拙劣的程序员担心代码。好的程序员担心数据结构及它们的关系。——林纳斯•托瓦兹</span></p><p style="text-align: start;"><br></p><p style="text-align: start;">Welcome to the Jishou University Programming Online Platform (Online Judge, abbreviated as OJ)!</p><p style="text-align: start;">This is a dedicated competition and learning platform for university faculty, students, and programming enthusiasts.</p><p style="text-align: start;">Here, you can:</p><p style="text-align: start;">• Enhance your learning experience: The online programming evaluation system provides instant feedback on code quality and correctness, allowing you to grow quickly through practice. Additionally, detailed problem solutions and learning materials will help you understand challenging problems.</p><p style="text-align: start;">• Interact with faculty and students: Here, you can exchange experiences and share insights with peers and teachers. We also offer discussion forums where you can ask questions, get answers, and learn from each other.</p><p style="text-align: start;">Join the platform to grow together and experience the fun and challenges of programming!</p><p style="text-align: start;">Whether you''re a beginner or an experienced developer, this is the ideal place for you to learn programming!</p><p style="text-align: start;"><br></p><p style="text-align: start;"><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255); font-size: 15px;">Bad programmers worry about the code. Good programmers worry about data structures and their relationships. &nbsp;</span><span style="color: rgb(25, 27, 31); background-color: rgb(255, 255, 255);">——</span><span style="color: rgb(13, 13, 13); background-color: rgb(255, 255, 255); font-size: 16px;">Linus Torvalds</span></p>', 'f', '2024-04-03 19:58:21+00', '2024-04-03 19:58:27+00', '<h3 style="text-align: start;">OJ</h3><p style="text-align: start;"><a href="https://www.luogu.com.cn/" target="_blank">洛谷</a> <a href="https://codeforces.com/" target="_blank">Codeforces </a> <a href="https://atcoder.jp/" target="_blank">AtCoder</a> <a href="https://leetcode.cn/" target="_blank">力扣</a> <a href="https://www.acwing.com/" target="_blank">AcWing</a></p><h3 style="text-align: start;">网站</h3><p style="text-align: start;"><a href="https://oi-wiki.org/" target="_blank">OI Wiki &nbsp;</a><a href="https://stackoverflow.com/" target="_blank">Stack Overflow </a></p><p style="text-align: start;"><a href="https://juejin.cn/" target="_blank">稀土掘金</a> <a href="https://github.com/" target="_blank">Github</a></p>');

-- ----------------------------
-- Table structure for domain_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."domain_user";
CREATE TABLE "public"."domain_user" (
  "id" int4 NOT NULL DEFAULT nextval('domain_user_id_seq'::regclass),
  "user_id" int4 NOT NULL,
  "domain_id" int4 NOT NULL,
  "role_id" int4 NOT NULL,
  "is_deleted" bool NOT NULL
)
;

-- ----------------------------
-- Records of domain_user
-- ----------------------------
INSERT INTO "public"."domain_user" VALUES (2, 1, 2, 2, 'f');
INSERT INTO "public"."domain_user" VALUES (1, 1, 1, 1, 'f');
INSERT INTO "public"."domain_user" VALUES (3, 10, 1, 2, 'f');
INSERT INTO "public"."domain_user" VALUES (6, 16, 1, 2, 'f');
INSERT INTO "public"."domain_user" VALUES (4, 15, 1, 2, 't');

-- ----------------------------
-- Table structure for homework
-- ----------------------------
DROP TABLE IF EXISTS "public"."homework";
CREATE TABLE "public"."homework" (
  "id" int4 NOT NULL DEFAULT nextval('homework_id_seq'::regclass),
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default" NOT NULL,
  "domain_id" int4 NOT NULL,
  "creator_id" int4 NOT NULL,
  "start_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "update_time" timestamptz(6) NOT NULL,
  "is_deleted" bool NOT NULL,
  "public" bool NOT NULL,
  "problem_ids" int4[] NOT NULL
)
;

-- ----------------------------
-- Records of homework
-- ----------------------------
INSERT INTO "public"."homework" VALUES (11, '最新的作业', '<p>4523452345</p>', 1, 1, '2024-04-20 22:36:00+00', '2024-04-20 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-19 22:55:29.47303+00', 'f', 't', '{16}');
INSERT INTO "public"."homework" VALUES (3, '第一次作业', '<p>123412341234</p>', 1, 1, '2024-04-12 00:34:00+00', '2024-04-19 23:37:00+00', '2024-04-12 15:34:21.895057+00', '2024-04-12 15:34:21.895057+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (5, '第二次作业', '<p>4523452345</p>', 1, 1, '2024-04-12 23:36:00+00', '2024-04-13 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-12 15:37:03.635013+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (6, '第三次作业', '<p>4523452345</p>', 1, 1, '2024-04-12 23:36:00+00', '2024-04-13 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-12 15:37:03.635013+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (7, '第四次作业', '<p>4523452345</p>', 1, 1, '2024-04-12 23:36:00+00', '2024-04-13 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-12 15:37:03.635013+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (8, '第五次作业', '<p>4523452345</p>', 1, 1, '2024-04-12 23:36:00+00', '2024-04-13 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-12 15:37:03.635013+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (9, '第六次作业', '<p>4523452345</p>', 1, 1, '2024-04-12 23:36:00+00', '2024-04-21 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-19 23:02:02.967957+00', 'f', 't', '{}');
INSERT INTO "public"."homework" VALUES (10, '第七次作业', '<p>4523452345</p>', 1, 1, '2024-04-19 23:36:00+00', '2024-04-23 23:37:00+00', '2024-04-12 15:37:03.635013+00', '2024-04-19 23:02:10.546296+00', 'f', 't', '{}');

-- ----------------------------
-- Table structure for notification
-- ----------------------------
DROP TABLE IF EXISTS "public"."notification";
CREATE TABLE "public"."notification" (
  "id" int4 NOT NULL DEFAULT nextval('notification_id_seq'::regclass),
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "domain_id" int4 NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "is_deleted" bool NOT NULL
)
;

-- ----------------------------
-- Records of notification
-- ----------------------------
INSERT INTO "public"."notification" VALUES (1, '第二题', '已重判', 1, '2024-04-18 22:45:52.872761+00', 'f');
INSERT INTO "public"."notification" VALUES (2, '已发布新比赛', '请同学们于周日前完成', 1, '2024-04-19 23:15:22.654814+00', 'f');

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS "public"."permission";
CREATE TABLE "public"."permission" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "bit" int4
)
;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO "public"."permission" VALUES ('创建题目', 0);
INSERT INTO "public"."permission" VALUES ('递交题目', 1);
INSERT INTO "public"."permission" VALUES ('修改题目', 2);
INSERT INTO "public"."permission" VALUES ('查看未公开题目', 3);
INSERT INTO "public"."permission" VALUES ('查看提交详情', 4);
INSERT INTO "public"."permission" VALUES ('重新判题', 5);
INSERT INTO "public"."permission" VALUES ('创建作业', 6);
INSERT INTO "public"."permission" VALUES ('修改作业', 7);
INSERT INTO "public"."permission" VALUES ('查看未公开作业', 8);
INSERT INTO "public"."permission" VALUES ('创建比赛', 9);
INSERT INTO "public"."permission" VALUES ('修改比赛', 10);
INSERT INTO "public"."permission" VALUES ('查看未公开比赛', 11);
INSERT INTO "public"."permission" VALUES ('创建讨论', 12);
INSERT INTO "public"."permission" VALUES ('删除其他人的讨论', 14);
INSERT INTO "public"."permission" VALUES ('删除其他人的评论', 15);
INSERT INTO "public"."permission" VALUES ('修改自己的讨论', 13);
INSERT INTO "public"."permission" VALUES ('删除通知', 17);
INSERT INTO "public"."permission" VALUES ('创建通知', 16);

-- ----------------------------
-- Table structure for problem
-- ----------------------------
DROP TABLE IF EXISTS "public"."problem";
CREATE TABLE "public"."problem" (
  "id" int4 NOT NULL DEFAULT nextval('question_id_seq'::regclass),
  "description" text COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamptz(6) NOT NULL,
  "memory_limit" int8 NOT NULL,
  "time_limit" int8 NOT NULL,
  "creator_id" int4 NOT NULL,
  "title" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "is_deleted" bool NOT NULL,
  "diff" int2 NOT NULL,
  "in_fmt" text COLLATE "pg_catalog"."default" NOT NULL,
  "out_fmt" text COLLATE "pg_catalog"."default" NOT NULL,
  "other" text COLLATE "pg_catalog"."default" NOT NULL,
  "public" bool NOT NULL,
  "domain_id" int4 NOT NULL,
  "update_time" timestamptz(6) NOT NULL,
  "judge_type" int2 NOT NULL,
  "test_cases" text COLLATE "pg_catalog"."default" NOT NULL,
  "special_code" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ac_num" int4 NOT NULL,
  "submit_num" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."problem"."diff" IS '0~4,0是没选难度';
COMMENT ON COLUMN "public"."problem"."public" IS '是否公开';
COMMENT ON COLUMN "public"."problem"."domain_id" IS '属于的域的ID';
COMMENT ON COLUMN "public"."problem"."judge_type" IS '0是普通，1是special judge';

-- ----------------------------
-- Records of problem
-- ----------------------------
INSERT INTO "public"."problem" VALUES (11, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (14, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (15, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (16, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (17, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (18, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (19, '<p>123</p>', '2024-04-12 15:27:56.887574+00', 131072, 1000, 1, '1234', 'f', 1, '<p>123</p>', '<p>123</p>', '<p>123</p>', 't', 1, '2024-04-12 15:27:56.887574+00', 0, '[]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (20, '<p>打印1234</p>', '2024-04-13 13:03:51.005078+00', 131072, 1000, 1, '私有题目', 'f', 2, '<p>无</p>', '<p>1234</p>', '<p><br></p>', 'f', 1, '2024-04-14 01:11:14.140492+00', 0, '[{"input":"","expect":"1234","isSample":true}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (29, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (30, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (21, '<p>打印偶数</p>', '2024-04-14 01:26:34.925839+00', 262144, 1000, 1, '新题目', 'f', 0, '<p>无</p>', '<p>a</p>', '<p><br></p>', 't', 1, '2024-04-16 02:28:21.012387+00', 1, '[{"input":"1","expect":"a","isSample":false}]', 'def judge(lines)->bool:
    flag=False
    for line in lines:
        if int(line)%2==1:
          return False
    return True', 3, 5);
INSERT INTO "public"."problem" VALUES (22, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (23, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (24, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (25, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (26, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (31, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (7, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (5, '<p>摩尔</p>', '2024-04-01 08:23:05.147392+00', 131072, 1000, 1, '膜2即可', 'f', 1, '<p>3</p>', '<p>2</p>', '<p><br></p>', 't', 1, '2024-04-20 00:15:19.583014+00', 1, '[{"input":"5","expect":"2\n4\n6","isSample":true},{"input":"7","expect":"8","isSample":false}]', 'def judge(lines):
  flag=True  
  for line in lines:
    if int(line)%2!=0:
      flag=False
      break
  return flag', 0, 0);
INSERT INTO "public"."problem" VALUES (4, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-20 00:12:24.863154+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 1);
INSERT INTO "public"."problem" VALUES (8, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (9, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (10, '<p>a+b</p>', '2024-03-19 17:50:26.482566+00', 131072, 1000, 1, '求a+b', 'f', 3, '<p>1 2</p>', '<p>3</p>', '<p>qwerqwer</p>', 't', 1, '2024-04-12 15:20:19+00', 0, '[{"input":"1 3","expect":"4","isSample":true},{"input":"5 1","expect":"6","isSample":true}]', '12341234', 0, 0);
INSERT INTO "public"."problem" VALUES (6, '<p>打印hello world</p>', '2024-04-09 13:22:18.437013+00', 131072, 1000, 1, '输出hello world', 'f', 1, '<p>无</p>', '<p>自己尝试</p>', '<p>哈哈哈签到题</p>', 't', 1, '2024-04-09 13:22:18.437013+00', 0, '[{"input":"","expect":"hello world","isSample":false}]', '', 3, 4);
INSERT INTO "public"."problem" VALUES (32, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (33, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (34, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (35, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (36, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (37, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (38, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);
INSERT INTO "public"."problem" VALUES (39, '<p>求一组数的和</p>', '2024-04-09 13:30:22.333127+00', 131072, 1000, 1, '累加', 'f', 1, '<p>第一行输入N，表示整数个数</p><p>第二行输入N个数，用空格分割</p>', '<p>打印累加的值</p>', '<p><br></p>', 't', 1, '2024-04-09 13:30:22.333127+00', 0, '[{"input":"3\n1 2 3","expect":"6","isSample":true},{"input":"5\n9 8 0 3 2","expect":"22","isSample":false},{"input":"1\n0","expect":"0","isSample":false}]', '', 0, 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS "public"."role";
CREATE TABLE "public"."role" (
  "id" int4 NOT NULL DEFAULT nextval('role_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "domain_id" int4 NOT NULL,
  "permission" int8 NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_deleted" bool,
  "create_time" timestamptz(6),
  "update_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."role"."domain_id" IS '为0的就是公有的';
COMMENT ON COLUMN "public"."role"."description" IS '角色的描述';

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO "public"."role" VALUES (1, 'owner', 0, 8388607, '拥有者', 'f', '2024-03-21 16:43:36+00', '2024-03-25 10:50:49.049352+00');
INSERT INTO "public"."role" VALUES (5, 'urlyy', 1, 9, 'qwerqwer', 't', '2024-04-13 23:40:47.78494+00', '2024-04-13 23:40:47.78494+00');
INSERT INTO "public"."role" VALUES (2, 'default', 0, 4988930, '默认用户', 'f', '2024-03-25 22:34:42+00', '2024-03-25 22:34:45+00');
INSERT INTO "public"."role" VALUES (3, 'student', 1, 15, '学生12341234', 'f', '2024-03-21 17:31:46+00', '2024-04-13 23:40:29.764074+00');
INSERT INTO "public"."role" VALUES (4, '1234', 1, 196610, 'qwer', 'f', '2024-03-25 10:56:46.178258+00', '2024-03-25 10:56:46.178258+00');

-- ----------------------------
-- Table structure for submission
-- ----------------------------
DROP TABLE IF EXISTS "public"."submission";
CREATE TABLE "public"."submission" (
  "id" int4 NOT NULL DEFAULT nextval('submission_id_seq'::regclass),
  "problem_id" int4 NOT NULL,
  "domain_id" int4 NOT NULL,
  "from_type" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "submit_time" timestamptz(6) NOT NULL,
  "status" int2 NOT NULL,
  "max_memory" int4 NOT NULL,
  "max_time" int4 NOT NULL,
  "lang" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "last_judge_time" timestamptz(6) NOT NULL,
  "code" text COLLATE "pg_catalog"."default" NOT NULL,
  "from_id" int4 NOT NULL,
  "pass_percent" float4 NOT NULL,
  "log" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."submission"."from_type" IS '作业还是参赛';
COMMENT ON COLUMN "public"."submission"."from_id" IS '从竞赛还是作业里的id';

-- ----------------------------
-- Records of submission
-- ----------------------------
INSERT INTO "public"."submission" VALUES (27, 4, 1, 'contest', 10, '2024-04-10 12:11:18.052443+00', 0, 0, 0, 'gcc', '2024-04-10 12:11:18.052443+00', '#include<stdio.h>
int main(){
printf("2")
}', 1, 0, 'data/code: In function ''main'':
data/code:3:12: error: expected '';'' before ''}'' token
    3 | printf("2")
      |            ^
      |            ;
    4 | }
      | ~           
');
INSERT INTO "public"."submission" VALUES (30, 18, 1, 'problem', 1, '2024-04-13 11:54:23.52441+00', 2, 1152, 2, 'gcc', '2024-04-13 11:54:23.52441+00', '#include<stdio.h>
int main(){
printf("4");
}', 0, 0.5, '');
INSERT INTO "public"."submission" VALUES (31, 18, 1, 'problem', 1, '2024-04-13 11:56:43.103654+00', 2, 1152, 3, 'gcc', '2024-04-13 11:56:43.103654+00', '#include<stdio.h>
int main(){
printf("4");
}', 0, 0.5, '');
INSERT INTO "public"."submission" VALUES (29, 4, 1, 'contest', 1, '2024-04-12 18:33:39.435328+00', 2, 0, 0, 'gcc', '2024-04-13 12:04:26.113965+00', '#include<stdio.h>
int main(){
printf("2");
}', 7, 0, '');
INSERT INTO "public"."submission" VALUES (42, 21, 1, 'problem', 1, '2024-04-14 01:28:05.495523+00', 0, 1152, 1, 'gcc', '2024-04-14 01:30:44.112788+00', '#include<stdio.h>
int main(){
printf("a");
}', 0, 1, '');
INSERT INTO "public"."submission" VALUES (32, 16, 1, 'problem', 1, '2024-04-13 12:05:26.947027+00', 2, 1152, 3, 'gcc', '2024-04-13 12:05:36.826316+00', '#include<stdio.h>
int main(){
  printf("4");
}', 0, 0.5, '');
INSERT INTO "public"."submission" VALUES (33, 6, 1, 'contest', 1, '2024-04-13 12:06:48.109801+00', 0, 1280, 3, 'gcc', '2024-04-13 12:06:48.109801+00', '#include<stdio.h>
int main(){
printf("hello world");
}', 9, 1, '');
INSERT INTO "public"."submission" VALUES (34, 4, 1, 'contest', 1, '2024-04-13 12:10:12.959119+00', 2, 0, 0, 'gcc', '2024-04-13 12:10:12.959119+00', '#include<stdio.h>
int main(){
printf("2");
}', 8, 0, '');
INSERT INTO "public"."submission" VALUES (38, 10, 1, 'problem', 1, '2024-04-13 18:03:22.420297+00', 2, 0, 0, 'gcc', '2024-04-13 18:03:22.420297+00', '#include<stdio.h>
int main(){
printf("2");
}', 0, 0, '');
INSERT INTO "public"."submission" VALUES (16, 7, 1, 'problem', 1, '2024-03-22 13:32:28+00', 0, 1408, 2, 'gcc', '2024-04-09 13:32:28.458362+00', '#include<stdio.h>
int main(){
  int N;
  scanf("%d",&N);
  int sum = 0;
  for(int i=0;i<N;i++){
    int a;
    scanf("%d",&a);
    sum+=a;
  }
  printf("%d",sum);
}', 0, 1, '');
INSERT INTO "public"."submission" VALUES (17, 7, 1, 'problem', 1, '2024-03-22 13:32:28+00', 0, 1408, 2, 'gcc', '2024-04-09 13:34:03.78394+00', '#include<stdio.h>
int main(){
  int N;
  scanf("%d",&N);
  int sum = 0;
  for(int i=0;i<N;i++){
    int a;
    scanf("%d",&a);
    sum+=a;
  }
  printf("%d",sum);
}', 0, 1, '');
INSERT INTO "public"."submission" VALUES (21, 4, 1, 'contest', 1, '2024-03-22 13:32:28+00', 4, 0, 0, 'gcc', '2024-04-09 14:39:36.296293+00', '#include<stdio.h>
int main(){
  while(1){}
printf("2");
}', 1, 0, '');
INSERT INTO "public"."submission" VALUES (22, 4, 1, 'contest', 1, '2024-03-22 13:32:28+00', 0, 1408, 2, 'gcc', '2024-04-09 15:07:01.611805+00', '#include<stdio.h>
int main(){
  int a,b;
  scanf("%d %d",&a,&b);
  printf("%d",a+b);
  return 0;
}', 1, 1, '');
INSERT INTO "public"."submission" VALUES (23, 4, 1, 'contest', 1, '2024-03-22 13:32:28+00', 0, 1408, 2, 'gcc', '2024-04-09 15:08:52.419587+00', '#include<stdio.h>
int main(){
  int a,b;
  scanf("%d %d",&a,&b);
  printf("%d",a+b);
  return 0;
}', 1, 1, '');
INSERT INTO "public"."submission" VALUES (24, 6, 1, 'contest', 1, '2024-03-22 16:32:28+00', 0, 1280, 2, 'gcc', '2024-04-09 16:03:00.946224+00', '#include<stdio.h>
int main(){
printf("hello world");
}', 1, 1, '');
INSERT INTO "public"."submission" VALUES (25, 4, 1, 'contest', 10, '2024-04-10 09:29:32.765548+00', 2, 1280, 3, 'gcc', '2024-04-10 09:29:32.765548+00', '#include<stdio.h>
int main(){
  printf("4");
}', 1, 0.5, '');
INSERT INTO "public"."submission" VALUES (26, 4, 1, 'contest', 10, '2024-04-10 10:03:05.019286+00', 2, 1152, 2, 'gcc', '2024-04-10 10:03:05.019286+00', '#include<stdio.h>
int main(){
printf("4");
}', 1, 0.5, '');
INSERT INTO "public"."submission" VALUES (39, 16, 1, 'homework', 1, '2024-04-13 18:04:26.481024+00', 2, 0, 0, 'gcc', '2024-04-13 18:22:16.865072+00', '#include<stdio.h>
int main(){
printf("2");
}', 11, 0, '');
INSERT INTO "public"."submission" VALUES (28, 4, 1, 'contest', 1, '2024-04-10 13:58:46.444411+00', 2, 0, 0, 'gcc', '2024-04-10 13:58:46.444411+00', '#include<stdio.h>
int main(){
printf("2");
}', 1, 0, '');
INSERT INTO "public"."submission" VALUES (18, 7, 1, 'problem', 1, '2024-03-23 13:32:28+00', 0, 1408, 3, 'gcc', '2024-04-09 13:35:28.820938+00', '#include<stdio.h>
int main(){
  int N;
  scanf("%d",&N);
  int sum = 0;
  for(int i=0;i<N;i++){
    int a;
    scanf("%d",&a);
    sum+=a;
  }
  printf("%d",sum);
}', 0, 1, '');
INSERT INTO "public"."submission" VALUES (48, 6, 1, 'contest', 1, '2024-04-20 00:20:37.451114+00', 0, 1280, 2, 'cpp', '2024-04-20 00:20:37.451114+00', '#include<stdio.h>
int main(){
printf("hello world");
}', 8, 1, '');
INSERT INTO "public"."submission" VALUES (40, 9, 1, 'problem', 1, '2024-04-13 18:05:54.975207+00', 2, 0, 0, 'gcc', '2024-04-13 18:25:24.059598+00', '#include<stdio.h>
int main(){
printf("2");
}', 0, 0, '');
INSERT INTO "public"."submission" VALUES (41, 16, 1, 'homework', 1, '2024-04-13 18:25:47.775746+00', 2, 0, 0, 'gcc', '2024-04-13 18:25:47.775746+00', '#include<stdio.h>
int main(){
printf("2");
}', 11, 0, '');
INSERT INTO "public"."submission" VALUES (45, 21, 1, 'problem', 1, '2024-04-15 20:49:38.486615+00', 0, 1280, 1, 'cpp', '2024-04-15 20:53:41.479621+00', '#include<stdio.h>
int main(){
printf("12");
}', 0, 1, '');
INSERT INTO "public"."submission" VALUES (49, 6, 1, 'contest', 16, '2024-04-20 00:29:28.611445+00', 0, 7296, 10, 'python', '2024-04-20 00:29:28.611445+00', 'print("hello world")', 8, 1, '');
INSERT INTO "public"."submission" VALUES (47, 4, 1, 'contest', 1, '2024-04-20 00:19:06.470466+00', 2, 0, 0, 'java', '2024-04-20 00:41:37.842014+00', 'public class Main{
    public static void main(String[] args){
        System.out.println("qwerwqer");
    }
}', 8, 0, '');
INSERT INTO "public"."submission" VALUES (50, 6, 1, 'contest', 1, '2024-04-20 00:44:38.508982+00', 0, 7296, 20, 'python', '2024-04-20 00:44:38.508982+00', 'print("hello world")', 8, 1, '');
INSERT INTO "public"."submission" VALUES (51, 6, 1, 'contest', 1, '2024-04-20 00:45:06.242753+00', 2, 0, 0, 'java', '2024-04-20 00:45:06.242753+00', 'public class Main{
    public static void main(String[] args){
        System.out.println("qwerwqer");
    }
}', 8, 0, '');
INSERT INTO "public"."submission" VALUES (44, 21, 1, 'problem', 1, '2024-04-15 20:46:44.980805+00', 2, 0, 0, 'cpp', '2024-04-15 21:44:50.834899+00', '#include<stdio.h>
int main(){
printf("4");
}', 0, 0, '');
INSERT INTO "public"."submission" VALUES (43, 21, 1, 'problem', 1, '2024-04-15 19:00:37.860092+00', 2, 0, 0, 'gcc', '2024-04-15 21:44:55.53127+00', '#include<stdio.h>
int main(){
printf("2");
}', 0, 0, '');
INSERT INTO "public"."submission" VALUES (46, 21, 1, 'problem', 1, '2024-04-15 20:54:36.767414+00', 0, 1152, 1, 'gcc', '2024-04-19 16:50:37.403579+00', '#include<stdio.h>
int main(){
printf("2");
}', 0, 1, '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int4 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
  "true_id" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "school" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "salt" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "gender" int2 NOT NULL,
  "is_deleted" bool NOT NULL,
  "last_login" timestamptz(6) NOT NULL,
  "website" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."user"."id" IS '主键';
COMMENT ON COLUMN "public"."user"."true_id" IS '学号或者工号';
COMMENT ON COLUMN "public"."user"."gender" IS '默认0,1是男，2是女';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "public"."user" VALUES (17, '2025520', 'xxxx', '1234', '吉首大学', '刘宇阳17@17', 'salt', 1, 'f', '2024-04-19 00:24:44.651359+00', '');
INSERT INTO "public"."user" VALUES (16, '2024420', 'qwerwqer', '1234', '吉首大学', '刘宇阳16@16', 'salt', 1, 'f', '2024-04-20 00:28:31.639274+00', '');
INSERT INTO "public"."user" VALUES (10, '2024', 'urlyy', '1234', '西安电子科技大学', '178520@16.com', 'qwer', 1, 'f', '2024-04-19 00:25:47.654732+00', '12342314');
INSERT INTO "public"."user" VALUES (1, '2003', '456', 'qwerasdf', '吉首大学', '170@163.com', '1234', 2, 'f', '2024-05-18 04:28:40.633611+00', 'http://localhost:8080');
INSERT INTO "public"."user" VALUES (15, '202040', '1234', '1234', '吉林大学', '刘宇阳15@15', 'salt', 1, 'f', '2024-05-12 12:55:59.216483+00', '');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."contest_id_seq"
OWNED BY "public"."contest"."id";
SELECT setval('"public"."contest_id_seq"', 9, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."discussion_comment_id_seq"
OWNED BY "public"."discussion_comment"."id";
SELECT setval('"public"."discussion_comment_id_seq"', 36, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."discussion_id_seq"
OWNED BY "public"."discussion"."id";
SELECT setval('"public"."discussion_id_seq"', 13, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."domain_id_seq"
OWNED BY "public"."domain"."id";
SELECT setval('"public"."domain_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."domain_user_id_seq"
OWNED BY "public"."domain_user"."id";
SELECT setval('"public"."domain_user_id_seq"', 6, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."homework_id_seq"
OWNED BY "public"."homework"."id";
SELECT setval('"public"."homework_id_seq"', 11, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."notification_id_seq"
OWNED BY "public"."notification"."id";
SELECT setval('"public"."notification_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."question_id_seq"
OWNED BY "public"."problem"."id";
SELECT setval('"public"."question_id_seq"', 39, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."role_id_seq"
OWNED BY "public"."role"."id";
SELECT setval('"public"."role_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."submission_id_seq"
OWNED BY "public"."submission"."id";
SELECT setval('"public"."submission_id_seq"', 51, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_id_seq"
OWNED BY "public"."user"."id";
SELECT setval('"public"."user_id_seq"', 17, true);

-- ----------------------------
-- Primary Key structure for table contest
-- ----------------------------
ALTER TABLE "public"."contest" ADD CONSTRAINT "contest_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table discussion_comment
-- ----------------------------
ALTER TABLE "public"."discussion_comment" ADD CONSTRAINT "discussion_comment_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table domain
-- ----------------------------
ALTER TABLE "public"."domain" ADD CONSTRAINT "domain_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table domain_user
-- ----------------------------
ALTER TABLE "public"."domain_user" ADD CONSTRAINT "domain_user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table homework
-- ----------------------------
ALTER TABLE "public"."homework" ADD CONSTRAINT "homework_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table notification
-- ----------------------------
ALTER TABLE "public"."notification" ADD CONSTRAINT "notification_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table permission
-- ----------------------------
ALTER TABLE "public"."permission" ADD CONSTRAINT "permission_bit_key" UNIQUE ("bit");

-- ----------------------------
-- Primary Key structure for table problem
-- ----------------------------
ALTER TABLE "public"."problem" ADD CONSTRAINT "question_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table role
-- ----------------------------
ALTER TABLE "public"."role" ADD CONSTRAINT "role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table submission
-- ----------------------------
ALTER TABLE "public"."submission" ADD CONSTRAINT "submission_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
