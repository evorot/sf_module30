DROP TABLE IF EXISTS tasks_labels,tasks, labels, users;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT

);
CREATE TABLE labels (
                       id SERIAL PRIMARY KEY,
                       name TEXT
);
CREATE TABLE tasks (
                        id SERIAL PRIMARY KEY,
                        opened BIGINT DEFAULT extract (epoch from now()),
                        closed BIGINT DEFAULT 0,
                        author_id INT REFERENCES users(id) DEFAULT 0,
                        assigned_id INT REFERENCES users(id) DEFAULT 0,
                        title TEXT,
                        content TEXT
);

CREATE TABLE tasks_labels (
                              task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE ,
                              label_id INTEGER REFERENCES labels(id) ON DELETE SET NULL
);

-------------------------------------------------------
CREATE OR REPLACE PROCEDURE populate()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..100 LOOP
            INSERT INTO users(name) VALUES ('User ' || i);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;
CALL populate();
SELECT * FROM users;

-------------------------------------------------------
CREATE OR REPLACE PROCEDURE create_labels()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..5 LOOP
            INSERT INTO labels(name) VALUES ('Label ' || i);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;
CALL create_labels();
SELECT * FROM labels;

------------------------------------------
CREATE OR REPLACE PROCEDURE create_tasks()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..100 LOOP
            INSERT INTO tasks(author_id, assigned_id, title, content) VALUES (i,i,'Заголовок '||i,'Создание текста №'||i);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;
CALL create_tasks();
------------------------------------------

CREATE OR REPLACE PROCEDURE create_tasks_labels()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..20 LOOP
            INSERT INTO tasks_labels(task_id,label_id) VALUES (i,1);
        END LOOP;
    FOR i IN 21..40 LOOP
            INSERT INTO tasks_labels(task_id,label_id) VALUES (i,2);
        END LOOP;
    FOR i IN 41..60 LOOP
            INSERT INTO tasks_labels(task_id,label_id) VALUES (i,3);
        END LOOP;
    FOR i IN 61..80 LOOP
            INSERT INTO tasks_labels(task_id,label_id) VALUES (i,4);
        END LOOP;
    FOR i IN 81..100 LOOP
            INSERT INTO tasks_labels(task_id,label_id) VALUES (i,5);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;
CALL create_tasks_labels();



