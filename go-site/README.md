TODO - Name of the chart
========================

Overview
--------

TODO - Give an overview of what the chart does

Configuration
-------------

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `example.password` | Set as the `PASSWORD` env var; This is an example password for the TODO app. | `''` |
| `example.username` | Set as the `USERNAME` env var; This is an example username for the TODO app. | `''` |
| `image.name` | The name of the docker image to install | `Docker_Image_name` |
| `image.tag` | The tag of the image to install | `Docker_image_tag` |
| `replica.max` | The maximum number of instances to run when auto scaling | `5` |
| `replica.min` | The minimum number of instances to run when auto scaling | `1` |
| `replica.targetCPUUsage` | The target CPU usage for the Horizontal Pod Autoscaler | `75` |
| `resources` | Set the resource requests and limits for the pod [Compute Resources Guide](https://kubernetes.io/docs/user-guide/compute-resources/) | `{requests: {cpu: "100m", memory: "128Mi"}, limits: {cpu: "200m", memory: "256Mi"}}` TODO - does this match the defaults in values.yaml? |
| `service.externalPort` | The port the service will listen on | `80` |
| `service.internalPort` | The port the TODO app running in the container will listen on | `8080` |

