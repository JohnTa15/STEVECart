const canvas = document.getElementById("minimap");
const ctx = canvas.getContext("2d");
const img = document.getElementById("map");

const width = canvas.width;
const height = canvas.height;

ctx.fillStyle = "black";
ctx.fillRect(0, 0, width, height);

let mouseX = 0;
