                           Table "public.archived_user_data"
   Column   |            Type             |                            Modifiers                            
------------+-----------------------------+-----------------------------------------------------------------
 id         | bigint                      | not null default nextval('archived_user_data_id_seq'::regclass)
 uid        | integer                     | not null
 media_url  | text                        | not null
 created_on | timestamp without time zone | default now()
Indexes:
    "archived_user_data_pkey" PRIMARY KEY, btree (id)


Table "public.cities"
   Column   |            Type             |                         Modifiers                          
------------+-----------------------------+------------------------------------------------------------
 id         | bigint                      | not null default nextval('cities_id_seq'::regclass)
 city_name  | character varying(50)       | not null
 state_id   | smallint                    | not null
 created_on | timestamp without time zone | default now()
 updated_on | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "cities_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "cities_state_id_fkey" FOREIGN KEY (state_id) REFERENCES states(id)




        Table "public.cityoffers"
     Column     |            Type             |                         Modifiers                          
----------------+-----------------------------+------------------------------------------------------------
 id             | bigint                      | not null default nextval('cityoffers_id_seq'::regclass)
 city_id        | smallint                    | 
 user_id        | smallint                    | 
 image_url      | character varying(250)      | default ''::character varying
 offer_title    | character varying(250)      | default ''::character varying
 offer_subtitle | character varying(250)      | default ''::character varying
 offer_message  | text                        | default ''::text
 status         | smallint                    | default 0
 deactivated_by | integer                     | default 0
 created_on     | timestamp without time zone | default now()
 updated_on     | timestamp without time zone | default '1990-01-01 12:00:00'::timestamp without time zone
Indexes:
    "cityoffers_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "cityoffers_city_id_fkey" FOREIGN KEY (city_id) REFERENCES cities(id)
    "cityoffers_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)


    

              Table "public.comments"
     Column      |            Type             |                           Modifiers                           
-----------------+-----------------------------+---------------------------------------------------------------
 comment_id      | bigint                      | not null default nextval('comments_comment_id_seq'::regclass)
 user_id         | integer                     | not null
 post_id         | integer                     | not null
 comment_message | text                        | not null
 image_url       | text                        | default ''::text
 status          | smallint                    | not null
 created_on      | timestamp without time zone | default now()
 updated_on      | timestamp without time zone | 
 deactivated_on  | timestamp without time zone | 
Indexes:
    "comments_pkey" PRIMARY KEY, btree (comment_id)
Foreign-key constraints:
    "comments_post_id_fkey" FOREIGN KEY (post_id) REFERENCES posts(id)
    "comments_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)










