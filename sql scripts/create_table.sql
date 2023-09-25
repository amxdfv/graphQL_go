CREATE  table if not EXISTS users (
                                      id char not null primary key UNIQUE,
                                      name char not null,
                                      age numeric,
                                      address char,
                                      document_type numeric not null,
                                      ducument numeric UNIQUE );

CREATE table if not EXISTS goods (
                                     id char not null unique PRIMARY key,
                                     name char not null,
                                     price NUMERIC not null,
                                     currency char not null,
                                     country_origin char);


CREATE TABLE if not EXISTS "transactions" (
                                              id	char NOT NULL UNIQUE primary key,
                                              rrn	NUMERIC NOT NULL UNIQUE,
                                              amount NUMERIC not null,
                                              currency char not null,
                                              user_id char not null,
                                              good_id char not null,
                                              place char,
                                              t_time datetime not null,
                                              FOREIGN KEY (user_id)  REFERENCES users (id),
    FOREIGN KEY (good_id)  REFERENCES goods (id)
    );

