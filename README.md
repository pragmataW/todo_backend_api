# todo_backend_api
<h1>Register:</h1>
<h3>endpoint =</h3> /register <br>
<h3>method =</h3> post <br>
<h3>json format =</h3> {
   "username": "pragmatas", <br>
   "password": "123456789",<br>
   "name": "yusuf",<br>
   "surname": "ciftci"<br>
}<br>

<hr>

<h1>Login:</h1>
<h3>endpoint = </h3>/login <br>
<h3>method =</h3> post <br>
<h3>json format =</h3> {
    "username": "pragmatas",<br>
    "password": "123456789"<br>
}<br>

<hr>

<h1>Logout:</h1>
<h3>endpoint = </h3>/logout<br>
<h3>method =</h3> get <br>
<h3>Need auth</h3> <br>

<hr>

<h1>Add Todo:</h1>
<h3>endpoint = </h3>/todos<br>
<h3>method =</h3> post <br>
<h3>Need auth</h3> <br>
<h3>json format =</h3>{
    "username": "pragmatas",<br>
    "title":"Deneme",<br>
    "content":"Deneme yazısı hihiha"<br>
}

<h1>Get user's all todos:</h1>
<h3>endpoint = </h3>/todos<br>
<h3>method =</h3> get <br>
<h3>Need auth</h3> <br>

<hr>

<h1>Get user's specific todo:</h1>
<h3>endpoint = </h3>/todos/id<br>
<h3>method =</h3> get <br>
<h3>Need auth</h3> <br>

<hr>

<h1>Update Content:</h1>
<h3>endpoint = </h3>/todos/id?content=new content<br>
<h3>method =</h3> patch <br>
<h3>Need auth</h3> <br>

<hr>

<h1>Update title:</h1>
<h3>endpoint = </h3>/todos/id?title=new content<br>
<h3>method =</h3> patch <br>
<h3>Need auth</h3> <br>

<hr>

<h1>Delete todo:</h1>
<h3>endpoint = </h3>/todos/id<br>
<h3>method =</h3> delete <br>
<h3>Need auth</h3> <br>
