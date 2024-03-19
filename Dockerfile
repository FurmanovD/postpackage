ARG IMAGE_TAG

FROM ${IMAGE_TAG}

# copy the config to a container
ADD ./config/config.local.json /app/config.json

WORKDIR /app
