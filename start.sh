while ! nc -z base_service_mq 5672; do
  echo "wait for control_center_mq"
  sleep 1
done
while ! nc -z base_service_mongodb 27017; do
  echo "wait for mongodb"
  sleep 1
done
while ! nc -z base_service_influxdb 8086; do
  echo "wait for redis"
  sleep 1
done
while ! nc -z base_service_redis 6379; do
  echo "wait for redis"
  sleep 1
done
/usr/src/ControlCenter $1 --config /usr/src/config_dev.json
