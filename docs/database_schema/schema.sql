/*************** COMMUNITY ******************/

                                   Table "public.community"
   Column   |            Type             |                       Modifiers                        
------------+-----------------------------+--------------------------------------------------------
 id         | bigint                      | not null default nextval('community_id_seq'::regclass)
 name       | character varying(50)       | 
 created_on | timestamp without time zone | default now()
 is_active  | boolean                     | default true
Indexes:
    "community_pkey" PRIMARY KEY, btree (id)


/****************CATEGORY *******************/

                                Table "public.category"
    Column     |            Type             |                       Modifiers                       
---------------+-----------------------------+-------------------------------------------------------
 id            | bigint                      | not null default nextval('category_id_seq'::regclass)
 category_name | character varying(50)       | not null
 status        | smallint                    | not null
 created_on    | timestamp without time zone | default now()
Indexes:
    "category_pkey" PRIMARY KEY, btree (id)


/****************AUTH_USERS *****************/


                                           Table "public.auth_users"
    Column     |            Type             |                         Modifiers                          
---------------+-----------------------------+------------------------------------------------------------
 id            | bigint                      | not null default nextval('auth_users_id_seq'::regclass)
 user_name     | character varying(50)       | default ''::character varying
 email         | character varying(250)      | not null
 first_name    | character varying(50)       | default ''::character varying
 last_name     | character varying(50)       | default ''::character varying
 dob           | date                        | default '1900-01-01'::date
 date_joined   | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
 community_id  | integer                     | not null
 profile_pic   | text                        | default ''::text
 profile_video | text                        | default ''::text
 token         | text                        | not null
 token_updated | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
 is_active     | boolean                     | default true
 is_superuser  | boolean                     | default false
 deleted_on    | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
 login_type    | smallint                    | default 1
 category_id   | integer                     | not null
Indexes:
    "auth_users_pkey" PRIMARY KEY, btree (id)
    "auth_users_email_key" UNIQUE CONSTRAINT, btree (email)
Foreign-key constraints:
    "auth_users_category_id_fkey" FOREIGN KEY (category_id) REFERENCES category(id)
    "auth_users_community_id_fkey" FOREIGN KEY (community_id) REFERENCES community(id)


