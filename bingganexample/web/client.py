import json
import urllib2

resp = urllib2.urlopen("http://127.0.0.1:8090/users")
data = json.loads(resp.read())
for user in  data['data']:
    print user["id"], user["name"], user["note"], user["isadmin"]
