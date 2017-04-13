/* COMMUNITY  */

create table community(id bigserial primary key not null,name varchar(50),
created_on timestamp default CURRENT_TIMESTAMP,is_active boolean DEFAULT TRUE);

/* CATEGORY */

create table category(id bigserial primary key not null,category_name varchar(50) not null,status smallint not null,created_on timestamp 
default CURRENT_TIMESTAMP);

/* AUTH USERS  */

create table auth_users(id bigserial primary key not null,user_name varchar(50) DEFAULT '',email varchar(250) not null UNIQUE,first_name varchar(50) DEFAULT '',
 last_name varchar(50) DEFAULT '',dob date default '1900-01-01',
 date_joined  timestamp  default '1900-01-01 12:00:00',community_id integer not null, profile_pic text default '',profile_video text default '',
 token text not null, token_updated timestamp default '1900-01-01 12:00:00',    
 is_active boolean default TRUE, is_superuser boolean default FALSE,
 deleted_on timestamp default '1900-01-01 12:00:00',
 login_type smallint default 1,category_id integer not null,
  FOREIGN KEY(community_id) REFERENCES community(id),FOREIGN KEY(category_id) REFERENCES category(id));


/*    ARCHIVED USERS DATA */

create table archived_user_data(id bigserial primary key not null,uid integer not null,media_url text not null,
created_on timestamp default CURRENT_TIMESTAMP ,FOREIGN KEY(uid) REFERENCES auth_users(id));



/* COUNTRIES */

 create table countries(id bigserial primary key not null,sequence integer not null,country_name varchar(50) not null,
 iso_code_2 varchar(50) not null,iso_code_3 varchar(50) not null,status smallint not null,created_on timestamp default  CURRENT_TIMESTAMP,
 updated_on timestamp default null);


/* FEEDBACK */

 create table feedback(id bigserial primary key not null,user_id integer not null,fback_message text not null,
 created_on timestamp default CURRENT_TIMESTAMP ,updated_on timestamp default null,FOREIGN KEY(user_id) REFERENCES auth_users(id));


/* STATES */

create table states(id bigserial primary key not null,statename varchar(50) not null,country_id integer not null,status smallint not null,
created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null);

/* COMPANIES */


create table companies(id bigserial primary key not null,company_name varchar(50) not null,web_url varchar(100) not null,address text not null,
state integer not null,pincode integer not null,created_on timestamp default CURRENT_TIMESTAMP);

/* CITIES */

create table cities(id bigserial primary key not null,city_name varchar(50) not null,state_id smallint not null,created_on timestamp default
CURRENT_TIMESTAMP,updated_on timestamp default null,FOREIGN KEY(state_id) REFERENCES states(id));


/* CITY OFFERS */

create table cityoffers(id bigserial primary key not null,city_id smallint ,user_id smallint,image_url varchar(250),offer_title varchar(250),
offer_subtitle varchar(250),offer_message text not null,status smallint,deactivated_by integer,created_on timestamp default CURRENT_TIMESTAMP,
updated_on timestamp default null,FOREIGN KEY(city_id) REFERENCES cities(id),FOREIGN KEY(user_id) REFERENCES auth_users(id));


/* LIKES */


create table likes (id bigserial primary key not null,user_id integer  not null,is_liked smallint not null,post_id bigint not null,
liked_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp   default null,FOREIGN KEY(post_id) REFERENCES posts(id),
FOREIGN KEY(user_id) REFERENCES auth_users(id));


/* POSTS */

create table posts(id bigserial primary key not null,base_postid integer not null,recent_postid integer not null,recent_userid integer not null,
 user_id integer not null,visibility_type smallint not null,community_type integer not null,media_type integer not null,post_type smallint not null,
 text_message text not null,status smallint not null,created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null,FOREIGN KEY
 (base_postid) REFERENCES posts(id),FOREIGN KEY(community_id) REFERENCES community(id),FOREIGN KEY(user_id) REFERENCES auth_users(id));



/* USERS_DETAILS */

 create table users_details(id bigserial primary key not null,user_id integer not null,current_address text not null,current_address_pincode integer
 not null,current_address_state integer not null,present_address text not null,present_address_pincode integer not null,present_address_state integer not null,
 created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null,FOREIGN KEY(user_id) REFERENCES auth_users(id));


/* REPORTAPROBLEM */

create table reportaproblem(id bigserial primary key not null,user_id integer not null,problem text not null,created_on timestamp default 
CURRENT_TIMESTAMP,updated_on timestamp default null);


/* SPAMPOSTS */

create table spamposts(id bigserial primary key not null,user_id integer not null,post_id integer not null,
created_on timestamp default CURRENT_TIMESTAMP,FOREIGN KEY(post_id) REFERENCES posts(id),FOREIGN KEY(user_id) REFERENCES auth_users(id));


/* NOTIFICATIONS */

 create table notifications(id bigserial primary key not null,post_id integer not null,recent_like_uid integer not null,like_count integer not null,
  recent_cmt_uid integer not null,cmt_count integer not null,created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null,
  FOREIGN KEY (post_id) REFERENCES posts(id),FOREIGN KEY(recent_cmt_uid) REFERENCES auth_users(id),FOREIGN KEY(recent_like_uid) REFERENCES auth_users(id));


/*  MEDIA */

 create table media(id bigserial primary key not null,media_type integer not null,user_id integer not null,post_id integer not null,
 status smallint not null,deactivated_by integer not null,created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null,
 FOREIGN KEY(post_id) REFERENCES posts(id),FOREIGN KEY (user_id) REFERENCES auth_users(id));


/* FRIENDS */


create table friends(id bigserial primary key not null,from_user_id integer not null,to_user_id integer not null,created_on timestamp default CURRENT_TIMESTAMP,
acc_rej_on timestamp default null,status smallint not null,unfriend_by integer not null,unfriend_on timestamp default null,
FOREIGN KEY (from_user_id) REFERENCES auth_users(id),FOREIGN KEY (to_user_id) REFERENCES auth_users(id));


/* COMMENTS */

 create table comments(comment_id bigserial primary key not null,user_id integer not null,post_id integer not null,comment_message text not null,
 image_url text not null,status smallint not null,created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null,
 deactivated_on timestamp default null); 


/* USERS_WORKEXPERIENCE */

 create table users_workexperience(id bigserial primary key not null,user_id integer not null,company_id integer not null,start_date date default '1900-01-01',
end_date date default null,created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default null);


/* JOBS */

 create table jobs(id bigserial primary key not null,company_id integer not null,uid integer not null,job_title varchar(50) not null,skills varchar(50) not null,
notice_period integer not null,salary_min integer not null,salary_max integer not null,experience_min integer not null,experience_max integer not null,
work_type integer not null,about_company text not null,interview_process text not null,created_on timestamp default CURRENT_TIMESTAMP,updated_on
timestamp default null,job_expire_on timestamp default null,FOREIGN KEY (company_id) REFERENCES companies(id),FOREIGN KEY (uid) REFERENCES auth_users(id));

