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
<h1>TimeLine</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="./assets/css/post.css" type="text/css"> 
<body>

<div class="box">
    <div class="box11">
    {{range $i, $v := .UserPosts}}
    <div class="box22">
    {{$image := index $.UserImages $i}}
    {{if $image}}
    <nobr><img src="./assets/img/{{$image}}" width="50" height="50">
    {{else}}
    <nobr><img src="./assets/img/noimage.png" width="50" height="50">
    {{end}}
    <h2>{{.}}</h2></nobr> <br><h3>&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;{{index $.Created_at $i}}</h3>
    <br>
    </div>
    {{end}}
    </div>
    <div class="postform">
        <form action="/timeline" method="post">
        <textarea name="post" rows="10" cols="20" wrap="hard" required></textarea>
    </div>
        <input type="submit" value="Post" class="button_postform">
        </form>
</div>

</body>
</html>
