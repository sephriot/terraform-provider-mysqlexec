provider "mysqlexec" {
  endpoint = "tcp(127.0.0.1)"
  username = "root"
  password = "root"
}

resource "mysqlexec_script" "file" {
  query = "CREATE DATABASE mysqlexec"
  #file_path = "example.sql"
}
