<html>
<head>
<ul id="nav">
  <li><a href="/top">Home</a></li>
  <li><a href="/profile">Profile</a></li>
  <li><a href="/post">Post</a></li>
  <li><a href="#">XSS</a></li>
  <li><a href="/logout">Logout</a></li>
<ul>
</head>
<div id="header_title">
<h1>New Post</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="./assets/css/post.css" type="text/css"> 
<body>

<div class="box">
    <div class="box11">
    {{range $i, $v := .UserPosts}}
    <p>{{.}}</p>
    <p>{{index $.Created_at $i}}</p>
    {{end}}
    {{range .Created_at}}
    <p>{{.}}</p>
    {{end}}
    </div>
    <div class="postform">
        <form action="/post" method="post">
        <textarea name="post" rows="10" cols="20"></textarea>
    </div>
        <input type="submit" value="Post" class="button_postform">
        </form>
</div>

</body>
</html>
