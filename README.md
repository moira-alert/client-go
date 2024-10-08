# Go client for Moira

If you're new here, better check out our main [README](https://github.com/moira-alert/moira/blob/master/README.md)

## Overview

This API client was generated by the [go-swagger](https://github.com/go-swagger/go-swagger) tool

Our current documentation is available on [swaggerhub](https://app.swaggerhub.com/apis/Moira/moira-alert/master).

## Installation

Add the following in import:

```golang

import "github.com/moira-alert/client-go"

```

## Getting started

Initialize Moira client:

```golang
import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/moira-alert/client-go/pkg/client"
)

func main() {
	config := client.DefaultTransportConfig().
		WithBasePath("api").
		WithHost("moira.example.com").
		WithSchemes([]string{"https"})

	transport := httptransport.New(config.Host, config.BasePath, config.Schemes)
	transport.Transport = newTransport() // Write a function here that creates an object that satisfies the http.RoundTripper interface

	client := client.New(transport, strfmt.Default)
}
```

## Examples

## Trigger

**Get all Triggers:**
```golang
client := client.New(transport, strfmt.Default)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

response, err := client.Trigger.GetAllTriggers(&trigger.GetAllTriggersParams{Context: ctx})

if err != nil {
	log.Fatal(err)
}

for _, trigger := range response.Payload.List {
	log.Println(trigger)
}
```

## Contact

**Create new contact:**
```golang
client := client.New(transport, strfmt.Default)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

contactDto := models.DtoContact{
	Type: "telegram",
	Value: "@moira_team",
}

response, err := client.Contact.CreateNewContact(&contact.CreateNewContactParams{
	Contact: &contactDto,
	Context: ctx,
})

if err != nil {
	log.Fatal(err)
}

log.Println(response.Payload)
```

## Subscription

**Update existing subscription:**
```golang
client := client.New(transport, strfmt.Default)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

subscriptionDto := &models.DtoSubscription{
	Tags: []string{"test"},
	Contacts: []string{"9fe937b0-721f-4b40-b849-ad1623f910a1"},
	Enabled: false,
}

response, err := client.Subscription.UpdateSubscription(&subscription.UpdateSubscriptionParams{
	Subscription: subscriptionDto,
	SubscriptionID: "2438f803-b278-47ee-91fb-e4345208dfde",
	Context: ctx,
})

if err != nil {
	log.Fatal(err)
}

log.Println(response.Payload)
```
