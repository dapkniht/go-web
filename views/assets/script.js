const time = document.getElementById("time");
function showCurrentTime() {
  let date = new Date();
  let hr = date.getHours();
  let min = date.getMinutes();
  let sec = date.getSeconds();
  time.textContent = `${hr}:${min}:${sec}`;
}
setInterval(showCurrentTime, 1000);
