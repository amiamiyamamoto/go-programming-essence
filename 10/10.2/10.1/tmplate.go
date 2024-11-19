{{define "index"}}
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Todo List</title>
<link rel="stylesheet" type="text/css" href="/static/style.css" media="all" />
</head>
<body>
<h1>TODO</h1>

<!-- エラーの一覧 -->
{{range .Errors}}
	<p class="error">{{.}}</p>
{{end}}

<!-- TODOの一覧 -->
<ul class="todo-overview-list">
{{range .Todos}}
<li>
<form action="/" method="post">
<input type="hidden" name="id" value="{{.ID}}" />
<input type="hidden" name="done" value="{{if .Done}}0{{else}}1{{end}}" />
<p class="{{if .Done}}todo-done{{end}}">{{.Content}}
	<span class="controls">
	{{if .Util}}{{FormatDateTime .Util}}{{end}}
	<input type="submit" id="update" name="update" class="update" value="{{if .Done}}未完了{{else}}完了{{end}}" />
	<input type="submit" id="delete" name="delete" class="delete" value="削除" />
	</span>
</p>
</form>
</li>
{{end}}
</ul>