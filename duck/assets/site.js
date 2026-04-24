var duck,
	wisdom,
	state = 1;

document.addEventListener("DOMContentLoaded", () => {
	duck = document.getElementById("duck");
	wisdom = document.getElementById("wisdom");
});

window.addEventListener("keydown", (e) => {
	if (e.key === "r") {
		duck.src = "/assets/duck-rubber.svg";
	}
});

// Mirror and animate duck, also dispense wisdom
function dispense(e) {
	// If client has JS disabled, this falls back to an href
	e.preventDefault();
	// Mirror and animate duck
	duck.style.transform = `scaleX(${-state})`;
	duck.classList.toggle("animated");
	duck.classList.toggle("flip-animated");
	// Fetch dispensed wisdom from API
	// On slow connections, animate
	wisdom.classList.toggle("visible");
	fetch("/api/wisdom/dispense" + window.location.search ?? "")
		.then((r) => r.text())
		.then((w) => (wisdom.innerText = `"${w}"`))
		.then(() => wisdom.classList.toggle("visible"));
	// Update state
	state = -state;
}
