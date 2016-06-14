package main

const index = `<!DOCTYPE html>

<html>
	<head>
		<title>kill-trigger</title>
	</head>
	<body>
		<div style="display:flex;">
			<div style="flex-grow: 1;">
				{{ range .commands }}
					<button id="{{.ID}}">{{.Name}}</button>
				{{ end }}
			</div>
			<div>
				<span>Last KeepAlive:</span>
				<input id="lastKeepAlive"></input>
				<span>Clients:</span>
				<input id="clientCount"></input>
			</div>
		</div>
		<pre id="console"></pre>
		<script>
			const allButtons = f => {
				for(const button of document.getElementsByTagName("button"))
					f(button);
			};

			const ajax = (method, url) => new Promise((resolve, reject) => {
				const xhr = new XMLHttpRequest();
				xhr.addEventListener("load", function() {
					resolve(this.response);
				});
				xhr.addEventListener("error", err => {
					reject(err);
				});
				xhr.open(method, url);
				xhr.send();
			});

			const clickListener = event => {
				allButtons(b => b.disabled = true);

				const id = event.target.id;
				console.log("Clicked", id);

				const url = "https://kill-trigger.herokuapp.com/api/" + id + "/send";
				ajax("POST", url).then(() => {
					console.log("OK");
					allButtons(b => b.disabled = false);
				}).catch(err => {
					console.error(err);
					allButtons(b => b.disabled = false);
				});
			};
			allButtons(b => b.addEventListener("click", clickListener));
		</script>
		<script>
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
			}

			const commands = {{ .commands }};
			const ws = new WebSocket("wss://kill-trigger.herokuapp.com/socket");
			ws.binaryType = "arraybuffer";
			ws.addEventListener("message", event => {
				const cmds = commands.filter(cmd => cmd.ID == Number.parseInt(event.data));
				let name = "unknown";
				if(cmds.length > 0) {
					if(cmds[0].ID === 0) {
						document.getElementById("lastKeepAlive").value = now();
						return;
					}

					name = cmds[0].Name;
				}
				const line = now() + "\t" + name + "(" + event.data + ")\n";
				document.getElementById("console").textContent += line;
			});
		</script>
		<script>
			window.setInterval(() => {
				ajax("GET", "https://kill-trigger.herokuapp.com/api/numUsers").then(num => {
					document.getElementById("clientCount").value = num;
				});
			}, 1000);
		</script>
	</body>
</html>`
