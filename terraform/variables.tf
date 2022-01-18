variable "yc_cloud" {
  type = string
  description = "Yandex Cloud ID"
}

variable "yc_folder" {
  type = string
  description = "Yandex Cloud folder"
}

variable "db_password" {
  description = "MySQL user pasword"
}

variable "count_instance_a" {
  description = "Number of instances in a zone"
}

variable "count_instance_b" {
  description = "Number of instances in b zone"
}

variable "count_instance_c" {
  description = "Number of instances in c zone"
}