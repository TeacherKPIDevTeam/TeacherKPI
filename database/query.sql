CREATE TABLE "users" (
	"UserID"	INTEGER NOT NULL,
	"Username"	text NOT NULL,
	PRIMARY KEY("UserID" AUTOINCREMENT)
);

CREATE TABLE "tasks" (
	"TaskID"	INTEGER NOT NULL,
	"TaskName"	text NOT NULL,
	"TaskDescription"	text NOT NULL,
	"User"	int NOT NULL,
	"Stage"	int NOT NULL,
	PRIMARY KEY("TaskID" AUTOINCREMENT),
	FOREIGN KEY("User") REFERENCES "users"("UserID")
    FOREIGN KEY("Stage") REFERENCES "stages"("StageID")
);

CREATE TABLE "stages" (
	"StageID"	INTEGER NOT NULL,
	"StageName"	text NOT NULL,
	PRIMARY KEY("StageID" AUTOINCREMENT)
);

--- Шаблонная запись
-- Для 'users'
INSERT INTO users VALUES(NULL,"Admin"); -- ID 1
INSERT INTO users VALUES(NULL,"Teacher"); -- ID 2
INSERT INTO users VALUES(NULL,"Guest 1"); -- ID 3
INSERT INTO users VALUES(NULL,"Guest 2"); -- ID 4
-- Для 'stages'
insert into stages values(null, 'Первая стадия');
insert into stages values(null, 'Вторая стадия');
insert into stages values(null, 'Третья стадия');
-- Для 'tasks'
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 3);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 3);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 3);

--- Выборка для users
select * from users where UserID = 3; -- Выведем юзера с ID 3
delete from users where UserID = 4; -- Удалием юзера с ID 4

-- Выборка для tasks
select * from tasks ORDER BY Stage ASC; -- Сортировка записей от меньшего к большему
select * from tasks ORDER BY Stage DESC; -- Сортировка от больш к меньш
-- Удаление по параметрам в tasks
delete from tasks where TaskID = 5; -- удаление по ID
delete from tasks where (User = 1) AND (Stage = 3); -- удаление по юзеру и стадии
