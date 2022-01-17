provider "yandex" {
  cloud_id  = var.yc_cloud
  folder_id = var.yc_folder
  service_account_key_file = "./key.json" # путь к файлу с ключами сервисного аккаунта
  #token = var.yc_token  не нужна, т.к. используем сервисный аккаунт
}

terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
}
