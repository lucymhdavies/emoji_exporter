# Prometheus Exporter for Twitter Emojis

![Aerial Tramway Emoji](images/aerial-tramway_1f6a1.png)

Proof of Concept, using Emojitracker APIs:

- [Emojitracker REST API](https://github.com/emojitracker/emojitrack-rest-api)
- [Emojitracker Streaming API](https://github.com/emojitracker/emojitrack-streamer-spec)



## Usage

Copy `example.env` to `.env`, and modify as necessary:

| Variable           | What it does                                                                    |
|--------------------|---------------------------------------------------------------------------------|
| `LOG_LEVEL`        | Log verbosity of the exporter. Options are: debug, info, warn, error            |

### Local

Run locally with:

```
go run *.go
```

### Docker

Launch the exporter, as well as a prometheus instance with:

```
docker-compose up -d
```

Prometheus will be available, on [http://localhost:9090](http://localhost:9090)

The docker-compose file also comes with Grafana, with a pre-configured dashboard, on [http://localhost:3000](http://localhost:3000) (username/password: admin/foobar)

![screenshot of Grafana dashboard](images/grafana.png)

An interactive version of the above screenshot is available [here](https://snapshot.raintank.io/dashboard/snapshot/p6FNOHU7kdn6jqY4uidWYg56yfbcu8dT?orgId=2)


## If you want to support me making silly things
https://ko-fi.com/lucydavinhart
