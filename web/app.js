const form = document.getElementById("login-form");
const message = document.getElementById("message");

form.addEventListener("submit", (event) => {
	event.preventDefault();
	
	const username = document.getElementById("username").value;
	const password = document.getElementById("password").value;
	
	fetch("/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ username, password }),
	})
		.then((response) => {
			if (response.ok) {
				message.innerHTML = "登录成功";
			} else {
				message.innerHTML = "登录失败";
			}
		})
		.catch((error) => {
			console.error(error);
		});
});
