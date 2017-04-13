 Table "public.companies"
    Column    |            Type             |                       Modifiers                        
--------------+-----------------------------+--------------------------------------------------------
 id           | bigint                      | not null default nextval('companies_id_seq'::regclass)
 company_name | character varying(50)       | not null
 web_url      | character varying(100)      | not null
 address      | text                        | not null
 state        | integer                     | not null
 pincode      | integer                     | not null
 created_on   | timestamp without time zone | default now()
Indexes:
    "companies_pkey" PRIMARY KEY, btree (id)



     Table "public.countries"
    Column    |            Type             |                       Modifiers                        
--------------+-----------------------------+--------------------------------------------------------
 id           | bigint                      | not null default nextval('countries_id_seq'::regclass)
 sequence     | integer                     | not null
 country_name | character varying(50)       | not null
 iso_code_2   | character varying(50)       | not null
 iso_code_3   | character varying(50)       | not null
 status       | smallint                    | not null
 created_on   | timestamp without time zone | default now()
 updated_on   | timestamp without time zone | 
Indexes:
    "countries_pkey" PRIMARY KEY, btree (id)


Table "public.feedback"
    Column     |            Type             |                         Modifiers                          
---------------+-----------------------------+------------------------------------------------------------
 id            | bigint                      | not null default nextval('feedback_id_seq'::regclass)
 user_id       | integer                     | not null
 fback_message | text                        | not null
 created_on    | timestamp without time zone | default now()
 updated_on    | timestamp without time zone | default '1990-01-01 12:00:00'::timestamp without time zone
Indexes:
    "feedback_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "feedback_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)




 Table "public.friends"
    Column    |            Type             |                         Modifiers                          
--------------+-----------------------------+------------------------------------------------------------
 id           | bigint                      | not null default nextval('friends_id_seq'::regclass)
 from_user_id | integer                     | not null
 to_user_id   | integer                     | not null
 created_on   | timestamp without time zone | default now()
 acc_rej_on   | timestamp without time zone | default '1990-01-01 12:00:00'::timestamp without time zone
 status       | smallint                    | default 0
 unfriend_by  | integer                     | default 0
 unfriend_on  | timestamp without time zone | default '1990-01-01 12:00:00'::timestamp without time zone
Indexes:
    "friends_pkey" PRIMARY KEY, btree (id)
    "friends_from_user_id_to_user_id_key" UNIQUE CONSTRAINT, btree (from_user_id, to_user_id)
Foreign-key constraints:
    "friends_from_user_id_fkey" FOREIGN KEY (from_user_id) REFERENCES auth_users(id)
    "friends_to_user_id_fkey" FOREIGN KEY (to_user_id) REFERENCES auth_users(id)
