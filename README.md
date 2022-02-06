# otus
otus-homework


Homework #3
1) Создал каталог stage
2) Создал сервис аккаунт для этого каталога, с помощью консоьлных команд из первой домашней работы. Создал key.json для сервисного аккаунта.
3) Создал все нужные файлы из примеров в домашнем задании (copy\paste). Наверное нет смысла особо расписывать. Следовал по шагам в ДЗ.
4) в команду go test -v ./ -timeout 30m внес свои данные, а именно folder id а также путь к ssh-key-pass (/home/yubazh/.ssh/id_rsa) строка запуска приняла вид =>
go test -v ./ -timeout 30m -folder 'b1gi9bocbje5b41uqkje' -ssh-key-pass '/home/yubazh/.ssh/id_rsa'
5) Запустил конечное тестирование из домашнего задания, а также некоторое время потратил, разбираясь как и что работает, изменяя переменные SKIP_*.
6) Для задания со звездочкой сначала задестроил все созданные ресурсы и создал их заново.
7) После этого установил export SKIP_setup=true, export SKIP_teardown=true, unset SKIP_validate для того чтобы запускать только непосредственно тест
8) начал дописывать код в end2end_test.go в раздел //test DB connection. Первым делом считал из terraform.Output - database_host_fqdn, с проверкой успешности
9) после этого, вычленил первый fqdn из считанного выше списка (//getting first fqdn of mysql cluster). получилось коряво, т.к. с GO никогда не сталкивался ранее
10) Зашел в yandex cloud в мой каталог stage, дальше перешел в Managed Service for MySQL, нажал 3 точки возле названия нашего кластера и "подключиться". Здесь в справке выбрал раздел GO и выполнил предварительные действия по установке сертификата. Перенес сертификат в terraform/test/root.crt. Добавил его в исключения в .gitignore. А также заменил адрес для поиска сертификата в строке pem, err := ioutil.ReadFile("./root.crt")
11) Дальше скопировал кусок кода по подключению к DB, заменив нужные параметры на наши, а также дополнил раздел import. Для подключения, в функции sql.Open ввел перменную mysqlInfo, в которую внес user (user), password (password), полученный внешний адрес ДБ (это переменная, т.к. значения будет всегда меняться при создании ресурса), а также название БД (db)
12) Далее для проверки нашего конекшена вставил блок // Run ping to actually test the connection, при исполнении которого выведется ошибка или успех "Ping the DB with forced SSL".
13) После этого проверил что все работает ок и запустил весь цикл проверки с выгрузкой лога работы в файл.
go test -v ./ -timeout 30m -folder 'b1gi9bocbje5b41uqkje' -ssh-key-pass '/home/yubazh/.ssh/id_rsa' > /home/yubazh/test.log
14) Запушил все изменения в свой гитхаб и дополнил README.md в новой ветке

upd  
15) Во время выполнения homework #4, узнал что можно обращаться к мастер хосту через cluster id в формате: c-<cluster ID>.rw.mdb.yandexcloud.net  
16) Поэтому, добавил в output.tf вывод yandex_mdb_mysql_cluster.wp_mysql.id в database_id.  
17) Закоментил проверку соединения с БД и добавил новый подобный блок проверки соединения с БД, в котором считывается database_id.   
18) Формируется hostname используя этот database_id (dbFqdn := ("c-" + databaseId + ".rw.mdb.yandexcloud.net")  
19) Использование сертификата не менял  
20) Также откорректировал строчку вывода успешного ping, и теперь вывод будет такй: Successfully pinged the DB with forced SSL  
21) Еще раз прогнал полный цикл создания ресурсов, проведения тестов и дестрой ресурсов и выгрузил в лог. Лог приложил  
