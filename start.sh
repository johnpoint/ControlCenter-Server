while ! nc -z control_center_mq 5672; do
  echo "wait for control_center_mq"
  sleep 1
done
while ! nc -z control_center_database 27017; do
  echo "wait for mongodb"
  sleep 1
done
while ! nc -z control_center_redis 6379; do
  echo "wait for redis"
  sleep 1
done
/usr/src/ControlCenter api --config /usr/src/config_dev.json
