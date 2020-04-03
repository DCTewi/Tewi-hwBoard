{{- define "layout" -}}
{{- template "header" . -}}

{{- if .UserInfo -}}

<br style="margin-bottom: 2px;" />
{{- $token := .Token -}}
{{- range .Tasks -}}
<!-- card -->
<div class="card" id="Post-{{.ID}}">
    <!-- card header (title) -->
    <div class="card-header">
        {{- with $x := checktime . -}}
        <b class="{{availdcolor $x}} mr-3">{{$x}}</b><span class="mr-2">{{- "TimeValue" | tr -}}:</span>
        {{- end -}}
        <span class="text-info mr-2">{{- .Start | time}} ~ {{.End | time -}}</span>
        <span class="float-right text-monospace text-muted">ID: {{.ID}}</span>
    </div>
    <!-- card body -->
    <div class="card-body">
        <!-- text here -->
        <b class="text-body mr-2">{{.Subject}}</b>
        <span class="text-body mr-2">{{.SubTitle}}</span>
        <code class="float-right">({{.FileType}})</code>

        {{- if checktime . | s2bool -}}
        <hr style="margin-bottom: 8px;" />
        <!-- submit button -->
        <button type="button" class="btn btn-light text-primary" data-toggle="modal" data-target="#SubmitModal-{{.ID}}">
            {{- "Submit" | tr -}}
        </button>
        <!-- submit modal -->
        <div class="modal fade" id="SubmitModal-{{.ID}}">
            <!-- modal content -->
            <div class="modal-dialog modal-lg">
                <div class="modal-content">
                    <!-- modal header -->
                    <div class="modal-header">
                        <h5 class="modal-title">
                            {{"Submit" | tr}} ID:
                            <span class="text-info text-monospace text-muted">{{.ID}}</span>
                        </h5>
                        <button type="button" class="close" data-dismiss="modal">&times;</button>
                    </div>

                    <!-- modal body -->
                    <div class="modal-body">
                        <div class="form-group">
                            <form action="/submit" enctype="multipart/form-data" method="POST" id="SubmitForm-{{.ID}}">
                                <div class="custom-file">
                                    <input type="file" class="custom-file-input" accept=".{{.FileType}}" name="file"
                                        id="FileUpload-{{.ID}}" />
                                    <label class="custom-file-label text-secondary"
                                        for="FileUpload-{{.ID}}">{{"SubmitPlaceHolder" | tr}}</label>
                                </div>
                                <hr>
                                <div>
                                    <button id="SubmitButton-{{.ID}}" type="submit"
                                        class="btn btn-primary mr-3">{{"Submit" | tr}}</button>
                                    <span class="text-warning">({{"SubmitWarning" | tr}})</span>
                                </div>
                                <input type="hidden" name="submitto" value="{{.ID}}">
                                <input type="hidden" id="Token-{{.ID}}" name="token" value="{{$token}}">
                                <script>$('#SubmitButton-{{.ID}}').on('click', function () {
                                        var maxsize = 30 * 1024 * 1024; //30M
                                        var errMsg = "上传的附件文件不能超过30M！";
                                        var tipMsg = "您的浏览器暂不支持计算上传文件的大小，确保上传文件不要超过30M，建议使用IE、FireFox、Chrome浏览器。";
                                        try {
                                            var fileobj = document.getElementById("FileUpload-{{.ID}}")
                                            if (fileobj.value == "") {
                                                alert("请先选择文件")
                                                return false
                                            }
                                            var filesize = fileobj.files[0].size
                                            if (filesize == -1) {
                                                alert(tipMsg)
                                                return false
                                            } else if (filesize > maxsize) {
                                                alert(errMsg)
                                                return false
                                            }
                                        } catch (e) {
                                            console.error(e)
                                        }
                                        return true
                                    })
                                    $('#FileUpload-{{.ID}}').on('change', function () {
                                        var filepath = $(this).val();
                                        filepath = filepath.replace(/\\/g, '/')
                                        console.log(filepath)
                                        var pos = filepath.lastIndexOf('/')
                                        $(this).next('.custom-file-label').html(filepath.substr(pos + 1));
                                    })</script>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        {{- end -}}
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