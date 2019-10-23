<html>
<head>
<ul id="nav">
  <li><a href="/top">Home</a></li>
  <li><a href="/profile">Profile</a></li>
  <li><a href="/timeline">TimeLine</a></li>
  <li><a href="/post">Post</a></li>
  <li><a href="/hints">Hints</a></li>
  <li><a href="/bonus">Bonus</a></li>
  <li><a href="/logout">Logout</a></li>
<ul>
</head>
<div id="header_title">
<h1>Confirm OK ?</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="../../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="../../assets/img/{{.Image}}" width="400" height="300">
    <h2>{{.Word}}</h2>
    {{ else }}
    <img src="../../assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <h1>Name : <t> <i>{{.UserName}}</i></h1>
    <h2>Age : <t> {{.Age}}</h2>
    <h2>Mail : <t> {{.Mail}}</h2>
    <h2>Address : <t> {{.Address}}</h2>
    <h2>Favorite Animal : <t> {{.Animal}}</h2>
    </div>
    <form action="/profile/edit/update" method="post">
    <input type="hidden" name="username" value="{{.UserName}}">
    <input type="hidden" name="age" value="{{.Age}}"> 
    <input type="hidden" name="mail" value="{{.Mail}}">
    <input type="hidden" name="address" value="{{.Address}}">
    <input type="hidden" name="animal" value="{{.Animal}}">
    <input type="hidden" name="word" value="{{.Word}}">
    </div>
    <input type="submit" class="button_edit" value="Confirmed">
</div>




</body>
</html>
