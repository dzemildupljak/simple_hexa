run-background-app:
	sudo nohup ./main -p 80 > hexa_app.log 2>&1 &

read-app-logs:
	tail -f app.log

check-app-process:
	ps aux | grep your-app