<!DOCTYPE html>
<html >
<head >
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.Title}}</title>
    <meta name="description" content="{{.Description}}">
    <meta name="keywords" content="{{.Keywords}}">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="format-detection" content="telephone=no">
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <link rel="icon" type="image/png" href="static/i/favicon.png">
    <link rel="apple-touch-icon-precomposed" href="static/i/app-icon72x72@2x.png">
    <meta name="apple-mobile-web-app-title" content="{{.Title}}" />
    <link rel="stylesheet" href="static/css/amazeui.min.css" />
    <link rel="stylesheet" href="static/css/admin.css">
    <link rel="stylesheet" href="static/css/app.css">
    <script src="static/js/echarts.min.js"></script>
    {{.Css}}
</head>
<body data-type="index">
<!--[if lte IE 9]>
<p class="browsehappy">你正在使用<strong>过时</strong>的浏览器，Amaze UI 暂不支持。 请 <a href="http://browsehappy.com/" target="_blank">升级浏览器</a>
    以获得更好的体验！</p>
<![endif]-->
{{.LayoutContent}}
</body>

<script src="static/js/jquery.min.js"></script>
<script src="static/js/amazeui.min.js"></script>
<script src="static/js/iscroll.js"></script>
<script src="static/js/app.js"></script>
{{.Js}}
</html>