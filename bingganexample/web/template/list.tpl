<html>
  <body>
    <table border=1, cellspacing=0>
      <tr>
        <th>Id</th>
        <th>名字</th>
        <th>备注</th>
        <th>是否admin</th>
        <th>操作</th>
        <th>操作</th>
      </tr>
      {{range .}}
      <tr>
        <td>{{.Id}}</td>
        <td>{{.Name}}</td>
        <td>{{.Note}}</td>
        <td>{{.Isadmin}}</td>
        <td><a href="/delete?id={{.Id}}">删除</a></td>
        <td><a href="/update?id={{.Id}}">更新</a></td>
      </tr>
      {{end}}
    </table>
    <p><a href="/add">新增</a></p>
  </body>
</html>
