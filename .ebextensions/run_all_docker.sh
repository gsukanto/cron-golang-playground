echo "run get all email recipients at `date`" >> /tmp/golang.log
docker run aws_beanstalk/staging-app go run email_recipient_main.go >> /tmp/golang.log
echo "run low inventory at `date`" >> /tmp/golang.log
docker run aws_beanstalk/staging-app go run low_inventory_main.go >> /tmp/golang.log
echo "run daily sales at `date`" >> /tmp/golang.log
docker run aws_beanstalk/staging-app go run daily_sales_main.go >> /tmp/golang.log
