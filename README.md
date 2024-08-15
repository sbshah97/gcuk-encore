# Talk - Go Fast, Ship Faster with Encore
Repo for content for Gophercon UK 2024
# REST API Starter

This is a RESTful API Starter with a single Hello World API endpoint.

## Developing locally

When you have [installed Encore](https://encore.dev/docs/install), you can create a new Encore application and clone this example with this command.

```bash
encore app create my-app-name --example=event-world
```

## Running locally
```bash
encore run
```

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to view Encore's [local developer dashboard](https://encore.dev/docs/observability/dev-dash).

## Using the API

To see that your app is running, you can ping the API.

```bash
<<<<<<< HEAD
curl http://localhost:4000/events/1
=======
curl http://localhost:4000/event/World
>>>>>>> 6124aa6 (Commit 1: Basic project setup)
```

## Deployment

Deploy your application to a staging environment in Encore's free development cloud:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

Then head over to the [Cloud Dashboard](https://app.encore.dev) to monitor your deployment and find your production URL.


## Testing

```bash
encore test ./...
```
