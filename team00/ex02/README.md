# Пояснение изменений для gRPC клиента:

## Импортирование пакетов:
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

Импортирован gorm и gorm.io/driver/postgres для работы с PostgreSQL и GORM.

## Константы:
```dsn``` — строка подключения к базе данных PostgreSQL. Замените  **dbname** и **password** host=localhost user=postgres dbname=```postgres``` password=```postgres``` port=5432 sslmode=disable``` на базу данных и пароль вашего пользователя PostgreSQL. 

Я использовал базу данных по умолчанию и пароль согласно "/etc/postgresql/16/main/pg_hba.conf" 
```
# Database administrative login by Unix domain socket
local   all             postgres                                peer
```
Аутентификация "peer" требует, чтобы имя пользователя в системе совпадало с именем пользователя в PostgreSQL.

## Структура Anomaly:

Определена структура Anomaly, которая будет использоваться для записи аномалий в базу данных.
Подключение к базе данных:

Используется ```gorm.Open``` для подключения к PostgreSQL.
Выполняется автоматическая миграция схемы базы данных для создания таблицы anomalies.
Запись аномалий:

Когда обнаруживается аномалия, создается новый объект Anomaly и записывается в базу данных с помощью db.Create.
Эти изменения добавят функциональность записи обнаруженных аномалий в базу данных PostgreSQL с использованием GORM. Убедитесь, что у вас настроен PostgreSQL и создана база данных, указанная в строке подключения dsn.

## Запуск сервера
```go build -o transmitter server.go
./transmitter
```
## Запуск клиента
```
go build -o anomaly_detector_with_db client.go
./anomaly_detector_with_db -k=3.0
```
---

# Пояснения к SQL-коду:
CREATE TABLE anomalies: Создает новую таблицу с именем anomalies.

id SERIAL PRIMARY KEY: Определяет столбец id как автоинкрементируемый и основной ключ таблицы.

session_id VARCHAR(36): Определяет столбец session_id для хранения идентификатора с типом данных VARCHAR(36), подходящим для хранения UUID.

frequency DOUBLE PRECISION: Определяет столбец frequency для хранения значения частоты с типом данных DOUBLE PRECISION.

timestamp TIMESTAMP: Определяет столбец timestamp для хранения метки времени.

---

# Пояснения к работе go test

# gRPC server_test.go

Этот проект демонстрирует использование gRPC для создания тестового сервиса, который предоставляет частотные данные через стриминг. В проекте также реализован тест для проверки работы метода стриминга частот.

## Структура проекта

- **main.go**: Основной файл, который содержит реализацию тестового gRPC сервера и тестов.
  - Реализует сервер gRPC для стриминга данных частот.
  - Включает тестовый метод для проверки корректности работы метода `StreamFrequencies`.

## Установка

1. Убедитесь, что у вас установлен Go и необходимые зависимости.

    ```sh
    go mod tidy
    ```

2. Установите необходимые пакеты для работы с gRPC:

    ```sh
    go get google.golang.org/grpc
    go get google.golang.org/grpc/test/bufconn
    ```

## Запуск тестов

Для запуска тестов используйте команду:

```sh
go test -v
```

## Описание кода

### `main.go`

1. **Инициализация сервера**:
   - В функции `init()` создается буферизованное соединение и запускается gRPC сервер. Сервер регистрирует сервис `FrequencyService` и слушает на созданном соединении.

2. **bufDialer**:
   - Эта функция используется для создания соединения с сервером через буферизованное соединение.

3. **Тест `TestStreamFrequencies`**:
   - Создает gRPC клиент и вызывает метод `StreamFrequencies`, который стримит частотные данные.
   - Получает сообщения из потока и проверяет, что частота не равна нулю.
   - Логирует полученные сообщения для проверки.

4. **Функция `TestMain`**:
   - Запускает все тесты, предварительно инициализируя сервер gRPC с помощью функции `init()`.

## Примечания

- **Bufconn Listener**: Используется для создания буферизованного соединения, что упрощает тестирование gRPC сервера без необходимости реального сетевого подключения.
- **Тестирование**: В тесте проверяется корректность работы метода стриминга и правильность получаемых данных.

---

# gRPC client_test.go

Этот проект демонстрирует использование gRPC для стриминга данных частот и их обработки в тестовой среде. Проект включает реализацию сервера gRPC, клиента и тестов, которые проверяют основную логику работы с данными частот, их запись в базу данных SQLite и проверку на аномалии.

## Структура проекта

- **main.go**: Основной файл, который содержит реализацию сервера gRPC и тестов.
- **pkg/frequency**: Пакет с сгенерированными файлами для работы с gRPC (здесь должен быть файл `frequency.pb.go`).

## Установка

1. Убедитесь, что у вас установлен Go и необходимые зависимости:

    ```sh
    go mod tidy
    ```

2. Установите необходимые пакеты для работы с gRPC и GORM:

    ```sh
    go get google.golang.org/grpc
    go get gorm.io/driver/sqlite
    go get gorm.io/gorm
    ```

## Запуск тестов

Тесты в проекте предназначены для проверки работы сервера gRPC и логики обработки данных частот. Тестовые функции используют буферное соединение для имитации взаимодействия клиента и сервера.

Для запуска тестов используйте команду:

```sh
go test -v
```

## Описание кода

### `main.go`

1. **Реализация сервера gRPC**:
   - Реализует сервер `FrequencyService`, который отправляет тестовые данные частот клиенту через метод `StreamFrequencies`.

2. **Инициализация сервера**:
   - Используется буферное соединение для создания и тестирования gRPC сервера.

3. **Функция `bufDialer`**:
   - Создает подключение к буферному серверу.

4. **Функция `TestMainFunction`**:
   - Выполняет тестирование клиента gRPC и взаимодействие с базой данных SQLite.
   - Осуществляет стриминг данных частот, вычисляет среднее значение и стандартное отклонение, проверяет наличие аномалий и записывает их в базу данных.
   - Проверяет количество записанных аномалий и выводит результат в лог.

## Примечания

- **Bufconn Listener**: Используется для создания буферного соединения, которое позволяет тестировать сервер gRPC без необходимости реального сетевого подключения.
- **SQLite**: Используется в качестве тестовой базы данных для хранения информации о аномалиях.

---