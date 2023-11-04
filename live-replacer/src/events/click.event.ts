import {ClickEventData, EventData} from "./types";

export class ClickEvent {
    private static api: WsApi | null = null;
    private static eventHandler = (e: Event) => {
        console.info(e);
        const event = e as PointerEvent;
        if (!event) {
            console.warn("click event is not of type PointerEvent", e);
            return;
        }
        if (!event.target) {
            console.warn("missing click event target", e);
            return;
        }
        ClickEvent.handle(event.target as HTMLElement, ClickEvent.typ, {
            typ: ClickEvent.typ,
            ctrlKey: event.ctrlKey,
        } as ClickEventData);
    }

    private static handle(target: HTMLElement | null, eventName: string, eventData: EventData) {
        if (!target) {
            console.warn(`no element on target event ${eventName}`);
            return;
        }
        const actionName = `lr${eventName}action`;
        const payloadName = `lr${eventName}payload`;
        const inputsName = `lr${eventName}inputs`;
        const action = target.getAttribute(actionName);
        if (!action) {
            console.warn(`missing ${actionName} on element`, target);
            return;
        }
        const elementId = target.getAttribute("id") ?? null;
        const payload = target.getAttribute(payloadName) ?? null;
        const inputSelectors: string | null = target.getAttribute(inputsName) ?? null;
        if (!payload && !inputSelectors) {
            ClickEvent.api?.send({
                action: {
                    elementId,
                    type: action,
                    payload: null,
                    inputs: {},
                },
                eventData,
            });
            return;
        }
        if (!inputSelectors) {
            ClickEvent.api?.send({
                action: {
                    elementId,
                    type: action,
                    payload: JSON.parse(payload ?? "null"),
                    inputs: {},
                },
                eventData,
            });
            return;
        }

        const selectors = inputSelectors.split("<=>");
        const inputData: { [key: string]: { [key: string]: string } } = {};
        for (const selector of selectors) {
            document.querySelectorAll(selector)?.forEach((el) => {
                el.querySelectorAll("input")?.forEach((input: HTMLInputElement) => {
                    const inputName = input.getAttribute("name");
                    if (!inputName) {
                        return;
                    }
                    if (!inputData[selector]) {
                        inputData[selector] = {};
                    }
                    inputData[selector][inputName] = input.value;
                });
            });
        }
        ClickEvent.api?.send({
            action: {
                elementId,
                type: action,
                payload: JSON.parse(payload ?? "null"),
                inputs: {},
            },
            eventData,
        });
    }

    static get typ() {
        return "click";
    }

    constructor(api: WsApi) {
        ClickEvent.api = api;
        document.querySelectorAll(`[lr${ClickEvent.typ}action]`).forEach((el) => {
            el.removeEventListener(ClickEvent.typ, ClickEvent.eventHandler);
            el.addEventListener(ClickEvent.typ, ClickEvent.eventHandler);
        });
    }
}