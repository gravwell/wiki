## Example code

The following example code that logs in, grabs auth headers, sets up a websocket, performs a search query parse, and then logs out.

```javascript
process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
var q = require('q');
var http = require('http');
var WebSocket = require('ws');
var ws;

var auth = {
	"User": "admin",
	"Pass": "changeme"
};


var jwt = "";

function login(newuser) {
	console.log("Logging in");
	const options = {
		hostname: '172.19.0.2',
		port: 80,
		path: '/api/login',
		method: 'POST'
	};

	var def = q.defer();

	var req = http.request(options, (res) => {
		res.on('data', (d) => {
			d = JSON.parse(d);
			jwt = d.JWT;
			def.resolve();
		});
	});

	req.on('error', (e) => {
		console.error(e);
	});

	req.write(JSON.stringify(auth));
	req.end();

	return def.promise;
}

function logout(msg) {
	console.log("Logging out");
	var def = q.defer();
	const options = {
		hostname: '172.19.0.2',
		port: 80,
		path: '/api/logout',
		method: 'PUT'
	};

	ws.terminate();

	var req = http.request(options, (res) => {
		res.on('end', () => {
			console.log('logged out', res.statusCode);
			def.resolve(msg);
		});
		res.on('error', (err) => {
			console.error("could not log out!");
			def.reject(err);
		});
	});

	//set auth headers
	req.setHeader("Authorization", "Bearer "+jwt);
	req.end();
	return def.promise;
}

function upgrade() {
	console.log("Upgrading to websocket");
	var def = q.defer();
	//set auth headers
	ws = new WebSocket("ws://172.19.0.2:80/api/ws/search", jwt)

	ws.on('open', () => {
		ws.send(JSON.stringify({
			Subs: ["PONG", "parse", "search", "attach"]
		}));
	});

	ws.on('message', function(message) {
		var msg = JSON.parse(message);
		if (msg.Resp === "ACK") {
			def.resolve();
			return;
		}
	});

	ws.on('close', function(code) {
		console.log('Disconnected: ' + code);
		def.resolve();
	});

	ws.on('error', function(error) {
		console.log('Error: ' + error);
		def.reject();
	});

	return def.promise;
}

function parse() {
	var searchString = "grep foo";
	console.log("Checking if query is good:", searchString);
	var def = q.defer();
	var json = JSON.stringify({
		type: 'parse',
		data: {
			SearchString: searchString
		}
	});

	ws.on('message', function(msg) {
		msg = JSON.parse(msg);
		if (msg.type !== 'parse') {
			return;
		}
		if (msg.data.GoodQuery === true) {
			console.log("Query OK!");
			def.resolve();
		} else {
			console.log("Query bad!");
			def.reject();
		}
	});

	ws.send(json);
	return def.promise;
}

login().then(upgrade).then(parse).catch(console.log).finally(logout);
```
