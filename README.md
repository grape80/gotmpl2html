# gotmpl2html
gotmpl2html is a cli tool to convert gotemplate to html.

## Usage
```sh
$ gotmpl2html input.gotmpl > output.html
```
For more information, see [usage](usage.go)

## Feature
- [x] JSON Binding

  Example

  - employeelist.json
    ```json
    {
        "title": "Employee List",
        "employees": [
            {
                "id": "001",
                "name": "James",
                "age": 20
            },
            {
                "id": "002",
                "name": "Mary",
                "age": 30
            },
            {
                "id": "003",
                "name": "Linda",
                "age": 40
            }
        ]
    }
    ```
  - input.gotmpl
    ```html
    {{data "/employeelist.json" }}
    <html>
    <head>
        <title>{{.title}}</title>
    </head>
    <body>
        {{template "/_header.partial" . }}
        <table>
            {{range .employees}}
            <tr>
                <td>{{.id}}</td>
                <td>{{.name}}</td>
                <td>{{.age}}</td>
            </tr>
            {{end}}
        </table>
        {{template "/_footer.partial" . }}
    </body>
    </html>
    ```
