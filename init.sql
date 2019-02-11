create extension citext;

create table if not exist user_answers (
    id serial,
    user_id integer,
    variant_id integer
    exercise_id integer,
    file_name citext,
    answers jsonb,
    socre integer,
    created timestamp
);


// 
