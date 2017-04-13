/* Surround strings with single quotes*/
insert into auth_user(id,email,user_name,fname,lname,dob,date_of_join,stoken,stoken_created_on, )  values(1,'rob@golang.org','RobertG','Robert','Griesemer','1970-01-01','2016-10-01 12:00:00','12ds@dsd#dsdds12','2016-10-01 12:00:00');

/*add a boolean column*/
alter table auth_user add is_active boolean not null default TRUE;

/*Change field name*/
alter table auth_user rename stoken_created_on to stoken_updated_on;

/*Insert 2 more rows*/
insert into auth_user(id,email,user_name,fname,lname,dob,date_of_join,stoken,stoken_created_on)  values(default,'r@golang.org','RobPike','Rob','Pike','1970-01-02','2016-10-02 12:00:02','$2ds@dsd$dsdds#$','2016-10-02 12:00:02'),(default,'ken@golang.org','KenT','Ken','Thompson','1970-01-03','2016-10-03 12:00:03','12ds@dsd#d$%dsDF','2016-10-03 12:00:03');


