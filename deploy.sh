GOOS=linux go build

scp -i /Users/tomek/.ssh/aws-snake.pem snakehero ec2-user@50.112.142.158:.

