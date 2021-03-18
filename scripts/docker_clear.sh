#!/bin/sh
docker stop watermark_container
docker rm watermark_container
docker image rm watermark_images