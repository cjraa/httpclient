# Sample application that uses go generic

Sample usage:
```
    client := httpclient.Client[map[string]any]{}

    // Data will have the type of map[string]any
    data, err := client.Get(url)
```