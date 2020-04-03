{{define "footer"}}
</div>
<footer id="Footer">
  <div class="jumbotron border-top text-center" id="FooterBlock">
    <p class="text-secondary small">
      Copyright <i class="fa fa-copyright"></i> 2020 DCTewi All Rights
      Reserved - dctewi@dctewi.com <br />
      Powered by <a target="_blank" href="https://github.com/dctewi/tewi-hwboard/">Tewi-hwBoard</a> <i
        class="fa fa-code"></i> <br />
      <i class="fa fa-connectdevelop"></i>
      <a href="/admin">
        {{- "AdminLogin" | tr -}}
        <i class="fa fa-unlock" aria-hidden="true"></i>
      </a>
    </p>
  </div>
</footer>
</div>
</body>

</html>
{{end}}