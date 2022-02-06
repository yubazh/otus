#!/bin/python

from operator import inv
import requests
import json
import subprocess
import sys

# env variable
env_template = "prod"
folder_template = "stage"

# Get token - yc iam create-token
# получаем токен с помощью команды yc iam create-token, кладем его в переменную iam_token
# редактируем строку, чтоб остался только токен
p = subprocess.Popen('yc iam create-token', shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
iam_token = str(p.stdout.readline().rstrip())
iam_token = iam_token[:-1]
iam_token = iam_token[2:]
#следующие 2 строчки для тестов
#iam_token = "t1.9euelZrIyp2el42XyZXNys6SlYydx-3rnpWaicjKnZOSipeLiszPlZrNm8nl9PcUTH9v-e8CYlLI3fT3VHp8b_nvAmJSyA.C8t-PO8v-EbPlAK8qYjHL7K1syX08PIe1M4PU4gQq3EsoYvyylMw4txdapoZcW3DrIQveU7ej8u3UYhx_9XpCw"
#print(iam_token)
# create header
headers = {'Authorization': 'Bearer ' + iam_token,'Content-Type': 'application/json'}

# get cloud id
# cloud id я внес собственноручно, т.к. он у меня не изменялся
url = 'https://resource-manager.api.cloud.yandex.net/resource-manager/v1/clouds'
response = requests.get(url, headers=headers)
result = response.json()
cloud_id = "b1g1mo7rdqq79hlp4b86"
#print(cloud_id)

# get folder id
# folder id я также внес собственноручно, т.к. он тоже у меня не изменялся
url = 'https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders?cloudId='+str(cloud_id)
response = requests.get(url, headers=headers)
result = response.json()
folder_id = "b1gi9bocbje5b41uqkje"
#print(folder_id)

# блоки работы с файлами остались после тестов вывода. попутно у меня выводится вся информация 
# в файл environments/prod/inventory, формируя нормальный инвентори в файле

f = open('environments/prod/inventory', 'w')
f.write("[wp_app]\n")

# get inventory
# обращаемся к API и получаем список инстансов
url = 'https://compute.api.cloud.yandex.net/compute/v1/instances?folderId='+str(folder_id)
response = requests.get(url, headers=headers)
result = response.json()
# print(json.dumps(result, indent=4) )
# inventory = {"_meta": {"hostvars": {}},}
# мой кривой код по рекдатированию полученных результатов
# а именно извлечение ip всех compute instance'ов
inventory = "{\n    \"wp_app\": [" # переменная для вывода json
#print(inventory)
if result:
    # get hosts with external ip, no any checks
    for instance in result["instances"]:
	# check host match by filter
	    # get app
        if "wp-app" in instance["name"]:
          inss = ([instance["networkInterfaces"][0]["primaryV4Address"]["oneToOneNat"]["address"]])
          for ind in inss:
            f.write(ind + '\n')
            inventory = inventory+ "\"" + ind + "\", "
            host = {"wp_app":{"hosts": [instance["networkInterfaces"][0]["primaryV4Address"]["oneToOneNat"]["address"]]}}
            #inventory.update(host)

# редактируем наш вывод, удаляя последние символы и добавляя закрывающие скобки
inventory = inventory[:-2]
inventory = inventory + "]\n}"
# выводим результат
print(inventory)
# закрываем файл (в который дублировали ивентори)
f.close          
         