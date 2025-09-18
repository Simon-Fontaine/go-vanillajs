import "./components/HomePage.js";
import "./components/AnimatedLoading.js";
import "./components/YouTubeEmbed.js";
import Router from "./services/Router.js";
import API from "./services/API.js";

window.app = {
  API,
  Router,
  showError: (
    message = "There was an error loading the page",
    goToHome = true
  ) => {
    document.querySelector("#alert-modal").showModal();
    document.querySelector("#alert-modal p").textContents = message;
    if (goToHome) app.Router.go("/");
    return;
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },
};

window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});
