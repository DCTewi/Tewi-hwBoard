{{- define "layout" -}}
{{- template "header" . -}}

{{- if .UserInfo -}}

<br style="margin-bottom: 2px;" />
{{- range .Historys -}}
<!-- card -->
<div class="card" id="History-{{.ID}}">
    <!-- card body -->
    <div class="card-body">
        <!-- text here -->
        <span class="text-info mr-2">{{- .Time | time}}</span>
        <span class="float-right text-monospace text-muted">{{- "SubmitTo" | tr}}: {{.SubmitTo}}</span>
        <b>MD5: </b>
        <code class="mr-2">{{.FileMD5}}</code>
    </div>
</div>
<br style="margin-bottom: 2px;" />
{{- end -}}
{{- else -}}
<br />
<h2 class="text-muted">{{- "PleaseLogin" | tr -}}</h2>
{{- end -}}

{{- template "footer" . -}}
{{- end -}}