# airtable-cli

Helps create customizable airtable cli apps.

## Usage

```
airtable-cli -n ${confname} ${args}
airtable-cli -n ${confname} add alias1 value1 alias2 value2
airtable-cli -n ${confname} list ${query}
airtable-cli -n ${confname} mod ${record-id} alias1 value1 alias2 value2
```

```
# list
airtable-cli -n ${confname} list 'date > 20200600 && date < 20200700'
```

## Conf

```toml
# confname.toml
apikey = "test"
baseid = "test"
table = "t1"
view = "v1"
[fields]
[fields.f1]
name = "field1"
alias = ["alias1"] # first alias as column name
```

## Interactive Mode(TODO)

```
airtable-cli -n ${name} -i
query bar, filter after pressing enter, move focus to lists.
hjkl to navigate.
enter to mod field, enter to commit mod.
/ to go back query bar.
v to batch select.
dd to batch del.
E to batch edit.
```

## Limit

Airtable only allows 5 requests per sec.
