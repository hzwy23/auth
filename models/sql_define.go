package models

var (
	sys_rdbms_001 = `select count(*) from sys_handle_logs t where user_id = ? order by handle_time desc`
	sys_rdbms_002 = `select user_id,user_passwd,status_id,continue_error_cnt from sys_sec_user where user_id = ?`
	sys_rdbms_003 = `select count(*) from sys_handle_logs t where handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc`
	sys_rdbms_004 = `update sys_sec_user set continue_error_cnt = ? where user_id = ?`
	sys_rdbms_005 = `update sys_resource_info set res_name = ?, res_up_id = ?, service_cd = ? where res_id = ?`
	sys_rdbms_006 = `select count(*) from sys_theme_resource where theme_id = ? and res_id = ?`
	sys_rdbms_007 = `delete from sys_user_info where user_id = ? and org_unit_id = ?`
	sys_rdbms_008 = `insert into sys_theme_resource(uuid,theme_id,res_id,res_url,res_open_type,res_bg_color,res_class,group_id,res_img,sort_id,new_iframe) value(uuid(),?,?,?,?,?,?,?,?,?,?)`
	sys_rdbms_009 = `update sys_theme_resource set res_url = ?, res_bg_color = ?, res_class = ?, res_img = ?, group_id = ?, sort_id = ?, res_open_type = ?, new_iframe = ? where theme_id = ? and res_id = ?`
	sys_rdbms_010 = `select org_unit_id from sys_user_info where user_id = ?`
	sys_rdbms_011 = `select distinct t2.res_url from sys_user_theme t1 inner join sys_theme_resource t2 on t1.theme_id = t2.theme_id inner join sys_resource_info t3 on t2.res_id = t3.res_id where t1.user_id = ? and t2.res_id = ? and t3.res_type = '0'`
	sys_rdbms_012 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t order by handle_time desc`
	sys_rdbms_013 = `select res_type from sys_resource_info where res_id = ?`
	sys_rdbms_014 = `update sys_sec_user set user_passwd = ? where user_id = ? and user_passwd = ?`
	sys_rdbms_015 = `update sys_sec_user set user_passwd = ? where user_id = ?`
	sys_rdbms_016 = `update sys_sec_user set status_id = ? ,continue_error_cnt = '0' where user_id = ?`
	sys_rdbms_017 = `select t.user_id,t.user_name,a.status_desc,t.create_time, t.create_user,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc, t.modify_time,t.modify_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id`
	sys_rdbms_018 = `insert into sys_user_info (user_id,user_name,create_time,create_user,user_email,user_phone,org_unit_id,modify_time,modify_user) values(?,?,now(),?,?,?,?,now(),?)`
	sys_rdbms_019 = `insert into sys_sec_user(user_id,user_passwd,status_id) values(?,?,?)`
	sys_rdbms_020 = `update sys_sec_user set user_passwd = ? where user_id = ?`
	sys_rdbms_021 = `update sys_user_info t set t.user_name = ?, t.user_phone = ?, t.user_email = ? ,t.modify_time = now(), t.modify_user = ?,t.org_unit_id = ? where t.user_id = ?`
	sys_rdbms_022 = `update sys_sec_user set status_id = 1 where user_id = ?`
	sys_rdbms_023 = `select t.user_id,t.user_name,a.status_desc,t.create_time, t.create_user,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc,t.modify_time,t.modify_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id where t.user_id = ?`
	sys_rdbms_024 = `update sys_user_theme set theme_id = ? where user_id = ?`
	sys_rdbms_025 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.create_time, t.create_user, t.modify_time, t.modify_user, t.domain_status_id from sys_domain_define t inner join sys_domain_status_attr s on t.domain_status_id = s.domain_status_id`
	sys_rdbms_026 = `insert into sys_role_define(role_id,role_name,create_user,create_time,role_status_id,modify_time,modify_user) values(?,?,?,now(),?,now(),?)`
	sys_rdbms_027 = `delete from sys_role_define where role_id = ?`
	sys_rdbms_028 = `select t.role_name,t.create_user,t.create_time,a.role_status_desc,a.role_status_id,t.modify_time,t.modify_user,t.role_id from sys_role_define t inner join sys_role_status_attr a on t.role_status_id = a.role_status_id`
	sys_rdbms_029 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t order by handle_time desc limit ?,?`
	sys_rdbms_030 = `select count(*) from sys_handle_logs`
	sys_rdbms_031 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where user_id = ? and handle_time >= str_to_date(?,'%Y-%m-%d') and handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc limit ?,?`
	sys_rdbms_032 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where user_id = ? and handle_time >= str_to_date(?,'%Y-%m-%d') order by handle_time desc limit ?,?`
	sys_rdbms_033 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where handle_time >= str_to_date(?,'%Y-%m-%d') and handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc limit ?,?`
	sys_rdbms_034 = `select count(*) from sys_handle_logs t where user_id = ? and handle_time >= str_to_date(?,'%Y-%m-%d') and handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc`
	sys_rdbms_035 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where handle_time >= str_to_date(?,'%Y-%m-%d') order by handle_time desc limit ?,?`
	sys_rdbms_036 = `insert into sys_domain_define(domain_id,domain_name,domain_status_id,create_time, create_user,modify_time,modify_user) values(?,?,?,now(),?,now(),?)`
	sys_rdbms_037 = `delete from sys_domain_define where domain_id = ?`
	sys_rdbms_038 = `update sys_domain_define set domain_name = ?, domain_status_id = ?, modify_time = now(), modify_user = ? where domain_id = ?`
	sys_rdbms_039 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc limit ?,?`
	sys_rdbms_040 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where user_id = ? order by handle_time desc limit ?,?`
	sys_rdbms_041 = `select org_unit_id,org_unit_desc,up_org_id,create_date,maintance_date,create_user,maintance_user from sys_org_info t`
	sys_rdbms_042 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t order by user_id,handle_time desc limit ?,?`
	sys_rdbms_043 = `insert into sys_org_info(org_unit_desc,up_org_id,create_date,maintance_date,create_user,maintance_user,org_unit_id) values(?,?,now(),now(),?,?,?)`
	sys_rdbms_044 = `delete from sys_org_info where org_unit_id = ?`
	sys_rdbms_045 = `insert into sys_user_theme(user_id,theme_id) values(?,?)`
	sys_rdbms_046 = `select res_id, res_name ,res_up_id from sys_resource_info where res_attr = '0'`
	sys_rdbms_047 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc, t.sys_flag, t.inner_flag from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type where t.res_type <> '2'`
	sys_rdbms_048 = `select service_cd from sys_resource_info where res_id = ?`
	sys_rdbms_049 = `select user_id,privilege_id,role_id,domain_id,permission,role_status_id from v_sys_privilege_user_domain where user_id = ? and domain_id = ?`
	sys_rdbms_050 = `update sys_role_define t set t.role_name = ? ,t.role_status_id = ?, modify_time = now(), modify_user = ? where t.role_id = ?`
	sys_rdbms_051 = `insert into sys_handle_logs(uuid,user_id,handle_time,client_ip,status_code,method,url,data) values(uuid(),?,now(),?,?,?,?,left(?,2999))`
	sys_rdbms_052 = `select distinct t.domain_id,d.domain_name from v_sys_privilege_user_domain t inner join sys_domain_define d on t.domain_id = d.domain_id where t.user_id = ?`
	sys_rdbms_053 = `select domain_id,domain_name from sys_domain_define`
	sys_rdbms_054 = `select org_unit_id,org_unit_desc,up_org_id,create_date,maintance_date,create_user,maintance_user from sys_org_info where org_unit_id = ?`
	sys_rdbms_055 = `select t.privilege_id, t.privilege_desc, t.create_user, t.create_time, t.modify_user, t.modify_time from sys_privilege t`
	sys_rdbms_056 = `insert into sys_privilege(privilege_id,privilege_desc,create_user,create_time,modify_user,modify_time) values(?,?,?,?,?,?)`
	sys_rdbms_057 = `update sys_privilege set privilege_desc = ?, modify_user = ?, modify_time = ? where privilege_id = ?`
	sys_rdbms_058 = `delete from sys_privilege where privilege_id = ?`
	sys_rdbms_059 = `select t.privilege_id, t.privilege_desc, t.create_user, t.create_time, t.modify_user, t.modify_time from sys_privilege t where t.privilege_id = ?`
	sys_rdbms_060 = `select d.uuid,t.privilege_id, t.privilege_desc, d.domain_id, e.domain_name, d.permission, d.create_user,d.create_time,d.modify_user,d.modify_time from sys_privilege t inner join sys_privilege_domain d on t.privilege_id = d.privilege_id inner join sys_domain_define e on d.domain_id = e.domain_id where t.privilege_id = ?`
	sys_rdbms_061 = `select uuid,privilege_id,privilege_desc,role_id,role_name,create_user,create_time from v_sys_privilege_role where privilege_id = ?`
	sys_rdbms_062 = `insert into sys_privilege_domain(uuid,privilege_id,domain_id,permission,create_user,create_time,modify_user,modify_time) values(?,?,?,?,?,?,?,?)`
	sys_rdbms_063 = `delete from sys_privilege_domain where uuid = ?`
	sys_rdbms_064 = `update sys_privilege_domain set domain_id = ?,permission = ?,modify_user = ?,modify_time = ? where uuid = ?`
	sys_rdbms_065 = `insert into sys_privilege_role(uuid,privilege_id,role_id,create_user,create_time,modify_user,modify_time) values(?,?,?,?,?,?,?)`
	sys_rdbms_066 = `delete from sys_privilege_role where uuid = ?`
	sys_rdbms_067 = `update sys_privilege_role set role_id = ?,modify_user = ?, modify_time = ? where uuid = ?`
	sys_rdbms_068 = `select t.role_id, t.role_name from sys_role_define t where role_status_id = '0' and not exists ( select 1 from sys_privilege_role r where t.role_id = r.role_id and r.privilege_id = ?)`
	sys_rdbms_069 = `update sys_org_info set org_unit_desc = ? ,up_org_id = ?, maintance_date = now(),maintance_user=? where org_unit_id = ?`
	sys_rdbms_070 = `select t.theme_id,i.theme_desc,res_id,res_url,res_open_type,res_bg_color,res_class,group_id,res_img,sort_id,t.new_iframe from sys_theme_resource t left join sys_theme_define i on t.theme_id = i.theme_id where t.theme_id = ? and t.res_id = ?`
	sys_rdbms_071 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc, t.sys_flag, t.inner_flag, t.service_cd from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type`
	sys_rdbms_072 = `insert into sys_resource_info(res_id,res_name,res_attr,res_up_id,res_type,inner_flag,service_cd) values(?,?,?,?,?,?,?)`
	sys_rdbms_073 = `delete from sys_resource_info where res_id = ? and res_type = '2'`
	sys_rdbms_074 = `insert into sys_role_resource(uuid,role_id,res_id) values(?,?,?)`
	sys_rdbms_075 = `delete from sys_role_resource where res_id = ?`
	sys_rdbms_076 = `delete from sys_theme_resource where res_id = ?`
	sys_rdbms_077 = `delete from sys_resource_info where res_id = ?`
	sys_rdbms_078 = `select t1.res_url from sys_index_page t1 inner join sys_user_theme t2 on t1.theme_id = t2.theme_id where t2.user_id = ?`
	sys_rdbms_079 = `select inner_flag from sys_resource_info where res_id = ?`
	sys_rdbms_080 = `select o.org_unit_id from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id where user_id = ?`
	sys_rdbms_083 = `select count(*) from sys_handle_logs t where user_id = ? and handle_time >= str_to_date(?,'%Y-%m-%d') order by handle_time desc`
	sys_rdbms_084 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.create_time, t.create_user as user_id,t.modify_time,t.modify_user from sys_domain_define t inner join sys_domain_status_attr s  on t.domain_status_id = s.domain_status_id where t.domain_id = ?`
	sys_rdbms_085 = `select t.res_id,i.res_name,i.res_up_id from v_sys_privilege_user_resource t inner join sys_resource_info i on t.res_id = i.res_id where res_type = 0 and t.user_id = ?`
	sys_rdbms_086 = `select count(*) from sys_handle_logs t where handle_time >= str_to_date(?,'%Y-%m-%d') and handle_time < str_to_date(?,'%Y-%m-%d') order by handle_time desc`
	sys_rdbms_087 = `select count(*) from sys_handle_logs t where handle_time >= str_to_date(?,'%Y-%m-%d') order by handle_time desc`
	sys_rdbms_088 = `select t.user_id,t.user_name,t.org_unit_id,i.org_unit_desc,? from sys_user_info t inner join sys_org_info i on t.org_unit_id = i.org_unit_id where not exists (select 1 from sys_role_user u where u.role_id = ? and u.user_id = t.user_id)`
	sys_rdbms_089 = `select t.res_id, t.res_name, t.res_attr, a.res_attr_desc, t.res_up_id, t.res_type, r.res_type_desc, t.sys_flag, t.inner_flag, t.service_cd from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type where res_id = ?`
	sys_rdbms_093 = `delete from sys_role_resource where role_id = ? and res_id = ?`
	sys_rdbms_094 = `select r.user_id, t.role_id, t.role_name,t.role_status_id from sys_role_define t inner join sys_role_user r on t.role_id = r.role_id where r.user_id = ? and t.role_status_id = '0'`
	sys_rdbms_095 = `select '',t.role_id, t.role_name from sys_role_define t where  not exists (select 1 from sys_role_user r where r.user_id = ? and r.role_id = t.role_id)`
	sys_rdbms_096 = `insert into sys_role_user(uuid,role_id,user_id,create_date,create_user) values(?,?,?,now(),?)`
	sys_rdbms_097 = `delete from sys_role_user where user_id = ? and role_id = ?`
	sys_rdbms_098 = `select count(*) from sys_role_user r inner join sys_role_resource e on r.role_id = e.role_id inner join sys_theme_resource v on e.res_id = v.res_id inner join sys_user_theme m on v.theme_id = m.theme_id and r.user_id = m.user_id where r.user_id = ? and v.res_url = ?`
	sys_rdbms_099 = `select t.user_id,i.user_name,i.org_unit_id,o.org_unit_desc,t.role_id,t.create_user,t.create_date from sys_role_user t inner join sys_user_info i on t.user_id = i.user_id inner join sys_org_info o on i.org_unit_id = o.org_unit_id where t.role_id = ?`
	sys_rdbms_100 = `select role_id,res_id from sys_role_resource where role_id = ?`
	sys_rdbms_101 = `select t.theme_id,i.theme_desc,res_id,res_url,res_open_type,res_bg_color,res_class,group_id,res_img,sort_id,t.new_iframe from sys_theme_resource t inner join sys_theme_define i on t.theme_id = i.theme_id where t.theme_id = ? order by group_id,sort_id asc`
	sys_rdbms_102 = `select t.user_id,t.theme_id,i.theme_desc from sys_user_theme t inner join sys_theme_define i on t.theme_id = i.theme_id where t.user_id = ?`
	sys_rdbms_103 = `select t.uuid,t.theme_id,t.res_id,d.res_name,t.res_url,t.res_open_type res_open_type,d.service_cd,d.res_up_id,t.new_iframe from sys_theme_resource t inner join sys_resource_info d on t.res_id = d.res_id where d.res_type = '2' and theme_id = ?`
	sys_rdbms_104 = `update sys_resource_info set res_name = ?, service_cd = ?, inner_flag = ? where res_id = ?`
	sys_rdbms_105 = `update sys_theme_resource set res_url = ?,new_iframe = ?,res_open_type = ?,theme_id = ? where uuid = ?`
	sys_rdbms_106 = `select count(*) from sys_handle_logs t order by user_id,handle_time desc`
	sys_rdbms_107 = `select t.domain_id,t.domain_name from sys_domain_define t where  t.domain_status_id = '0' and not exists (select 1 from  sys_privilege_domain d where t.domain_id = d.domain_id and d.privilege_id = ?)`
	sys_rdbms_108 = `select count(*) from sys_role_user t inner join sys_role_resource e on t.role_id = e.role_id where t.user_id = ? and e.res_id = ?`
)
