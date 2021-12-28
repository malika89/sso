CREATE TABLE dept (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      name VARCHAR(50) NOT NULL DEFAULT '',
      code varchar(30) NOT NULL DEFAULT '', status SMALLINT NOT NULL DEFAULT 1,
      employee_no VARCHAR(30) NOT NULL DEFAULT '',
      setup_time TIMESTAMP NOT NULL, slevel SMALLINT NOT NULL DEFAULT 1,
      parent_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "dept_fk" FOREIGN KEY ("parent_id") REFERENCES "dept" ("id") ON DELETE CASCADE
);
ALTER TABLE dept ADD CONSTRAINT code_idx UNIQUE("code");


CREATE TABLE users (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      name VARCHAR(50) NOT NULL DEFAULT '',
      join_time TIMESTAMP without time zone  NOT NULL DEFAULT CURRENT_TIMESTAMP,
      login_ip VARCHAR(200) NOT NULL DEFAULT '',
      status SMALLINT NOT NULL DEFAULT 1,
      dept_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "users_dept_fk" FOREIGN KEY ("dept_id") REFERENCES "dept" ("id") ON DELETE CASCADE
);
ALTER TABLE users ADD CONSTRAINT user_name_idx UNIQUE("name");

CREATE TABLE application (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      name VARCHAR(50) NOT NULL DEFAULT '',
      domain VARCHAR(50) NOT NULL DEFAULT '',
      status SMALLINT NOT NULL DEFAULT 1,
      chargeman_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "application_chargeman_fk" FOREIGN KEY ("chargeman_id") REFERENCES "users" ("id") ON DELETE CASCADE
);
ALTER TABLE application ADD CONSTRAINT application_name_idx UNIQUE("name");


CREATE TABLE organization (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      name VARCHAR(50) NOT NULL DEFAULT '',
      address VARCHAR(200) NOT NULL DEFAULT '',
      status SMALLINT NOT NULL DEFAULT 1,
      dept_id UUID NOT NULL,
      PRIMARY KEY ("id")
);
alter table organization add constraint  "organization_name_idx" unique ("name");
alter table organization add constraint "org_dept_fk" FOREIGN KEY ("dept_id") REFERENCES "dept" ("id") ON DELETE CASCADE

CREATE TABLE role (
     id UUID NOT NULL DEFAULT uuid_generate_v1(),
     name VARCHAR(50) NOT NULL DEFAULT '',
     is_super SMALLINT NOT NULL DEFAULT 0,
     status SMALLINT NOT NULL DEFAULT 1,
     PRIMARY KEY ("id")
);
alter table role add constraint  "role_name_idx" unique ("name");


CREATE TABLE menu (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      name VARCHAR(50) NOT NULL DEFAULT '', status SMALLINT NOT NULL DEFAULT 1,
      category VARCHAR(30) NOT NULL DEFAULT 'page',
      app_id UUID NOT NULL,
      parent_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "menu_app_fk" FOREIGN KEY ("app_id") REFERENCES "application" ("id") ON
          DELETE CASCADE,
      CONSTRAINT "menu_fk" FOREIGN KEY ("parent_id") REFERENCES "menu" ("id") ON
                          DELETE CASCADE
);
ALTER TABLE menu ADD CONSTRAINT menu_app_idx UNIQUE ("name","app_id");


CREATE TABLE role_menu (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      menu_id UUID NOT NULL,
      role_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "role_id_fk" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE CASCADE,
      CONSTRAINT "menu_id_fk" FOREIGN KEY ("menu_id") REFERENCES "menu" ("id") ON DELETE CASCADE
);
ALTER TABLE role_menu ADD CONSTRAINT role_menu_idx UNIQUE ("menu_id","role_id");



CREATE TABLE user_dept (
      id UUID NOT NULL DEFAULT uuid_generate_v1(),
      user_id UUID NOT NULL,
      dept_id UUID NOT NULL,
      PRIMARY KEY ("id"),
      CONSTRAINT "user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
      CONSTRAINT "dept_id_fk" FOREIGN KEY ("dept_id") REFERENCES "dept" ("id") ON DELETE CASCADE
);
ALTER TABLE user_dept ADD CONSTRAINT user_dept_idx UNIQUE ("user_id","dept_id");


