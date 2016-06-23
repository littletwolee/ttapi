/*
Navicat PGSQL Data Transfer

Source Server         : 10.157.193.19
Source Server Version : 90305
Source Host           : 10.157.193.19:5432
Source Database       : ttapi
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 90305
File Encoding         : 65001

Date: 2016-06-23 20:13:27
*/


-- ----------------------------
-- Sequence structure for relationships_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "relationships_id_seq";
CREATE SEQUENCE "relationships_id_seq"
 INCREMENT 1
 MINVALUE 1
 MAXVALUE 9223372036854775807
 START 4
 CACHE 1;
SELECT setval('"public"."relationships_id_seq"', 4, true);

-- ----------------------------
-- Sequence structure for relationships_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "relationships_user_id_seq";
CREATE SEQUENCE "relationships_user_id_seq"
 INCREMENT 1
 MINVALUE 1
 MAXVALUE 9223372036854775807
 START 1
 CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "users_id_seq";
CREATE SEQUENCE "users_id_seq"
 INCREMENT 1
 MINVALUE 1
 MAXVALUE 9223372036854775807
 START 5
 CACHE 1;
SELECT setval('"public"."users_id_seq"', 5, true);

-- ----------------------------
-- Table structure for relationships
-- ----------------------------
DROP TABLE IF EXISTS "relationships";
CREATE TABLE "relationships" (
"id" int4 DEFAULT nextval('relationships_id_seq'::regclass) NOT NULL,
"user_id" int4 DEFAULT nextval('relationships_user_id_seq'::regclass) NOT NULL,
"state" varchar COLLATE "default" NOT NULL,
"type" varchar COLLATE "default" NOT NULL,
"other_user_id" int4 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
"id" int4 DEFAULT nextval('users_id_seq'::regclass) NOT NULL,
"name" varchar COLLATE "default",
"type" varchar COLLATE "default"
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Function structure for relationship_liked
-- ----------------------------
CREATE OR REPLACE FUNCTION "relationship_liked"("input_user_id" int4, "input_other_user_id" int4)
  RETURNS "pg_catalog"."void" AS $BODY$
                            Declare
                            ustate VARCHAR;
                            ostate VARCHAR;
                            _Liked VARCHAR := 'liked';
                            _Unliked VARCHAR := 'unliked';
                            _Matched VARCHAR := 'matched';
                            _RType VARCHAR := 'relationship';
                            Begin
                                     SELECT  relationships.state INTO ostate FROM relationships WHERE  relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
	                                    IF ostate = _Liked THEN
                                                  UPDATE relationships SET state =  _Matched  WHERE relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
                                                  SELECT state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                  IF ustate IS NULL THEN
                                                                INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, _Matched, _RType, input_other_user_id);
																									ELSE 
                                                                UPDATE relationships SET state = _Matched WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
                                                   END IF;
                                      ELSE
                                                                                       
                                                   SELECT relationships.state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                   IF ustate = _Unliked THEN
																														     UPDATE relationships SET state = _Liked WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
                                                   ELSEIF ustate IS NULL THEN
                                                                 INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, _Liked, _RType, input_other_user_id);
																									 END IF;
                                       END IF;
                                        
														END;
																								$BODY$
  LANGUAGE 'plpgsql' VOLATILE COST 100
;

-- ----------------------------
-- Function structure for relationship_unliked
-- ----------------------------
CREATE OR REPLACE FUNCTION "relationship_unliked"("input_user_id" int4, "input_other_user_id" int4)
  RETURNS "pg_catalog"."void" AS $BODY$
                            Declare
                            ustate VARCHAR;
                            ostate VARCHAR;
                            _Liked VARCHAR := 'liked';
                            _Unliked VARCHAR := 'unliked';
                            _Matched VARCHAR := 'matched';
                            _RType VARCHAR := 'relationship';
                            Begin
                                     SELECT  relationships.state INTO ostate FROM relationships WHERE  relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
	                                    IF ostate = _Matched THEN
                                                  UPDATE relationships SET state =  _Liked  WHERE relationships.other_user_id = input_user_id AND relationships.user_id = input_other_user_id;
                                                  SELECT state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                  IF ustate IS NULL THEN
                                                                INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, _Unliked, _RType, input_other_user_id);
																									ELSE
                                                                UPDATE relationships SET state = _Unliked WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
                                                   END IF;
                                      ELSE                                      
                                                   SELECT relationships.state INTO ustate FROM relationships WHERE  relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id;
                                                   IF ustate IS NULL THEN
                                                                 INSERT INTO relationships(user_id, state, type, other_user_id) VALUES(input_user_id, _Unliked, _RType, input_other_user_id);
																														     
                                                   ELSE
                                                                 UPDATE relationships SET state = _Unliked WHERE relationships.user_id = input_user_id AND relationships.other_user_id = input_other_user_id ;
																									 END IF;
                                       END IF;
                                        
														END;
																								$BODY$
  LANGUAGE 'plpgsql' VOLATILE COST 100
;

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------
ALTER SEQUENCE "relationships_id_seq" OWNED BY "relationships"."id";
ALTER SEQUENCE "relationships_user_id_seq" OWNED BY "relationships"."user_id";
ALTER SEQUENCE "users_id_seq" OWNED BY "users"."id";

-- ----------------------------
-- Primary Key structure for table relationships
-- ----------------------------
ALTER TABLE "relationships" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "users" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Key structure for table "relationships"
-- ----------------------------
ALTER TABLE "relationships" ADD FOREIGN KEY ("other_user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "relationships" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
