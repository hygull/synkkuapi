      Table "public.states"
   Column   |            Type             |                      Modifiers                      
------------+-----------------------------+-----------------------------------------------------
 id         | bigint                      | not null default nextval('states_id_seq'::regclass)
 statename  | character varying(50)       | not null
 country_id | integer                     | not null
 status     | smallint                    | not null
 created_on | timestamp without time zone | default now()
 updated_on | timestamp without time zone | 
Indexes:
    "states_pkey" PRIMARY KEY, btree (id)


 Table "public.users_details"
         Column          |            Type             |                         Modifiers                          
-------------------------+-----------------------------+------------------------------------------------------------
 id                      | bigint                      | not null default nextval('users_details_id_seq'::regclass)
 user_id                 | integer                     | not null
 current_address         | text                        | default ''::text
 current_address_pincode | character varying(6)        | default '0'::character varying
 current_address_state   | integer                     | default 0
 present_address         | text                        | default ''::text
 present_address_pincode | character varying(6)        | default '0'::character varying
 present_address_state   | integer                     | default 0
 created_on              | timestamp without time zone | default now()
 updated_on              | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "users_details_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "users_details_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)


                         Table "public.users_workexperience"
   Column   |            Type             |                             Modifiers                             
------------+-----------------------------+-------------------------------------------------------------------
 id         | bigint                      | not null default nextval('users_workexperience_id_seq'::regclass)
 user_id    | integer                     | not null
 company_id | integer                     | not null
 start_date | date                        | default '1900-01-01'::date
 end_date   | date                        | default '1900-01-01'::date
 created_on | timestamp without time zone | default now()
 updated_on | timestamp without time zone | default '1900-01-01 12:00:00'::timestamp without time zone
Indexes:
    "users_workexperience_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "users_workexperience_company_id_fkey" FOREIGN KEY (company_id) REFERENCES companies(id)
    "users_workexperience_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth_users(id)



    