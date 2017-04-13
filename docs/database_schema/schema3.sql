:                                         Table "public.jobs"
      Column       |            Type             |                     Modifiers                     
-------------------+-----------------------------+---------------------------------------------------
 id                | bigint                      | not null default nextval('jobs_id_seq'::regclass)
 company_id        | integer                     | not null
 uid               | integer                     | not null
 job_title         | character varying(50)       | not null
 skills            | character varying(50)       | not null
 notice_period     | integer                     | not null
 salary_min        | integer                     | not null
 salary_max        | integer                     | not null
 experience_min    | integer                     | not null
 experience_max    | integer                     | not null
 work_type         | integer                     | not null
 about_company     | text                        | not null
 interview_process | text                        | not null
 created_on        | timestamp without time zone | default now()
 updated_on        | timestamp without time zone | 
 job_expire_on     | timestamp without time zone | 
Indexes:
    "jobs_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "jobs_company_id_fkey" FOREIGN KEY (company_id) REFERENCES companies(id)
    "jobs_uid_fkey" FOREIGN KEY (uid) REFERENCES auth_users(id)


Table "public.likes"
   Column   |            Type             |                         Modifiers                          
------------+-----------------------------+------------------------------------------------------------
 id         | bigint                      | not null default nextval('likes_id_seq'::regclass)
 user_id    | integer                     | not null
 is_liked   | boolean                     | default false
 post_id    | bigint                      | not null
 liked_on   | timestamp without time zone | default now()
 updated_on | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "likes_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "likes_post_id_fkey" FOREIGN KEY (post_id) REFERENCES posts(id)
    "likes_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)

             Table "public.media"
     Column     |            Type             |                         Modifiers                          
----------------+-----------------------------+------------------------------------------------------------
 id             | bigint                      | not null default nextval('media_id_seq'::regclass)
 media_type     | integer                     | not null
 user_id        | integer                     | not null
 post_id        | integer                     | not null
 status         | smallint                    | default 0
 deactivated_by | integer                     | not null
 created_on     | timestamp without time zone | default now()
 updated_on     | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
 media_url      | text                        | default ''::text
Indexes:
    "media_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "media_post_id_fkey" FOREIGN KEY (post_id) REFERENCES posts(id)
    "media_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)