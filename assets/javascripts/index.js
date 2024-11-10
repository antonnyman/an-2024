function attachEvent(selector, event, fn) {
  const matches = document.querySelectorAll(selector)
  if (matches?.length > 0) {
    for (const element of matches) {
      element.addEventListener(
        event,
        () => {
          fn(element)
        },
        false
      )
    }
  }
}
window.onload = () => {
  attachEvent("[data-click-to-copy]", "click", (element) => {
    const email = "anton@antonnyman.se"
    navigator.clipboard.writeText(email).then(() => {
      document.getElementById("copy").classList.add("hidden")
      document.getElementById("copy").classList.remove("copy")
      document.getElementById("copied").classList.remove("hidden")
      document.getElementById("copied").classList.add("copy")
    })
    console.log("Click")
  })
}
