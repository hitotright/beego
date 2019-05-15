<!DOCTYPE html>
<html >
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <meta name="keywords" content="{{.Keywords}}">
    <meta name="description" content="{{.Description}}">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="format-detection" content="telephone=no">
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <link rel="alternate icon" type="image/png" href="static/i/favicon.png">
    <link rel="stylesheet" href="static/css/amazeui.min.css"/>
    {{.Css}}
</head>
<body>
<!--[if lte IE 9]>
<p class="browsehappy">你正在使用<strong>过时</strong>的浏览器，Amaze UI 暂不支持。 请 <a href="http://browsehappy.com/" target="_blank">升级浏览器</a>
    以获得更好的体验！</p>
<![endif]-->
{{.LayoutContent}}
</body>
{{.Js}}
</html>