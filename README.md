# grafana2UFO - Grafana webhook Alert to Dynatrace UFO

## UFO? WTF

I have a [Dynatrace UFO](https://github.com/Dynatrace/ufo/tree/master/quickstart). Yes, really. I'm borrowing this one from work but I plan on adding a (modified) 3D print version to my collection soon. I love the versatility of this device - you can radiate a lot of information from the small package.

I've always been a fan of physical status indicators. I know in a 100% remote world it makes less sense but even at home I like the idea that if I'm not paying attention to Slack that I'll see a physical light go on if something is wrong with a project.

The only problem with the UFO is that it isn't compatible with Grafana or Prometheus alerts (webhooks) out of the box. I was contemplating re-writing the firmware but decided a translation layer was probably easier.

## Webhook to UFO

All this app does it take a webhook call from Grafana and send a command to the Dynatrace UFO API. Nothing fancy. Quick and dirty implementation that I may expand upon someday.

## Go

```shell
go run main.go
```

## Docker

```shell
docker build . -t grafana2ufo
```

```shell
docker run -p 5022:5022 grafana2ufo
```

## Testing

```shell
cd examples
sh send_alerts.sh
sh send_recovery.sh
cd ..
```
