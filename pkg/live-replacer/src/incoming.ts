export async function replaceElements(action: Action<string>) {
    if (action.type.startsWith("replaceHtml::")) {
        const selector = action.type.split("::")[1];
        const elements = document.querySelectorAll(selector);
        if (!elements) {
            return;
        }
        elements.forEach((element) => {
            if (!element) {
                return;
            }
            element.innerHTML = action.payload ?? "";
        });
    }
}