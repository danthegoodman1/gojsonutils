# gojsonutils

Go version of https://github.com/danthegoodman1/DansJSONUtils

Example:

```json
{
  "hey": "ho",
  "lets": 1,
  "arr": [1,2],
  "obj": {
    "this": "key",
    "val": [1,2]
  },
  "objar": [
    {
      "a": "val",
      "b": 2,
      "c": [1],
      "d": [{"e": 1}, {"e": 2}],
      "dstr": [{"e": 1}, {"e": 2}, 2],
      "darr": [{"e": 1}, {"e": 2}, [2]],
      "f": {
        "g": 1,
        "h": [2]
      }
    },
    {
      "a": "val2",
      "new": "nvale"
    }
  ],
  "nestedar": [[1], [2]]
}
```

flattens to:

```json
{
  "hey": "ho",
  "lets": 1,
  "arr": [
    1,
    2
  ],
  "obj__this": "key",
  "obj__val": [
    1,
    2
  ],
  "objar__a": [
    "val",
    "val2"
  ],
  "objar__b": [
    2,
    null
  ],
  "objar__c": [
    [
      1
    ],
    null
  ],
  "objar__d__e": [
    [
      1,
      2
    ],
    null
  ],
  "objar__dstr": [
    "[{\"e\":1},{\"e\":2},2]",
    null
  ],
  "objar__darr": [
    "[{\"e\":1},{\"e\":2},[2]]",
    null
  ],
  "objar__f__g": [
    1,
    null
  ],
  "objar__f__h": [
    [
      2
    ],
    null
  ],
  "objar__new": [
    null,
    "nvale"
  ],
  "nestedar": [
    [
      1
    ],
    [
      2
    ]
  ]
}
```
