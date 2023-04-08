# tsarka-test-tasks

Тестовое задание на позицию junior Golang developer в ЦАРКА (bugbounty)

# доступные эндпоинты

1) POST /rest/substr/find — endpoint для нахождения заданной подстроки. Строка находится
в теле HTTP запроса.

2) POST /rest/email/check — endpoint, которая анализирует тело HTTP запроса и выдает все
найденные email адреса.

# run locally
```bash
https://github.com/Levap123/tsarka-test-tasks.git 
```
```bash
cd tsarka-test-tasks
```

```bash
make go
```

Либо докер 
```bash
make build && make run
```
Далее можно открыть сваггер, (ссылка в терминале), либо http://localhost:8080/swagger
# how to test

для запуска тестов

```bash
make tests
```