# PagerDuty Service for Cloud Foundry

The PagerDuty Service for Cloud Foundry are a series of Cloud Foundry applications and service brokers that provide users with helpful tools to integrate their Cloud Foundry environment with PagerDuty.

The PagerDuty Service Broker for Pivotal Cloud Foundry can be found on [Pivotal Network](https://network.pivotal.io/products/pagerduty-service-broker).

## Service Broker

The PagerDuty Service Broker allows operations managers to provide their application developers with access to different documentation available that will allow application developers to integrate their CF apps with PagerDuty.

### Installation

1. Log in to your deployment and target your desired `ORG` and `SPACE` for the PagerDuty Service Broker:

```
cf login
cf target -o ORG -s SPACE
```

1. Change directories to the service broker's directory:

```
cd /path/to/cf-pagerduty-service/servicebroker
```

1. Push the service broker application to Cloud Foundry:

```
cf push
```

1. Register your service broker using the `USERNAME` and `PASSWORD` for your broker:

```
cf create-service-broker SERVICE_BROKER_NAME USERNAME PASSWORD APP_URI
```

1. Create an instance of the PagerDuty service:

```
cf create-service p-pagerduty standard SERVICE_NAME
```

1. Bind the PagerDuty service instance to an app:

```
cf bind-service APP_NAME SERVICE_NAME
```

1. Restage the app for the service binding to take effect:

```
cf restage APP_NAME
```

1. The `VCAP_SERVICES` environment variable of your app now contains links to documentation to help you to connect the app to PagerDuty. The documentation is also available directly on the [PagerDuty website](https://www.pagerduty.com/integrations/).

## PagerDuty API

The PagerDuty API is a Cloud Foundry application that will expose an API to your application developers that they can use to easily produce PagerDuty events. This API is hosted on Cloud Foundry, making it easy for you to manage and audit any events sent by your application developers to PagerDuty. Furthermore, the API uses basic authentication to ensure that users accessing your API are trusted, internal developers.

### Installation

1. Log in to your deployment and target your desired `ORG` and `SPACE` for the PagerDuty API:

```
cf login
cf target -o ORG -s SPACE
```

1. Change directories to the app's directory:

```
cd /path/to/cf-pagerduty-service/pagerdutyapi
```

1. Push the application to Cloud Foundry:

```
cf push
```

## Contributing

1. Fork the GitHub repo.

1. Create your feature branch:

```
git checkout -b my-new-feature
```

1. Commit your changes:

```
git commit -m 'Add some feature'
```

1. Push the branch:

```
git push origin my-new-feature
```

1. Create a new Pull Request

## License

[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)
