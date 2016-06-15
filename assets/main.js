// Library Functions

const ajax = (method, url, data) => new Promise((resolve, reject) => {
	const xhr = new XMLHttpRequest();
	xhr.addEventListener("load", function() {
		if(this.status !== 200)
			reject(new Error(this.statusText));
		else
			resolve(this.response);
	});
	xhr.addEventListener("error", err => {
		reject(err);
	});
	xhr.open(method, url);
	xhr.send(data);
});

const allButtons = f => {
	for(const button of document.querySelectorAll("button:first-of-type"))
		f(button);
};

const jsonSHelper = (k, v) => {
	if(typeof v === "string" && v.length > 25)
		return v.substring(0, 10) + " ... " + v.substring(v.length - 10);
	return v;
};

const leftPadder = row => {
	const str = row[0].toString();
	const num = row[1] + 1;
	if(str.length > num) return str;
	return new Array(num - str.length).join("0") + str;
};

const now = () => {
	const date = new Date();
	return [
		[date.getFullYear(), 4],
		[date.getMonth(),    2],
		[date.getDate(),     2],
	].map(leftPadder).join("-") + " " + [
		[date.getHours(),   2],
		[date.getMinutes(), 2],
		[date.getSeconds(), 2],
	].map(leftPadder).join(":");
};

// Commands

allButtons(b => b.addEventListener("click", event => {
	allButtons(b => b.disabled = true);
	let ctx = event.target;
	while(ctx.tagName !== "FIELDSET")
		ctx = ctx.parentElement;

	const data = {name: ctx.id};
	for(const e of ctx.querySelectorAll("[name]"))
		data[e.name] = e.value;

	ajax("POST", "/api/send", JSON.stringify(data)).then(res => {
		console.log(">", res);
	}).catch(err => {
		console.error("!", err);
	}).then(() => {
		allButtons(b => b.disabled = false);
	});
}));

// Agora Testing

document.getElementById("agora-test").addEventListener("click", () => {
	const code = document.querySelector("#agora [name=code]").value;
	agora(code)
		.then(value => JSON.stringify(value, null, "\t"))
		.then(value => document.getElementById("agora-console").textContent = value)
		.then(id => id, err => console.error(err));
	// TODO https://github.com/augustoroman/promise/issues/1
});

// Console

const ws = new WebSocket("wss://" + location.host + "/socket");
ws.binaryType = "arraybuffer";
ws.addEventListener("message", event => {
	console.log("<", event.data);
	const data = JSON.parse(event.data);
	if(data.name === "keepalive") {
		document.getElementById("lastKeepAlive").value = now();
	} else {
		const line = now() + "\t" + JSON.stringify(data, jsonSHelper) + "\n";
		document.getElementById("console").textContent += line;
	}
});

// clientCount

window.setInterval(() => {
	ajax("GET", "/api/numUsers").then(num => {
		document.getElementById("clientCount").value = num;
	});
}, 1000);
