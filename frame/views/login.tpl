{{- define "layout" -}}
{{- template "header" . -}}

<div class="mt-3" id="PageHead">
<h2>{{"UserLogin" | tr}}</h2>
{{if .Message}}
<div class="alert alert-danger alert-dismissible fade show mt-3" role="alert">
    <strong>{{"Error" | tr}}:</strong> {{.Message -}}
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
        <span aria-hidden="true">&times;</span>
    </button>
</div>
{{end}}
</div>
<script src="/js/util.js"></script>
<script>function checkForm(){return(checkbyid('loginform','userq',/^[0-9]{5,15}$/g)&&checkbyid('loginform','stuid',/^20182410[0-9]{4}$/g)&&checkbyid('loginform','cikey',/^[0-9a-zA-Z_=]{6}$/g));}</script>
<form action="/login" method="POST" id="LoginForm" name="loginform" onsubmit="return checkForm()">
    <div class="input-group mt-3 mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"Email" | tr}}:</span>
        </div>
        <input type="text" name="userq" id="userq" class="form-control" placeholder="{{"EmailPlaceHolder" | tr}}">
        <div class="input-group-append">
            <span class="input-group-text">@{{"EmailConstr" | tr}}</span>
        </div>
    </div>
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"StuID" | tr}}:</span>
        </div>
        <input type="text" id="stuid" name="stuid" class="form-control" placeholder="{{"StuIDPlaceHolder" | tr}}">
    </div>
    <input type="hidden" name="token" value="{{.Token}}">
    <div class="input-group mb-3">
        <div class="input-group-prepend">
            <span class="input-group-text">{{"ConfirmKey" | tr}}:</span>
        </div>
        <input type="text" id="cikey" name="cikey" class="form-control" placeholder="{{"ConfirmKeyPlaceHolder" | tr}}">
        <div class="input-group-append">
            <button class="btn btn-outline-secondary" type="button" id="QueryCKey">{{"SendConfirmKey" | tr}}</button>
        </div>
    </div>
    <hr>
    <div>
        <button type="submit" class="btn btn-primary mb-2 mr-2">{{"UserLogin" | tr}}</button>
        <span class="text-warning">({{"LoginWarning" | tr}})</span>
    </div>
</form>
<script>$(function(){$("#QueryCKey").click(function(){if(checkbyid('loginform','userq',/^[0-9]{5,15}$/g)&&checkbyid('loginform','stuid',/^20182410[0-9]{4}$/g)){var qq=document.forms["loginform"]["userq"].value;var id=document.forms["loginform"]["stuid"].value;$.ajax({type:'POST',url:'/login?q=ckey',data:{"userq":qq,"stuid":id},success:function(data){console.log(data);if(data!=""){$("#PageHead").append('<div class="alert alert-danger alert-dismissible fade show mt-3"role="alert"><strong>错误:</strong>'+data+'<button type="button"class="close"data-dismiss="alert"aria-label="Close"><span aria-hidden="true">×</span></button></div>')}}});$("#QueryCKey").attr("disabled",true);$("#QueryCKey").addClass("disabled")}})})</script>

{{- template "footer" . -}}
{{- end -}}