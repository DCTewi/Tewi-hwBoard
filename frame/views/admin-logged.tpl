{{define "layout"}}
{{template "header" . -}}

<div class="mt-3" id="PageHead">
    <h2>Welcome, admin</h2>
    {{if .Message}}
    <div class="alert alert-danger alert-dismissible fade show mt-3" role="alert">
        <strong>{{"Error" | tr}}:</strong> {{.Message -}}
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
    {{end}}
</div>
<hr class="mb-3" />
<h3 class="mb-3">New task</h3>
<script src="/js/util.js"></script>
<script>
    function checkTask() {
        return (checkbyid('taskform', 'sub', /^\S{2,}$/g) && checkbyid('taskform', 'ttl', /^\S{2,}$/g) &&
            checkbyid('taskform', 'fmt', /^[a-zA-Z0-9]{1,}$/g) &&
            checkbyid('taskform', 'dat', /^[0-9]{4}-[0-9]{2}-[0-9]{2}$/g));
    }
</script>
<form action="/admin" method="POST" id="TaskForm" name="taskform" onsubmit="return checkTask()">
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"Subject" | tr}}:</span>
        </div>
        <input type="text" id="sub" name="sub" class="form-control" placeholder="{{"SubjectPlaceHolder" | tr}}">
    </div>
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"Subtitle" | tr}}:</span>
        </div>
        <textarea id="ttl" name="ttl" class="form-control" placeholder="{{"SubtitlePlaceHolder" | tr}}"></textarea>
    </div>
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"FileFormat" | tr}}:</span>
        </div>
        <input type="text" id="fmt" name="fmt" class="form-control" placeholder="{{"FileFormatPlaceHolder" | tr}}">
    </div>
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"Date" | tr}}:</span>
        </div>
        <input type="date" value="2020-04-01" id="dat" name="dat" class="form-control"
            placeholder="{{"DatePlaceHolder" | tr}}" required pattern="[0-9]{4}-[0-9]{2}-[0-9]{2}">
    </div>
    <input type="hidden" name="token" value="{{.Token}}">
    <hr>
    <div>
        <button type="submit" class="btn btn-primary mb-2 mr-2">{{"AdminSubmit" | tr}}</button>
        <span class="text-warning">({{"AdminWarning" | tr}})</span>
    </div>
</form>

<hr class="mb-3" />
<h3 class="mb-3">Get result of task</h3>
<div class="input-group mb-3">
    <div class="input-group-prepend">
        <span class="input-group-text">{{"SubmitTo" | tr}}:</span>
    </div>
    <input type="text" id="QueryID" class="form-control" placeholder="{{- "SubmitTo" | tr -}}" />
    <div class="input-group-append">
        <button class="btn btn-outline-primary" type="button" id="QueryFile"><a
                id="QueryFileA">{{"Submit" | tr}}</a></button>
    </div>
</div>
<script>
    var updateHref = function () {
        var r = $("#QueryID").val();
        var url = "/admin?get=" + r;
        $("#QueryFileA").attr('href', url);
    }

    $("#QueryID").on('change', updateHref);
    $("#QueryFile").hover(updateHref)
</script>

{{template "footer" .}}
{{end}}