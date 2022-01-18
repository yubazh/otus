resource "yandex_compute_instance" "wp-app-1" {
  zone = "ru-central1-a"
  count = "${var.count_instance_a}"
  name = "wp-app-1-${(count.index)}"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd80viupr3qjr5g6g9du"
    }
  }

  network_interface {
    # Указан id подсети default-ru-central1-a
    subnet_id = yandex_vpc_subnet.wp-subnet-a.id
    nat       = true
  }

  metadata = {
    ssh-keys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
  }
}


resource "yandex_compute_instance" "wp-app-2" {
  zone = "ru-central1-b"
  count = "${var.count_instance_b}"
  name = "wp-app-2-${(count.index)}"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd80viupr3qjr5g6g9du"
    }
  }

  network_interface {
    # Указан id подсети default-ru-central1-b
    subnet_id = yandex_vpc_subnet.wp-subnet-b.id
    nat       = true
  }

  metadata = {
    ssh-keys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
  }
}

resource "yandex_compute_instance" "wp-app-3" {
  zone = "ru-central1-c"
  count = "${var.count_instance_c}"
  name = "wp-app-3-${(count.index)}"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd80viupr3qjr5g6g9du"
    }
  }

  network_interface {
    # Указан id подсети default-ru-central1-c
    subnet_id = yandex_vpc_subnet.wp-subnet-c.id
    nat       = true
  }

  metadata = {
    ssh-keys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
  }
}  