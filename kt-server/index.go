package main

const index = `<!DOCTYPE html>

<html>
	<head>
		<title>kill-trigger</title>
	</head>
	<body>
		{{ range .commands }}
			<button id="{{.ID}}">{{.Name}}</button>
		{{ end }}
		<pre id="console"></pre>
		<script>
			const allButtons = f => {
				for(const button of document.getElementsByTagName("button"))
					f(button);
			};

			const clickListener = event => {
				allButtons(b => b.disabled = true);

				const id = event.target.id;
				console.log("Clicked", id);

				const xhr = new XMLHttpRequest();
				xhr.addEventListener("load", function() {
					console.log(this.statusText);
					allButtons(b => b.disabled = false);
				});
				xhr.open("POST", "https://kill-trigger.herokuapp.com/api/" + id + "/send");
				xhr.send();
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

			const commands = {{ .commands }};
			const ws = new WebSocket("wss://kill-trigger.herokuapp.com/socket");
			ws.binaryType = "arraybuffer";
			ws.addEventListener("message", event => {
				const cmds = commands.filter(cmd => cmd.ID == Number.parseInt(event.data));
				let name = "unknown";
				if(cmds.length > 0) {
					name = cmds[0].Name;
				}
				const now = new Date();
				const line = [
					[now.getFullYear(), 4],
					[now.getMonth(),    2],
					[now.getDate(),     2],
				].map(leftPadder).join("-") + " " + [
					[now.getHours(),   2],
					[now.getMinutes(), 2],
					[now.getSeconds(), 2],
				].map(leftPadder).join(":") + "\t" + name + "(" + event.data + ")\n";
				document.getElementById("console").textContent += line;
			});
		</script>
	</body>
</html>`
