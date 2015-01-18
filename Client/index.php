<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
  <script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.7/jquery.min.js"></script>
    <script type="text/javascript" src="scripts/jquery.rpc.js"></script>
     <script type="text/javascript">
        var user = new $.JsonRpcClient({
            //ajaxUrl : 'http://127.0.0.1:8080/jsonrpc',
            socketUrl: 'ws://127.0.0.1:8080/jsonrpc',
            onmessage : function(result){
               var results = JSON.parse(result.data);
    
                console.log(result.result);
            },
            onopen : function(o){
                console.log("opened");
            },
            onclose : function(c){
                console.log("closed");
            },
            onerror : function(e){
                console.log(e);
            },
        });

        function cb_success(data){
            console.log(data);
        }
        function Profile(){
            user.call(
                'UserMethod.Profile', [{"ID":"316a743e-49bf-491d-988a-7c60cd227f2d"}],cb_success, function(){
                    console.log("loi");
                }
      
            );
        }
        function Create(){
            user.call(
                'UserMethod.CreateUser', [{"Email": $("#email").val(), "Password" : $("#password").val()}]
            );
        }
        function List(){
            user.call(
                'UserMethod.ListUser', []
            );
        }
      
    
    </script>
<title>Insert title here</title>
</head>
<body>

<fieldset>
	<legend>Simple JSON-RPC - Plus</legend>
    <div style="margin-bottom:10px;">
        <input type="text" name="email" id="email" placeholder="email" />
        <input type="password" name="password" id="password" placeholder="password" />
    </div>
    <input type="button" value="Profile" onclick='Profile()'>
    <input type="button" value="Create" onclick='Create()'>
    <input type="button" value="List" onclick='List()'>
</fieldset>

<fieldset>
    <legend>List User</legend>
    <table>
        <th>ID</th>
        <th>Email</th>
        <th>Password</th>
        <tr>
            <td></td>
            <td></td>
            <td></td>
        </tr>
    </table>
</fieldset>
</body>
</html>