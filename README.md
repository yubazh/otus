# otus
otus-homework

Вторая домашняя работа otus.
1) создал каталог otus в яндекс клауд
2) создал репозиторий otus в своем github https://github.com/yubazh/otus
3) через заявку увеличил квоту на сеть
4) создал директорию terraform и поместил в нее все файлы из домашнего задания: db.tf, lb.tf, network.tf, output.tf, provider.tf, variables.tf, wp-app.tf, wp.auto.tfvars
5) дополнительно создал файл .gitignore и добавил в него исключения файла key.json, содержащего данные для работы с сервисным аккаунтом, поэтому в файле provider.tf я на него ссылаюсь (./key.json) но в репозиторий github его нет
6) в файле provider.tf закомментил token и добавил service_account_key_file = "./key.json" для работы через сервисный аккаунт
7) в файле variables.tf удалил переменную yc_token и добавил три переменные count_instance_a, count_instance_b, count_instance_c, каждая из которых, отвечает за количество поднимаемых инстансов в каждой из зон
8) в файле wp.auto.tfvars удалил переменную yc_folder и добавил значения трех переменных count_instance_a, count_instance_b, count_instance_c
9) для задания со звездочкой добавил сервисный аккаунт через командную строку и сформировал файл с его ключом -> key.json. команды взял из прошлого домашнего задания
10) в файл network.tf изменения не вносил
11) в файле wp-app.tf добавил в каждый из ресурсов строчку с количество создаваемых инстансов count = "${var.count_instance_a}", а также изменил имя, чтобы в имени была отражена нумерация конкретного инстанса конкретной зоны name = "wp-app-1-${(count.index)}" (wp-app-1, wp-app-2, wp-app3)
12) в файле lb.tf закомментил блок со старыми target'ами и добавил динамичные блоки для каждого из target каждого subnet'a. в блоке через конструкцию for_each проходимся по каждой созданной копии инстанса ---.wp-app-1.*.--- (wp-app-1, wp-app-2, wp-app3) и в address помещаем значение "network_interface.0.ip_address" каджого инстанса
13) описание решения задачи с двумя звездочками выше
14) в файлы bd.tf и output.tf изменений не вносил
15) проверил в дашборде визуально созданные ресурсы, через ssh ubuntu@instance-public-ip проверил подключение к хостам.
16) после чего удалил созданные ресурсы terraform destroy --auto-approve
