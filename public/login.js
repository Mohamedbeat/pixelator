function isLogged() {
  const strdata = localStorage.getItem("user");
  const data = JSON.parse(strdata);
  if (data) {
    window.location.href = "/";
  }
}
isLogged();

var user = {
  name: "",
  color: "",
};

const handleColorChange = (input) => {
  console.log(input.value);
};

const handleLogin = () => {
  const name = document.getElementById("name").value;
  const color = document.getElementById("color").value;

  const data = {
    name,
    color,
  };
  if (!data.name || !data.color) {
    alert("please fill all infos");
    return;
  }

  const strData = JSON.stringify(data);

  localStorage.setItem("user", strData);

  console.log(name, color);

  window.location.href = "/";
};
