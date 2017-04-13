/* community */
create table community(id serial primary key not null,name varchar(50),
	created_on timestamp default CURRENT_TIMESTAMP,is_active boolean DEFAULT TRUE);


/* auth_users */
create table auth_users(id bigserial primary key not null,user_name varchar(50) email text not null UNIQUE,fname varchar(50) DEFAULT '',
	lname varchar(50) DEFAULT '',dob date default '1900-01-01',
	doj timestamp  not null,community_id integer not null,p_pic text default '',p_video text default '',
	token text not null, token_updated_on timestamp default '1900-01-01 12:00:00',	
	created_on timestamp default CURRENT_TIMESTAMP,updated_on timestamp default '1900-01-01 12:00:00',
	is_active boolean default TRUE,is_superuser boolean default FALSE,
	deleted_on timestamp default '1900-01-01 12:00:00',
	details_updated_on timestamp default  '1900-01-01 12:00:00', ltype smallint default 1,cat_name  
	FOREIGN KEY(community_id) REFERENCES community(id) );

/* category */
create table category(id serial primary key not null,category_name varchar(50) not null,
	status smallint not null,created_on timestamp default CURRENT_TIMESTAMP);

/* countries */
create table countries(id serial primary key not null,sequence integer not null,
	country_name varchar(50) not null,iso_code_2 varchar(50) not null,
	iso_code_3 varchar(50) not null,status smallint not null,created_on timestamp default CURRENT_TIMESTAMP,
    updated_on timestamp default '1900-01-01 12:00:00');
