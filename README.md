# Ciprand

Random String Generator API

### Parameters :

- len => length of the strings (default 10)
- count => number of random strings to be generated (default 1)
---

```
curl https://ciprand.p3p.repl.co/api?len=20&count=10
```

```json
{
  "Strings": [
    "05015e9007c5c4942c808",
    "18c0230fc85315417cff6",
    "226144acf040e899ad691",
    "e8875cd059d3f8f563be4",
    "178bb8144094bce6a6d22",
    "7591f21389c28838a6bc3",
    "6c0332ed81ec1e41be773",
    "2a476b18fb6ab7c19b346",
    "08bdea1f1ee6a1e8952f4",
    "1143f5d633c67a40d3a52"
  ],
  "Count": 10,
  "Length": 20
}
```



##### Environment Variables :

- PORT => Listen Port


This API uses reverse proxy for handling traffic. Make sure to pass the URL parameters.
