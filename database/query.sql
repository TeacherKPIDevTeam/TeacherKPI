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
	"TaskType"	INTEGER NOT NULL,
	"StageType"	INTEGER NOT NULL,
	PRIMARY KEY("TaskID" AUTOINCREMENT),
	FOREIGN KEY("User") REFERENCES "users"("UserID"),
	FOREIGN KEY("Stage") REFERENCES "stages"("StageID"),
	FOREIGN KEY("StageType") REFERENCES "stagetypes"("TypeID"),
	FOREIGN KEY("TaskType") REFERENCES "tasktypes"("TypeID")
);

CREATE TABLE "stages" (
	"StageID"	INTEGER NOT NULL,
	"StageName"	text NOT NULL,
	PRIMARY KEY("StageID")
);

CREATE TABLE "tasktypes" (
	"TypeID"	int NOT NULL,
	"TypeName"	text NOT NULL,
	PRIMARY KEY("TypeID")
);

CREATE TABLE "stagetypes" (
	"TypeID"	int NOT NULL,
	"TypeName"	text NOT NULL,
	PRIMARY KEY("TypeID")
);

--- Шаблонная запись
-- Для 'users'
INSERT INTO users VALUES(NULL,"Admin"); -- ID 1
INSERT INTO users VALUES(NULL,"Teacher"); -- ID 2
INSERT INTO users VALUES(NULL,"Guest 1"); -- ID 3
INSERT INTO users VALUES(NULL,"Guest 2"); -- ID 4
-- Для 'stages'
insert into stages values(1, 'Первая стадия');
insert into stages values(2, 'Вторая стадия');
insert into stages values(3, 'Третья стадия');
-- Для 'stagetypes'
insert into stagetypes values(1, 'Первый тип');
insert into stagetypes values(2, 'Второй тип');
insert into stagetypes values(3, 'Третий тип');
insert into stagetypes values(4, 'Четвертый тип');
insert into stagetypes values(5, 'Пятый тип');
-- Для 'tasktypes'
insert into tasktypes values(1, 'Первый тип задания');
insert into tasktypes values(2, 'Второй тип задания');
insert into tasktypes values(3, 'Третий тип задания');
insert into tasktypes values(4, 'Четвертый тип задания');
insert into tasktypes values(5, 'Пятый тип задания');
-- Для 'tasks'
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 1, 1, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 2, 1, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 3, 1, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 1, 2, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 2, 3, 1);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 2, 3, 3, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 1, 3, 3);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 2, 4, 2);
INSERT INTO tasks VALUES(null, 'Test', 'Test', 3, 3, 4, 2);

---- Выборка
-- For users
INSERT INTO users VALUES(NULL, 'Пользователь'); -- Добавление юзера
select * from users where UserID = 1; -- Выведем юзера с ID 1
delete from users where UserID = 1; -- Удалим юзера с ID 1
select TaskID from tasks where User = 1; -- ID заданий первого юзера

-- For Task
INSERT INTO tasks VALUES(null, 'Test', 'Test', 1, 1, 1, 1); -- Добавим тест запись
select * from tasks where TaskID = 2; -- Вывод строки с ID 2
delete from tasks where TaskID = 2; -- удалим запись с ID 2

-- For Stage
insert into stages values(10, 'Имя стадии');
delete from stages where StageID = 4; -- Удаление по ID
select * from stages where StageID = 4; -- Вывод по StageID
select Stage from tasks where TaskID = 3; -- Получение Stage от имени ID ??
UPDATE Stage SET StageName = 'Тестовое состояние' where StageID = 3 -- Обновление имени состояния

-- For TaskType
insert into tasktypes values (10, 'TaskTypeName'); -- Вставка данных
select * from tasktypes where TypeID = 4; -- Вывод по TypeID
delete from tasktypes where TypeID = 4; -- Удаление по TypeID

-- For StageType
insert into stagetypes values (10, 'StageTypeName'); -- Вставка
select * from stagetypes where TypeID = 4; -- Выборка по TypeID
delete from stagetypes where TypeID = 4; -- Удаление по TypeID