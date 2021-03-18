#!/bin/sh
docker build -t watermark_images . -f images/watermark/Dockerfile
docker run -d  -p 3333:8081 --name watermark_container watermark_images