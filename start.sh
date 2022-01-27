while ! nc -z base_service_rabbitmq 5672; do
  echo "wait for base_service_rabbitmq"
  sleep 1
done
while ! nc -z base_service_mongodb 27017; do
  echo "wait for base_service_mongodb"
  sleep 1
done
while ! nc -z base_service_influxdb 8086; do
  echo "wait for base_service_influxdb"
  sleep 1
done
while ! nc -z base_service_redis 6379; do
  echo "wait for base_service_redis"
  sleep 1
done
/usr/src/ControlCenter $1 --config /usr/src/config_dev.yaml
