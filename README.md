# Datadog Scaler

A webhook to execute a script to scale up or down

## Build it

1. Make sure you have `glide` installed
1. `glide install`
1. go build

## Running it

1. `datadog-scaler --help`
1. `./datadog-scaler`

## Using it

Currently listens to `127.0.0.1:8000` by default.

## HTTP Basic Auth

username: circleci
password: test

## Datadog Alert

`EventTitle` from Datadog webhook payload must contain `[Scale Up Trusty]`, `[Scale Up Precise]`, `[Scale Down Trusty]`, `[Scale Down Precise]` to trigger scaling
