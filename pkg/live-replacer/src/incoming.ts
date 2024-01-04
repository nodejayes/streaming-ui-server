function addClass(el: Element, classDescription: string) {
  if (!el || classDescription?.length <= 0) {
    return;
  }
  for (const cl of classDescription.split(" ")) {
    el.classList.add(cl);
  }
}

function removeClass(el: Element, classDescription: string) {
  if (!el || classDescription?.length <= 0) {
    return;
  }
  for (const cl of classDescription.split(" ")) {
    el.classList.remove(cl);
  }
}

export async function replaceElements(action: Action<string>) {
  if (action.type.startsWith("replaceHtml::")) {
    const options = action.type.split("::");
    const selector = options[1];
    const elements = document.querySelectorAll(selector);
    if (!elements) {
      return;
    }
    const animation = JSON.parse(options[2] ?? "null");
    elements.forEach((element) => {
      if (!element) {
        return;
      }
      if (animation?.name && animation?.duration) {
        if (animation?.moment === "start") {
          addClass(element, animation.name);
          setTimeout(() => {
            element.innerHTML = action.payload ?? "";
            removeClass(element, animation.name);
          }, animation.duration);
          return;
        }
        if (animation?.moment === "end") {
          element.innerHTML = action.payload ?? "";
          addClass(element, animation.name);
          setTimeout(() => {
            removeClass(element, animation.name);
          }, animation.duration);
          return;
        }
      }
      element.innerHTML = action.payload ?? "";
    });
  }
  if (action.type.startsWith("redirect::")) {
    const url = action.payload;
    if (url) {
      window.location.href = url;
    }
  }
  if (action.type.startsWith("alert::")) {
    alert(action.payload ?? "unknown alert!");
  }
}
