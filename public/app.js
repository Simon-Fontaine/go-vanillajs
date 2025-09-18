import API from "./services/API.js";
import "./components/HomePage.js";
import "./components/AnimatedLoading.js";
import "./components/YouTubeEmbed.js";

window.app = {
  search: (event) => {
    event.preventDefault();
    const query = document.querySelector("input[type=search]").value;
  },
  api: API,
};
