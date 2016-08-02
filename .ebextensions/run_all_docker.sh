docker run aws_beanstalk/staging-app go run email_recipient_main.go &&
docker run aws_beanstalk/staging-app go run low_inventory_main.go &&
docker run aws_beanstalk/staging-app go run daily_sales_main.go