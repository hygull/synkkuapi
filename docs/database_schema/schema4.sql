                Table "public.notifications"
     Column      |            Type             |                         Modifiers                          
-----------------+-----------------------------+------------------------------------------------------------
 id              | bigint                      | not null default nextval('notifications_id_seq'::regclass)
 post_id         | integer                     | not null
 recent_like_uid | integer                     | not null
 like_count      | integer                     | default 0
 recent_cmt_uid  | integer                     | not null
 cmt_count       | integer                     | default 0
 created_on      | timestamp without time zone | default now()
 updated_on      | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "notifications_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "notifications_post_id_fkey" FOREIGN KEY (post_id) REFERENCES posts(id)
    "notifications_recent_cmt_uid_fkey" FOREIGN KEY (recent_cmt_uid) REFERENCES auth_users(id)
    "notifications_recent_like_uid_fkey" FOREIGN KEY (recent_like_uid) REFERENCES auth_users(id)
  

  Table "public.posts"
     Column      |            Type             |                     Modifiers                      
-----------------+-----------------------------+----------------------------------------------------
 id              | bigint                      | not null default nextval('posts_id_seq'::regclass)
 base_postid     | integer                     | not null
 recent_postid   | integer                     | not null
 recent_userid   | integer                     | not null
 user_id         | integer                     | not null
 visibility_type | smallint                    | not null
 community_id    | integer                     | not null
 media_type      | integer                     | not null
 post_type       | smallint                    | not null
 text_message    | text                        | not null
 status          | smallint                    | not null
 created_on      | timestamp without time zone | default now()
 updated_on      | timestamp without time zone | 
Indexes:
    "posts_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "posts_base_postid_fkey" FOREIGN KEY (base_postid) REFERENCES posts(id)
    "posts_community_id_fkey" FOREIGN KEY (community_id) REFERENCES community(id)
    "posts_recent_postid_fkey" FOREIGN KEY (recent_postid) REFERENCES posts(id)
    "posts_recent_userid_fkey" FOREIGN KEY (recent_userid) REFERENCES auth_users(id)
    "posts_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)



     Table "public.reportaproblem"
   Column   |            Type             |                          Modifiers                          
------------+-----------------------------+-------------------------------------------------------------
 id         | bigint                      | not null default nextval('reportaproblem_id_seq'::regclass)
 user_id    | integer                     | not null
 problem    | text                        | not null
 created_on | timestamp without time zone | default now()
 updated_on | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "reportaproblem_pkey" PRIMARY KEY, btree (id)


             Table "public.spamposts"
   Column   |            Type             |                       Modifiers                        
------------+-----------------------------+--------------------------------------------------------
 id         | bigint                      | not null default nextval('spamposts_id_seq'::regclass)
 user_id    | integer                     | not null
 post_id    | integer                     | not null
 created_on | timestamp without time zone | default now()
Indexes:
    "spamposts_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "spamposts_post_id_fkey" FOREIGN KEY (post_id) REFERENCES posts(id)
    "spamposts_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)

 