output "cloud_sql_instance" {
  value = "${ google_sql_database_instance.instance.connection_name}"
}