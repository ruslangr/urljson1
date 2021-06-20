resource "aws_db_instance" "rds_db" {
  identifier             = "rdsdb"
  instance_class         = "db.t2.micro"
  allocated_storage      = 5
  engine                 = "postgres"
  engine_version         = "12.5"
  username               = var.dbuser
  password               = var.dbpass
  db_subnet_group_name   = aws_db_subnet_group.ruslangr-net-gr.name
  vpc_security_group_ids = [aws_security_group.worker_group.id]
  #parameter_group_name   = aws_db_parameter_group.education.name
  publicly_accessible    = true
  skip_final_snapshot    = true
}