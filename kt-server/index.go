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
		<script>
			const allButtons = f => {
				for(const button of document.getElementsByTagName("button"))
					f(button);
			};

			const clickListener = event => {
				allButtons(b => b.disabled = true);

				const id = event.target.id;
				console.log("Clicked ", id);

				const xhr = new XMLHttpRequest();
				xhr.addEventListener("load", function() {
					console.log(this);
					allButtons(b => b.disabled = false);
				});
				xhr.open("POST", "/api/" + id + "/send");
				xhr.send();
			};
			allButtons(b => b.addEventListener("click", clickListener));
		</script>
	</body>
</html>`
