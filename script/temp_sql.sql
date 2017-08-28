create table sys_privilege_type_attr(
    privilege_type       varchar(10) primary key,
    privilege_type_desc  varchar(100)
);

create table sys_privilege(
    privilege_id   varchar(30)  primary key,
	privilege_desc varchar(200),
    privilege_type varchar(10),
    create_user    varchar(30),
    create_time    datetime,
    constraint pk_sys_privilege01 foreign key(privilege_type) references sys_privilege_type_attr(privilege_type)
);

create table sys_privilege_role(
    uuid         varchar(66) primary key,
	privilege_id varchar(66),
    role_id      varchar(66),
    create_user  varchar(30),
    create_time  datetime,
    constraint pk_sys_privilege_role01 foreign key(privilege_id) references sys_privilege(privilege_id),
    constraint pk_sys_privilege_role02 foreign key(role_id)  references sys_role_define(role_id)
);

create table sys_privilege_domain(
    uuid          varchar(66) primary key,
    privilege_id  varchar(66),
    domain_id     varchar(30),
	permission    varchar(5),
    create_user   varchar(30),
    create_time   datetime,
    constraint pk_sys_privilege_domain01 foreign key(privilege_id) references sys_privilege(privilege_id),
    constraint pk_sys_privilege_domain02 foreign key(domain_id) references sys_domain_define(domain_id)
);

