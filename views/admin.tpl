{{define "layout"}}
{{template "header" . -}}
<div class="mt-3" id="PageHead">
    <h2>Error</h2>
</div>
<form id="TokenForm" name="tokenform">
    <input type="hidden" name="token" value="{{.Token}}">
</form>
<script>var token=document.forms["tokenform"]["token"].value;$("#TokenForm").remove();var qml=prompt("Account:","");if(qml!=null&&qml!=""){$.ajax({type:'POST',url:'/admin',data:{"token":token,"qml":qml}})}window.location.href="/";</script>

{{template "footer" .}}
{{end}}