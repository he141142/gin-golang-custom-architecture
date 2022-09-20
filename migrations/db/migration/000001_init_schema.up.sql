CREATE SEQUENCE table_name_id_seq;


CREATE TABLE restaurants
(
    id integer NOT NULL DEFAULT nextval('table_name_id_seq') PRIMARY KEY ,
    owner_id            int          default NULL,
    name                varchar(50)  NOT NULL,
    address                varchar(255) not null,
    city_id             int                   default null,
    lat                 double precision      default null,
    lng                 double precision      default null,
    cover               json         default null,
    logo                json         default null,
    shipping_fee_per_km double precision      default 0,
    status              int          not null default 1,
    created_at          timestamp    null     default CURRENT_TIMESTAMP,
    updated_at          timestamp    null     default CURRENT_TIMESTAMP
);


CREATE RULE log_shoelace AS ON UPDATE TO restaurants
    DO
    update restaurants
    set updated_at = (select now());

CREATE INDEX index_name
    ON restaurants (owner_id, city_id, status);
