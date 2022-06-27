resource "google_sql_database_instance" "instance" {
  name             = var.name
  region           = var.location
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-custom-1-4096"
    availability_type = "ZONAL"
    disk_size = "10"
    disk_type = "PD_HDD"
  }

  deletion_protection  = "true"
}

resource "google_sql_user" "database-user" {
    name = var.database_user
    instance = google_sql_database_instance.instance.name
    password = var.database_password
}