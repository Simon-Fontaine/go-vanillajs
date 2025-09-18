import API from "./services/API.js";
import "./components/HomePage.js";

window.app = {
  search: (event) => {
    event.preventDefault();
    const query = document.querySelector("input[type=search]").value;
  },
  api: API,
};
