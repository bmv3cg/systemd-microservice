Systemd-microservice
====================

Systemd microservice api to manage a systemd unit which is running in user space. Normally systemd services are configured with root user account and requires sudo permissions to restart. Systemd microservice uses systemd with user configuration and enables user to manage systemd service with minimal permissions. You can find details of configuring systemd as user [here](https://wiki.archlinux.org/index.php/Systemd/User)

Configuration
-------------

You can configure the systemd service which is managed by micorservice in this [config file](config.yaml) 

Service health
--------------

You can check the health of systemd service using the following get reqeust. Replace systemdd-service.local with your hostname.

```json
curl  -H "Accept: application/json"  https://systemd-service.local:8080/api/v1/service/health
{
  "statuscode": 200,
  "unitname": "systemd-test-user.service",
  "unitstatus": "systemd unit running"
} 
```

Starting service
----------------

You can start the systemd service using the following post reqeust. Replace systemdd-service.local with your hostname.

```json
curl -X POST -H "Accept: application/json"  https://systemd-service.local:8080/api/v1/service/start  
{
  "statuscode": 200,
  "unitname": "systemd-test-user.service",
  "unitstatus": "systemd unit started"
```

Stopping service
----------------

You can stop the systemd service using the following post reqeust. Replace systemdd-service.local with your hostname.

```json
curl -X POST -H "Accept: application/json"  https://systemd-service.local:8080/api/v1/service/stop 
{
  "statuscode": 200,
  "unitname": "systemd-test-user.service",
  "unitstatus": "systemd unit stopped"
}
```

Restarting service
----------------

You can restart the systemd service using the following post reqeust. Replace systemdd-service.local with your hostname. 

```json
curl -X POST -H "Accept: application/json"  https://systemd-service.local:8080/api/v1/service/restart
{
  "code": 200,
  "unitname": "systemd-test-user.service",
  "unitstatus": "systemd unit restarted"
} 