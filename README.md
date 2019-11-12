![xl-api](https://img.shields.io/badge/Kigaliz--analytics-red)[![serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)

API based on GCF

This should enable Kigaliz to develop a collection of polyglot reusable functions.
The usecase in focus for this repo is for developing API.

References : 
1. https://serverless.com/
2. GCP : https://serverless.com/framework/docs/providers/google/guide/quick-start/
3. AWS : https://serverless.com/framework/docs/providers/aws/guide/quick-start/
4. https://github.com/serverless/examples/tree/master/google-node-simple-http-endpoint

***
### Using the SLS

* Install the SLS framework 
* For Google Cloud function use the following substeps.
 * `sls create --template google-python --path xlapi-someFunctionalityName`
 * Navigate the the created folder > cd xlapi-someFunctionalityName
 * Edit the `serverless.yml` as below
```yml
service: xlapi-someFunctionalityName

provider:
  name: google
  stage: dev
  runtime: python37
  region: us-central1
  project: xl-insight-app
  credentials: ~/.gcloud/keyfile.json  ## This should be created in the first step of sls setup

plugins:
  - serverless-google-cloudfunctions

package:
  exclude:
    - node_modules/**
    - .gitignore
    - .git/**

functions:
  first:
    handler: main  ## This is the GCF entry function that will be called when the GCF is invoked
    events:
      - http: path
  # NOTE: the following uses an "event" event (pubSub event in this case).
  # Please create the corresponding resources in the Google Cloud
  # before deploying this service through Serverless
  #second:
  #  handler: event
  #  events:
  #    - event:
  #        eventType: providers/cloud.pubsub/eventTypes/topic.publish
  #        resource: projects/*/topics/my-topic
```


 * run the following command `npm install` to install dependencies and whatnots.
 * Go nuts and code the function out and deploy it using `sls deploy`
   Note: maintain `requirements.txt` using either pipenv or manually

***
