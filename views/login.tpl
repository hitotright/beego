<div class="header">
    <div class="am-g">
        <div class="am-u-lg-offset-2 am-u-lg-2 am-u-sm-centered">
            <h1>登录</h1>
        </div>
    </div>
    <hr />
</div>
<div class="am-g">
    <div class="am-u-lg-6 am-u-md-8 am-u-sm-centered">
        <br>
        <form method="post" class="am-form" >
            <label for="login_name">登录名:</label>
            <input type="text" name="login_name" id="login_name" value="">
            <br>
            <label for="password">密码:</label>
            <input type="password" name="password" id="password" value="">
            <br>
            <label for="remember-me">
                <input id="remember-me" type="checkbox">
                记住密码
            </label>
            <br />
            <div class="am-cf">
                <input type="button" onclick="ajaxLogin()" value="登 录" class="am-btn am-btn-primary am-btn-sm am-fl">
                <a target="_blank" class="am-btn am-btn-default am-btn-sm am-fr">忘记密码 ^_^? </a>
            </div>
        </form>
        <hr>
        <div class="am-btn-group">
            <a href="#" class="am-btn am-btn-secondary am-btn-sm"><i class="am-icon-github am-icon-sm"></i> Github</a>
            <a href="#" class="am-btn am-btn-success am-btn-sm"><i class="am-icon-google-plus-square am-icon-sm"></i> Google+</a>
            <a href="#" class="am-btn am-btn-primary am-btn-sm"><i class="am-icon-stack-overflow am-icon-sm"></i> stackOverflow</a>
        </div>
        <br>
    </div>
</div>
<footer>
    <hr>
    <p class="am-padding-left">© 2014 AllMobilize, Inc. Licensed under MIT license.</p>
</footer>
<script>
    function ajaxLogin() {
       var data = $('.am-form').serializeArray();
       $.post('/login',data).then(function (res) {
           if(res.status){
               window.location ='/'
           }else{

           }
       })
    }
</script>