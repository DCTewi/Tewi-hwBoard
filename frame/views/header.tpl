{{- define "header" -}}
<!DOCTYPE html>
<html lang="zh-cn">

<head>
  <!-- meta -->
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
  <meta name="description" content="Tewi Homeword Board" />
  <meta name="keyword" content="homework, DCTewi, tewi, hwboard, 冻葱Tewi" />
  <meta name="robot" content="index, follow" />
  <meta name="application-name" content="tewi-hwboard" />
  <meta name="author" content="dctewi@dctewi.com" />
  <!-- link -->
  <link rel="stylesheet" href="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/css/bootstrap.css" />
  <link rel="stylesheet" href="https://cdn.staticfile.org/font-awesome/4.7.0/css/font-awesome.min.css" />
  <link rel="shortcut icon" href="/img/favicon.ico">
  <link rel="stylesheet" href="/css/global.css" />
  <!-- script -->
  <script src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
  <script src="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/js/bootstrap.min.js"></script>
  <!-- title -->
  <title>{{- "Title" | tr -}}</title>
</head>

<body>
  <div class="wrapper">
    <!-- START HERE -->
    <header id="Header">
      <!-- title text -->
      <div class="jumbotron" id="HeaderContainer">
        <h1 id="TitleText">
          <a class="text-muted" href="/">{{- "Title" | tr -}}</a>
        </h1>
      </div>
      <!-- nav bar -->
      <nav class="navbar navbar-expand-md navbar-light bg-light border-bottom">

        <!-- nav bar toggler -->
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#MainNavbar">
          <span class="navbar-toggler-icon"></span>
        </button>

        <!-- nav bar items -->
        <div class="collapse navbar-collapse" id="MainNavbar">
          <!-- nav bar items list -->
          <ul class="navbar-nav mr-auto">
            <!-- genres -->
            <li class="nav-item">
              <a class="nav-link" href="/">
                {{- "TaskList" | tr -}}
              </a>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link" href="/history">
                {{- "SubmitHistory" | tr -}}
              </a>
            </li>
          </ul>
          <!-- login -->
          <div class="input-group-prepend my-2 my-lg-0">
            {{- if .UserInfo -}}
            <label class="align-middle text-center mt-2 mr-2" for="LogOutButton">Hello, {{.UserInfo.Email -}}</label>
            <button class="btn btn-danger" type="button" id="LogOutButton">
              <a class="text-light" href="/login?logout=true">{{- "UserLogout" | tr}} <i class="fa fa-sign-out"></i></a>
            </button>
            {{- else -}}
            <button class="btn btn-outline-primary" type="button">
              <a href="/login">{{- "UserLogin" | tr}} <i class="fa fa-sign-in"></i></a>
            </button>
            {{- end -}}
          </div>
        </div>
      </nav>
    </header>
    <!-- END HERE -->

    <div class="container" id="MainContainer">
      {{- end -}}