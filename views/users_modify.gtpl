<html>
<head>
</head>
<div id="header_title">
<h1>Profile Edit</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="../assets/img/{{.Image}}" width="400" height="300">
    <h2>{{.Word}}</h2>
    {{ else }}
    <img src="../assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <h1>Name : <t> <input type="text" name="username" value="{{.UserName}}"> </h1>
    <h2>Age : <t> <input type="text" name="username" value="{{.Age}}"> </h2>
    <h2>Mail : <t> <input type="text" name="username" value="{{.Mail}}" disabled="disable"> </h2>
    <h2>Address : <t> <input type="text" name="username" value="{{.Address}}"></h2>
    <h2>Favorite Animal : <t> <input type="text" name="username" value="{{.Animal}}"> </h2>
    </div>
</div>

</body>
</html>
