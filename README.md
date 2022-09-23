# Library Service

### Run Application with Docker

for build image, run:
```bash
docker build . -t ${IMAGE_TAG_NAME} -f Dockerfile

 ```
for push image, run:
```bash
docker push ${IMAGE_TAG_NAME}
 ```

for run application, run:
```bash
docker compose up -d && docker compose start 
 ```

for stop application, run:
```bash
docker compose stop
 ```

for stop and clean application, run:
```bash
docker compose down --rmi local --remove-orphans -v 
 ```







