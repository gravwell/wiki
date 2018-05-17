## Example code

The following example code that logs in, grabs auth headers, sets up a websocket, performs a search query parse, and then logs out.

```javascript
process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
var q = require('q');
var https = require('https');
var WebSocket = require('ws');
var ws;

var auth = {
	"User": "admin",
	"Pass": "changeme"
};


var cookie = {};
var csrf = {};

function login(newuser) {
	console.log("Logging in");
	const options = {
		hostname: '127.0.0.1',
		port: 8080,
		path: '/api/login',
		method: 'POST'
	};

	var def = q.defer();

	var req = https.request(options, (res) => {
		res.on('data', (d) => {
			d = JSON.parse(d);
			cookie.name = d.CookieName;
			cookie.value = d.Cookie;
			csrf.name = d.CSRFName;
			csrf.value = d.CSRFToken;
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
		hostname: '127.0.0.1',
		port: 8080,
		path: '/api/logout',
		method: 'PUT'
	};

	ws.terminate();

	var req = https.request(options, (res) => {
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
	req.setHeader('Cookie', cookie.name + '=' + cookie.value);
	req.setHeader(csrf.name, csrf.value);
	req.end();
	return def.promise;
}

function upgrade() {
	console.log("Upgrading to websocket");
	var def = q.defer();
	//set auth headers
	var headers = {
		'Cookie': cookie.name + '=' + cookie.value
	};
	headers[csrf.name] = csrf.value;

	ws = new WebSocket("wss://localhost:8080/api/ws/search", {
		headers: headers
	});

	ws.on('open', () => {
		ws.send(JSON.stringify({
			Subs: ["PONG", "parse", "search", "attach"]
		}));
		// console.log("open");
	});

	ws.on('message', function(message) {
		var msg = JSON.parse(message);
		if (msg.Resp === "ACK") {
			def.resolve();
			return;
		}
		// console.log('Received: ' + message);
	});

	ws.on('close', function(code) {
		// console.log('Disconnected: ' + code);
		def.resolve();
	});

	ws.on('error', function(error) {
		// console.log('Error: ' + error);
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

login().then(upgrade).then(parse).finally(logout);
```
