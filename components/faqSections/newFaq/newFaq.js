document.addEventListener("DOMContentLoaded", function () {
  const toggleButtons = document.querySelectorAll(".qa-toggle-button");

  toggleButtons.forEach(function (button) {
    button.addEventListener("click", function () {
      const parentItem = this.closest(".qa-item");
      const answer = parentItem.querySelector(".qa-answer");
      const toggleIcon = this.querySelector(".qa-toggle-icon-svg");

      if (answer) {
        // Toggle the "hidden" class for the answer
        answer.classList.toggle("hidden");

        // Toggle the button icon
        const iconPath = toggleIcon.querySelector("path");
        if (answer.classList.contains("hidden")) {
          iconPath.setAttribute("d", "M12 6v12m6-6H6");
        } else {
          iconPath.setAttribute("d", "M18 12H6");
        }
      }
    });
  });

  // Hide all answers initially
  const allAnswers = document.querySelectorAll(".qa-answer");
  allAnswers.forEach(function (answer) {
    answer.classList.add("hidden");
  });
});
